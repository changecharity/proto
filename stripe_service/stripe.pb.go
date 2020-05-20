// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.3
// source: stripe.proto

//option go_package = "github.com/chezky/changeproto/db_service";

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

type CreateCusOrAcc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StripeToken string `protobuf:"bytes,1,opt,name=stripe_token,json=stripeToken,proto3" json:"stripe_token,omitempty"`
	Email       string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *CreateCusOrAcc) Reset() {
	*x = CreateCusOrAcc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stripe_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCusOrAcc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCusOrAcc) ProtoMessage() {}

func (x *CreateCusOrAcc) ProtoReflect() protoreflect.Message {
	mi := &file_stripe_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCusOrAcc.ProtoReflect.Descriptor instead.
func (*CreateCusOrAcc) Descriptor() ([]byte, []int) {
	return file_stripe_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCusOrAcc) GetStripeToken() string {
	if x != nil {
		return x.StripeToken
	}
	return ""
}

func (x *CreateCusOrAcc) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type CreateCusReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *CreateCusReply) Reset() {
	*x = CreateCusReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stripe_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCusReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCusReply) ProtoMessage() {}

func (x *CreateCusReply) ProtoReflect() protoreflect.Message {
	mi := &file_stripe_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCusReply.ProtoReflect.Descriptor instead.
func (*CreateCusReply) Descriptor() ([]byte, []int) {
	return file_stripe_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCusReply) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

type CreateAccReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *CreateAccReply) Reset() {
	*x = CreateAccReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stripe_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccReply) ProtoMessage() {}

func (x *CreateAccReply) ProtoReflect() protoreflect.Message {
	mi := &file_stripe_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccReply.ProtoReflect.Descriptor instead.
func (*CreateAccReply) Descriptor() ([]byte, []int) {
	return file_stripe_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAccReply) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type WHCreateTransferData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	AccountId  string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Amount     int32  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *WHCreateTransferData) Reset() {
	*x = WHCreateTransferData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stripe_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WHCreateTransferData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WHCreateTransferData) ProtoMessage() {}

func (x *WHCreateTransferData) ProtoReflect() protoreflect.Message {
	mi := &file_stripe_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WHCreateTransferData.ProtoReflect.Descriptor instead.
func (*WHCreateTransferData) Descriptor() ([]byte, []int) {
	return file_stripe_proto_rawDescGZIP(), []int{3}
}

func (x *WHCreateTransferData) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *WHCreateTransferData) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *WHCreateTransferData) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type WHCreateTransferDataReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChargeId   string `protobuf:"bytes,1,opt,name=charge_id,json=chargeId,proto3" json:"charge_id,omitempty"`
	TransferId string `protobuf:"bytes,2,opt,name=transfer_id,json=transferId,proto3" json:"transfer_id,omitempty"`
}

func (x *WHCreateTransferDataReply) Reset() {
	*x = WHCreateTransferDataReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stripe_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WHCreateTransferDataReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WHCreateTransferDataReply) ProtoMessage() {}

func (x *WHCreateTransferDataReply) ProtoReflect() protoreflect.Message {
	mi := &file_stripe_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WHCreateTransferDataReply.ProtoReflect.Descriptor instead.
func (*WHCreateTransferDataReply) Descriptor() ([]byte, []int) {
	return file_stripe_proto_rawDescGZIP(), []int{4}
}

func (x *WHCreateTransferDataReply) GetChargeId() string {
	if x != nil {
		return x.ChargeId
	}
	return ""
}

func (x *WHCreateTransferDataReply) GetTransferId() string {
	if x != nil {
		return x.TransferId
	}
	return ""
}

var File_stripe_proto protoreflect.FileDescriptor

var file_stripe_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x75, 0x73, 0x4f, 0x72, 0x41, 0x63, 0x63, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x70,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73,
	0x74, 0x72, 0x69, 0x70, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x22, 0x31, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x64, 0x22, 0x6e, 0x0a, 0x14, 0x57, 0x48, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x59, 0x0a, 0x19, 0x57, 0x48, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x64, 0x32,
	0x4f, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73,
	0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x75, 0x73, 0x4f, 0x72, 0x41, 0x63, 0x63, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x32, 0x4f, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x75, 0x73, 0x4f, 0x72, 0x41, 0x63, 0x63, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x32, 0x72, 0x0a, 0x19, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x48, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55,
	0x0a, 0x12, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x57, 0x48, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x48, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x48, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stripe_proto_rawDescOnce sync.Once
	file_stripe_proto_rawDescData = file_stripe_proto_rawDesc
)

func file_stripe_proto_rawDescGZIP() []byte {
	file_stripe_proto_rawDescOnce.Do(func() {
		file_stripe_proto_rawDescData = protoimpl.X.CompressGZIP(file_stripe_proto_rawDescData)
	})
	return file_stripe_proto_rawDescData
}

var file_stripe_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_stripe_proto_goTypes = []interface{}{
	(*CreateCusOrAcc)(nil),            // 0: proto.CreateCusOrAcc
	(*CreateCusReply)(nil),            // 1: proto.CreateCusReply
	(*CreateAccReply)(nil),            // 2: proto.CreateAccReply
	(*WHCreateTransferData)(nil),      // 3: proto.WHCreateTransferData
	(*WHCreateTransferDataReply)(nil), // 4: proto.WHCreateTransferDataReply
}
var file_stripe_proto_depIdxs = []int32{
	0, // 0: proto.CreateCusService.CreateCus:input_type -> proto.CreateCusOrAcc
	0, // 1: proto.CreateAccService.CreateAcc:input_type -> proto.CreateCusOrAcc
	3, // 2: proto.FinalTransactionWHService.FinalTransactionWH:input_type -> proto.WHCreateTransferData
	1, // 3: proto.CreateCusService.CreateCus:output_type -> proto.CreateCusReply
	2, // 4: proto.CreateAccService.CreateAcc:output_type -> proto.CreateAccReply
	4, // 5: proto.FinalTransactionWHService.FinalTransactionWH:output_type -> proto.WHCreateTransferDataReply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stripe_proto_init() }
func file_stripe_proto_init() {
	if File_stripe_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stripe_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCusOrAcc); i {
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
		file_stripe_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCusReply); i {
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
		file_stripe_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccReply); i {
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
		file_stripe_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WHCreateTransferData); i {
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
		file_stripe_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WHCreateTransferDataReply); i {
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
			RawDescriptor: file_stripe_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_stripe_proto_goTypes,
		DependencyIndexes: file_stripe_proto_depIdxs,
		MessageInfos:      file_stripe_proto_msgTypes,
	}.Build()
	File_stripe_proto = out.File
	file_stripe_proto_rawDesc = nil
	file_stripe_proto_goTypes = nil
	file_stripe_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CreateCusServiceClient is the client API for CreateCusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CreateCusServiceClient interface {
	CreateCus(ctx context.Context, in *CreateCusOrAcc, opts ...grpc.CallOption) (*CreateCusReply, error)
}

type createCusServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCreateCusServiceClient(cc grpc.ClientConnInterface) CreateCusServiceClient {
	return &createCusServiceClient{cc}
}

func (c *createCusServiceClient) CreateCus(ctx context.Context, in *CreateCusOrAcc, opts ...grpc.CallOption) (*CreateCusReply, error) {
	out := new(CreateCusReply)
	err := c.cc.Invoke(ctx, "/proto.CreateCusService/CreateCus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreateCusServiceServer is the server API for CreateCusService service.
type CreateCusServiceServer interface {
	CreateCus(context.Context, *CreateCusOrAcc) (*CreateCusReply, error)
}

// UnimplementedCreateCusServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCreateCusServiceServer struct {
}

func (*UnimplementedCreateCusServiceServer) CreateCus(context.Context, *CreateCusOrAcc) (*CreateCusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCus not implemented")
}

func RegisterCreateCusServiceServer(s *grpc.Server, srv CreateCusServiceServer) {
	s.RegisterService(&_CreateCusService_serviceDesc, srv)
}

func _CreateCusService_CreateCus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCusOrAcc)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateCusServiceServer).CreateCus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CreateCusService/CreateCus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateCusServiceServer).CreateCus(ctx, req.(*CreateCusOrAcc))
	}
	return interceptor(ctx, in, info, handler)
}

var _CreateCusService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CreateCusService",
	HandlerType: (*CreateCusServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCus",
			Handler:    _CreateCusService_CreateCus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stripe.proto",
}

// CreateAccServiceClient is the client API for CreateAccService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CreateAccServiceClient interface {
	CreateAcc(ctx context.Context, in *CreateCusOrAcc, opts ...grpc.CallOption) (*CreateAccReply, error)
}

type createAccServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCreateAccServiceClient(cc grpc.ClientConnInterface) CreateAccServiceClient {
	return &createAccServiceClient{cc}
}

func (c *createAccServiceClient) CreateAcc(ctx context.Context, in *CreateCusOrAcc, opts ...grpc.CallOption) (*CreateAccReply, error) {
	out := new(CreateAccReply)
	err := c.cc.Invoke(ctx, "/proto.CreateAccService/CreateAcc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreateAccServiceServer is the server API for CreateAccService service.
type CreateAccServiceServer interface {
	CreateAcc(context.Context, *CreateCusOrAcc) (*CreateAccReply, error)
}

// UnimplementedCreateAccServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCreateAccServiceServer struct {
}

func (*UnimplementedCreateAccServiceServer) CreateAcc(context.Context, *CreateCusOrAcc) (*CreateAccReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAcc not implemented")
}

func RegisterCreateAccServiceServer(s *grpc.Server, srv CreateAccServiceServer) {
	s.RegisterService(&_CreateAccService_serviceDesc, srv)
}

func _CreateAccService_CreateAcc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCusOrAcc)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateAccServiceServer).CreateAcc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CreateAccService/CreateAcc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateAccServiceServer).CreateAcc(ctx, req.(*CreateCusOrAcc))
	}
	return interceptor(ctx, in, info, handler)
}

var _CreateAccService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CreateAccService",
	HandlerType: (*CreateAccServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAcc",
			Handler:    _CreateAccService_CreateAcc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stripe.proto",
}

// FinalTransactionWHServiceClient is the client API for FinalTransactionWHService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FinalTransactionWHServiceClient interface {
	FinalTransactionWH(ctx context.Context, in *WHCreateTransferData, opts ...grpc.CallOption) (*WHCreateTransferDataReply, error)
}

type finalTransactionWHServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFinalTransactionWHServiceClient(cc grpc.ClientConnInterface) FinalTransactionWHServiceClient {
	return &finalTransactionWHServiceClient{cc}
}

func (c *finalTransactionWHServiceClient) FinalTransactionWH(ctx context.Context, in *WHCreateTransferData, opts ...grpc.CallOption) (*WHCreateTransferDataReply, error) {
	out := new(WHCreateTransferDataReply)
	err := c.cc.Invoke(ctx, "/proto.FinalTransactionWHService/FinalTransactionWH", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FinalTransactionWHServiceServer is the server API for FinalTransactionWHService service.
type FinalTransactionWHServiceServer interface {
	FinalTransactionWH(context.Context, *WHCreateTransferData) (*WHCreateTransferDataReply, error)
}

// UnimplementedFinalTransactionWHServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFinalTransactionWHServiceServer struct {
}

func (*UnimplementedFinalTransactionWHServiceServer) FinalTransactionWH(context.Context, *WHCreateTransferData) (*WHCreateTransferDataReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinalTransactionWH not implemented")
}

func RegisterFinalTransactionWHServiceServer(s *grpc.Server, srv FinalTransactionWHServiceServer) {
	s.RegisterService(&_FinalTransactionWHService_serviceDesc, srv)
}

func _FinalTransactionWHService_FinalTransactionWH_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WHCreateTransferData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FinalTransactionWHServiceServer).FinalTransactionWH(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.FinalTransactionWHService/FinalTransactionWH",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FinalTransactionWHServiceServer).FinalTransactionWH(ctx, req.(*WHCreateTransferData))
	}
	return interceptor(ctx, in, info, handler)
}

var _FinalTransactionWHService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.FinalTransactionWHService",
	HandlerType: (*FinalTransactionWHServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FinalTransactionWH",
			Handler:    _FinalTransactionWHService_FinalTransactionWH_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stripe.proto",
}
