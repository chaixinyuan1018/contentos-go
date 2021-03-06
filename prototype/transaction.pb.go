// Code generated by protoc-gen-go. DO NOT EDIT.
// source: prototype/transaction.proto

package prototype

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Operation struct {
	// Types that are valid to be assigned to Op:
	//	*Operation_Op1
	//	*Operation_Op2
	//	*Operation_Op3
	//	*Operation_Op4
	//	*Operation_Op5
	//	*Operation_Op6
	//	*Operation_Op7
	//	*Operation_Op8
	//	*Operation_Op9
	//	*Operation_Op10
	//	*Operation_Op13
	//	*Operation_Op14
	//	*Operation_Op16
	//	*Operation_Op17
	//	*Operation_Op18
	//	*Operation_Op19
	//	*Operation_Op20
	//	*Operation_Op21
	//	*Operation_Op22
	//	*Operation_Op23
	//	*Operation_Op24
	Op                   isOperation_Op `protobuf_oneof:"op"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Operation) Reset()         { *m = Operation{} }
func (m *Operation) String() string { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()    {}
func (*Operation) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3aa2bc02ae1e20c, []int{0}
}

func (m *Operation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Operation.Unmarshal(m, b)
}
func (m *Operation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Operation.Marshal(b, m, deterministic)
}
func (m *Operation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Operation.Merge(m, src)
}
func (m *Operation) XXX_Size() int {
	return xxx_messageInfo_Operation.Size(m)
}
func (m *Operation) XXX_DiscardUnknown() {
	xxx_messageInfo_Operation.DiscardUnknown(m)
}

var xxx_messageInfo_Operation proto.InternalMessageInfo

type isOperation_Op interface {
	isOperation_Op()
}

type Operation_Op1 struct {
	Op1 *AccountCreateOperation `protobuf:"bytes,1,opt,name=op1,proto3,oneof"`
}

type Operation_Op2 struct {
	Op2 *TransferOperation `protobuf:"bytes,2,opt,name=op2,proto3,oneof"`
}

type Operation_Op3 struct {
	Op3 *BpRegisterOperation `protobuf:"bytes,3,opt,name=op3,proto3,oneof"`
}

type Operation_Op4 struct {
	Op4 *BpEnableOperation `protobuf:"bytes,4,opt,name=op4,proto3,oneof"`
}

type Operation_Op5 struct {
	Op5 *BpVoteOperation `protobuf:"bytes,5,opt,name=op5,proto3,oneof"`
}

type Operation_Op6 struct {
	Op6 *PostOperation `protobuf:"bytes,6,opt,name=op6,proto3,oneof"`
}

type Operation_Op7 struct {
	Op7 *ReplyOperation `protobuf:"bytes,7,opt,name=op7,proto3,oneof"`
}

type Operation_Op8 struct {
	Op8 *FollowOperation `protobuf:"bytes,8,opt,name=op8,proto3,oneof"`
}

type Operation_Op9 struct {
	Op9 *VoteOperation `protobuf:"bytes,9,opt,name=op9,proto3,oneof"`
}

type Operation_Op10 struct {
	Op10 *TransferToVestOperation `protobuf:"bytes,10,opt,name=op10,proto3,oneof"`
}

type Operation_Op13 struct {
	Op13 *ContractDeployOperation `protobuf:"bytes,13,opt,name=op13,proto3,oneof"`
}

type Operation_Op14 struct {
	Op14 *ContractApplyOperation `protobuf:"bytes,14,opt,name=op14,proto3,oneof"`
}

type Operation_Op16 struct {
	Op16 *ConvertVestOperation `protobuf:"bytes,16,opt,name=op16,proto3,oneof"`
}

type Operation_Op17 struct {
	Op17 *StakeOperation `protobuf:"bytes,17,opt,name=op17,proto3,oneof"`
}

type Operation_Op18 struct {
	Op18 *UnStakeOperation `protobuf:"bytes,18,opt,name=op18,proto3,oneof"`
}

type Operation_Op19 struct {
	Op19 *BpUpdateOperation `protobuf:"bytes,19,opt,name=op19,proto3,oneof"`
}

type Operation_Op20 struct {
	Op20 *AccountUpdateOperation `protobuf:"bytes,20,opt,name=op20,proto3,oneof"`
}

type Operation_Op21 struct {
	Op21 *AcquireTicketOperation `protobuf:"bytes,21,opt,name=op21,proto3,oneof"`
}

type Operation_Op22 struct {
	Op22 *VoteByTicketOperation `protobuf:"bytes,22,opt,name=op22,proto3,oneof"`
}

type Operation_Op23 struct {
	Op23 *DelegateVestOperation `protobuf:"bytes,23,opt,name=op23,proto3,oneof"`
}

type Operation_Op24 struct {
	Op24 *UnDelegateVestOperation `protobuf:"bytes,24,opt,name=op24,proto3,oneof"`
}

func (*Operation_Op1) isOperation_Op() {}

func (*Operation_Op2) isOperation_Op() {}

func (*Operation_Op3) isOperation_Op() {}

func (*Operation_Op4) isOperation_Op() {}

func (*Operation_Op5) isOperation_Op() {}

func (*Operation_Op6) isOperation_Op() {}

func (*Operation_Op7) isOperation_Op() {}

func (*Operation_Op8) isOperation_Op() {}

func (*Operation_Op9) isOperation_Op() {}

func (*Operation_Op10) isOperation_Op() {}

func (*Operation_Op13) isOperation_Op() {}

func (*Operation_Op14) isOperation_Op() {}

func (*Operation_Op16) isOperation_Op() {}

func (*Operation_Op17) isOperation_Op() {}

func (*Operation_Op18) isOperation_Op() {}

func (*Operation_Op19) isOperation_Op() {}

func (*Operation_Op20) isOperation_Op() {}

func (*Operation_Op21) isOperation_Op() {}

func (*Operation_Op22) isOperation_Op() {}

func (*Operation_Op23) isOperation_Op() {}

func (*Operation_Op24) isOperation_Op() {}

func (m *Operation) GetOp() isOperation_Op {
	if m != nil {
		return m.Op
	}
	return nil
}

func (m *Operation) GetOp1() *AccountCreateOperation {
	if x, ok := m.GetOp().(*Operation_Op1); ok {
		return x.Op1
	}
	return nil
}

func (m *Operation) GetOp2() *TransferOperation {
	if x, ok := m.GetOp().(*Operation_Op2); ok {
		return x.Op2
	}
	return nil
}

func (m *Operation) GetOp3() *BpRegisterOperation {
	if x, ok := m.GetOp().(*Operation_Op3); ok {
		return x.Op3
	}
	return nil
}

func (m *Operation) GetOp4() *BpEnableOperation {
	if x, ok := m.GetOp().(*Operation_Op4); ok {
		return x.Op4
	}
	return nil
}

func (m *Operation) GetOp5() *BpVoteOperation {
	if x, ok := m.GetOp().(*Operation_Op5); ok {
		return x.Op5
	}
	return nil
}

func (m *Operation) GetOp6() *PostOperation {
	if x, ok := m.GetOp().(*Operation_Op6); ok {
		return x.Op6
	}
	return nil
}

func (m *Operation) GetOp7() *ReplyOperation {
	if x, ok := m.GetOp().(*Operation_Op7); ok {
		return x.Op7
	}
	return nil
}

func (m *Operation) GetOp8() *FollowOperation {
	if x, ok := m.GetOp().(*Operation_Op8); ok {
		return x.Op8
	}
	return nil
}

func (m *Operation) GetOp9() *VoteOperation {
	if x, ok := m.GetOp().(*Operation_Op9); ok {
		return x.Op9
	}
	return nil
}

func (m *Operation) GetOp10() *TransferToVestOperation {
	if x, ok := m.GetOp().(*Operation_Op10); ok {
		return x.Op10
	}
	return nil
}

func (m *Operation) GetOp13() *ContractDeployOperation {
	if x, ok := m.GetOp().(*Operation_Op13); ok {
		return x.Op13
	}
	return nil
}

func (m *Operation) GetOp14() *ContractApplyOperation {
	if x, ok := m.GetOp().(*Operation_Op14); ok {
		return x.Op14
	}
	return nil
}

func (m *Operation) GetOp16() *ConvertVestOperation {
	if x, ok := m.GetOp().(*Operation_Op16); ok {
		return x.Op16
	}
	return nil
}

func (m *Operation) GetOp17() *StakeOperation {
	if x, ok := m.GetOp().(*Operation_Op17); ok {
		return x.Op17
	}
	return nil
}

func (m *Operation) GetOp18() *UnStakeOperation {
	if x, ok := m.GetOp().(*Operation_Op18); ok {
		return x.Op18
	}
	return nil
}

func (m *Operation) GetOp19() *BpUpdateOperation {
	if x, ok := m.GetOp().(*Operation_Op19); ok {
		return x.Op19
	}
	return nil
}

func (m *Operation) GetOp20() *AccountUpdateOperation {
	if x, ok := m.GetOp().(*Operation_Op20); ok {
		return x.Op20
	}
	return nil
}

func (m *Operation) GetOp21() *AcquireTicketOperation {
	if x, ok := m.GetOp().(*Operation_Op21); ok {
		return x.Op21
	}
	return nil
}

func (m *Operation) GetOp22() *VoteByTicketOperation {
	if x, ok := m.GetOp().(*Operation_Op22); ok {
		return x.Op22
	}
	return nil
}

func (m *Operation) GetOp23() *DelegateVestOperation {
	if x, ok := m.GetOp().(*Operation_Op23); ok {
		return x.Op23
	}
	return nil
}

func (m *Operation) GetOp24() *UnDelegateVestOperation {
	if x, ok := m.GetOp().(*Operation_Op24); ok {
		return x.Op24
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Operation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Operation_Op1)(nil),
		(*Operation_Op2)(nil),
		(*Operation_Op3)(nil),
		(*Operation_Op4)(nil),
		(*Operation_Op5)(nil),
		(*Operation_Op6)(nil),
		(*Operation_Op7)(nil),
		(*Operation_Op8)(nil),
		(*Operation_Op9)(nil),
		(*Operation_Op10)(nil),
		(*Operation_Op13)(nil),
		(*Operation_Op14)(nil),
		(*Operation_Op16)(nil),
		(*Operation_Op17)(nil),
		(*Operation_Op18)(nil),
		(*Operation_Op19)(nil),
		(*Operation_Op20)(nil),
		(*Operation_Op21)(nil),
		(*Operation_Op22)(nil),
		(*Operation_Op23)(nil),
		(*Operation_Op24)(nil),
	}
}

// transaction
type Transaction struct {
	RefBlockNum          uint32        `protobuf:"varint,1,opt,name=ref_block_num,json=refBlockNum,proto3" json:"ref_block_num,omitempty"`
	RefBlockPrefix       uint32        `protobuf:"varint,2,opt,name=ref_block_prefix,json=refBlockPrefix,proto3" json:"ref_block_prefix,omitempty"`
	Expiration           *TimePointSec `protobuf:"bytes,3,opt,name=expiration,proto3" json:"expiration,omitempty"`
	Operations           []*Operation  `protobuf:"bytes,4,rep,name=operations,proto3" json:"operations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3aa2bc02ae1e20c, []int{1}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetRefBlockNum() uint32 {
	if m != nil {
		return m.RefBlockNum
	}
	return 0
}

func (m *Transaction) GetRefBlockPrefix() uint32 {
	if m != nil {
		return m.RefBlockPrefix
	}
	return 0
}

func (m *Transaction) GetExpiration() *TimePointSec {
	if m != nil {
		return m.Expiration
	}
	return nil
}

func (m *Transaction) GetOperations() []*Operation {
	if m != nil {
		return m.Operations
	}
	return nil
}

type SignedTransaction struct {
	Trx                  *Transaction   `protobuf:"bytes,1,opt,name=trx,proto3" json:"trx,omitempty"`
	Signature            *SignatureType `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SignedTransaction) Reset()         { *m = SignedTransaction{} }
func (m *SignedTransaction) String() string { return proto.CompactTextString(m) }
func (*SignedTransaction) ProtoMessage()    {}
func (*SignedTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3aa2bc02ae1e20c, []int{2}
}

func (m *SignedTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedTransaction.Unmarshal(m, b)
}
func (m *SignedTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedTransaction.Marshal(b, m, deterministic)
}
func (m *SignedTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedTransaction.Merge(m, src)
}
func (m *SignedTransaction) XXX_Size() int {
	return xxx_messageInfo_SignedTransaction.Size(m)
}
func (m *SignedTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_SignedTransaction proto.InternalMessageInfo

func (m *SignedTransaction) GetTrx() *Transaction {
	if m != nil {
		return m.Trx
	}
	return nil
}

func (m *SignedTransaction) GetSignature() *SignatureType {
	if m != nil {
		return m.Signature
	}
	return nil
}

type OperationReceiptWithInfo struct {
	Status               uint32   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	GasUsage             uint64   `protobuf:"varint,2,opt,name=gas_usage,json=gasUsage,proto3" json:"gas_usage,omitempty"`
	VmConsole            string   `protobuf:"bytes,3,opt,name=vm_console,json=vmConsole,proto3" json:"vm_console,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperationReceiptWithInfo) Reset()         { *m = OperationReceiptWithInfo{} }
func (m *OperationReceiptWithInfo) String() string { return proto.CompactTextString(m) }
func (*OperationReceiptWithInfo) ProtoMessage()    {}
func (*OperationReceiptWithInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3aa2bc02ae1e20c, []int{3}
}

func (m *OperationReceiptWithInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperationReceiptWithInfo.Unmarshal(m, b)
}
func (m *OperationReceiptWithInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperationReceiptWithInfo.Marshal(b, m, deterministic)
}
func (m *OperationReceiptWithInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperationReceiptWithInfo.Merge(m, src)
}
func (m *OperationReceiptWithInfo) XXX_Size() int {
	return xxx_messageInfo_OperationReceiptWithInfo.Size(m)
}
func (m *OperationReceiptWithInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OperationReceiptWithInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OperationReceiptWithInfo proto.InternalMessageInfo

func (m *OperationReceiptWithInfo) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *OperationReceiptWithInfo) GetGasUsage() uint64 {
	if m != nil {
		return m.GasUsage
	}
	return 0
}

func (m *OperationReceiptWithInfo) GetVmConsole() string {
	if m != nil {
		return m.VmConsole
	}
	return ""
}

type TransactionReceiptWithInfo struct {
	Status               uint32                      `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	NetUsage             uint64                      `protobuf:"varint,2,opt,name=net_usage,json=netUsage,proto3" json:"net_usage,omitempty"`
	CpuUsage             uint64                      `protobuf:"varint,3,opt,name=cpu_usage,json=cpuUsage,proto3" json:"cpu_usage,omitempty"`
	ErrorInfo            string                      `protobuf:"bytes,4,opt,name=error_info,json=errorInfo,proto3" json:"error_info,omitempty"`
	OpResults            []*OperationReceiptWithInfo `protobuf:"bytes,5,rep,name=op_results,json=opResults,proto3" json:"op_results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *TransactionReceiptWithInfo) Reset()         { *m = TransactionReceiptWithInfo{} }
func (m *TransactionReceiptWithInfo) String() string { return proto.CompactTextString(m) }
func (*TransactionReceiptWithInfo) ProtoMessage()    {}
func (*TransactionReceiptWithInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3aa2bc02ae1e20c, []int{4}
}

func (m *TransactionReceiptWithInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionReceiptWithInfo.Unmarshal(m, b)
}
func (m *TransactionReceiptWithInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionReceiptWithInfo.Marshal(b, m, deterministic)
}
func (m *TransactionReceiptWithInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionReceiptWithInfo.Merge(m, src)
}
func (m *TransactionReceiptWithInfo) XXX_Size() int {
	return xxx_messageInfo_TransactionReceiptWithInfo.Size(m)
}
func (m *TransactionReceiptWithInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionReceiptWithInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionReceiptWithInfo proto.InternalMessageInfo

func (m *TransactionReceiptWithInfo) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *TransactionReceiptWithInfo) GetNetUsage() uint64 {
	if m != nil {
		return m.NetUsage
	}
	return 0
}

func (m *TransactionReceiptWithInfo) GetCpuUsage() uint64 {
	if m != nil {
		return m.CpuUsage
	}
	return 0
}

func (m *TransactionReceiptWithInfo) GetErrorInfo() string {
	if m != nil {
		return m.ErrorInfo
	}
	return ""
}

func (m *TransactionReceiptWithInfo) GetOpResults() []*OperationReceiptWithInfo {
	if m != nil {
		return m.OpResults
	}
	return nil
}

type TransactionWrapperWithInfo struct {
	SigTrx               *SignedTransaction          `protobuf:"bytes,1,opt,name=sig_trx,json=sigTrx,proto3" json:"sig_trx,omitempty"`
	Receipt              *TransactionReceiptWithInfo `protobuf:"bytes,2,opt,name=receipt,proto3" json:"receipt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *TransactionWrapperWithInfo) Reset()         { *m = TransactionWrapperWithInfo{} }
func (m *TransactionWrapperWithInfo) String() string { return proto.CompactTextString(m) }
func (*TransactionWrapperWithInfo) ProtoMessage()    {}
func (*TransactionWrapperWithInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3aa2bc02ae1e20c, []int{5}
}

func (m *TransactionWrapperWithInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionWrapperWithInfo.Unmarshal(m, b)
}
func (m *TransactionWrapperWithInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionWrapperWithInfo.Marshal(b, m, deterministic)
}
func (m *TransactionWrapperWithInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionWrapperWithInfo.Merge(m, src)
}
func (m *TransactionWrapperWithInfo) XXX_Size() int {
	return xxx_messageInfo_TransactionWrapperWithInfo.Size(m)
}
func (m *TransactionWrapperWithInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionWrapperWithInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionWrapperWithInfo proto.InternalMessageInfo

func (m *TransactionWrapperWithInfo) GetSigTrx() *SignedTransaction {
	if m != nil {
		return m.SigTrx
	}
	return nil
}

func (m *TransactionWrapperWithInfo) GetReceipt() *TransactionReceiptWithInfo {
	if m != nil {
		return m.Receipt
	}
	return nil
}

type TransactionReceipt struct {
	Status               uint32   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	NetUsage             uint64   `protobuf:"varint,2,opt,name=net_usage,json=netUsage,proto3" json:"net_usage,omitempty"`
	CpuUsage             uint64   `protobuf:"varint,3,opt,name=cpu_usage,json=cpuUsage,proto3" json:"cpu_usage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionReceipt) Reset()         { *m = TransactionReceipt{} }
func (m *TransactionReceipt) String() string { return proto.CompactTextString(m) }
func (*TransactionReceipt) ProtoMessage()    {}
func (*TransactionReceipt) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3aa2bc02ae1e20c, []int{6}
}

func (m *TransactionReceipt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionReceipt.Unmarshal(m, b)
}
func (m *TransactionReceipt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionReceipt.Marshal(b, m, deterministic)
}
func (m *TransactionReceipt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionReceipt.Merge(m, src)
}
func (m *TransactionReceipt) XXX_Size() int {
	return xxx_messageInfo_TransactionReceipt.Size(m)
}
func (m *TransactionReceipt) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionReceipt.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionReceipt proto.InternalMessageInfo

func (m *TransactionReceipt) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *TransactionReceipt) GetNetUsage() uint64 {
	if m != nil {
		return m.NetUsage
	}
	return 0
}

func (m *TransactionReceipt) GetCpuUsage() uint64 {
	if m != nil {
		return m.CpuUsage
	}
	return 0
}

type TransactionWrapper struct {
	SigTrx               *SignedTransaction  `protobuf:"bytes,1,opt,name=sig_trx,json=sigTrx,proto3" json:"sig_trx,omitempty"`
	Receipt              *TransactionReceipt `protobuf:"bytes,2,opt,name=receipt,proto3" json:"receipt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *TransactionWrapper) Reset()         { *m = TransactionWrapper{} }
func (m *TransactionWrapper) String() string { return proto.CompactTextString(m) }
func (*TransactionWrapper) ProtoMessage()    {}
func (*TransactionWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3aa2bc02ae1e20c, []int{7}
}

func (m *TransactionWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionWrapper.Unmarshal(m, b)
}
func (m *TransactionWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionWrapper.Marshal(b, m, deterministic)
}
func (m *TransactionWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionWrapper.Merge(m, src)
}
func (m *TransactionWrapper) XXX_Size() int {
	return xxx_messageInfo_TransactionWrapper.Size(m)
}
func (m *TransactionWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionWrapper proto.InternalMessageInfo

func (m *TransactionWrapper) GetSigTrx() *SignedTransaction {
	if m != nil {
		return m.SigTrx
	}
	return nil
}

func (m *TransactionWrapper) GetReceipt() *TransactionReceipt {
	if m != nil {
		return m.Receipt
	}
	return nil
}

// block
type BlockHeader struct {
	Previous              *Sha256       `protobuf:"bytes,1,opt,name=previous,proto3" json:"previous,omitempty"`
	Timestamp             *TimePointSec `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	BlockProducer         *AccountName  `protobuf:"bytes,3,opt,name=block_producer,json=blockProducer,proto3" json:"block_producer,omitempty"`
	TransactionMerkleRoot *Sha256       `protobuf:"bytes,4,opt,name=transaction_merkle_root,json=transactionMerkleRoot,proto3" json:"transaction_merkle_root,omitempty"`
	PrevApplyHash         uint64        `protobuf:"varint,5,opt,name=prev_apply_hash,json=prevApplyHash,proto3" json:"prev_apply_hash,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}      `json:"-"`
	XXX_unrecognized      []byte        `json:"-"`
	XXX_sizecache         int32         `json:"-"`
}

func (m *BlockHeader) Reset()         { *m = BlockHeader{} }
func (m *BlockHeader) String() string { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()    {}
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3aa2bc02ae1e20c, []int{8}
}

func (m *BlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHeader.Unmarshal(m, b)
}
func (m *BlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHeader.Marshal(b, m, deterministic)
}
func (m *BlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeader.Merge(m, src)
}
func (m *BlockHeader) XXX_Size() int {
	return xxx_messageInfo_BlockHeader.Size(m)
}
func (m *BlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeader proto.InternalMessageInfo

func (m *BlockHeader) GetPrevious() *Sha256 {
	if m != nil {
		return m.Previous
	}
	return nil
}

func (m *BlockHeader) GetTimestamp() *TimePointSec {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *BlockHeader) GetBlockProducer() *AccountName {
	if m != nil {
		return m.BlockProducer
	}
	return nil
}

func (m *BlockHeader) GetTransactionMerkleRoot() *Sha256 {
	if m != nil {
		return m.TransactionMerkleRoot
	}
	return nil
}

func (m *BlockHeader) GetPrevApplyHash() uint64 {
	if m != nil {
		return m.PrevApplyHash
	}
	return 0
}

type SignedBlockHeader struct {
	Header                 *BlockHeader   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	BlockProducerSignature *SignatureType `protobuf:"bytes,2,opt,name=block_producer_signature,json=blockProducerSignature,proto3" json:"block_producer_signature,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}       `json:"-"`
	XXX_unrecognized       []byte         `json:"-"`
	XXX_sizecache          int32          `json:"-"`
}

func (m *SignedBlockHeader) Reset()         { *m = SignedBlockHeader{} }
func (m *SignedBlockHeader) String() string { return proto.CompactTextString(m) }
func (*SignedBlockHeader) ProtoMessage()    {}
func (*SignedBlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3aa2bc02ae1e20c, []int{9}
}

func (m *SignedBlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedBlockHeader.Unmarshal(m, b)
}
func (m *SignedBlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedBlockHeader.Marshal(b, m, deterministic)
}
func (m *SignedBlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedBlockHeader.Merge(m, src)
}
func (m *SignedBlockHeader) XXX_Size() int {
	return xxx_messageInfo_SignedBlockHeader.Size(m)
}
func (m *SignedBlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedBlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_SignedBlockHeader proto.InternalMessageInfo

func (m *SignedBlockHeader) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SignedBlockHeader) GetBlockProducerSignature() *SignatureType {
	if m != nil {
		return m.BlockProducerSignature
	}
	return nil
}

type SignedBlock struct {
	SignedHeader         *SignedBlockHeader    `protobuf:"bytes,1,opt,name=signed_header,json=signedHeader,proto3" json:"signed_header,omitempty"`
	Transactions         []*TransactionWrapper `protobuf:"bytes,2,rep,name=transactions,proto3" json:"transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *SignedBlock) Reset()         { *m = SignedBlock{} }
func (m *SignedBlock) String() string { return proto.CompactTextString(m) }
func (*SignedBlock) ProtoMessage()    {}
func (*SignedBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3aa2bc02ae1e20c, []int{10}
}

func (m *SignedBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedBlock.Unmarshal(m, b)
}
func (m *SignedBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedBlock.Marshal(b, m, deterministic)
}
func (m *SignedBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedBlock.Merge(m, src)
}
func (m *SignedBlock) XXX_Size() int {
	return xxx_messageInfo_SignedBlock.Size(m)
}
func (m *SignedBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedBlock.DiscardUnknown(m)
}

var xxx_messageInfo_SignedBlock proto.InternalMessageInfo

func (m *SignedBlock) GetSignedHeader() *SignedBlockHeader {
	if m != nil {
		return m.SignedHeader
	}
	return nil
}

func (m *SignedBlock) GetTransactions() []*TransactionWrapper {
	if m != nil {
		return m.Transactions
	}
	return nil
}

type EmptySignedBlock struct {
	SignedHeader         *SignedBlockHeader `protobuf:"bytes,1,opt,name=signed_header,json=signedHeader,proto3" json:"signed_header,omitempty"`
	TrxCount             uint32             `protobuf:"varint,2,opt,name=trx_count,json=trxCount,proto3" json:"trx_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *EmptySignedBlock) Reset()         { *m = EmptySignedBlock{} }
func (m *EmptySignedBlock) String() string { return proto.CompactTextString(m) }
func (*EmptySignedBlock) ProtoMessage()    {}
func (*EmptySignedBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3aa2bc02ae1e20c, []int{11}
}

func (m *EmptySignedBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptySignedBlock.Unmarshal(m, b)
}
func (m *EmptySignedBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptySignedBlock.Marshal(b, m, deterministic)
}
func (m *EmptySignedBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptySignedBlock.Merge(m, src)
}
func (m *EmptySignedBlock) XXX_Size() int {
	return xxx_messageInfo_EmptySignedBlock.Size(m)
}
func (m *EmptySignedBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptySignedBlock.DiscardUnknown(m)
}

var xxx_messageInfo_EmptySignedBlock proto.InternalMessageInfo

func (m *EmptySignedBlock) GetSignedHeader() *SignedBlockHeader {
	if m != nil {
		return m.SignedHeader
	}
	return nil
}

func (m *EmptySignedBlock) GetTrxCount() uint32 {
	if m != nil {
		return m.TrxCount
	}
	return 0
}

func init() {
	proto.RegisterType((*Operation)(nil), "prototype.operation")
	proto.RegisterType((*Transaction)(nil), "prototype.transaction")
	proto.RegisterType((*SignedTransaction)(nil), "prototype.signed_transaction")
	proto.RegisterType((*OperationReceiptWithInfo)(nil), "prototype.operation_receipt_with_info")
	proto.RegisterType((*TransactionReceiptWithInfo)(nil), "prototype.transaction_receipt_with_info")
	proto.RegisterType((*TransactionWrapperWithInfo)(nil), "prototype.transaction_wrapper_with_info")
	proto.RegisterType((*TransactionReceipt)(nil), "prototype.transaction_receipt")
	proto.RegisterType((*TransactionWrapper)(nil), "prototype.transaction_wrapper")
	proto.RegisterType((*BlockHeader)(nil), "prototype.block_header")
	proto.RegisterType((*SignedBlockHeader)(nil), "prototype.signed_block_header")
	proto.RegisterType((*SignedBlock)(nil), "prototype.signed_block")
	proto.RegisterType((*EmptySignedBlock)(nil), "prototype.empty_signed_block")
}

func init() { proto.RegisterFile("prototype/transaction.proto", fileDescriptor_f3aa2bc02ae1e20c) }

var fileDescriptor_f3aa2bc02ae1e20c = []byte{
	// 1113 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0xdf, 0x6e, 0xdb, 0xb6,
	0x17, 0xc7, 0x7f, 0x4e, 0xdc, 0x34, 0x62, 0xea, 0xfe, 0x5a, 0x36, 0x4d, 0xd8, 0x78, 0x09, 0x32,
	0x6d, 0x2b, 0x82, 0x01, 0xb1, 0x13, 0xd9, 0xb1, 0x9d, 0x0d, 0x18, 0xb0, 0x04, 0x03, 0xd2, 0x8b,
	0x0d, 0x85, 0xba, 0xdd, 0xec, 0x86, 0xa0, 0x65, 0xda, 0x16, 0x62, 0x8b, 0x2c, 0x49, 0x39, 0xf6,
	0x13, 0x0c, 0xd8, 0x0b, 0xec, 0x62, 0x4f, 0xb0, 0x77, 0xd9, 0xfd, 0x5e, 0x67, 0x20, 0xa5, 0xc8,
	0xd4, 0x9f, 0x14, 0xe8, 0xd0, 0x3b, 0x8b, 0xfc, 0x7e, 0xcc, 0xef, 0x39, 0x3c, 0xe7, 0x48, 0xa0,
	0xc9, 0x05, 0x53, 0x4c, 0xad, 0x38, 0x6d, 0x2b, 0x41, 0x22, 0x49, 0x02, 0x15, 0xb2, 0xa8, 0x65,
	0x56, 0xa1, 0x93, 0x6d, 0x1e, 0xec, 0x5a, 0xba, 0x15, 0xa7, 0x89, 0xe0, 0xe0, 0xd5, 0x7a, 0x95,
	0x71, 0x2a, 0xc8, 0x9a, 0x75, 0x7f, 0x77, 0x80, 0x93, 0xad, 0xc1, 0x3e, 0xd8, 0x64, 0xfc, 0x1c,
	0xd5, 0x8e, 0x6b, 0x27, 0x3b, 0xde, 0x17, 0xad, 0x0c, 0x6b, 0x91, 0x20, 0x60, 0x71, 0xa4, 0x70,
	0x20, 0x28, 0x51, 0x14, 0x67, 0xc4, 0xcd, 0xff, 0x7c, 0x4d, 0xc0, 0x73, 0x0d, 0x7a, 0x68, 0xc3,
	0x80, 0x87, 0x16, 0x68, 0xdc, 0x8e, 0xa9, 0x28, 0x22, 0x1e, 0xec, 0x6a, 0xa4, 0x83, 0x36, 0x0d,
	0x72, 0x6c, 0x21, 0x43, 0x8e, 0x05, 0x9d, 0x84, 0x52, 0x95, 0xa9, 0x0e, 0xf4, 0x34, 0xd5, 0x45,
	0x75, 0x43, 0x1d, 0xe5, 0x29, 0x1a, 0x91, 0xe1, 0xac, 0x64, 0xae, 0x0b, 0xcf, 0x34, 0x73, 0x81,
	0x1e, 0x19, 0xe6, 0xb3, 0x3c, 0xb3, 0x60, 0xe5, 0x70, 0x2e, 0xe0, 0xa9, 0x26, 0x7a, 0x68, 0xcb,
	0x10, 0xaf, 0x2c, 0x82, 0x33, 0xa9, 0x8a, 0xf2, 0x1e, 0x6c, 0x69, 0x79, 0x1f, 0x3d, 0x36, 0xf2,
	0x03, 0x4b, 0x2e, 0x28, 0x9f, 0xad, 0x8a, 0xfa, 0x3e, 0x6c, 0x6b, 0xfd, 0x00, 0x6d, 0x1b, 0x7d,
	0xd3, 0xd2, 0x8f, 0xd9, 0x6c, 0xc6, 0xee, 0x8a, 0xc0, 0x20, 0xf1, 0x73, 0x89, 0x9c, 0x92, 0x9f,
	0x2a, 0xfb, 0x97, 0xf0, 0x5b, 0x50, 0x67, 0xfc, 0xfc, 0x0c, 0x01, 0xa3, 0xff, 0xaa, 0xea, 0x3a,
	0x14, 0xc3, 0x0b, 0x5a, 0x88, 0xc5, 0x40, 0xf0, 0x1b, 0x03, 0x77, 0x50, 0xc3, 0xc0, 0x5f, 0x5a,
	0x70, 0xc0, 0x22, 0x25, 0x48, 0xa0, 0xf0, 0x88, 0xf2, 0x19, 0x5b, 0x95, 0xd8, 0x0e, 0xbc, 0x34,
	0x6c, 0x17, 0x3d, 0x2d, 0x15, 0x50, 0xc6, 0x12, 0x5e, 0x4c, 0x89, 0x41, 0x60, 0xdf, 0xa0, 0x3d,
	0xf4, 0xcc, 0xa0, 0x9f, 0xe7, 0xd1, 0x05, 0x15, 0xaa, 0xda, 0x6f, 0x0f, 0x9e, 0x19, 0xb0, 0x8f,
	0x9e, 0x97, 0xb2, 0x2f, 0x15, 0xb9, 0xa5, 0x25, 0xa2, 0x0f, 0x3b, 0x86, 0x18, 0x20, 0x58, 0xaa,
	0xd6, 0x38, 0xc2, 0xd5, 0xd0, 0x00, 0x76, 0x0d, 0x74, 0x89, 0x5e, 0x54, 0x55, 0x5e, 0xcc, 0x47,
	0xc5, 0xb6, 0x30, 0xea, 0x24, 0x21, 0xde, 0x19, 0xda, 0x7d, 0xb0, 0xa3, 0xaa, 0x51, 0xef, 0x2c,
	0x45, 0xcf, 0xd1, 0xcb, 0x0a, 0xf4, 0x7d, 0x1c, 0x0a, 0x8a, 0x55, 0x18, 0xdc, 0xd2, 0x52, 0x4a,
	0xbc, 0xf3, 0x14, 0xf5, 0xd0, 0x5e, 0x09, 0x35, 0xf5, 0x32, 0x5c, 0x3d, 0x84, 0x7a, 0x70, 0x60,
	0xd0, 0x0e, 0xda, 0x37, 0xa8, 0x6b, 0xa1, 0x23, 0x3a, 0xa3, 0x13, 0x6d, 0xb5, 0xea, 0x1e, 0xbc,
	0x4e, 0x52, 0x74, 0x5e, 0x17, 0xa1, 0x52, 0xd1, 0xc5, 0x11, 0xfe, 0x30, 0xdc, 0xbd, 0xaa, 0x83,
	0x0d, 0xc6, 0xdd, 0xbf, 0x6b, 0x60, 0xc7, 0x1a, 0x6f, 0xd0, 0x05, 0x0d, 0x41, 0xc7, 0x78, 0x38,
	0x63, 0xc1, 0x2d, 0x8e, 0xe2, 0xb9, 0x19, 0x4c, 0x0d, 0x7f, 0x47, 0xd0, 0xf1, 0x95, 0x5e, 0xfb,
	0x29, 0x9e, 0xc3, 0x13, 0xf0, 0x6c, 0xad, 0xe1, 0x82, 0x8e, 0xc3, 0xa5, 0x19, 0x43, 0x0d, 0xff,
	0xe9, 0xbd, 0xec, 0xad, 0x59, 0x85, 0x97, 0x00, 0xd0, 0x25, 0x0f, 0x93, 0x93, 0xd3, 0xb9, 0x63,
	0xf7, 0x92, 0x0a, 0xe7, 0x14, 0x73, 0x16, 0x46, 0x0a, 0x4b, 0x1a, 0xf8, 0x96, 0x18, 0x76, 0x01,
	0xc8, 0x3c, 0x4b, 0x54, 0x3f, 0xde, 0x3c, 0xd9, 0xf1, 0x76, 0x2d, 0x34, 0xdb, 0xf4, 0x2d, 0x9d,
	0x7b, 0x07, 0xa0, 0x0c, 0x27, 0x11, 0x1d, 0x61, 0x3b, 0xa8, 0x13, 0xb0, 0xa9, 0xc4, 0x32, 0x9d,
	0xb1, 0x7b, 0xc5, 0xde, 0x4c, 0x44, 0xbe, 0x96, 0xc0, 0x3e, 0x70, 0x34, 0x4f, 0x54, 0x2c, 0x68,
	0x3a, 0x5a, 0x6d, 0xbf, 0xd9, 0x1e, 0xd6, 0x8f, 0xfe, 0x5a, 0xeb, 0xbe, 0x07, 0xcd, 0xcc, 0x06,
	0x16, 0x34, 0xa0, 0x21, 0x57, 0xf8, 0x2e, 0x54, 0x53, 0x1c, 0x46, 0x63, 0x06, 0xf7, 0xc0, 0x96,
	0x54, 0x44, 0xc5, 0x32, 0xcd, 0x67, 0xfa, 0x04, 0x9b, 0xc0, 0x99, 0x10, 0x89, 0x63, 0x49, 0x26,
	0xc9, 0x79, 0x75, 0x7f, 0x7b, 0x42, 0xe4, 0x2f, 0xfa, 0x19, 0x1e, 0x02, 0xb0, 0x98, 0xe3, 0x80,
	0x45, 0x92, 0xcd, 0xa8, 0xc9, 0x9e, 0xe3, 0x3b, 0x8b, 0xf9, 0x75, 0xb2, 0xe0, 0xfe, 0x53, 0x03,
	0x87, 0x56, 0x00, 0x1f, 0x77, 0x6a, 0x44, 0x55, 0xfe, 0xd4, 0x88, 0xaa, 0xe4, 0xd4, 0x26, 0x70,
	0x02, 0x1e, 0xa7, 0x9b, 0x9b, 0xc9, 0x66, 0xc0, 0xe3, 0xcc, 0x12, 0x15, 0x82, 0x09, 0xf3, 0xff,
	0xe6, 0x95, 0xe0, 0xf8, 0x8e, 0x59, 0x79, 0xa3, 0x0f, 0xfc, 0x41, 0x5f, 0x1a, 0x16, 0x54, 0xc6,
	0x33, 0x25, 0xd1, 0x23, 0x73, 0x69, 0xaf, 0xab, 0x2e, 0xad, 0x6c, 0xd6, 0x77, 0x18, 0xf7, 0x13,
	0xd0, 0xfd, 0xb3, 0x10, 0xd9, 0x9d, 0x20, 0x9c, 0x53, 0x61, 0x45, 0xd6, 0x03, 0x8f, 0x65, 0x38,
	0xc1, 0xeb, 0x5b, 0x3d, 0x2c, 0xdc, 0x52, 0xbe, 0x02, 0xfc, 0x2d, 0x19, 0x4e, 0x7e, 0x16, 0x4b,
	0x78, 0x05, 0x1e, 0xa7, 0x27, 0xa7, 0xb7, 0x7b, 0x52, 0x5d, 0x0d, 0x15, 0xfe, 0xee, 0x41, 0x77,
	0x02, 0x5e, 0x54, 0x28, 0x3f, 0x7d, 0xb2, 0xdd, 0xdf, 0x6a, 0xf9, 0x93, 0xd2, 0x34, 0xfc, 0xe7,
	0xe0, 0x07, 0xc5, 0xe0, 0x8f, 0x3e, 0x1c, 0xfc, 0x3a, 0xe4, 0xbf, 0x36, 0xc0, 0x93, 0xa4, 0xdd,
	0xa7, 0x94, 0x8c, 0xa8, 0x80, 0xa7, 0x60, 0x9b, 0x0b, 0xba, 0x08, 0x59, 0x1a, 0xee, 0x8e, 0xf7,
	0xdc, 0xf6, 0x30, 0x25, 0xde, 0x45, 0xcf, 0xcf, 0x24, 0xba, 0xad, 0x74, 0xab, 0x4b, 0x45, 0xe6,
	0xbc, 0xa2, 0xad, 0x0a, 0x63, 0x60, 0xad, 0x85, 0xdf, 0x81, 0xa7, 0xf7, 0x63, 0x86, 0x8d, 0xe2,
	0x80, 0x8a, 0x74, 0x88, 0xec, 0x57, 0x8c, 0xf5, 0x88, 0xcc, 0xa9, 0xdf, 0x18, 0x26, 0xe3, 0x27,
	0x51, 0xc3, 0x37, 0x60, 0xdf, 0x0e, 0x6c, 0x4e, 0xc5, 0xed, 0x8c, 0x62, 0xc1, 0x98, 0x4a, 0xbf,
	0x67, 0x2a, 0x6c, 0xbf, 0xb4, 0x88, 0x1f, 0x0d, 0xe0, 0x33, 0xa6, 0xe0, 0x6b, 0xf0, 0x7f, 0x1d,
	0x4f, 0xfa, 0x36, 0x9d, 0x12, 0x39, 0x35, 0x9f, 0x37, 0x75, 0xbf, 0xa1, 0x97, 0xbf, 0xd7, 0xab,
	0x37, 0x44, 0x4e, 0x75, 0xf1, 0xbe, 0x48, 0x2f, 0x21, 0x97, 0xb2, 0x36, 0xd8, 0x4a, 0x7e, 0xa5,
	0x09, 0xb3, 0x43, 0xb0, 0x85, 0x7e, 0x2a, 0x83, 0xef, 0x00, 0xca, 0xc7, 0x8e, 0x3f, 0x62, 0x34,
	0xed, 0xe5, 0xf2, 0xf0, 0x2e, 0x9b, 0x53, 0x7f, 0xd4, 0xc0, 0x13, 0xdb, 0x1d, 0xbc, 0x06, 0x8d,
	0xf4, 0x39, 0xe7, 0xee, 0xa8, 0x5c, 0x52, 0x39, 0x93, 0xe9, 0x9f, 0xdc, 0x24, 0x56, 0xaf, 0xc0,
	0x13, 0x2b, 0x69, 0x12, 0x6d, 0x98, 0xce, 0x7f, 0xa8, 0xbc, 0xd2, 0x3a, 0xf6, 0x73, 0x8c, 0xbb,
	0x00, 0x90, 0xce, 0xb9, 0x5a, 0xe1, 0x4f, 0x6f, 0xaf, 0x09, 0x1c, 0x25, 0x96, 0xd8, 0x94, 0x49,
	0xfa, 0xa6, 0xda, 0x56, 0x62, 0x79, 0xad, 0x9f, 0xaf, 0xde, 0x02, 0x37, 0x64, 0xe6, 0x63, 0x89,
	0x46, 0x8a, 0xc9, 0x16, 0x89, 0x46, 0x82, 0x85, 0xa3, 0x96, 0x1c, 0xdd, 0xae, 0x0f, 0xf9, 0xf5,
	0xeb, 0x49, 0xa8, 0xa6, 0xf1, 0xb0, 0x15, 0xb0, 0x79, 0x3b, 0x60, 0x32, 0x98, 0x92, 0x30, 0x6a,
	0x67, 0xd0, 0xe9, 0x84, 0xb5, 0x33, 0xed, 0x70, 0xcb, 0xfc, 0xec, 0xfc, 0x1b, 0x00, 0x00, 0xff,
	0xff, 0xe6, 0xd8, 0xac, 0x23, 0x42, 0x0c, 0x00, 0x00,
}
