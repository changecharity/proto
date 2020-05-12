// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.3
// source: aws.proto

//option go_package = ".;proto";

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type QRData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URL string `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`
}

func (x *QRData) Reset() {
	*x = QRData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aws_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QRData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QRData) ProtoMessage() {}

func (x *QRData) ProtoReflect() protoreflect.Message {
	mi := &file_aws_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QRData.ProtoReflect.Descriptor instead.
func (*QRData) Descriptor() ([]byte, []int) {
	return file_aws_proto_rawDescGZIP(), []int{0}
}

func (x *QRData) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

type QRDataReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
}

func (x *QRDataReply) Reset() {
	*x = QRDataReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aws_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QRDataReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QRDataReply) ProtoMessage() {}

func (x *QRDataReply) ProtoReflect() protoreflect.Message {
	mi := &file_aws_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QRDataReply.ProtoReflect.Descriptor instead.
func (*QRDataReply) Descriptor() ([]byte, []int) {
	return file_aws_proto_rawDescGZIP(), []int{1}
}

func (x *QRDataReply) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

var File_aws_proto protoreflect.FileDescriptor

var file_aws_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x1a, 0x0a, 0x06, 0x51, 0x52, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03,
	0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x4c, 0x22, 0x29,
	0x0a, 0x0b, 0x51, 0x52, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x4a, 0x0a, 0x13, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x33, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x52, 0x44, 0x61, 0x74, 0x61, 0x1a,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x52, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aws_proto_rawDescOnce sync.Once
	file_aws_proto_rawDescData = file_aws_proto_rawDesc
)

func file_aws_proto_rawDescGZIP() []byte {
	file_aws_proto_rawDescOnce.Do(func() {
		file_aws_proto_rawDescData = protoimpl.X.CompressGZIP(file_aws_proto_rawDescData)
	})
	return file_aws_proto_rawDescData
}

var file_aws_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_aws_proto_goTypes = []interface{}{
	(*QRData)(nil),      // 0: proto.QRData
	(*QRDataReply)(nil), // 1: proto.QRDataReply
}
var file_aws_proto_depIdxs = []int32{
	0, // 0: proto.CreateQRCodeService.CreateQRCode:input_type -> proto.QRData
	1, // 1: proto.CreateQRCodeService.CreateQRCode:output_type -> proto.QRDataReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_aws_proto_init() }
func file_aws_proto_init() {
	if File_aws_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aws_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QRData); i {
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
		file_aws_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QRDataReply); i {
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
			RawDescriptor: file_aws_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aws_proto_goTypes,
		DependencyIndexes: file_aws_proto_depIdxs,
		MessageInfos:      file_aws_proto_msgTypes,
	}.Build()
	File_aws_proto = out.File
	file_aws_proto_rawDesc = nil
	file_aws_proto_goTypes = nil
	file_aws_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CreateQRCodeServiceClient is the client API for CreateQRCodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CreateQRCodeServiceClient interface {
	CreateQRCode(ctx context.Context, in *QRData, opts ...grpc.CallOption) (*QRDataReply, error)
}

type createQRCodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCreateQRCodeServiceClient(cc grpc.ClientConnInterface) CreateQRCodeServiceClient {
	return &createQRCodeServiceClient{cc}
}

func (c *createQRCodeServiceClient) CreateQRCode(ctx context.Context, in *QRData, opts ...grpc.CallOption) (*QRDataReply, error) {
	out := new(QRDataReply)
	err := c.cc.Invoke(ctx, "/proto.CreateQRCodeService/CreateQRCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreateQRCodeServiceServer is the server API for CreateQRCodeService service.
type CreateQRCodeServiceServer interface {
	CreateQRCode(context.Context, *QRData) (*QRDataReply, error)
}

// UnimplementedCreateQRCodeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCreateQRCodeServiceServer struct {
}

func (*UnimplementedCreateQRCodeServiceServer) CreateQRCode(context.Context, *QRData) (*QRDataReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQRCode not implemented")
}

func RegisterCreateQRCodeServiceServer(s *grpc.Server, srv CreateQRCodeServiceServer) {
	s.RegisterService(&_CreateQRCodeService_serviceDesc, srv)
}

func _CreateQRCodeService_CreateQRCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QRData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateQRCodeServiceServer).CreateQRCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CreateQRCodeService/CreateQRCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateQRCodeServiceServer).CreateQRCode(ctx, req.(*QRData))
	}
	return interceptor(ctx, in, info, handler)
}

var _CreateQRCodeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CreateQRCodeService",
	HandlerType: (*CreateQRCodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateQRCode",
			Handler:    _CreateQRCodeService_CreateQRCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aws.proto",
}
