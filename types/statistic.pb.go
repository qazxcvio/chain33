// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: statistic.proto

package types

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//手续费
type TotalFee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fee     int64 `protobuf:"varint,1,opt,name=fee,proto3" json:"fee,omitempty"`
	TxCount int64 `protobuf:"varint,2,opt,name=txCount,proto3" json:"txCount,omitempty"`
}

func (x *TotalFee) Reset() {
	*x = TotalFee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TotalFee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TotalFee) ProtoMessage() {}

func (x *TotalFee) ProtoReflect() protoreflect.Message {
	mi := &file_statistic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TotalFee.ProtoReflect.Descriptor instead.
func (*TotalFee) Descriptor() ([]byte, []int) {
	return file_statistic_proto_rawDescGZIP(), []int{0}
}

func (x *TotalFee) GetFee() int64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *TotalFee) GetTxCount() int64 {
	if x != nil {
		return x.TxCount
	}
	return 0
}

//查询symbol代币总额
type ReqGetTotalCoins struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol    string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	StateHash []byte `protobuf:"bytes,2,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	StartKey  []byte `protobuf:"bytes,3,opt,name=startKey,proto3" json:"startKey,omitempty"`
	Count     int64  `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	Execer    string `protobuf:"bytes,5,opt,name=execer,proto3" json:"execer,omitempty"`
}

func (x *ReqGetTotalCoins) Reset() {
	*x = ReqGetTotalCoins{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqGetTotalCoins) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqGetTotalCoins) ProtoMessage() {}

func (x *ReqGetTotalCoins) ProtoReflect() protoreflect.Message {
	mi := &file_statistic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqGetTotalCoins.ProtoReflect.Descriptor instead.
func (*ReqGetTotalCoins) Descriptor() ([]byte, []int) {
	return file_statistic_proto_rawDescGZIP(), []int{1}
}

func (x *ReqGetTotalCoins) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *ReqGetTotalCoins) GetStateHash() []byte {
	if x != nil {
		return x.StateHash
	}
	return nil
}

func (x *ReqGetTotalCoins) GetStartKey() []byte {
	if x != nil {
		return x.StartKey
	}
	return nil
}

func (x *ReqGetTotalCoins) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ReqGetTotalCoins) GetExecer() string {
	if x != nil {
		return x.Execer
	}
	return ""
}

//查询symbol代币总额应答
type ReplyGetTotalCoins struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count   int64  `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Num     int64  `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	Amount  int64  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	NextKey []byte `protobuf:"bytes,4,opt,name=nextKey,proto3" json:"nextKey,omitempty"`
}

func (x *ReplyGetTotalCoins) Reset() {
	*x = ReplyGetTotalCoins{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyGetTotalCoins) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyGetTotalCoins) ProtoMessage() {}

func (x *ReplyGetTotalCoins) ProtoReflect() protoreflect.Message {
	mi := &file_statistic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyGetTotalCoins.ProtoReflect.Descriptor instead.
func (*ReplyGetTotalCoins) Descriptor() ([]byte, []int) {
	return file_statistic_proto_rawDescGZIP(), []int{2}
}

func (x *ReplyGetTotalCoins) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ReplyGetTotalCoins) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *ReplyGetTotalCoins) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *ReplyGetTotalCoins) GetNextKey() []byte {
	if x != nil {
		return x.NextKey
	}
	return nil
}

//迭代查询symbol代币总额
type IterateRangeByStateHash struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StateHash []byte `protobuf:"bytes,1,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	Start     []byte `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty"`
	End       []byte `protobuf:"bytes,3,opt,name=end,proto3" json:"end,omitempty"`
	Count     int64  `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *IterateRangeByStateHash) Reset() {
	*x = IterateRangeByStateHash{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistic_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IterateRangeByStateHash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IterateRangeByStateHash) ProtoMessage() {}

func (x *IterateRangeByStateHash) ProtoReflect() protoreflect.Message {
	mi := &file_statistic_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IterateRangeByStateHash.ProtoReflect.Descriptor instead.
func (*IterateRangeByStateHash) Descriptor() ([]byte, []int) {
	return file_statistic_proto_rawDescGZIP(), []int{3}
}

func (x *IterateRangeByStateHash) GetStateHash() []byte {
	if x != nil {
		return x.StateHash
	}
	return nil
}

func (x *IterateRangeByStateHash) GetStart() []byte {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *IterateRangeByStateHash) GetEnd() []byte {
	if x != nil {
		return x.End
	}
	return nil
}

func (x *IterateRangeByStateHash) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type TotalAmount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 统计的总数
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *TotalAmount) Reset() {
	*x = TotalAmount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistic_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TotalAmount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TotalAmount) ProtoMessage() {}

func (x *TotalAmount) ProtoReflect() protoreflect.Message {
	mi := &file_statistic_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TotalAmount.ProtoReflect.Descriptor instead.
func (*TotalAmount) Descriptor() ([]byte, []int) {
	return file_statistic_proto_rawDescGZIP(), []int{4}
}

func (x *TotalAmount) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

//查询symbol在合约中的代币总额，如果execAddr为空，则为查询symbol在所有合约中的代币总额
type ReqGetExecBalance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol    string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	StateHash []byte `protobuf:"bytes,2,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	Addr      []byte `protobuf:"bytes,3,opt,name=addr,proto3" json:"addr,omitempty"`
	ExecAddr  []byte `protobuf:"bytes,4,opt,name=execAddr,proto3" json:"execAddr,omitempty"`
	Execer    string `protobuf:"bytes,5,opt,name=execer,proto3" json:"execer,omitempty"`
	Count     int64  `protobuf:"varint,6,opt,name=count,proto3" json:"count,omitempty"`
	NextKey   []byte `protobuf:"bytes,7,opt,name=nextKey,proto3" json:"nextKey,omitempty"`
}

func (x *ReqGetExecBalance) Reset() {
	*x = ReqGetExecBalance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistic_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqGetExecBalance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqGetExecBalance) ProtoMessage() {}

func (x *ReqGetExecBalance) ProtoReflect() protoreflect.Message {
	mi := &file_statistic_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqGetExecBalance.ProtoReflect.Descriptor instead.
func (*ReqGetExecBalance) Descriptor() ([]byte, []int) {
	return file_statistic_proto_rawDescGZIP(), []int{5}
}

func (x *ReqGetExecBalance) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *ReqGetExecBalance) GetStateHash() []byte {
	if x != nil {
		return x.StateHash
	}
	return nil
}

func (x *ReqGetExecBalance) GetAddr() []byte {
	if x != nil {
		return x.Addr
	}
	return nil
}

func (x *ReqGetExecBalance) GetExecAddr() []byte {
	if x != nil {
		return x.ExecAddr
	}
	return nil
}

func (x *ReqGetExecBalance) GetExecer() string {
	if x != nil {
		return x.Execer
	}
	return ""
}

func (x *ReqGetExecBalance) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ReqGetExecBalance) GetNextKey() []byte {
	if x != nil {
		return x.NextKey
	}
	return nil
}

type ExecBalanceItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExecAddr []byte `protobuf:"bytes,1,opt,name=execAddr,proto3" json:"execAddr,omitempty"`
	Frozen   int64  `protobuf:"varint,2,opt,name=frozen,proto3" json:"frozen,omitempty"`
	Active   int64  `protobuf:"varint,3,opt,name=active,proto3" json:"active,omitempty"`
}

func (x *ExecBalanceItem) Reset() {
	*x = ExecBalanceItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistic_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecBalanceItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecBalanceItem) ProtoMessage() {}

func (x *ExecBalanceItem) ProtoReflect() protoreflect.Message {
	mi := &file_statistic_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecBalanceItem.ProtoReflect.Descriptor instead.
func (*ExecBalanceItem) Descriptor() ([]byte, []int) {
	return file_statistic_proto_rawDescGZIP(), []int{6}
}

func (x *ExecBalanceItem) GetExecAddr() []byte {
	if x != nil {
		return x.ExecAddr
	}
	return nil
}

func (x *ExecBalanceItem) GetFrozen() int64 {
	if x != nil {
		return x.Frozen
	}
	return 0
}

func (x *ExecBalanceItem) GetActive() int64 {
	if x != nil {
		return x.Active
	}
	return 0
}

//查询symbol在合约中的代币总额应答
type ReplyGetExecBalance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount       int64              `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	AmountFrozen int64              `protobuf:"varint,2,opt,name=amountFrozen,proto3" json:"amountFrozen,omitempty"`
	AmountActive int64              `protobuf:"varint,3,opt,name=amountActive,proto3" json:"amountActive,omitempty"`
	NextKey      []byte             `protobuf:"bytes,4,opt,name=nextKey,proto3" json:"nextKey,omitempty"`
	Items        []*ExecBalanceItem `protobuf:"bytes,5,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ReplyGetExecBalance) Reset() {
	*x = ReplyGetExecBalance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statistic_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyGetExecBalance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyGetExecBalance) ProtoMessage() {}

func (x *ReplyGetExecBalance) ProtoReflect() protoreflect.Message {
	mi := &file_statistic_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyGetExecBalance.ProtoReflect.Descriptor instead.
func (*ReplyGetExecBalance) Descriptor() ([]byte, []int) {
	return file_statistic_proto_rawDescGZIP(), []int{7}
}

func (x *ReplyGetExecBalance) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *ReplyGetExecBalance) GetAmountFrozen() int64 {
	if x != nil {
		return x.AmountFrozen
	}
	return 0
}

func (x *ReplyGetExecBalance) GetAmountActive() int64 {
	if x != nil {
		return x.AmountActive
	}
	return 0
}

func (x *ReplyGetExecBalance) GetNextKey() []byte {
	if x != nil {
		return x.NextKey
	}
	return nil
}

func (x *ReplyGetExecBalance) GetItems() []*ExecBalanceItem {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_statistic_proto protoreflect.FileDescriptor

var file_statistic_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0x36, 0x0a, 0x08, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x46, 0x65, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x66, 0x65, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x78, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x78, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x92, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x71, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x43, 0x6f, 0x69, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x78, 0x65, 0x63, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65,
	0x78, 0x65, 0x63, 0x65, 0x72, 0x22, 0x6e, 0x0a, 0x12, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x47, 0x65,
	0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x6e, 0x75, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6e,
	0x65, 0x78, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6e, 0x65,
	0x78, 0x74, 0x4b, 0x65, 0x79, 0x22, 0x75, 0x0a, 0x17, 0x49, 0x74, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x42, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x23, 0x0a, 0x0b,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x22, 0xc1, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x47, 0x65, 0x74, 0x45, 0x78, 0x65, 0x63,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a,
	0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x61, 0x64, 0x64,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x65, 0x63, 0x41, 0x64, 0x64, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x08, 0x65, 0x78, 0x65, 0x63, 0x41, 0x64, 0x64, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x78, 0x65, 0x63, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65,
	0x78, 0x65, 0x63, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6e,
	0x65, 0x78, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6e, 0x65,
	0x78, 0x74, 0x4b, 0x65, 0x79, 0x22, 0x5d, 0x0a, 0x0f, 0x45, 0x78, 0x65, 0x63, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x65, 0x63,
	0x41, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x65, 0x78, 0x65, 0x63,
	0x41, 0x64, 0x64, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x66, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x22, 0xbd, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x47, 0x65,
	0x74, 0x45, 0x78, 0x65, 0x63, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x72,
	0x6f, 0x7a, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x46, 0x72, 0x6f, 0x7a, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6e, 0x65, 0x78, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6e,
	0x65, 0x78, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x78,
	0x65, 0x63, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x33, 0x33, 0x63, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x33, 0x33, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_statistic_proto_rawDescOnce sync.Once
	file_statistic_proto_rawDescData = file_statistic_proto_rawDesc
)

func file_statistic_proto_rawDescGZIP() []byte {
	file_statistic_proto_rawDescOnce.Do(func() {
		file_statistic_proto_rawDescData = protoimpl.X.CompressGZIP(file_statistic_proto_rawDescData)
	})
	return file_statistic_proto_rawDescData
}

var file_statistic_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_statistic_proto_goTypes = []interface{}{
	(*TotalFee)(nil),                // 0: types.TotalFee
	(*ReqGetTotalCoins)(nil),        // 1: types.ReqGetTotalCoins
	(*ReplyGetTotalCoins)(nil),      // 2: types.ReplyGetTotalCoins
	(*IterateRangeByStateHash)(nil), // 3: types.IterateRangeByStateHash
	(*TotalAmount)(nil),             // 4: types.TotalAmount
	(*ReqGetExecBalance)(nil),       // 5: types.ReqGetExecBalance
	(*ExecBalanceItem)(nil),         // 6: types.ExecBalanceItem
	(*ReplyGetExecBalance)(nil),     // 7: types.ReplyGetExecBalance
}
var file_statistic_proto_depIdxs = []int32{
	6, // 0: types.ReplyGetExecBalance.items:type_name -> types.ExecBalanceItem
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_statistic_proto_init() }
func file_statistic_proto_init() {
	if File_statistic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_statistic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TotalFee); i {
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
		file_statistic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqGetTotalCoins); i {
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
		file_statistic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyGetTotalCoins); i {
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
		file_statistic_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IterateRangeByStateHash); i {
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
		file_statistic_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TotalAmount); i {
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
		file_statistic_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqGetExecBalance); i {
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
		file_statistic_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecBalanceItem); i {
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
		file_statistic_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyGetExecBalance); i {
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
			RawDescriptor: file_statistic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_statistic_proto_goTypes,
		DependencyIndexes: file_statistic_proto_depIdxs,
		MessageInfos:      file_statistic_proto_msgTypes,
	}.Build()
	File_statistic_proto = out.File
	file_statistic_proto_rawDesc = nil
	file_statistic_proto_goTypes = nil
	file_statistic_proto_depIdxs = nil
}
