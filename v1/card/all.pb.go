// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/openbank/openbank/v1/card/all.proto

package card

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	types "github.com/openbank/openbank/v1/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Card holds all details about a card.
type Card struct {
	// ID is a unique identifier of a card.
	ID string `protobuf:"bytes,1,opt,name=ID,json=id,proto3" json:"ID,omitempty"`
	// Account is the ID of the account associated with the card.
	Account string `protobuf:"bytes,2,opt,name=Account,json=account,proto3" json:"Account,omitempty"`
	// OwnerName is the name of the card owner.
	OwnerName string `protobuf:"bytes,3,opt,name=OwnerName,json=owner_name,proto3" json:"OwnerName,omitempty"`
	// ContactNumber is the contact number of the card owner.
	ContactNumber string `protobuf:"bytes,4,opt,name=ContactNumber,json=contact_number,proto3" json:"ContactNumber,omitempty"`
	// FirstName is the first name of card owner.
	FirstName string `protobuf:"bytes,5,opt,name=FirstName,json=first_name,proto3" json:"FirstName,omitempty"`
	// LastName is the last name of the card owner.
	LastName string `protobuf:"bytes,6,opt,name=LastName,json=last_name,proto3" json:"LastName,omitempty"`
	// Expiry is an expiry date of the card.
	Expiry *timestamp.Timestamp `protobuf:"bytes,7,opt,name=Expiry,json=expiry,proto3" json:"Expiry,omitempty"`
	// Status is the status of the card.
	Status types.CardStatus `protobuf:"varint,8,opt,name=Status,json=status,proto3,enum=types.CardStatus" json:"Status,omitempty"`
	// AccessStatus is the access status of the card.
	AccessStatus         types.CardAccessStatus `protobuf:"varint,9,opt,name=AccessStatus,json=access_status,proto3,enum=types.CardAccessStatus" json:"AccessStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Card) Reset()         { *m = Card{} }
func (m *Card) String() string { return proto.CompactTextString(m) }
func (*Card) ProtoMessage()    {}
func (*Card) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d7396a66589e9e, []int{0}
}

func (m *Card) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Card.Unmarshal(m, b)
}
func (m *Card) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Card.Marshal(b, m, deterministic)
}
func (m *Card) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Card.Merge(m, src)
}
func (m *Card) XXX_Size() int {
	return xxx_messageInfo_Card.Size(m)
}
func (m *Card) XXX_DiscardUnknown() {
	xxx_messageInfo_Card.DiscardUnknown(m)
}

var xxx_messageInfo_Card proto.InternalMessageInfo

func (m *Card) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Card) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *Card) GetOwnerName() string {
	if m != nil {
		return m.OwnerName
	}
	return ""
}

func (m *Card) GetContactNumber() string {
	if m != nil {
		return m.ContactNumber
	}
	return ""
}

func (m *Card) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Card) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Card) GetExpiry() *timestamp.Timestamp {
	if m != nil {
		return m.Expiry
	}
	return nil
}

func (m *Card) GetStatus() types.CardStatus {
	if m != nil {
		return m.Status
	}
	return types.CardStatus_UnknownCardStatus
}

func (m *Card) GetAccessStatus() types.CardAccessStatus {
	if m != nil {
		return m.AccessStatus
	}
	return types.CardAccessStatus_UnknownCardAccessStatus
}

// GetCardRequest is the request envelope to retrieve card information.
type GetCardRequest struct {
	// CardToken is the identifier to get the card information.
	CardToken            string   `protobuf:"bytes,1,opt,name=CardToken,json=card_token,proto3" json:"CardToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCardRequest) Reset()         { *m = GetCardRequest{} }
func (m *GetCardRequest) String() string { return proto.CompactTextString(m) }
func (*GetCardRequest) ProtoMessage()    {}
func (*GetCardRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d7396a66589e9e, []int{1}
}

func (m *GetCardRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCardRequest.Unmarshal(m, b)
}
func (m *GetCardRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCardRequest.Marshal(b, m, deterministic)
}
func (m *GetCardRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCardRequest.Merge(m, src)
}
func (m *GetCardRequest) XXX_Size() int {
	return xxx_messageInfo_GetCardRequest.Size(m)
}
func (m *GetCardRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCardRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCardRequest proto.InternalMessageInfo

func (m *GetCardRequest) GetCardToken() string {
	if m != nil {
		return m.CardToken
	}
	return ""
}

// UpdateCardStatusRequest is the request envelope to update card status.
type UpdateCardStatusRequest struct {
	// CardToken is the identifier to get the card information.
	CardToken string `protobuf:"bytes,1,opt,name=CardToken,json=card_token,proto3" json:"CardToken,omitempty"`
	// Status is the new card status.
	Status               types.CardStatus `protobuf:"varint,2,opt,name=Status,json=status,proto3,enum=types.CardStatus" json:"Status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UpdateCardStatusRequest) Reset()         { *m = UpdateCardStatusRequest{} }
func (m *UpdateCardStatusRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateCardStatusRequest) ProtoMessage()    {}
func (*UpdateCardStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d7396a66589e9e, []int{2}
}

func (m *UpdateCardStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCardStatusRequest.Unmarshal(m, b)
}
func (m *UpdateCardStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCardStatusRequest.Marshal(b, m, deterministic)
}
func (m *UpdateCardStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCardStatusRequest.Merge(m, src)
}
func (m *UpdateCardStatusRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateCardStatusRequest.Size(m)
}
func (m *UpdateCardStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCardStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCardStatusRequest proto.InternalMessageInfo

func (m *UpdateCardStatusRequest) GetCardToken() string {
	if m != nil {
		return m.CardToken
	}
	return ""
}

func (m *UpdateCardStatusRequest) GetStatus() types.CardStatus {
	if m != nil {
		return m.Status
	}
	return types.CardStatus_UnknownCardStatus
}

// UpdateCardAccessStatusRequest is the request envelope to update card access status.
type UpdateCardAccessStatusRequest struct {
	// CardToken is the identifier to get the card information.
	CardToken string `protobuf:"bytes,1,opt,name=CardToken,json=card_token,proto3" json:"CardToken,omitempty"`
	// AccessStatus is the new card access status.
	AccessStatus         types.CardAccessStatus `protobuf:"varint,2,opt,name=AccessStatus,json=access_status,proto3,enum=types.CardAccessStatus" json:"AccessStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *UpdateCardAccessStatusRequest) Reset()         { *m = UpdateCardAccessStatusRequest{} }
func (m *UpdateCardAccessStatusRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateCardAccessStatusRequest) ProtoMessage()    {}
func (*UpdateCardAccessStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d7396a66589e9e, []int{3}
}

func (m *UpdateCardAccessStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCardAccessStatusRequest.Unmarshal(m, b)
}
func (m *UpdateCardAccessStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCardAccessStatusRequest.Marshal(b, m, deterministic)
}
func (m *UpdateCardAccessStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCardAccessStatusRequest.Merge(m, src)
}
func (m *UpdateCardAccessStatusRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateCardAccessStatusRequest.Size(m)
}
func (m *UpdateCardAccessStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCardAccessStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCardAccessStatusRequest proto.InternalMessageInfo

func (m *UpdateCardAccessStatusRequest) GetCardToken() string {
	if m != nil {
		return m.CardToken
	}
	return ""
}

func (m *UpdateCardAccessStatusRequest) GetAccessStatus() types.CardAccessStatus {
	if m != nil {
		return m.AccessStatus
	}
	return types.CardAccessStatus_UnknownCardAccessStatus
}

// Result is result of the update operation.
type Result struct {
	// Success is a boolean indicating the success of the operation.
	Success bool `protobuf:"varint,1,opt,name=Success,json=success,proto3" json:"Success,omitempty"`
	// ErrorCode is the code of the error.
	ErrorCode string `protobuf:"bytes,2,opt,name=ErrorCode,json=error_code,proto3" json:"ErrorCode,omitempty"`
	// ErrorMessage is the message of the error.
	ErrorMessage         string   `protobuf:"bytes,3,opt,name=ErrorMessage,json=error_message,proto3" json:"ErrorMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_41d7396a66589e9e, []int{4}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Result) GetErrorCode() string {
	if m != nil {
		return m.ErrorCode
	}
	return ""
}

func (m *Result) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*Card)(nil), "card.Card")
	proto.RegisterType((*GetCardRequest)(nil), "card.GetCardRequest")
	proto.RegisterType((*UpdateCardStatusRequest)(nil), "card.UpdateCardStatusRequest")
	proto.RegisterType((*UpdateCardAccessStatusRequest)(nil), "card.UpdateCardAccessStatusRequest")
	proto.RegisterType((*Result)(nil), "card.Result")
}

func init() {
	proto.RegisterFile("github.com/openbank/openbank/v1/card/all.proto", fileDescriptor_41d7396a66589e9e)
}

var fileDescriptor_41d7396a66589e9e = []byte{
	// 1254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0xcf, 0x6f, 0x1b, 0x45,
	0x14, 0xc7, 0x77, 0xd7, 0x89, 0xe3, 0x4c, 0x7f, 0x90, 0x4e, 0xab, 0xd4, 0xac, 0x5a, 0x18, 0xa5,
	0x48, 0x0d, 0x55, 0xbb, 0x76, 0xdc, 0xf6, 0x40, 0x04, 0x45, 0xae, 0x49, 0x7f, 0x44, 0xa5, 0x8d,
	0xdc, 0x82, 0x50, 0x2f, 0xd1, 0x78, 0xf7, 0xd9, 0x5e, 0xba, 0x9e, 0x59, 0x66, 0x66, 0x93, 0x06,
	0x84, 0x84, 0x40, 0xaa, 0x72, 0xac, 0xc2, 0x15, 0xc4, 0x5f, 0x81, 0xc4, 0x11, 0x09, 0x71, 0xe2,
	0x00, 0x12, 0x17, 0x4e, 0x9c, 0xe0, 0x80, 0xc4, 0x8d, 0x03, 0x1c, 0xd1, 0xcc, 0xac, 0x7f, 0xc5,
	0x8e, 0xaa, 0x14, 0x4e, 0xf6, 0xce, 0xfb, 0xbe, 0x37, 0x6f, 0x3e, 0xef, 0xbd, 0xd9, 0x45, 0x41,
	0x27, 0x56, 0xdd, 0xac, 0x15, 0x84, 0xbc, 0x57, 0xe1, 0x29, 0xb0, 0x16, 0x65, 0x8f, 0x86, 0x7f,
	0xb6, 0x56, 0x2a, 0x21, 0x15, 0x51, 0x85, 0x26, 0x49, 0x90, 0x0a, 0xae, 0x38, 0x9e, 0xd1, 0xcf,
	0xfe, 0xcb, 0x1d, 0xce, 0x3b, 0x09, 0x54, 0xcc, 0x5a, 0x2b, 0x6b, 0x57, 0x54, 0xdc, 0x03, 0xa9,
	0x68, 0x2f, 0xb5, 0x32, 0xff, 0x4c, 0x2e, 0xa0, 0x69, 0x5c, 0xa1, 0x8c, 0x71, 0x45, 0x55, 0xcc,
	0x99, 0xcc, 0xad, 0x17, 0xcd, 0x4f, 0x78, 0xa9, 0x03, 0xec, 0x92, 0xdc, 0xa6, 0x9d, 0x0e, 0x88,
	0x0a, 0x4f, 0x8d, 0x62, 0x8a, 0xba, 0xf2, 0xac, 0x14, 0xd5, 0x4e, 0x0a, 0x72, 0x98, 0xe3, 0xd2,
	0x77, 0x05, 0x34, 0xd3, 0xa0, 0x22, 0xc2, 0x3e, 0xf2, 0x6e, 0xbf, 0x55, 0x76, 0x89, 0xbb, 0x3c,
	0x7f, 0x1d, 0x95, 0x9c, 0xb2, 0xb3, 0xec, 0x54, 0x9d, 0x0d, 0xa7, 0xe9, 0xc5, 0x11, 0x7e, 0x05,
	0xcd, 0xd5, 0xc3, 0x90, 0x67, 0x4c, 0x95, 0xbd, 0x09, 0xc1, 0x1c, 0xb5, 0x26, 0xfc, 0x2a, 0x9a,
	0xbf, 0xb7, 0xcd, 0x40, 0xdc, 0xa5, 0x3d, 0x28, 0x17, 0x26, 0x74, 0x88, 0x6b, 0xe3, 0x26, 0xa3,
	0x3d, 0xc0, 0x2b, 0xe8, 0x58, 0x83, 0x33, 0x45, 0x43, 0x75, 0x37, 0xeb, 0xb5, 0x40, 0x94, 0x67,
	0x26, 0xe4, 0xc7, 0x43, 0x2b, 0xd8, 0x64, 0x46, 0xa1, 0xa3, 0xdf, 0x88, 0x85, 0x54, 0x26, 0xfa,
	0xec, 0x64, 0xf4, 0xb6, 0x36, 0xda, 0xe8, 0xe7, 0x51, 0xe9, 0x0e, 0xcd, 0x95, 0xc5, 0x09, 0xe5,
	0x7c, 0x42, 0xfb, 0xc2, 0x6b, 0xa8, 0xb8, 0xf6, 0x38, 0x8d, 0xc5, 0x4e, 0x79, 0x8e, 0xb8, 0xcb,
	0x47, 0x6a, 0x7e, 0x60, 0x4b, 0x11, 0xf4, 0x6b, 0x15, 0x3c, 0xe8, 0xd7, 0x6a, 0x2c, 0x44, 0x11,
	0x8c, 0x17, 0xbe, 0x8a, 0x8a, 0xf7, 0x15, 0x55, 0x99, 0x2c, 0x97, 0x88, 0xbb, 0x7c, 0xbc, 0x76,
	0x22, 0x30, 0x78, 0x03, 0x0d, 0xd4, 0x1a, 0xc6, 0xdd, 0xa4, 0x59, 0xc3, 0xb7, 0xd0, 0xd1, 0x7a,
	0x18, 0x82, 0x94, 0xb9, 0xf3, 0xbc, 0x71, 0x3e, 0x3d, 0xe2, 0x3c, 0x6a, 0x1e, 0x0b, 0x71, 0x8c,
	0x1a, 0xcb, 0xa6, 0x8d, 0xb4, 0x5a, 0x2c, 0x39, 0x0b, 0x4e, 0xd9, 0x59, 0x6a, 0xa0, 0xe3, 0x37,
	0x41, 0x69, 0xcf, 0x26, 0x7c, 0x90, 0x81, 0x34, 0xc5, 0xd0, 0x8f, 0x0f, 0xf8, 0x23, 0x60, 0x53,
	0xaa, 0x8a, 0x74, 0x6b, 0x6e, 0x2a, 0x6d, 0x1d, 0x04, 0xf9, 0xcc, 0x45, 0xa7, 0xdf, 0x49, 0x23,
	0xaa, 0x60, 0x98, 0xff, 0xe1, 0xc3, 0x8d, 0x40, 0xf1, 0x0e, 0x01, 0x65, 0x90, 0xc5, 0x17, 0x2e,
	0x3a, 0x3b, 0xcc, 0x62, 0x14, 0xc4, 0x73, 0xe4, 0xb2, 0x9f, 0xb4, 0xf7, 0x9f, 0x49, 0xef, 0xb9,
	0xa8, 0xd8, 0x04, 0x99, 0x25, 0x4a, 0x4f, 0xc5, 0xfd, 0xcc, 0x88, 0x4c, 0x16, 0xa5, 0xf1, 0xa9,
	0x90, 0xd6, 0xa4, 0xb3, 0x5d, 0x13, 0x82, 0x8b, 0x06, 0x8f, 0x60, 0xca, 0xf4, 0x20, 0xd0, 0xc6,
	0xcd, 0x90, 0x47, 0x80, 0x2b, 0xe8, 0xa8, 0x91, 0xbe, 0x0d, 0x52, 0xd2, 0xce, 0xb4, 0x19, 0x3a,
	0x66, 0xd5, 0x3d, 0x2b, 0xe8, 0x27, 0x55, 0xfb, 0x7a, 0x0e, 0x1d, 0x31, 0x78, 0x41, 0x6c, 0xc5,
	0x21, 0xe0, 0x1f, 0x3d, 0x34, 0x97, 0xf7, 0x03, 0x3e, 0x15, 0x68, 0x1e, 0xc1, 0x78, 0x7b, 0xf8,
	0xc8, 0xae, 0xea, 0xa5, 0xa5, 0x2f, 0xbd, 0xbd, 0xfa, 0xdf, 0x6e, 0x7e, 0x0d, 0xbc, 0xd8, 0x04,
	0x25, 0x62, 0xd8, 0x02, 0xa2, 0x05, 0x24, 0x66, 0x6d, 0x2e, 0x7a, 0xe6, 0x8a, 0xf1, 0x6f, 0xf5,
	0x4d, 0x92, 0xd0, 0x24, 0x21, 0x11, 0x55, 0x94, 0xb4, 0x05, 0xef, 0x11, 0x6a, 0xb4, 0x17, 0x89,
	0x84, 0x04, 0x42, 0x05, 0x11, 0x69, 0xed, 0x10, 0xd5, 0xb5, 0x11, 0x6c, 0x21, 0xc8, 0x0e, 0xcf,
	0x88, 0xcc, 0xd2, 0x34, 0x89, 0x21, 0x0a, 0xd6, 0x1b, 0xa8, 0x50, 0xab, 0x56, 0xf1, 0xeb, 0xe8,
	0xa5, 0x3c, 0x1f, 0x02, 0x8f, 0x21, 0xcc, 0xb4, 0x6b, 0x8e, 0xad, 0x9d, 0x25, 0xc9, 0x4e, 0x80,
	0x7d, 0x54, 0xf6, 0x17, 0xcf, 0x55, 0x22, 0x68, 0xc7, 0x2c, 0xb6, 0x77, 0x9e, 0x0e, 0xaa, 0x33,
	0x5d, 0x5f, 0x41, 0x85, 0x2b, 0xd5, 0x2b, 0xf8, 0x02, 0x5a, 0x6e, 0x82, 0xca, 0x04, 0x83, 0x88,
	0x6c, 0x77, 0x81, 0x99, 0x9d, 0x05, 0x48, 0x9e, 0x89, 0x10, 0x48, 0x2c, 0x09, 0xe3, 0x8a, 0xb4,
	0x79, 0xc6, 0xa2, 0xa0, 0x85, 0xd1, 0x02, 0x2a, 0xde, 0xab, 0x67, 0xaa, 0x5b, 0xc3, 0x45, 0x34,
	0x23, 0x80, 0x46, 0x9f, 0xfe, 0xfc, 0xdb, 0xe7, 0xde, 0x22, 0x3e, 0x35, 0xb8, 0xbe, 0x3f, 0x1a,
	0xf4, 0xd6, 0xc7, 0xbb, 0x9e, 0xf3, 0xd4, 0x33, 0xe8, 0xf1, 0x13, 0x0f, 0x2d, 0xec, 0x9f, 0x0d,
	0x7c, 0xd6, 0x42, 0x3c, 0x60, 0x66, 0xfc, 0xa3, 0xd6, 0x6c, 0xbb, 0x65, 0xe9, 0x5b, 0x77, 0xaf,
	0xfe, 0x55, 0x9f, 0x32, 0xb6, 0x3e, 0x96, 0xb1, 0xed, 0x30, 0xff, 0x64, 0xbe, 0xd6, 0x18, 0xae,
	0x05, 0xcf, 0x71, 0xc8, 0xf5, 0xf3, 0x16, 0x2e, 0x79, 0x16, 0xdc, 0x03, 0x69, 0x9c, 0x58, 0x7a,
	0x61, 0x40, 0xc3, 0xe6, 0x31, 0x02, 0xe2, 0x7b, 0x0f, 0x2d, 0x4e, 0x1f, 0x4f, 0x7c, 0x6e, 0x3f,
	0x8e, 0x29, 0xc3, 0xbb, 0x0f, 0xca, 0x5f, 0xee, 0x5e, 0xfd, 0x87, 0x41, 0xeb, 0x8d, 0x42, 0xb1,
	0xd3, 0xd7, 0x67, 0xe3, 0x1f, 0x68, 0x0a, 0xd6, 0xd7, 0xec, 0x79, 0xaf, 0x3d, 0xb3, 0x99, 0xce,
	0x20, 0xdf, 0x2f, 0x4f, 0x36, 0x93, 0x4d, 0xe5, 0xff, 0x6c, 0xa7, 0xf2, 0xd2, 0xe2, 0xf0, 0x6b,
	0x60, 0xf4, 0x16, 0x19, 0x72, 0xf4, 0x0b, 0xbb, 0x9e, 0x73, 0x7d, 0xb7, 0xb8, 0x57, 0xff, 0x7d,
	0x16, 0x15, 0x6a, 0x41, 0x15, 0xdf, 0x41, 0x25, 0x53, 0xf9, 0xfa, 0xc6, 0x6d, 0xfc, 0xda, 0x86,
	0xe0, 0x5b, 0x71, 0x04, 0x92, 0x84, 0x02, 0xf4, 0xb9, 0x29, 0x8b, 0x88, 0x0e, 0x4f, 0x78, 0x0a,
	0xc2, 0xbe, 0xf0, 0x09, 0x67, 0x83, 0x21, 0x1b, 0x24, 0x18, 0xd4, 0x66, 0x57, 0x82, 0x6a, 0x50,
	0xbd, 0xe0, 0x7a, 0xb5, 0x05, 0xaa, 0xc7, 0x2d, 0x34, 0xea, 0xca, 0xfb, 0x92, 0xb3, 0xd5, 0x89,
	0x95, 0xe6, 0xa6, 0x3e, 0x74, 0x15, 0xbf, 0x87, 0xde, 0x9d, 0x76, 0x68, 0x4b, 0xb3, 0xc5, 0xa3,
	0x1d, 0x7d, 0xf0, 0x1e, 0x4d, 0xec, 0x5d, 0xa0, 0xc9, 0x72, 0x41, 0x22, 0x0e, 0x96, 0x46, 0x8f,
	0xaa, 0xb0, 0x6b, 0x5c, 0xe0, 0x71, 0x6a, 0x2f, 0x80, 0xdc, 0x37, 0x68, 0xde, 0xd1, 0x1b, 0xac,
	0xe0, 0x35, 0xd4, 0x38, 0x78, 0x83, 0x41, 0x20, 0xf3, 0xc6, 0x8f, 0x99, 0x34, 0xd6, 0x4c, 0x82,
	0x38, 0x6f, 0x00, 0x44, 0xc0, 0x54, 0x4c, 0x13, 0x19, 0x34, 0x37, 0x74, 0xb4, 0xcb, 0xf8, 0x36,
	0xba, 0x39, 0x19, 0x4d, 0xeb, 0x87, 0xa1, 0xba, 0x74, 0x0b, 0x48, 0x0a, 0xa2, 0x17, 0x4b, 0x19,
	0x6b, 0x52, 0xbc, 0xdf, 0x36, 0xa3, 0xf5, 0x0c, 0x9a, 0x87, 0xaf, 0x7a, 0xf3, 0x06, 0x2a, 0x5c,
	0xad, 0x56, 0xf1, 0x9b, 0xe8, 0x8d, 0x71, 0x17, 0xca, 0x48, 0xc6, 0x06, 0x04, 0xcc, 0x45, 0x4d,
	0x78, 0x18, 0x66, 0x42, 0xe3, 0xb2, 0x11, 0x25, 0x88, 0x2d, 0x10, 0x44, 0xc6, 0x11, 0x04, 0x0f,
	0xff, 0x74, 0xd1, 0x1f, 0xee, 0xa0, 0x7f, 0x7e, 0x75, 0x4b, 0x05, 0xfc, 0xc4, 0xad, 0xe7, 0x49,
	0x72, 0xa2, 0x04, 0x65, 0x92, 0x86, 0xb6, 0xd6, 0xfd, 0x54, 0xa4, 0xce, 0x45, 0x80, 0x54, 0x22,
	0x36, 0xbb, 0xe8, 0x63, 0x65, 0xaa, 0xab, 0x01, 0x85, 0x54, 0x2f, 0x68, 0x0a, 0x32, 0x20, 0x0f,
	0xba, 0x30, 0x6a, 0xd0, 0x04, 0x52, 0xc1, 0x4d, 0xe8, 0x36, 0x4f, 0x12, 0xbe, 0x6d, 0x39, 0xe8,
	0xad, 0xb9, 0x88, 0x3f, 0xb4, 0x0a, 0xfd, 0x86, 0x22, 0xed, 0x84, 0x6f, 0x07, 0xcb, 0x33, 0xb5,
	0x92, 0xee, 0x60, 0x1d, 0x62, 0x75, 0xde, 0x7c, 0x36, 0xea, 0x0b, 0xf1, 0xfa, 0x2a, 0xf2, 0x6d,
	0x9b, 0x63, 0x7c, 0x53, 0x50, 0xa6, 0xa4, 0x6d, 0x4a, 0x4b, 0x16, 0x9d, 0x41, 0xb3, 0xdb, 0x22,
	0x56, 0x80, 0x4f, 0xe6, 0x46, 0xf3, 0x94, 0x5b, 0x6f, 0xb9, 0x1b, 0xce, 0x43, 0xf3, 0x41, 0xfc,
	0x89, 0xeb, 0xec, 0xba, 0xce, 0x53, 0xd7, 0xf9, 0xc6, 0x75, 0x7e, 0x71, 0x9d, 0x7f, 0x5c, 0xe7,
	0x27, 0xcf, 0x69, 0x15, 0xcd, 0x27, 0xd7, 0xe5, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x9a, 0xea,
	0x89, 0x4e, 0x64, 0x0b, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CardServiceClient is the client API for CardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CardServiceClient interface {
	// GetCard retrieves the detail of a card information, selected by its token.
	GetCard(ctx context.Context, in *GetCardRequest, opts ...grpc.CallOption) (*Card, error)
	// UpdateCardStatus update the card status.
	UpdateCardStatus(ctx context.Context, in *UpdateCardStatusRequest, opts ...grpc.CallOption) (*Result, error)
	// UpdateCardAccessStatus update the card access status.
	UpdateCardAccessStatus(ctx context.Context, in *UpdateCardAccessStatusRequest, opts ...grpc.CallOption) (*Result, error)
}

type cardServiceClient struct {
	cc *grpc.ClientConn
}

func NewCardServiceClient(cc *grpc.ClientConn) CardServiceClient {
	return &cardServiceClient{cc}
}

func (c *cardServiceClient) GetCard(ctx context.Context, in *GetCardRequest, opts ...grpc.CallOption) (*Card, error) {
	out := new(Card)
	err := c.cc.Invoke(ctx, "/card.CardService/GetCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardServiceClient) UpdateCardStatus(ctx context.Context, in *UpdateCardStatusRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/card.CardService/UpdateCardStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardServiceClient) UpdateCardAccessStatus(ctx context.Context, in *UpdateCardAccessStatusRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/card.CardService/UpdateCardAccessStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CardServiceServer is the server API for CardService service.
type CardServiceServer interface {
	// GetCard retrieves the detail of a card information, selected by its token.
	GetCard(context.Context, *GetCardRequest) (*Card, error)
	// UpdateCardStatus update the card status.
	UpdateCardStatus(context.Context, *UpdateCardStatusRequest) (*Result, error)
	// UpdateCardAccessStatus update the card access status.
	UpdateCardAccessStatus(context.Context, *UpdateCardAccessStatusRequest) (*Result, error)
}

// UnimplementedCardServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCardServiceServer struct {
}

func (*UnimplementedCardServiceServer) GetCard(ctx context.Context, req *GetCardRequest) (*Card, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCard not implemented")
}
func (*UnimplementedCardServiceServer) UpdateCardStatus(ctx context.Context, req *UpdateCardStatusRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCardStatus not implemented")
}
func (*UnimplementedCardServiceServer) UpdateCardAccessStatus(ctx context.Context, req *UpdateCardAccessStatusRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCardAccessStatus not implemented")
}

func RegisterCardServiceServer(s *grpc.Server, srv CardServiceServer) {
	s.RegisterService(&_CardService_serviceDesc, srv)
}

func _CardService_GetCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardServiceServer).GetCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/card.CardService/GetCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardServiceServer).GetCard(ctx, req.(*GetCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CardService_UpdateCardStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCardStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardServiceServer).UpdateCardStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/card.CardService/UpdateCardStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardServiceServer).UpdateCardStatus(ctx, req.(*UpdateCardStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CardService_UpdateCardAccessStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCardAccessStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardServiceServer).UpdateCardAccessStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/card.CardService/UpdateCardAccessStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardServiceServer).UpdateCardAccessStatus(ctx, req.(*UpdateCardAccessStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CardService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "card.CardService",
	HandlerType: (*CardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCard",
			Handler:    _CardService_GetCard_Handler,
		},
		{
			MethodName: "UpdateCardStatus",
			Handler:    _CardService_UpdateCardStatus_Handler,
		},
		{
			MethodName: "UpdateCardAccessStatus",
			Handler:    _CardService_UpdateCardAccessStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/openbank/openbank/v1/card/all.proto",
}
