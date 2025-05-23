// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: account/v1/address.proto

package accountv1

import (
	common "github.com/shopnexus/shopnexus-protobuf-gen-go/pb/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ADDRESS ENTITY
type AddressEntity struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        int64                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FullName      string                 `protobuf:"bytes,3,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Phone         string                 `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Address       string                 `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	City          string                 `protobuf:"bytes,6,opt,name=city,proto3" json:"city,omitempty"`
	Province      string                 `protobuf:"bytes,7,opt,name=province,proto3" json:"province,omitempty"`
	Country       string                 `protobuf:"bytes,8,opt,name=country,proto3" json:"country,omitempty"`
	DateCreated   int64                  `protobuf:"varint,9,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty"`
	DateUpdated   int64                  `protobuf:"varint,10,opt,name=date_updated,json=dateUpdated,proto3" json:"date_updated,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddressEntity) Reset() {
	*x = AddressEntity{}
	mi := &file_account_v1_address_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddressEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressEntity) ProtoMessage() {}

func (x *AddressEntity) ProtoReflect() protoreflect.Message {
	mi := &file_account_v1_address_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressEntity.ProtoReflect.Descriptor instead.
func (*AddressEntity) Descriptor() ([]byte, []int) {
	return file_account_v1_address_proto_rawDescGZIP(), []int{0}
}

func (x *AddressEntity) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AddressEntity) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AddressEntity) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *AddressEntity) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *AddressEntity) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AddressEntity) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *AddressEntity) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *AddressEntity) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *AddressEntity) GetDateCreated() int64 {
	if x != nil {
		return x.DateCreated
	}
	return 0
}

func (x *AddressEntity) GetDateUpdated() int64 {
	if x != nil {
		return x.DateUpdated
	}
	return 0
}

type GetAddressRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAddressRequest) Reset() {
	*x = GetAddressRequest{}
	mi := &file_account_v1_address_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAddressRequest) ProtoMessage() {}

func (x *GetAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_v1_address_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAddressRequest.ProtoReflect.Descriptor instead.
func (*GetAddressRequest) Descriptor() ([]byte, []int) {
	return file_account_v1_address_proto_rawDescGZIP(), []int{1}
}

func (x *GetAddressRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetAddressResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          *AddressEntity         `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAddressResponse) Reset() {
	*x = GetAddressResponse{}
	mi := &file_account_v1_address_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAddressResponse) ProtoMessage() {}

func (x *GetAddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_v1_address_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAddressResponse.ProtoReflect.Descriptor instead.
func (*GetAddressResponse) Descriptor() ([]byte, []int) {
	return file_account_v1_address_proto_rawDescGZIP(), []int{2}
}

func (x *GetAddressResponse) GetData() *AddressEntity {
	if x != nil {
		return x.Data
	}
	return nil
}

type ListAddressesRequest struct {
	state           protoimpl.MessageState    `protogen:"open.v1"`
	Pagination      *common.PaginationRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	UserId          *int64                    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3,oneof" json:"user_id,omitempty"`
	FullName        *string                   `protobuf:"bytes,3,opt,name=full_name,json=fullName,proto3,oneof" json:"full_name,omitempty"`
	Phone           *string                   `protobuf:"bytes,4,opt,name=phone,proto3,oneof" json:"phone,omitempty"`
	Address         *string                   `protobuf:"bytes,5,opt,name=address,proto3,oneof" json:"address,omitempty"`
	City            *string                   `protobuf:"bytes,6,opt,name=city,proto3,oneof" json:"city,omitempty"`
	Province        *string                   `protobuf:"bytes,7,opt,name=province,proto3,oneof" json:"province,omitempty"`
	Country         *string                   `protobuf:"bytes,8,opt,name=country,proto3,oneof" json:"country,omitempty"`
	DateCreatedFrom *int64                    `protobuf:"varint,9,opt,name=date_created_from,json=dateCreatedFrom,proto3,oneof" json:"date_created_from,omitempty"`
	DateCreatedTo   *int64                    `protobuf:"varint,10,opt,name=date_created_to,json=dateCreatedTo,proto3,oneof" json:"date_created_to,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ListAddressesRequest) Reset() {
	*x = ListAddressesRequest{}
	mi := &file_account_v1_address_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAddressesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAddressesRequest) ProtoMessage() {}

func (x *ListAddressesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_v1_address_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAddressesRequest.ProtoReflect.Descriptor instead.
func (*ListAddressesRequest) Descriptor() ([]byte, []int) {
	return file_account_v1_address_proto_rawDescGZIP(), []int{3}
}

func (x *ListAddressesRequest) GetPagination() *common.PaginationRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ListAddressesRequest) GetUserId() int64 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

func (x *ListAddressesRequest) GetFullName() string {
	if x != nil && x.FullName != nil {
		return *x.FullName
	}
	return ""
}

func (x *ListAddressesRequest) GetPhone() string {
	if x != nil && x.Phone != nil {
		return *x.Phone
	}
	return ""
}

func (x *ListAddressesRequest) GetAddress() string {
	if x != nil && x.Address != nil {
		return *x.Address
	}
	return ""
}

func (x *ListAddressesRequest) GetCity() string {
	if x != nil && x.City != nil {
		return *x.City
	}
	return ""
}

func (x *ListAddressesRequest) GetProvince() string {
	if x != nil && x.Province != nil {
		return *x.Province
	}
	return ""
}

func (x *ListAddressesRequest) GetCountry() string {
	if x != nil && x.Country != nil {
		return *x.Country
	}
	return ""
}

func (x *ListAddressesRequest) GetDateCreatedFrom() int64 {
	if x != nil && x.DateCreatedFrom != nil {
		return *x.DateCreatedFrom
	}
	return 0
}

func (x *ListAddressesRequest) GetDateCreatedTo() int64 {
	if x != nil && x.DateCreatedTo != nil {
		return *x.DateCreatedTo
	}
	return 0
}

type ListAddressesResponse struct {
	state         protoimpl.MessageState     `protogen:"open.v1"`
	Data          []*AddressEntity           `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Pagination    *common.PaginationResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAddressesResponse) Reset() {
	*x = ListAddressesResponse{}
	mi := &file_account_v1_address_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAddressesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAddressesResponse) ProtoMessage() {}

func (x *ListAddressesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_v1_address_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAddressesResponse.ProtoReflect.Descriptor instead.
func (*ListAddressesResponse) Descriptor() ([]byte, []int) {
	return file_account_v1_address_proto_rawDescGZIP(), []int{4}
}

func (x *ListAddressesResponse) GetData() []*AddressEntity {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ListAddressesResponse) GetPagination() *common.PaginationResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type CreateAddressRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FullName      string                 `protobuf:"bytes,1,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Phone         string                 `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Address       string                 `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	City          string                 `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	Province      string                 `protobuf:"bytes,5,opt,name=province,proto3" json:"province,omitempty"`
	Country       string                 `protobuf:"bytes,6,opt,name=country,proto3" json:"country,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAddressRequest) Reset() {
	*x = CreateAddressRequest{}
	mi := &file_account_v1_address_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAddressRequest) ProtoMessage() {}

func (x *CreateAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_v1_address_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAddressRequest.ProtoReflect.Descriptor instead.
func (*CreateAddressRequest) Descriptor() ([]byte, []int) {
	return file_account_v1_address_proto_rawDescGZIP(), []int{5}
}

func (x *CreateAddressRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *CreateAddressRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateAddressRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CreateAddressRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *CreateAddressRequest) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *CreateAddressRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type CreateAddressResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          *AddressEntity         `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAddressResponse) Reset() {
	*x = CreateAddressResponse{}
	mi := &file_account_v1_address_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAddressResponse) ProtoMessage() {}

func (x *CreateAddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_v1_address_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAddressResponse.ProtoReflect.Descriptor instead.
func (*CreateAddressResponse) Descriptor() ([]byte, []int) {
	return file_account_v1_address_proto_rawDescGZIP(), []int{6}
}

func (x *CreateAddressResponse) GetData() *AddressEntity {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateAddressRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        *int64                 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3,oneof" json:"user_id,omitempty"`
	FullName      *string                `protobuf:"bytes,3,opt,name=full_name,json=fullName,proto3,oneof" json:"full_name,omitempty"`
	Phone         *string                `protobuf:"bytes,4,opt,name=phone,proto3,oneof" json:"phone,omitempty"`
	Address       *string                `protobuf:"bytes,5,opt,name=address,proto3,oneof" json:"address,omitempty"`
	City          *string                `protobuf:"bytes,6,opt,name=city,proto3,oneof" json:"city,omitempty"`
	Province      *string                `protobuf:"bytes,7,opt,name=province,proto3,oneof" json:"province,omitempty"`
	Country       *string                `protobuf:"bytes,8,opt,name=country,proto3,oneof" json:"country,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateAddressRequest) Reset() {
	*x = UpdateAddressRequest{}
	mi := &file_account_v1_address_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAddressRequest) ProtoMessage() {}

func (x *UpdateAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_v1_address_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAddressRequest.ProtoReflect.Descriptor instead.
func (*UpdateAddressRequest) Descriptor() ([]byte, []int) {
	return file_account_v1_address_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateAddressRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateAddressRequest) GetUserId() int64 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

func (x *UpdateAddressRequest) GetFullName() string {
	if x != nil && x.FullName != nil {
		return *x.FullName
	}
	return ""
}

func (x *UpdateAddressRequest) GetPhone() string {
	if x != nil && x.Phone != nil {
		return *x.Phone
	}
	return ""
}

func (x *UpdateAddressRequest) GetAddress() string {
	if x != nil && x.Address != nil {
		return *x.Address
	}
	return ""
}

func (x *UpdateAddressRequest) GetCity() string {
	if x != nil && x.City != nil {
		return *x.City
	}
	return ""
}

func (x *UpdateAddressRequest) GetProvince() string {
	if x != nil && x.Province != nil {
		return *x.Province
	}
	return ""
}

func (x *UpdateAddressRequest) GetCountry() string {
	if x != nil && x.Country != nil {
		return *x.Country
	}
	return ""
}

type UpdateAddressResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateAddressResponse) Reset() {
	*x = UpdateAddressResponse{}
	mi := &file_account_v1_address_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAddressResponse) ProtoMessage() {}

func (x *UpdateAddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_v1_address_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAddressResponse.ProtoReflect.Descriptor instead.
func (*UpdateAddressResponse) Descriptor() ([]byte, []int) {
	return file_account_v1_address_proto_rawDescGZIP(), []int{8}
}

type DeleteAddressRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAddressRequest) Reset() {
	*x = DeleteAddressRequest{}
	mi := &file_account_v1_address_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAddressRequest) ProtoMessage() {}

func (x *DeleteAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_v1_address_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAddressRequest.ProtoReflect.Descriptor instead.
func (*DeleteAddressRequest) Descriptor() ([]byte, []int) {
	return file_account_v1_address_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteAddressRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteAddressResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAddressResponse) Reset() {
	*x = DeleteAddressResponse{}
	mi := &file_account_v1_address_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAddressResponse) ProtoMessage() {}

func (x *DeleteAddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_v1_address_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAddressResponse.ProtoReflect.Descriptor instead.
func (*DeleteAddressResponse) Descriptor() ([]byte, []int) {
	return file_account_v1_address_proto_rawDescGZIP(), []int{10}
}

var File_account_v1_address_proto protoreflect.FileDescriptor

const file_account_v1_address_proto_rawDesc = "" +
	"\n" +
	"\x18account/v1/address.proto\x12\n" +
	"account.v1\x1a\x17common/pagination.proto\"\x95\x02\n" +
	"\rAddressEntity\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\x03R\x06userId\x12\x1b\n" +
	"\tfull_name\x18\x03 \x01(\tR\bfullName\x12\x14\n" +
	"\x05phone\x18\x04 \x01(\tR\x05phone\x12\x18\n" +
	"\aaddress\x18\x05 \x01(\tR\aaddress\x12\x12\n" +
	"\x04city\x18\x06 \x01(\tR\x04city\x12\x1a\n" +
	"\bprovince\x18\a \x01(\tR\bprovince\x12\x18\n" +
	"\acountry\x18\b \x01(\tR\acountry\x12!\n" +
	"\fdate_created\x18\t \x01(\x03R\vdateCreated\x12!\n" +
	"\fdate_updated\x18\n" +
	" \x01(\x03R\vdateUpdated\"#\n" +
	"\x11GetAddressRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\"C\n" +
	"\x12GetAddressResponse\x12-\n" +
	"\x04data\x18\x01 \x01(\v2\x19.account.v1.AddressEntityR\x04data\"\xfe\x03\n" +
	"\x14ListAddressesRequest\x129\n" +
	"\n" +
	"pagination\x18\x01 \x01(\v2\x19.common.PaginationRequestR\n" +
	"pagination\x12\x1c\n" +
	"\auser_id\x18\x02 \x01(\x03H\x00R\x06userId\x88\x01\x01\x12 \n" +
	"\tfull_name\x18\x03 \x01(\tH\x01R\bfullName\x88\x01\x01\x12\x19\n" +
	"\x05phone\x18\x04 \x01(\tH\x02R\x05phone\x88\x01\x01\x12\x1d\n" +
	"\aaddress\x18\x05 \x01(\tH\x03R\aaddress\x88\x01\x01\x12\x17\n" +
	"\x04city\x18\x06 \x01(\tH\x04R\x04city\x88\x01\x01\x12\x1f\n" +
	"\bprovince\x18\a \x01(\tH\x05R\bprovince\x88\x01\x01\x12\x1d\n" +
	"\acountry\x18\b \x01(\tH\x06R\acountry\x88\x01\x01\x12/\n" +
	"\x11date_created_from\x18\t \x01(\x03H\aR\x0fdateCreatedFrom\x88\x01\x01\x12+\n" +
	"\x0fdate_created_to\x18\n" +
	" \x01(\x03H\bR\rdateCreatedTo\x88\x01\x01B\n" +
	"\n" +
	"\b_user_idB\f\n" +
	"\n" +
	"_full_nameB\b\n" +
	"\x06_phoneB\n" +
	"\n" +
	"\b_addressB\a\n" +
	"\x05_cityB\v\n" +
	"\t_provinceB\n" +
	"\n" +
	"\b_countryB\x14\n" +
	"\x12_date_created_fromB\x12\n" +
	"\x10_date_created_to\"\x82\x01\n" +
	"\x15ListAddressesResponse\x12-\n" +
	"\x04data\x18\x01 \x03(\v2\x19.account.v1.AddressEntityR\x04data\x12:\n" +
	"\n" +
	"pagination\x18\x02 \x01(\v2\x1a.common.PaginationResponseR\n" +
	"pagination\"\xad\x01\n" +
	"\x14CreateAddressRequest\x12\x1b\n" +
	"\tfull_name\x18\x01 \x01(\tR\bfullName\x12\x14\n" +
	"\x05phone\x18\x02 \x01(\tR\x05phone\x12\x18\n" +
	"\aaddress\x18\x03 \x01(\tR\aaddress\x12\x12\n" +
	"\x04city\x18\x04 \x01(\tR\x04city\x12\x1a\n" +
	"\bprovince\x18\x05 \x01(\tR\bprovince\x12\x18\n" +
	"\acountry\x18\x06 \x01(\tR\acountry\"F\n" +
	"\x15CreateAddressResponse\x12-\n" +
	"\x04data\x18\x01 \x01(\v2\x19.account.v1.AddressEntityR\x04data\"\xcb\x02\n" +
	"\x14UpdateAddressRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x1c\n" +
	"\auser_id\x18\x02 \x01(\x03H\x00R\x06userId\x88\x01\x01\x12 \n" +
	"\tfull_name\x18\x03 \x01(\tH\x01R\bfullName\x88\x01\x01\x12\x19\n" +
	"\x05phone\x18\x04 \x01(\tH\x02R\x05phone\x88\x01\x01\x12\x1d\n" +
	"\aaddress\x18\x05 \x01(\tH\x03R\aaddress\x88\x01\x01\x12\x17\n" +
	"\x04city\x18\x06 \x01(\tH\x04R\x04city\x88\x01\x01\x12\x1f\n" +
	"\bprovince\x18\a \x01(\tH\x05R\bprovince\x88\x01\x01\x12\x1d\n" +
	"\acountry\x18\b \x01(\tH\x06R\acountry\x88\x01\x01B\n" +
	"\n" +
	"\b_user_idB\f\n" +
	"\n" +
	"_full_nameB\b\n" +
	"\x06_phoneB\n" +
	"\n" +
	"\b_addressB\a\n" +
	"\x05_cityB\v\n" +
	"\t_provinceB\n" +
	"\n" +
	"\b_country\"\x17\n" +
	"\x15UpdateAddressResponse\"&\n" +
	"\x14DeleteAddressRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\"\x17\n" +
	"\x15DeleteAddressResponseB\xaf\x01\n" +
	"\x0ecom.account.v1B\fAddressProtoP\x01ZFgithub.com/shopnexus/shopnexus-protobuf-gen-go/pb/account/v1;accountv1\xa2\x02\x03AXX\xaa\x02\n" +
	"Account.V1\xca\x02\n" +
	"Account\\V1\xe2\x02\x16Account\\V1\\GPBMetadata\xea\x02\vAccount::V1b\x06proto3"

var (
	file_account_v1_address_proto_rawDescOnce sync.Once
	file_account_v1_address_proto_rawDescData []byte
)

func file_account_v1_address_proto_rawDescGZIP() []byte {
	file_account_v1_address_proto_rawDescOnce.Do(func() {
		file_account_v1_address_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_account_v1_address_proto_rawDesc), len(file_account_v1_address_proto_rawDesc)))
	})
	return file_account_v1_address_proto_rawDescData
}

var file_account_v1_address_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_account_v1_address_proto_goTypes = []any{
	(*AddressEntity)(nil),             // 0: account.v1.AddressEntity
	(*GetAddressRequest)(nil),         // 1: account.v1.GetAddressRequest
	(*GetAddressResponse)(nil),        // 2: account.v1.GetAddressResponse
	(*ListAddressesRequest)(nil),      // 3: account.v1.ListAddressesRequest
	(*ListAddressesResponse)(nil),     // 4: account.v1.ListAddressesResponse
	(*CreateAddressRequest)(nil),      // 5: account.v1.CreateAddressRequest
	(*CreateAddressResponse)(nil),     // 6: account.v1.CreateAddressResponse
	(*UpdateAddressRequest)(nil),      // 7: account.v1.UpdateAddressRequest
	(*UpdateAddressResponse)(nil),     // 8: account.v1.UpdateAddressResponse
	(*DeleteAddressRequest)(nil),      // 9: account.v1.DeleteAddressRequest
	(*DeleteAddressResponse)(nil),     // 10: account.v1.DeleteAddressResponse
	(*common.PaginationRequest)(nil),  // 11: common.PaginationRequest
	(*common.PaginationResponse)(nil), // 12: common.PaginationResponse
}
var file_account_v1_address_proto_depIdxs = []int32{
	0,  // 0: account.v1.GetAddressResponse.data:type_name -> account.v1.AddressEntity
	11, // 1: account.v1.ListAddressesRequest.pagination:type_name -> common.PaginationRequest
	0,  // 2: account.v1.ListAddressesResponse.data:type_name -> account.v1.AddressEntity
	12, // 3: account.v1.ListAddressesResponse.pagination:type_name -> common.PaginationResponse
	0,  // 4: account.v1.CreateAddressResponse.data:type_name -> account.v1.AddressEntity
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_account_v1_address_proto_init() }
func file_account_v1_address_proto_init() {
	if File_account_v1_address_proto != nil {
		return
	}
	file_account_v1_address_proto_msgTypes[3].OneofWrappers = []any{}
	file_account_v1_address_proto_msgTypes[7].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_account_v1_address_proto_rawDesc), len(file_account_v1_address_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_account_v1_address_proto_goTypes,
		DependencyIndexes: file_account_v1_address_proto_depIdxs,
		MessageInfos:      file_account_v1_address_proto_msgTypes,
	}.Build()
	File_account_v1_address_proto = out.File
	file_account_v1_address_proto_goTypes = nil
	file_account_v1_address_proto_depIdxs = nil
}
