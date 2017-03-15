// Code generated by protoc-gen-go.
// source: persist_fs_commitlog.proto
// DO NOT EDIT!

package schema

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CommitLogInfo struct {
	Start    int64 `protobuf:"varint,1,opt,name=start" json:"start,omitempty"`
	Duration int64 `protobuf:"varint,2,opt,name=duration" json:"duration,omitempty"`
	Index    int64 `protobuf:"varint,3,opt,name=index" json:"index,omitempty"`
}

func (m *CommitLogInfo) Reset()                    { *m = CommitLogInfo{} }
func (m *CommitLogInfo) String() string            { return proto.CompactTextString(m) }
func (*CommitLogInfo) ProtoMessage()               {}
func (*CommitLogInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type CommitLog struct {
	Created    int64   `protobuf:"varint,1,opt,name=created" json:"created,omitempty"`
	Index      uint64  `protobuf:"varint,2,opt,name=index" json:"index,omitempty"`
	Metadata   []byte  `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Timestamp  int64   `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
	Value      float64 `protobuf:"fixed64,5,opt,name=value" json:"value,omitempty"`
	Unit       uint32  `protobuf:"varint,6,opt,name=unit" json:"unit,omitempty"`
	Annotation []byte  `protobuf:"bytes,7,opt,name=annotation,proto3" json:"annotation,omitempty"`
}

func (m *CommitLog) Reset()                    { *m = CommitLog{} }
func (m *CommitLog) String() string            { return proto.CompactTextString(m) }
func (*CommitLog) ProtoMessage()               {}
func (*CommitLog) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type CommitLogMetadata struct {
	Id        []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Namespace []byte `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Shard     uint32 `protobuf:"varint,3,opt,name=shard" json:"shard,omitempty"`
}

func (m *CommitLogMetadata) Reset()                    { *m = CommitLogMetadata{} }
func (m *CommitLogMetadata) String() string            { return proto.CompactTextString(m) }
func (*CommitLogMetadata) ProtoMessage()               {}
func (*CommitLogMetadata) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func init() {
	proto.RegisterType((*CommitLogInfo)(nil), "schema.CommitLogInfo")
	proto.RegisterType((*CommitLog)(nil), "schema.CommitLog")
	proto.RegisterType((*CommitLogMetadata)(nil), "schema.CommitLogMetadata")
}

var fileDescriptor1 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0x3f, 0x6b, 0xc3, 0x30,
	0x10, 0xc5, 0x51, 0xfe, 0x38, 0xcd, 0x61, 0x17, 0x2a, 0x3a, 0x88, 0x50, 0x8a, 0xf1, 0xe4, 0xa9,
	0x4b, 0x3f, 0x42, 0xa7, 0x42, 0xbb, 0x68, 0xc9, 0x18, 0xae, 0x96, 0x92, 0x08, 0x22, 0xc9, 0x48,
	0xe7, 0xd2, 0x6f, 0xd7, 0xaf, 0x56, 0x2c, 0x25, 0x76, 0x37, 0xfd, 0xee, 0x78, 0xef, 0xe9, 0x49,
	0xb0, 0xeb, 0x75, 0x88, 0x26, 0xd2, 0xe1, 0x18, 0x0f, 0x9d, 0xb7, 0xd6, 0xd0, 0xc5, 0x9f, 0x5e,
	0xfa, 0xe0, 0xc9, 0xf3, 0x22, 0x76, 0x67, 0x6d, 0xb1, 0xd9, 0x43, 0xf5, 0x96, 0x56, 0x1f, 0xfe,
	0xf4, 0xee, 0x8e, 0x9e, 0x3f, 0xc2, 0x3a, 0x12, 0x06, 0x12, 0xac, 0x66, 0xed, 0x52, 0x66, 0xe0,
	0x3b, 0xb8, 0x53, 0x43, 0x40, 0x32, 0xde, 0x89, 0x45, 0x5a, 0x4c, 0x3c, 0x2a, 0x8c, 0x53, 0xfa,
	0x47, 0x2c, 0xb3, 0x22, 0x41, 0xf3, 0xcb, 0x60, 0x3b, 0x39, 0x73, 0x01, 0x9b, 0x2e, 0x68, 0x24,
	0xad, 0xae, 0xbe, 0x37, 0x9c, 0xd5, 0xa3, 0xed, 0xea, 0xaa, 0x1e, 0xf3, 0xac, 0x26, 0x54, 0x48,
	0x98, 0x6c, 0x4b, 0x39, 0x31, 0x7f, 0x82, 0x2d, 0x19, 0xab, 0x23, 0xa1, 0xed, 0xc5, 0x2a, 0xb9,
	0xcd, 0x83, 0xd1, 0xef, 0x1b, 0x2f, 0x83, 0x16, 0xeb, 0x9a, 0xb5, 0x4c, 0x66, 0xe0, 0x1c, 0x56,
	0x83, 0x33, 0x24, 0x8a, 0x9a, 0xb5, 0x95, 0x4c, 0x67, 0xfe, 0x0c, 0x80, 0xce, 0x79, 0xca, 0xad,
	0x36, 0x29, 0xe5, 0xdf, 0xa4, 0xd9, 0xc3, 0xc3, 0x54, 0xe0, 0xf3, 0x16, 0x7e, 0x0f, 0x0b, 0x93,
	0x3b, 0x94, 0x72, 0x61, 0xd4, 0x78, 0x19, 0x87, 0x56, 0xc7, 0x1e, 0x3b, 0x9d, 0x2a, 0x94, 0x72,
	0x1e, 0xa4, 0xc7, 0x3c, 0x63, 0x50, 0xa9, 0x43, 0x25, 0x33, 0x7c, 0x15, 0xe9, 0x0b, 0x5e, 0xff,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xcd, 0xb3, 0x9e, 0xe2, 0xa0, 0x01, 0x00, 0x00,
}
