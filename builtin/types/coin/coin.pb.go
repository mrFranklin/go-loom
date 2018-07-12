// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/loomnetwork/go-loom/builtin/types/coin/coin.proto

/*
Package coin is a generated protocol buffer package.

It is generated from these files:
	github.com/loomnetwork/go-loom/builtin/types/coin/coin.proto

It has these top-level messages:
	Economy
	Account
	Allowance
	InitialAccount
	InitRequest
	MintRequest
	TotalSupplyRequest
	TotalSupplyResponse
	BalanceOfRequest
	BalanceOfResponse
	AllowanceRequest
	AllowanceResponse
	ApproveRequest
	ApproveResponse
	TransferRequest
	TransferResponse
	TransferFromRequest
	TransferFromResponse
*/
package coin

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/loomnetwork/go-loom/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Economy struct {
	TotalSupply *types.BigUInt `protobuf:"bytes,1,opt,name=total_supply,json=totalSupply" json:"total_supply,omitempty"`
}

func (m *Economy) Reset()                    { *m = Economy{} }
func (m *Economy) String() string            { return proto.CompactTextString(m) }
func (*Economy) ProtoMessage()               {}
func (*Economy) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{0} }

func (m *Economy) GetTotalSupply() *types.BigUInt {
	if m != nil {
		return m.TotalSupply
	}
	return nil
}

type Account struct {
	Owner   *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Balance *types.BigUInt `protobuf:"bytes,2,opt,name=balance" json:"balance,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{1} }

func (m *Account) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Account) GetBalance() *types.BigUInt {
	if m != nil {
		return m.Balance
	}
	return nil
}

type Allowance struct {
	Owner   *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Spender *types.Address `protobuf:"bytes,2,opt,name=spender" json:"spender,omitempty"`
	Amount  *types.BigUInt `protobuf:"bytes,3,opt,name=amount" json:"amount,omitempty"`
}

func (m *Allowance) Reset()                    { *m = Allowance{} }
func (m *Allowance) String() string            { return proto.CompactTextString(m) }
func (*Allowance) ProtoMessage()               {}
func (*Allowance) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{2} }

func (m *Allowance) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Allowance) GetSpender() *types.Address {
	if m != nil {
		return m.Spender
	}
	return nil
}

func (m *Allowance) GetAmount() *types.BigUInt {
	if m != nil {
		return m.Amount
	}
	return nil
}

type InitialAccount struct {
	Owner   *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Balance uint64         `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (m *InitialAccount) Reset()                    { *m = InitialAccount{} }
func (m *InitialAccount) String() string            { return proto.CompactTextString(m) }
func (*InitialAccount) ProtoMessage()               {}
func (*InitialAccount) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{3} }

func (m *InitialAccount) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *InitialAccount) GetBalance() uint64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

type InitRequest struct {
	Accounts []*InitialAccount `protobuf:"bytes,1,rep,name=accounts" json:"accounts,omitempty"`
}

func (m *InitRequest) Reset()                    { *m = InitRequest{} }
func (m *InitRequest) String() string            { return proto.CompactTextString(m) }
func (*InitRequest) ProtoMessage()               {}
func (*InitRequest) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{4} }

func (m *InitRequest) GetAccounts() []*InitialAccount {
	if m != nil {
		return m.Accounts
	}
	return nil
}

type MintRequest struct {
	To     *types.Address `protobuf:"bytes,1,opt,name=to" json:"to,omitempty"`
	Amount *types.BigUInt `protobuf:"bytes,2,opt,name=amount" json:"amount,omitempty"`
}

func (m *MintRequest) Reset()                    { *m = MintRequest{} }
func (m *MintRequest) String() string            { return proto.CompactTextString(m) }
func (*MintRequest) ProtoMessage()               {}
func (*MintRequest) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{5} }

func (m *MintRequest) GetTo() *types.Address {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *MintRequest) GetAmount() *types.BigUInt {
	if m != nil {
		return m.Amount
	}
	return nil
}

// read only
type TotalSupplyRequest struct {
}

func (m *TotalSupplyRequest) Reset()                    { *m = TotalSupplyRequest{} }
func (m *TotalSupplyRequest) String() string            { return proto.CompactTextString(m) }
func (*TotalSupplyRequest) ProtoMessage()               {}
func (*TotalSupplyRequest) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{6} }

type TotalSupplyResponse struct {
	TotalSupply *types.BigUInt `protobuf:"bytes,1,opt,name=total_supply,json=totalSupply" json:"total_supply,omitempty"`
}

func (m *TotalSupplyResponse) Reset()                    { *m = TotalSupplyResponse{} }
func (m *TotalSupplyResponse) String() string            { return proto.CompactTextString(m) }
func (*TotalSupplyResponse) ProtoMessage()               {}
func (*TotalSupplyResponse) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{7} }

func (m *TotalSupplyResponse) GetTotalSupply() *types.BigUInt {
	if m != nil {
		return m.TotalSupply
	}
	return nil
}

type BalanceOfRequest struct {
	Owner *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
}

func (m *BalanceOfRequest) Reset()                    { *m = BalanceOfRequest{} }
func (m *BalanceOfRequest) String() string            { return proto.CompactTextString(m) }
func (*BalanceOfRequest) ProtoMessage()               {}
func (*BalanceOfRequest) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{8} }

func (m *BalanceOfRequest) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

type BalanceOfResponse struct {
	Balance *types.BigUInt `protobuf:"bytes,1,opt,name=balance" json:"balance,omitempty"`
}

func (m *BalanceOfResponse) Reset()                    { *m = BalanceOfResponse{} }
func (m *BalanceOfResponse) String() string            { return proto.CompactTextString(m) }
func (*BalanceOfResponse) ProtoMessage()               {}
func (*BalanceOfResponse) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{9} }

func (m *BalanceOfResponse) GetBalance() *types.BigUInt {
	if m != nil {
		return m.Balance
	}
	return nil
}

type AllowanceRequest struct {
	Owner   *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Spender *types.Address `protobuf:"bytes,2,opt,name=spender" json:"spender,omitempty"`
}

func (m *AllowanceRequest) Reset()                    { *m = AllowanceRequest{} }
func (m *AllowanceRequest) String() string            { return proto.CompactTextString(m) }
func (*AllowanceRequest) ProtoMessage()               {}
func (*AllowanceRequest) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{10} }

func (m *AllowanceRequest) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *AllowanceRequest) GetSpender() *types.Address {
	if m != nil {
		return m.Spender
	}
	return nil
}

type AllowanceResponse struct {
	Amount *types.BigUInt `protobuf:"bytes,1,opt,name=amount" json:"amount,omitempty"`
}

func (m *AllowanceResponse) Reset()                    { *m = AllowanceResponse{} }
func (m *AllowanceResponse) String() string            { return proto.CompactTextString(m) }
func (*AllowanceResponse) ProtoMessage()               {}
func (*AllowanceResponse) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{11} }

func (m *AllowanceResponse) GetAmount() *types.BigUInt {
	if m != nil {
		return m.Amount
	}
	return nil
}

// volatile
type ApproveRequest struct {
	Spender *types.Address `protobuf:"bytes,1,opt,name=spender" json:"spender,omitempty"`
	Amount  *types.BigUInt `protobuf:"bytes,2,opt,name=amount" json:"amount,omitempty"`
}

func (m *ApproveRequest) Reset()                    { *m = ApproveRequest{} }
func (m *ApproveRequest) String() string            { return proto.CompactTextString(m) }
func (*ApproveRequest) ProtoMessage()               {}
func (*ApproveRequest) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{12} }

func (m *ApproveRequest) GetSpender() *types.Address {
	if m != nil {
		return m.Spender
	}
	return nil
}

func (m *ApproveRequest) GetAmount() *types.BigUInt {
	if m != nil {
		return m.Amount
	}
	return nil
}

type ApproveResponse struct {
}

func (m *ApproveResponse) Reset()                    { *m = ApproveResponse{} }
func (m *ApproveResponse) String() string            { return proto.CompactTextString(m) }
func (*ApproveResponse) ProtoMessage()               {}
func (*ApproveResponse) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{13} }

type TransferRequest struct {
	To     *types.Address `protobuf:"bytes,1,opt,name=to" json:"to,omitempty"`
	Amount *types.BigUInt `protobuf:"bytes,2,opt,name=amount" json:"amount,omitempty"`
}

func (m *TransferRequest) Reset()                    { *m = TransferRequest{} }
func (m *TransferRequest) String() string            { return proto.CompactTextString(m) }
func (*TransferRequest) ProtoMessage()               {}
func (*TransferRequest) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{14} }

func (m *TransferRequest) GetTo() *types.Address {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *TransferRequest) GetAmount() *types.BigUInt {
	if m != nil {
		return m.Amount
	}
	return nil
}

type TransferResponse struct {
}

func (m *TransferResponse) Reset()                    { *m = TransferResponse{} }
func (m *TransferResponse) String() string            { return proto.CompactTextString(m) }
func (*TransferResponse) ProtoMessage()               {}
func (*TransferResponse) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{15} }

type TransferFromRequest struct {
	From   *types.Address `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	To     *types.Address `protobuf:"bytes,2,opt,name=to" json:"to,omitempty"`
	Amount *types.BigUInt `protobuf:"bytes,3,opt,name=amount" json:"amount,omitempty"`
}

func (m *TransferFromRequest) Reset()                    { *m = TransferFromRequest{} }
func (m *TransferFromRequest) String() string            { return proto.CompactTextString(m) }
func (*TransferFromRequest) ProtoMessage()               {}
func (*TransferFromRequest) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{16} }

func (m *TransferFromRequest) GetFrom() *types.Address {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *TransferFromRequest) GetTo() *types.Address {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *TransferFromRequest) GetAmount() *types.BigUInt {
	if m != nil {
		return m.Amount
	}
	return nil
}

type TransferFromResponse struct {
}

func (m *TransferFromResponse) Reset()                    { *m = TransferFromResponse{} }
func (m *TransferFromResponse) String() string            { return proto.CompactTextString(m) }
func (*TransferFromResponse) ProtoMessage()               {}
func (*TransferFromResponse) Descriptor() ([]byte, []int) { return fileDescriptorCoin, []int{17} }

func init() {
	proto.RegisterType((*Economy)(nil), "Economy")
	proto.RegisterType((*Account)(nil), "Account")
	proto.RegisterType((*Allowance)(nil), "Allowance")
	proto.RegisterType((*InitialAccount)(nil), "InitialAccount")
	proto.RegisterType((*InitRequest)(nil), "InitRequest")
	proto.RegisterType((*MintRequest)(nil), "MintRequest")
	proto.RegisterType((*TotalSupplyRequest)(nil), "TotalSupplyRequest")
	proto.RegisterType((*TotalSupplyResponse)(nil), "TotalSupplyResponse")
	proto.RegisterType((*BalanceOfRequest)(nil), "BalanceOfRequest")
	proto.RegisterType((*BalanceOfResponse)(nil), "BalanceOfResponse")
	proto.RegisterType((*AllowanceRequest)(nil), "AllowanceRequest")
	proto.RegisterType((*AllowanceResponse)(nil), "AllowanceResponse")
	proto.RegisterType((*ApproveRequest)(nil), "ApproveRequest")
	proto.RegisterType((*ApproveResponse)(nil), "ApproveResponse")
	proto.RegisterType((*TransferRequest)(nil), "TransferRequest")
	proto.RegisterType((*TransferResponse)(nil), "TransferResponse")
	proto.RegisterType((*TransferFromRequest)(nil), "TransferFromRequest")
	proto.RegisterType((*TransferFromResponse)(nil), "TransferFromResponse")
}

func init() {
	proto.RegisterFile("github.com/loomnetwork/go-loom/builtin/types/coin/coin.proto", fileDescriptorCoin)
}

var fileDescriptorCoin = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xed, 0x6b, 0xd3, 0x40,
	0x18, 0x27, 0xd9, 0x5c, 0xe7, 0x13, 0x59, 0xdb, 0x6c, 0x48, 0x10, 0x91, 0x72, 0x9f, 0x06, 0xc3,
	0x46, 0x2a, 0x2a, 0x88, 0x5f, 0x5a, 0x50, 0xa8, 0x50, 0x84, 0x3a, 0xf7, 0x55, 0xd2, 0xf4, 0x5a,
	0x0f, 0x2f, 0xf7, 0xdc, 0xee, 0x2e, 0x96, 0xfe, 0xf7, 0x92, 0xf7, 0x64, 0x25, 0x36, 0xc2, 0xbe,
	0x1c, 0xdc, 0xf3, 0xdc, 0xef, 0xe5, 0xb9, 0xfb, 0x25, 0xf0, 0x69, 0xcb, 0xcc, 0xaf, 0x78, 0x35,
	0x0e, 0x31, 0xf2, 0x39, 0x62, 0x24, 0xa8, 0xd9, 0xa1, 0xfa, 0xed, 0x6f, 0xf1, 0x75, 0xb2, 0xf5,
	0x57, 0x31, 0xe3, 0x86, 0x09, 0xdf, 0xec, 0x25, 0xd5, 0x7e, 0x88, 0x4c, 0xa4, 0xcb, 0x58, 0x2a,
	0x34, 0xf8, 0xe2, 0xcd, 0x11, 0x74, 0x86, 0x4a, 0xd7, 0x0c, 0x41, 0xde, 0x43, 0xef, 0x73, 0x88,
	0x02, 0xa3, 0xbd, 0x7b, 0x03, 0xcf, 0x0c, 0x9a, 0x80, 0xff, 0xd4, 0xb1, 0x94, 0x7c, 0xef, 0x59,
	0x23, 0xeb, 0xda, 0x99, 0x9c, 0x8f, 0x67, 0x6c, 0xfb, 0x63, 0x2e, 0xcc, 0xd2, 0x49, 0xbb, 0xdf,
	0xd3, 0x26, 0x59, 0x40, 0x6f, 0x1a, 0x86, 0x18, 0x0b, 0xe3, 0xbe, 0x82, 0x27, 0xb8, 0x13, 0x54,
	0x95, 0x80, 0xe9, 0x7a, 0xad, 0xa8, 0xd6, 0xcb, 0xac, 0xec, 0x12, 0xe8, 0xad, 0x02, 0x1e, 0x88,
	0x90, 0x7a, 0xf6, 0x03, 0xca, 0xa2, 0x41, 0xee, 0xe1, 0xe9, 0x94, 0x73, 0xdc, 0x25, 0x9b, 0x2e,
	0x84, 0x5a, 0x52, 0xb1, 0xa6, 0xaa, 0x24, 0x2c, 0x4e, 0x14, 0x0d, 0x77, 0x04, 0x67, 0x41, 0x94,
	0xd8, 0xf3, 0x4e, 0x1e, 0x68, 0xe6, 0x75, 0xf2, 0x15, 0x2e, 0xe6, 0x82, 0x19, 0x16, 0xf0, 0xae,
	0x83, 0x78, 0xcd, 0x41, 0x4e, 0x2b, 0xfb, 0x1f, 0xc1, 0x49, 0xb8, 0x96, 0xf4, 0x3e, 0xa6, 0xda,
	0xb8, 0x37, 0x70, 0x1e, 0x64, 0x9c, 0xda, 0xb3, 0x46, 0x27, 0xd7, 0xce, 0xa4, 0x3f, 0x6e, 0x6a,
	0x2d, 0xcb, 0x03, 0x64, 0x0e, 0xce, 0x82, 0x89, 0x12, 0xeb, 0x81, 0x6d, 0xf0, 0xc0, 0x81, 0x6d,
	0xb0, 0x36, 0x92, 0xdd, 0x32, 0xd2, 0x15, 0xb8, 0xb7, 0xd5, 0x1b, 0xe5, 0x8c, 0x64, 0x06, 0x97,
	0x8d, 0xaa, 0x96, 0x28, 0x34, 0xfd, 0xbf, 0xe7, 0x9e, 0xc0, 0x60, 0x96, 0xcd, 0xfa, 0x6d, 0x53,
	0x38, 0x3d, 0x72, 0x5d, 0xe4, 0x03, 0x0c, 0x6b, 0x98, 0x5c, 0xb5, 0x16, 0x06, 0xab, 0x2d, 0x0c,
	0x77, 0x30, 0x28, 0xc3, 0xd0, 0x51, 0xac, 0x4b, 0x26, 0xc8, 0x3b, 0x18, 0xd6, 0x78, 0x73, 0x43,
	0xd5, 0xad, 0x5a, 0x2d, 0xb7, 0x7a, 0x07, 0x17, 0x53, 0x29, 0x15, 0xfe, 0x29, 0xcd, 0xd4, 0xc4,
	0xac, 0xe3, 0x01, 0x6c, 0x7b, 0xad, 0x21, 0xf4, 0x4b, 0xde, 0xcc, 0x0c, 0x59, 0x40, 0xff, 0x56,
	0x05, 0x42, 0x6f, 0xa8, 0x7a, 0x8c, 0x3c, 0xb8, 0x30, 0xa8, 0xe8, 0x72, 0x09, 0x84, 0xcb, 0xa2,
	0xf6, 0x45, 0x61, 0x54, 0xc8, 0xbc, 0x84, 0xd3, 0x8d, 0xc2, 0xe8, 0x40, 0x28, 0xad, 0xe6, 0x26,
	0xec, 0x7f, 0x9a, 0x68, 0xfb, 0xce, 0x9e, 0xc3, 0x55, 0x53, 0x30, 0x33, 0xb2, 0x3a, 0x4b, 0x7f,
	0x40, 0x6f, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xa9, 0xa8, 0x9e, 0x6e, 0xf2, 0x04, 0x00, 0x00,
}
