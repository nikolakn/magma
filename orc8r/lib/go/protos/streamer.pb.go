//
//Copyright 2020 The Magma Authors.
//
//This source code is licensed under the BSD-style license found in the
//LICENSE file in the root directory of this source tree.
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.10.0
// source: orc8r/protos/streamer.proto

package protos

import (
	context "context"
	any "github.com/golang/protobuf/ptypes/any"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GatewayId string `protobuf:"bytes,1,opt,name=gatewayId,proto3" json:"gatewayId,omitempty"`
	// stream_name to attach to.
	// E.g., subscriberdb, config, etc.
	StreamName string `protobuf:"bytes,2,opt,name=stream_name,json=streamName,proto3" json:"stream_name,omitempty"`
	// extra_args contain any extra data to send up with the stream request.
	// This value will be different per stream provider.
	ExtraArgs *any.Any `protobuf:"bytes,3,opt,name=extra_args,json=extraArgs,proto3" json:"extra_args,omitempty"`
}

func (x *StreamRequest) Reset() {
	*x = StreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_streamer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamRequest) ProtoMessage() {}

func (x *StreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_streamer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamRequest.ProtoReflect.Descriptor instead.
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_streamer_proto_rawDescGZIP(), []int{0}
}

func (x *StreamRequest) GetGatewayId() string {
	if x != nil {
		return x.GatewayId
	}
	return ""
}

func (x *StreamRequest) GetStreamName() string {
	if x != nil {
		return x.StreamName
	}
	return ""
}

func (x *StreamRequest) GetExtraArgs() *any.Any {
	if x != nil {
		return x.ExtraArgs
	}
	return nil
}

type DataUpdateBatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// updates to config values
	Updates []*DataUpdate `protobuf:"bytes,1,rep,name=updates,proto3" json:"updates,omitempty"`
	// resync is true iff the updates would be a snapshot of all the contents
	// in the cloud.
	Resync bool `protobuf:"varint,2,opt,name=resync,proto3" json:"resync,omitempty"`
}

func (x *DataUpdateBatch) Reset() {
	*x = DataUpdateBatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_streamer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataUpdateBatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataUpdateBatch) ProtoMessage() {}

func (x *DataUpdateBatch) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_streamer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataUpdateBatch.ProtoReflect.Descriptor instead.
func (*DataUpdateBatch) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_streamer_proto_rawDescGZIP(), []int{1}
}

func (x *DataUpdateBatch) GetUpdates() []*DataUpdate {
	if x != nil {
		return x.Updates
	}
	return nil
}

func (x *DataUpdateBatch) GetResync() bool {
	if x != nil {
		return x.Resync
	}
	return false
}

type DataUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// key is the unique key for each item
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// value can be file contents, protobuf serialized message, etc.
	// For key deletions, the value field would be absent.
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *DataUpdate) Reset() {
	*x = DataUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_streamer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataUpdate) ProtoMessage() {}

func (x *DataUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_streamer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataUpdate.ProtoReflect.Descriptor instead.
func (*DataUpdate) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_streamer_proto_rawDescGZIP(), []int{2}
}

func (x *DataUpdate) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *DataUpdate) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_orc8r_protos_streamer_proto protoreflect.FileDescriptor

var file_orc8r_protos_streamer_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d,
	0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f,
	0x61, 0x72, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x52, 0x09, 0x65, 0x78, 0x74, 0x72, 0x61, 0x41, 0x72, 0x67, 0x73, 0x22, 0x5c, 0x0a, 0x0f, 0x44,
	0x61, 0x74, 0x61, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x31,
	0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x79, 0x6e, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x79, 0x6e, 0x63, 0x22, 0x34, 0x0a, 0x0a, 0x44, 0x61, 0x74,
	0x61, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32,
	0x56, 0x0a, 0x08, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x2e, 0x6d, 0x61, 0x67, 0x6d,
	0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72,
	0x63, 0x38, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x22, 0x00, 0x30, 0x01, 0x42, 0x1b, 0x5a, 0x19, 0x6d, 0x61, 0x67, 0x6d, 0x61,
	0x2f, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orc8r_protos_streamer_proto_rawDescOnce sync.Once
	file_orc8r_protos_streamer_proto_rawDescData = file_orc8r_protos_streamer_proto_rawDesc
)

func file_orc8r_protos_streamer_proto_rawDescGZIP() []byte {
	file_orc8r_protos_streamer_proto_rawDescOnce.Do(func() {
		file_orc8r_protos_streamer_proto_rawDescData = protoimpl.X.CompressGZIP(file_orc8r_protos_streamer_proto_rawDescData)
	})
	return file_orc8r_protos_streamer_proto_rawDescData
}

var file_orc8r_protos_streamer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_orc8r_protos_streamer_proto_goTypes = []interface{}{
	(*StreamRequest)(nil),   // 0: magma.orc8r.StreamRequest
	(*DataUpdateBatch)(nil), // 1: magma.orc8r.DataUpdateBatch
	(*DataUpdate)(nil),      // 2: magma.orc8r.DataUpdate
	(*any.Any)(nil),         // 3: google.protobuf.Any
}
var file_orc8r_protos_streamer_proto_depIdxs = []int32{
	3, // 0: magma.orc8r.StreamRequest.extra_args:type_name -> google.protobuf.Any
	2, // 1: magma.orc8r.DataUpdateBatch.updates:type_name -> magma.orc8r.DataUpdate
	0, // 2: magma.orc8r.Streamer.GetUpdates:input_type -> magma.orc8r.StreamRequest
	1, // 3: magma.orc8r.Streamer.GetUpdates:output_type -> magma.orc8r.DataUpdateBatch
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_orc8r_protos_streamer_proto_init() }
func file_orc8r_protos_streamer_proto_init() {
	if File_orc8r_protos_streamer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orc8r_protos_streamer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orc8r_protos_streamer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataUpdateBatch); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orc8r_protos_streamer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataUpdate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_orc8r_protos_streamer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orc8r_protos_streamer_proto_goTypes,
		DependencyIndexes: file_orc8r_protos_streamer_proto_depIdxs,
		MessageInfos:      file_orc8r_protos_streamer_proto_msgTypes,
	}.Build()
	File_orc8r_protos_streamer_proto = out.File
	file_orc8r_protos_streamer_proto_rawDesc = nil
	file_orc8r_protos_streamer_proto_goTypes = nil
	file_orc8r_protos_streamer_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StreamerClient is the client API for Streamer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamerClient interface {
	// GetUpdates streams config updates from the cloud.
	// The RPC call would be kept open to push new updates as they happen.
	GetUpdates(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (Streamer_GetUpdatesClient, error)
}

type streamerClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamerClient(cc grpc.ClientConnInterface) StreamerClient {
	return &streamerClient{cc}
}

func (c *streamerClient) GetUpdates(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (Streamer_GetUpdatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Streamer_serviceDesc.Streams[0], "/magma.orc8r.Streamer/GetUpdates", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamerGetUpdatesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Streamer_GetUpdatesClient interface {
	Recv() (*DataUpdateBatch, error)
	grpc.ClientStream
}

type streamerGetUpdatesClient struct {
	grpc.ClientStream
}

func (x *streamerGetUpdatesClient) Recv() (*DataUpdateBatch, error) {
	m := new(DataUpdateBatch)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamerServer is the server API for Streamer service.
type StreamerServer interface {
	// GetUpdates streams config updates from the cloud.
	// The RPC call would be kept open to push new updates as they happen.
	GetUpdates(*StreamRequest, Streamer_GetUpdatesServer) error
}

// UnimplementedStreamerServer can be embedded to have forward compatible implementations.
type UnimplementedStreamerServer struct {
}

func (*UnimplementedStreamerServer) GetUpdates(*StreamRequest, Streamer_GetUpdatesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUpdates not implemented")
}

func RegisterStreamerServer(s *grpc.Server, srv StreamerServer) {
	s.RegisterService(&_Streamer_serviceDesc, srv)
}

func _Streamer_GetUpdates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamerServer).GetUpdates(m, &streamerGetUpdatesServer{stream})
}

type Streamer_GetUpdatesServer interface {
	Send(*DataUpdateBatch) error
	grpc.ServerStream
}

type streamerGetUpdatesServer struct {
	grpc.ServerStream
}

func (x *streamerGetUpdatesServer) Send(m *DataUpdateBatch) error {
	return x.ServerStream.SendMsg(m)
}

var _Streamer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.Streamer",
	HandlerType: (*StreamerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUpdates",
			Handler:       _Streamer_GetUpdates_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "orc8r/protos/streamer.proto",
}
