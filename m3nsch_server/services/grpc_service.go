package services

import (
	"github.com/m3db/m3nsch"
	"github.com/m3db/m3nsch/rpc"
	convert "github.com/m3db/m3nsch/rpc/convert"

	xlog "github.com/m3db/m3x/log"

	"github.com/uber-go/tally"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// NewGRPCService returns a new GRPCService wrapping a m3nsch.Agent
func NewGRPCService(
	agent m3nsch.Agent,
	metricsScope tally.Scope,
	logger xlog.Logger,
) (rpc.MenschServer, error) {
	return &menschServer{
		agent:  agent,
		scope:  metricsScope.SubScope("grpc-server"),
		logger: logger,
	}, nil
}

type menschServer struct {
	scope  tally.Scope
	logger xlog.Logger
	agent  m3nsch.Agent
}

func (ms *menschServer) Status(ctx context.Context, req *rpc.StatusRequest) (*rpc.StatusResponse, error) {
	status := ms.agent.Status()
	workload := convert.ToProtoWorkload(ms.agent.Workload())
	response := &rpc.StatusResponse{
		Token:    status.Token,
		Status:   convert.ToProtoStatus(status.Status),
		MaxQPS:   ms.agent.MaxQPS(),
		Workload: &workload,
	}
	return response, nil
}

func (ms *menschServer) Init(ctx context.Context, req *rpc.InitRequest) (*rpc.InitResponse, error) {
	if req == nil {
		return nil, grpc.Errorf(codes.InvalidArgument, "nil request")
	}

	ms.logger.Debugf("received init request: %v", req.String())
	workload, err := convert.ToM3nschWorkload(req.GetWorkload())
	if err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, "unable to parse workload: %v", err)
	}

	err = ms.agent.Init(req.GetToken(), workload, req.GetForce(),
		req.GetTargetZone(), req.GetTargetEnv())
	if err != nil {
		return nil, grpc.Errorf(codes.Unavailable, err.Error())
	}

	return &rpc.InitResponse{}, nil
}

func (ms *menschServer) Start(context.Context, *rpc.StartRequest) (*rpc.StartResponse, error) {
	ms.logger.Debugf("received Start() request")
	if err := ms.agent.Start(); err != nil {
		return nil, grpc.Errorf(codes.Unknown, err.Error())
	}
	return &rpc.StartResponse{}, nil
}

func (ms *menschServer) Stop(context.Context, *rpc.StopRequest) (*rpc.StopResponse, error) {
	ms.logger.Debugf("received Stop() request")
	if err := ms.agent.Stop(); err != nil {
		return nil, grpc.Errorf(codes.Unknown, err.Error())
	}
	return &rpc.StopResponse{}, nil
}

func (ms *menschServer) Modify(_ context.Context, req *rpc.ModifyRequest) (*rpc.ModifyResponse, error) {
	ms.logger.Debugf("received Modify() request: %v", req.String())
	if req == nil {
		return nil, grpc.Errorf(codes.InvalidArgument, "nil request")
	}
	workload, err := convert.ToM3nschWorkload(req.GetWorkload())
	if err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, "unable to parse workload: %v", err)
	}
	ms.agent.SetWorkload(workload)
	return &rpc.ModifyResponse{}, nil
}