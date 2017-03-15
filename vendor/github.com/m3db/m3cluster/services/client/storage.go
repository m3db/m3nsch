// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package client

import (
	placementproto "github.com/m3db/m3cluster/generated/proto/placement"
	"github.com/m3db/m3cluster/services"
	"github.com/m3db/m3cluster/services/placement"
)

// newPlacementStorage returns a client of placement.Storage
func newPlacementStorage(opts Options) (placement.Storage, error) {
	if opts.KVGen() == nil {
		return nil, errNoKVGen
	}

	return &client{
		kvManagers: map[string]*kvManager{},
		opts:       opts,
	}, nil
}

func (s *client) CheckAndSet(sid services.ServiceID, p services.ServicePlacement, version int) error {
	if err := validateServiceID(sid); err != nil {
		return err
	}

	placementProto, err := PlacementToProto(p)
	if err != nil {
		return err
	}

	kvm, err := s.getKVManager(sid.Zone())
	if err != nil {
		return err
	}

	_, err = kvm.kv.CheckAndSet(
		placementKey(sid),
		version,
		&placementProto,
	)
	return err
}

func (s *client) SetIfNotExist(sid services.ServiceID, p services.ServicePlacement) error {
	if err := validateServiceID(sid); err != nil {
		return err
	}

	placementProto, err := PlacementToProto(p)
	if err != nil {
		return err
	}

	kvm, err := s.getKVManager(sid.Zone())
	if err != nil {
		return err
	}

	_, err = kvm.kv.SetIfNotExists(
		placementKey(sid),
		&placementProto,
	)
	return err
}

func (s *client) Delete(sid services.ServiceID) error {
	if err := validateServiceID(sid); err != nil {
		return err
	}

	kvm, err := s.getKVManager(sid.Zone())
	if err != nil {
		return err
	}

	_, err = kvm.kv.Delete(placementKey(sid))
	return err
}

func (s *client) Placement(sid services.ServiceID) (services.ServicePlacement, int, error) {
	if err := validateServiceID(sid); err != nil {
		return nil, 0, err
	}

	v, err := s.getPlacementValue(sid)
	if err != nil {
		return nil, 0, err
	}

	var placementProto placementproto.Placement
	if err := v.Unmarshal(&placementProto); err != nil {
		return nil, 0, err
	}

	p, err := PlacementFromProto(placementProto)
	return p, v.Version(), err
}