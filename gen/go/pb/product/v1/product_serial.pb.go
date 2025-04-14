// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: product/v1/product_serial.proto

package productv1

import (
	common "github.com/shopnexus/shopnexus-protobuf-gen-go/pb/common"
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

type ProductSerialEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SerialId    string `protobuf:"bytes,1,opt,name=serial_id,json=serialId,proto3" json:"serial_id,omitempty"`
	ProductId   int64  `protobuf:"varint,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	IsSold      bool   `protobuf:"varint,3,opt,name=is_sold,json=isSold,proto3" json:"is_sold,omitempty"`
	IsActive    bool   `protobuf:"varint,4,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	DateCreated int64  `protobuf:"varint,5,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty"`
	DateUpdated int64  `protobuf:"varint,6,opt,name=date_updated,json=dateUpdated,proto3" json:"date_updated,omitempty"`
}

func (x *ProductSerialEntity) Reset() {
	*x = ProductSerialEntity{}
	mi := &file_product_v1_product_serial_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductSerialEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductSerialEntity) ProtoMessage() {}

func (x *ProductSerialEntity) ProtoReflect() protoreflect.Message {
	mi := &file_product_v1_product_serial_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductSerialEntity.ProtoReflect.Descriptor instead.
func (*ProductSerialEntity) Descriptor() ([]byte, []int) {
	return file_product_v1_product_serial_proto_rawDescGZIP(), []int{0}
}

func (x *ProductSerialEntity) GetSerialId() string {
	if x != nil {
		return x.SerialId
	}
	return ""
}

func (x *ProductSerialEntity) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *ProductSerialEntity) GetIsSold() bool {
	if x != nil {
		return x.IsSold
	}
	return false
}

func (x *ProductSerialEntity) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *ProductSerialEntity) GetDateCreated() int64 {
	if x != nil {
		return x.DateCreated
	}
	return 0
}

func (x *ProductSerialEntity) GetDateUpdated() int64 {
	if x != nil {
		return x.DateUpdated
	}
	return 0
}

type GetProductSerialRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SerialId string `protobuf:"bytes,1,opt,name=serial_id,json=serialId,proto3" json:"serial_id,omitempty"`
}

func (x *GetProductSerialRequest) Reset() {
	*x = GetProductSerialRequest{}
	mi := &file_product_v1_product_serial_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductSerialRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductSerialRequest) ProtoMessage() {}

func (x *GetProductSerialRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_v1_product_serial_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductSerialRequest.ProtoReflect.Descriptor instead.
func (*GetProductSerialRequest) Descriptor() ([]byte, []int) {
	return file_product_v1_product_serial_proto_rawDescGZIP(), []int{1}
}

func (x *GetProductSerialRequest) GetSerialId() string {
	if x != nil {
		return x.SerialId
	}
	return ""
}

type GetProductSerialResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *ProductSerialEntity `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetProductSerialResponse) Reset() {
	*x = GetProductSerialResponse{}
	mi := &file_product_v1_product_serial_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductSerialResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductSerialResponse) ProtoMessage() {}

func (x *GetProductSerialResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_v1_product_serial_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductSerialResponse.ProtoReflect.Descriptor instead.
func (*GetProductSerialResponse) Descriptor() ([]byte, []int) {
	return file_product_v1_product_serial_proto_rawDescGZIP(), []int{2}
}

func (x *GetProductSerialResponse) GetData() *ProductSerialEntity {
	if x != nil {
		return x.Data
	}
	return nil
}

type ListProductSerialsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination      *common.PaginationRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	SerialId        *string                   `protobuf:"bytes,2,opt,name=serial_id,json=serialId,proto3,oneof" json:"serial_id,omitempty"`
	ProductId       *int64                    `protobuf:"varint,3,opt,name=product_id,json=productId,proto3,oneof" json:"product_id,omitempty"`
	IsSold          *bool                     `protobuf:"varint,4,opt,name=is_sold,json=isSold,proto3,oneof" json:"is_sold,omitempty"`
	IsActive        *bool                     `protobuf:"varint,5,opt,name=is_active,json=isActive,proto3,oneof" json:"is_active,omitempty"`
	DateCreatedFrom *int64                    `protobuf:"varint,6,opt,name=date_created_from,json=dateCreatedFrom,proto3,oneof" json:"date_created_from,omitempty"`
	DateCreatedTo   *int64                    `protobuf:"varint,7,opt,name=date_created_to,json=dateCreatedTo,proto3,oneof" json:"date_created_to,omitempty"`
}

func (x *ListProductSerialsRequest) Reset() {
	*x = ListProductSerialsRequest{}
	mi := &file_product_v1_product_serial_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListProductSerialsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductSerialsRequest) ProtoMessage() {}

func (x *ListProductSerialsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_v1_product_serial_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductSerialsRequest.ProtoReflect.Descriptor instead.
func (*ListProductSerialsRequest) Descriptor() ([]byte, []int) {
	return file_product_v1_product_serial_proto_rawDescGZIP(), []int{3}
}

func (x *ListProductSerialsRequest) GetPagination() *common.PaginationRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ListProductSerialsRequest) GetSerialId() string {
	if x != nil && x.SerialId != nil {
		return *x.SerialId
	}
	return ""
}

func (x *ListProductSerialsRequest) GetProductId() int64 {
	if x != nil && x.ProductId != nil {
		return *x.ProductId
	}
	return 0
}

func (x *ListProductSerialsRequest) GetIsSold() bool {
	if x != nil && x.IsSold != nil {
		return *x.IsSold
	}
	return false
}

func (x *ListProductSerialsRequest) GetIsActive() bool {
	if x != nil && x.IsActive != nil {
		return *x.IsActive
	}
	return false
}

func (x *ListProductSerialsRequest) GetDateCreatedFrom() int64 {
	if x != nil && x.DateCreatedFrom != nil {
		return *x.DateCreatedFrom
	}
	return 0
}

func (x *ListProductSerialsRequest) GetDateCreatedTo() int64 {
	if x != nil && x.DateCreatedTo != nil {
		return *x.DateCreatedTo
	}
	return 0
}

type ListProductSerialsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data       []*ProductSerialEntity     `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Pagination *common.PaginationResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListProductSerialsResponse) Reset() {
	*x = ListProductSerialsResponse{}
	mi := &file_product_v1_product_serial_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListProductSerialsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductSerialsResponse) ProtoMessage() {}

func (x *ListProductSerialsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_v1_product_serial_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductSerialsResponse.ProtoReflect.Descriptor instead.
func (*ListProductSerialsResponse) Descriptor() ([]byte, []int) {
	return file_product_v1_product_serial_proto_rawDescGZIP(), []int{4}
}

func (x *ListProductSerialsResponse) GetData() []*ProductSerialEntity {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ListProductSerialsResponse) GetPagination() *common.PaginationResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type CreateProductSerialRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SerialId  string `protobuf:"bytes,1,opt,name=serial_id,json=serialId,proto3" json:"serial_id,omitempty"`
	ProductId int64  `protobuf:"varint,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	IsSold    bool   `protobuf:"varint,3,opt,name=is_sold,json=isSold,proto3" json:"is_sold,omitempty"`
	IsActive  bool   `protobuf:"varint,4,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
}

func (x *CreateProductSerialRequest) Reset() {
	*x = CreateProductSerialRequest{}
	mi := &file_product_v1_product_serial_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateProductSerialRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductSerialRequest) ProtoMessage() {}

func (x *CreateProductSerialRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_v1_product_serial_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductSerialRequest.ProtoReflect.Descriptor instead.
func (*CreateProductSerialRequest) Descriptor() ([]byte, []int) {
	return file_product_v1_product_serial_proto_rawDescGZIP(), []int{5}
}

func (x *CreateProductSerialRequest) GetSerialId() string {
	if x != nil {
		return x.SerialId
	}
	return ""
}

func (x *CreateProductSerialRequest) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *CreateProductSerialRequest) GetIsSold() bool {
	if x != nil {
		return x.IsSold
	}
	return false
}

func (x *CreateProductSerialRequest) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

type CreateProductSerialResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *ProductSerialEntity `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateProductSerialResponse) Reset() {
	*x = CreateProductSerialResponse{}
	mi := &file_product_v1_product_serial_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateProductSerialResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductSerialResponse) ProtoMessage() {}

func (x *CreateProductSerialResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_v1_product_serial_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductSerialResponse.ProtoReflect.Descriptor instead.
func (*CreateProductSerialResponse) Descriptor() ([]byte, []int) {
	return file_product_v1_product_serial_proto_rawDescGZIP(), []int{6}
}

func (x *CreateProductSerialResponse) GetData() *ProductSerialEntity {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateProductSerialRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SerialId string `protobuf:"bytes,1,opt,name=serial_id,json=serialId,proto3" json:"serial_id,omitempty"`
	IsSold   *bool  `protobuf:"varint,2,opt,name=is_sold,json=isSold,proto3,oneof" json:"is_sold,omitempty"`
	IsActive *bool  `protobuf:"varint,3,opt,name=is_active,json=isActive,proto3,oneof" json:"is_active,omitempty"`
}

func (x *UpdateProductSerialRequest) Reset() {
	*x = UpdateProductSerialRequest{}
	mi := &file_product_v1_product_serial_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateProductSerialRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductSerialRequest) ProtoMessage() {}

func (x *UpdateProductSerialRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_v1_product_serial_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductSerialRequest.ProtoReflect.Descriptor instead.
func (*UpdateProductSerialRequest) Descriptor() ([]byte, []int) {
	return file_product_v1_product_serial_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateProductSerialRequest) GetSerialId() string {
	if x != nil {
		return x.SerialId
	}
	return ""
}

func (x *UpdateProductSerialRequest) GetIsSold() bool {
	if x != nil && x.IsSold != nil {
		return *x.IsSold
	}
	return false
}

func (x *UpdateProductSerialRequest) GetIsActive() bool {
	if x != nil && x.IsActive != nil {
		return *x.IsActive
	}
	return false
}

type UpdateProductSerialResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateProductSerialResponse) Reset() {
	*x = UpdateProductSerialResponse{}
	mi := &file_product_v1_product_serial_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateProductSerialResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductSerialResponse) ProtoMessage() {}

func (x *UpdateProductSerialResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_v1_product_serial_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductSerialResponse.ProtoReflect.Descriptor instead.
func (*UpdateProductSerialResponse) Descriptor() ([]byte, []int) {
	return file_product_v1_product_serial_proto_rawDescGZIP(), []int{8}
}

type DeleteProductSerialRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SerialId string `protobuf:"bytes,1,opt,name=serial_id,json=serialId,proto3" json:"serial_id,omitempty"`
}

func (x *DeleteProductSerialRequest) Reset() {
	*x = DeleteProductSerialRequest{}
	mi := &file_product_v1_product_serial_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteProductSerialRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductSerialRequest) ProtoMessage() {}

func (x *DeleteProductSerialRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_v1_product_serial_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductSerialRequest.ProtoReflect.Descriptor instead.
func (*DeleteProductSerialRequest) Descriptor() ([]byte, []int) {
	return file_product_v1_product_serial_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteProductSerialRequest) GetSerialId() string {
	if x != nil {
		return x.SerialId
	}
	return ""
}

type DeleteProductSerialResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteProductSerialResponse) Reset() {
	*x = DeleteProductSerialResponse{}
	mi := &file_product_v1_product_serial_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteProductSerialResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductSerialResponse) ProtoMessage() {}

func (x *DeleteProductSerialResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_v1_product_serial_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductSerialResponse.ProtoReflect.Descriptor instead.
func (*DeleteProductSerialResponse) Descriptor() ([]byte, []int) {
	return file_product_v1_product_serial_proto_rawDescGZIP(), []int{10}
}

var File_product_v1_product_serial_proto protoreflect.FileDescriptor

var file_product_v1_product_serial_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x01, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73,
	0x5f, 0x73, 0x6f, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x53,
	0x6f, 0x6c, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x36, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x22, 0x4f,
	0x0a, 0x18, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x9b, 0x03, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0a, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x73,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1c,
	0x0a, 0x07, 0x69, 0x73, 0x5f, 0x73, 0x6f, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48,
	0x02, 0x52, 0x06, 0x69, 0x73, 0x53, 0x6f, 0x6c, 0x64, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09,
	0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x48,
	0x03, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2f,
	0x0a, 0x11, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x66,
	0x72, 0x6f, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x48, 0x04, 0x52, 0x0f, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x88, 0x01, 0x01, 0x12,
	0x2b, 0x0a, 0x0f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x74, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x48, 0x05, 0x52, 0x0d, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x6f, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a,
	0x5f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x69, 0x73,
	0x5f, 0x73, 0x6f, 0x6c, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x22, 0x8d, 0x01,
	0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x3a, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x8e, 0x01,
	0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x73,
	0x6f, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x53, 0x6f, 0x6c,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x52,
	0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x93, 0x01, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x07, 0x69, 0x73, 0x5f, 0x73, 0x6f, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48,
	0x00, 0x52, 0x06, 0x69, 0x73, 0x53, 0x6f, 0x6c, 0x64, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09,
	0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48,
	0x01, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0a,
	0x0a, 0x08, 0x5f, 0x69, 0x73, 0x5f, 0x73, 0x6f, 0x6c, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x69,
	0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x1d, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x49, 0x64, 0x22, 0x1d, 0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0xb5, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x76, 0x31, 0x42, 0x12, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x6e, 0x65, 0x78, 0x75, 0x73,
	0x2f, 0x73, 0x68, 0x6f, 0x70, 0x6e, 0x65, 0x78, 0x75, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x16, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_product_v1_product_serial_proto_rawDescOnce sync.Once
	file_product_v1_product_serial_proto_rawDescData = file_product_v1_product_serial_proto_rawDesc
)

func file_product_v1_product_serial_proto_rawDescGZIP() []byte {
	file_product_v1_product_serial_proto_rawDescOnce.Do(func() {
		file_product_v1_product_serial_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_v1_product_serial_proto_rawDescData)
	})
	return file_product_v1_product_serial_proto_rawDescData
}

var file_product_v1_product_serial_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_product_v1_product_serial_proto_goTypes = []any{
	(*ProductSerialEntity)(nil),         // 0: product.v1.ProductSerialEntity
	(*GetProductSerialRequest)(nil),     // 1: product.v1.GetProductSerialRequest
	(*GetProductSerialResponse)(nil),    // 2: product.v1.GetProductSerialResponse
	(*ListProductSerialsRequest)(nil),   // 3: product.v1.ListProductSerialsRequest
	(*ListProductSerialsResponse)(nil),  // 4: product.v1.ListProductSerialsResponse
	(*CreateProductSerialRequest)(nil),  // 5: product.v1.CreateProductSerialRequest
	(*CreateProductSerialResponse)(nil), // 6: product.v1.CreateProductSerialResponse
	(*UpdateProductSerialRequest)(nil),  // 7: product.v1.UpdateProductSerialRequest
	(*UpdateProductSerialResponse)(nil), // 8: product.v1.UpdateProductSerialResponse
	(*DeleteProductSerialRequest)(nil),  // 9: product.v1.DeleteProductSerialRequest
	(*DeleteProductSerialResponse)(nil), // 10: product.v1.DeleteProductSerialResponse
	(*common.PaginationRequest)(nil),    // 11: common.PaginationRequest
	(*common.PaginationResponse)(nil),   // 12: common.PaginationResponse
}
var file_product_v1_product_serial_proto_depIdxs = []int32{
	0,  // 0: product.v1.GetProductSerialResponse.data:type_name -> product.v1.ProductSerialEntity
	11, // 1: product.v1.ListProductSerialsRequest.pagination:type_name -> common.PaginationRequest
	0,  // 2: product.v1.ListProductSerialsResponse.data:type_name -> product.v1.ProductSerialEntity
	12, // 3: product.v1.ListProductSerialsResponse.pagination:type_name -> common.PaginationResponse
	0,  // 4: product.v1.CreateProductSerialResponse.data:type_name -> product.v1.ProductSerialEntity
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_product_v1_product_serial_proto_init() }
func file_product_v1_product_serial_proto_init() {
	if File_product_v1_product_serial_proto != nil {
		return
	}
	file_product_v1_product_serial_proto_msgTypes[3].OneofWrappers = []any{}
	file_product_v1_product_serial_proto_msgTypes[7].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_product_v1_product_serial_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_product_v1_product_serial_proto_goTypes,
		DependencyIndexes: file_product_v1_product_serial_proto_depIdxs,
		MessageInfos:      file_product_v1_product_serial_proto_msgTypes,
	}.Build()
	File_product_v1_product_serial_proto = out.File
	file_product_v1_product_serial_proto_rawDesc = nil
	file_product_v1_product_serial_proto_goTypes = nil
	file_product_v1_product_serial_proto_depIdxs = nil
}
