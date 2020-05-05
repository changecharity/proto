// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.3
// source: plaid.proto

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

type ExchangePlaid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Publictoken    string `protobuf:"bytes,1,opt,name=publictoken,proto3" json:"publictoken,omitempty"`
	Plaidaccountid string `protobuf:"bytes,2,opt,name=plaidaccountid,proto3" json:"plaidaccountid,omitempty"`
}

func (x *ExchangePlaid) Reset() {
	*x = ExchangePlaid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plaid_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangePlaid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangePlaid) ProtoMessage() {}

func (x *ExchangePlaid) ProtoReflect() protoreflect.Message {
	mi := &file_plaid_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangePlaid.ProtoReflect.Descriptor instead.
func (*ExchangePlaid) Descriptor() ([]byte, []int) {
	return file_plaid_proto_rawDescGZIP(), []int{0}
}

func (x *ExchangePlaid) GetPublictoken() string {
	if x != nil {
		return x.Publictoken
	}
	return ""
}

func (x *ExchangePlaid) GetPlaidaccountid() string {
	if x != nil {
		return x.Plaidaccountid
	}
	return ""
}

type ExchangePlaidReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plaidaccesstoken string `protobuf:"bytes,1,opt,name=plaidaccesstoken,proto3" json:"plaidaccesstoken,omitempty"`
	Itemid           string `protobuf:"bytes,2,opt,name=itemid,proto3" json:"itemid,omitempty"`
	Stripetoken      string `protobuf:"bytes,3,opt,name=stripetoken,proto3" json:"stripetoken,omitempty"`
}

func (x *ExchangePlaidReply) Reset() {
	*x = ExchangePlaidReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plaid_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangePlaidReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangePlaidReply) ProtoMessage() {}

func (x *ExchangePlaidReply) ProtoReflect() protoreflect.Message {
	mi := &file_plaid_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangePlaidReply.ProtoReflect.Descriptor instead.
func (*ExchangePlaidReply) Descriptor() ([]byte, []int) {
	return file_plaid_proto_rawDescGZIP(), []int{1}
}

func (x *ExchangePlaidReply) GetPlaidaccesstoken() string {
	if x != nil {
		return x.Plaidaccesstoken
	}
	return ""
}

func (x *ExchangePlaidReply) GetItemid() string {
	if x != nil {
		return x.Itemid
	}
	return ""
}

func (x *ExchangePlaidReply) GetStripetoken() string {
	if x != nil {
		return x.Stripetoken
	}
	return ""
}

type TransactionWHData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plaidtoken string `protobuf:"bytes,1,opt,name=plaidtoken,proto3" json:"plaidtoken,omitempty"`
	Threshold  int32  `protobuf:"varint,2,opt,name=threshold,proto3" json:"threshold,omitempty"`
}

func (x *TransactionWHData) Reset() {
	*x = TransactionWHData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plaid_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionWHData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionWHData) ProtoMessage() {}

func (x *TransactionWHData) ProtoReflect() protoreflect.Message {
	mi := &file_plaid_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionWHData.ProtoReflect.Descriptor instead.
func (*TransactionWHData) Descriptor() ([]byte, []int) {
	return file_plaid_proto_rawDescGZIP(), []int{2}
}

func (x *TransactionWHData) GetPlaidtoken() string {
	if x != nil {
		return x.Plaidtoken
	}
	return ""
}

func (x *TransactionWHData) GetThreshold() int32 {
	if x != nil {
		return x.Threshold
	}
	return 0
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// float 64
	Amount float64 `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Dot    string  `protobuf:"bytes,4,opt,name=dot,proto3" json:"dot,omitempty"`
	Change int32   `protobuf:"varint,5,opt,name=change,proto3" json:"change,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plaid_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_plaid_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_plaid_proto_rawDescGZIP(), []int{3}
}

func (x *Transaction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Transaction) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Transaction) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Transaction) GetDot() string {
	if x != nil {
		return x.Dot
	}
	return ""
}

func (x *Transaction) GetChange() int32 {
	if x != nil {
		return x.Change
	}
	return 0
}

type TransactionWHDataReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transactions []*Transaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *TransactionWHDataReply) Reset() {
	*x = TransactionWHDataReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plaid_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionWHDataReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionWHDataReply) ProtoMessage() {}

func (x *TransactionWHDataReply) ProtoReflect() protoreflect.Message {
	mi := &file_plaid_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionWHDataReply.ProtoReflect.Descriptor instead.
func (*TransactionWHDataReply) Descriptor() ([]byte, []int) {
	return file_plaid_proto_rawDescGZIP(), []int{4}
}

func (x *TransactionWHDataReply) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

var File_plaid_proto protoreflect.FileDescriptor

var file_plaid_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x0d, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x50, 0x6c, 0x61, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x6c, 0x61, 0x69, 0x64,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x70, 0x6c, 0x61, 0x69, 0x64, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x64, 0x22,
	0x7a, 0x0a, 0x12, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x6c, 0x61, 0x69, 0x64,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2a, 0x0a, 0x10, 0x70, 0x6c, 0x61, 0x69, 0x64, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x70, 0x6c, 0x61, 0x69, 0x64, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x74, 0x72,
	0x69, 0x70, 0x65, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x51, 0x0a, 0x11, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x48, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x69, 0x64, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x69, 0x64, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x22, 0x73,
	0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x6f, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x6f, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x22, 0x50, 0x0a, 0x16, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x57, 0x48, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x36, 0x0a,
	0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0x66, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x6c, 0x61, 0x69, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x6c, 0x61, 0x69, 0x64, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x6c, 0x61, 0x69, 0x64,
	0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x50, 0x6c, 0x61, 0x69, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x32, 0x6e, 0x0a,
	0x1a, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x57, 0x48, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x13, 0x53,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x57, 0x48, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x48, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x1d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x57, 0x48, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plaid_proto_rawDescOnce sync.Once
	file_plaid_proto_rawDescData = file_plaid_proto_rawDesc
)

func file_plaid_proto_rawDescGZIP() []byte {
	file_plaid_proto_rawDescOnce.Do(func() {
		file_plaid_proto_rawDescData = protoimpl.X.CompressGZIP(file_plaid_proto_rawDescData)
	})
	return file_plaid_proto_rawDescData
}

var file_plaid_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_plaid_proto_goTypes = []interface{}{
	(*ExchangePlaid)(nil),          // 0: proto.ExchangePlaid
	(*ExchangePlaidReply)(nil),     // 1: proto.ExchangePlaidReply
	(*TransactionWHData)(nil),      // 2: proto.TransactionWHData
	(*Transaction)(nil),            // 3: proto.Transaction
	(*TransactionWHDataReply)(nil), // 4: proto.TransactionWHDataReply
}
var file_plaid_proto_depIdxs = []int32{
	3, // 0: proto.TransactionWHDataReply.transactions:type_name -> proto.Transaction
	0, // 1: proto.CreateExchangePlaidService.CreateExchangePlaid:input_type -> proto.ExchangePlaid
	2, // 2: proto.SecondTransactionWHService.SecondTransactionWH:input_type -> proto.TransactionWHData
	1, // 3: proto.CreateExchangePlaidService.CreateExchangePlaid:output_type -> proto.ExchangePlaidReply
	4, // 4: proto.SecondTransactionWHService.SecondTransactionWH:output_type -> proto.TransactionWHDataReply
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_plaid_proto_init() }
func file_plaid_proto_init() {
	if File_plaid_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_plaid_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangePlaid); i {
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
		file_plaid_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangePlaidReply); i {
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
		file_plaid_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionWHData); i {
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
		file_plaid_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_plaid_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionWHDataReply); i {
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
			RawDescriptor: file_plaid_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_plaid_proto_goTypes,
		DependencyIndexes: file_plaid_proto_depIdxs,
		MessageInfos:      file_plaid_proto_msgTypes,
	}.Build()
	File_plaid_proto = out.File
	file_plaid_proto_rawDesc = nil
	file_plaid_proto_goTypes = nil
	file_plaid_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CreateExchangePlaidServiceClient is the client API for CreateExchangePlaidService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CreateExchangePlaidServiceClient interface {
	CreateExchangePlaid(ctx context.Context, in *ExchangePlaid, opts ...grpc.CallOption) (*ExchangePlaidReply, error)
}

type createExchangePlaidServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCreateExchangePlaidServiceClient(cc grpc.ClientConnInterface) CreateExchangePlaidServiceClient {
	return &createExchangePlaidServiceClient{cc}
}

func (c *createExchangePlaidServiceClient) CreateExchangePlaid(ctx context.Context, in *ExchangePlaid, opts ...grpc.CallOption) (*ExchangePlaidReply, error) {
	out := new(ExchangePlaidReply)
	err := c.cc.Invoke(ctx, "/proto.CreateExchangePlaidService/CreateExchangePlaid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreateExchangePlaidServiceServer is the server API for CreateExchangePlaidService service.
type CreateExchangePlaidServiceServer interface {
	CreateExchangePlaid(context.Context, *ExchangePlaid) (*ExchangePlaidReply, error)
}

// UnimplementedCreateExchangePlaidServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCreateExchangePlaidServiceServer struct {
}

func (*UnimplementedCreateExchangePlaidServiceServer) CreateExchangePlaid(context.Context, *ExchangePlaid) (*ExchangePlaidReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExchangePlaid not implemented")
}

func RegisterCreateExchangePlaidServiceServer(s *grpc.Server, srv CreateExchangePlaidServiceServer) {
	s.RegisterService(&_CreateExchangePlaidService_serviceDesc, srv)
}

func _CreateExchangePlaidService_CreateExchangePlaid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExchangePlaid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateExchangePlaidServiceServer).CreateExchangePlaid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CreateExchangePlaidService/CreateExchangePlaid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateExchangePlaidServiceServer).CreateExchangePlaid(ctx, req.(*ExchangePlaid))
	}
	return interceptor(ctx, in, info, handler)
}

var _CreateExchangePlaidService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CreateExchangePlaidService",
	HandlerType: (*CreateExchangePlaidServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateExchangePlaid",
			Handler:    _CreateExchangePlaidService_CreateExchangePlaid_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plaid.proto",
}

// SecondTransactionWHServiceClient is the client API for SecondTransactionWHService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SecondTransactionWHServiceClient interface {
	SecondTransactionWH(ctx context.Context, in *TransactionWHData, opts ...grpc.CallOption) (*TransactionWHDataReply, error)
}

type secondTransactionWHServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSecondTransactionWHServiceClient(cc grpc.ClientConnInterface) SecondTransactionWHServiceClient {
	return &secondTransactionWHServiceClient{cc}
}

func (c *secondTransactionWHServiceClient) SecondTransactionWH(ctx context.Context, in *TransactionWHData, opts ...grpc.CallOption) (*TransactionWHDataReply, error) {
	out := new(TransactionWHDataReply)
	err := c.cc.Invoke(ctx, "/proto.SecondTransactionWHService/SecondTransactionWH", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecondTransactionWHServiceServer is the server API for SecondTransactionWHService service.
type SecondTransactionWHServiceServer interface {
	SecondTransactionWH(context.Context, *TransactionWHData) (*TransactionWHDataReply, error)
}

// UnimplementedSecondTransactionWHServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSecondTransactionWHServiceServer struct {
}

func (*UnimplementedSecondTransactionWHServiceServer) SecondTransactionWH(context.Context, *TransactionWHData) (*TransactionWHDataReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SecondTransactionWH not implemented")
}

func RegisterSecondTransactionWHServiceServer(s *grpc.Server, srv SecondTransactionWHServiceServer) {
	s.RegisterService(&_SecondTransactionWHService_serviceDesc, srv)
}

func _SecondTransactionWHService_SecondTransactionWH_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionWHData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecondTransactionWHServiceServer).SecondTransactionWH(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SecondTransactionWHService/SecondTransactionWH",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecondTransactionWHServiceServer).SecondTransactionWH(ctx, req.(*TransactionWHData))
	}
	return interceptor(ctx, in, info, handler)
}

var _SecondTransactionWHService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SecondTransactionWHService",
	HandlerType: (*SecondTransactionWHServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SecondTransactionWH",
			Handler:    _SecondTransactionWHService_SecondTransactionWH_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plaid.proto",
}
