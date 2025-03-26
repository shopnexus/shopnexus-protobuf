// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: payment/v1/payment.proto

package paymentv1

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

type PaymentMethod int32

const (
	PaymentMethod_PAYMENT_METHOD_UNSPECIFIED PaymentMethod = 0
	PaymentMethod_PAYMENT_METHOD_CASH        PaymentMethod = 1
	PaymentMethod_PAYMENT_METHOD_MOMO        PaymentMethod = 2
	PaymentMethod_PAYMENT_METHOD_VNPAY       PaymentMethod = 3
)

// Enum value maps for PaymentMethod.
var (
	PaymentMethod_name = map[int32]string{
		0: "PAYMENT_METHOD_UNSPECIFIED",
		1: "PAYMENT_METHOD_CASH",
		2: "PAYMENT_METHOD_MOMO",
		3: "PAYMENT_METHOD_VNPAY",
	}
	PaymentMethod_value = map[string]int32{
		"PAYMENT_METHOD_UNSPECIFIED": 0,
		"PAYMENT_METHOD_CASH":        1,
		"PAYMENT_METHOD_MOMO":        2,
		"PAYMENT_METHOD_VNPAY":       3,
	}
)

func (x PaymentMethod) Enum() *PaymentMethod {
	p := new(PaymentMethod)
	*p = x
	return p
}

func (x PaymentMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PaymentMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_payment_v1_payment_proto_enumTypes[0].Descriptor()
}

func (PaymentMethod) Type() protoreflect.EnumType {
	return &file_payment_v1_payment_proto_enumTypes[0]
}

func (x PaymentMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PaymentMethod.Descriptor instead.
func (PaymentMethod) EnumDescriptor() ([]byte, []int) {
	return file_payment_v1_payment_proto_rawDescGZIP(), []int{0}
}

type ProductOnPayment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemQuantity *common.ItemQuantityInt64 `protobuf:"bytes,1,opt,name=item_quantity,json=itemQuantity,proto3" json:"item_quantity,omitempty"`
	SerialIds    []string                  `protobuf:"bytes,2,rep,name=serial_ids,json=serialIds,proto3" json:"serial_ids,omitempty"`
	Price        int64                     `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	TotalPrice   int64                     `protobuf:"varint,4,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
}

func (x *ProductOnPayment) Reset() {
	*x = ProductOnPayment{}
	mi := &file_payment_v1_payment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductOnPayment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductOnPayment) ProtoMessage() {}

func (x *ProductOnPayment) ProtoReflect() protoreflect.Message {
	mi := &file_payment_v1_payment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductOnPayment.ProtoReflect.Descriptor instead.
func (*ProductOnPayment) Descriptor() ([]byte, []int) {
	return file_payment_v1_payment_proto_rawDescGZIP(), []int{0}
}

func (x *ProductOnPayment) GetItemQuantity() *common.ItemQuantityInt64 {
	if x != nil {
		return x.ItemQuantity
	}
	return nil
}

func (x *ProductOnPayment) GetSerialIds() []string {
	if x != nil {
		return x.SerialIds
	}
	return nil
}

func (x *ProductOnPayment) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ProductOnPayment) GetTotalPrice() int64 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

type Payment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      int64               `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Method      PaymentMethod       `protobuf:"varint,3,opt,name=method,proto3,enum=payment.v1.PaymentMethod" json:"method,omitempty"`
	Status      common.Status       `protobuf:"varint,4,opt,name=status,proto3,enum=common.Status" json:"status,omitempty"`
	Address     string              `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Total       int64               `protobuf:"varint,6,opt,name=total,proto3" json:"total,omitempty"`
	DateCreated int64               `protobuf:"varint,7,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty"`
	Products    []*ProductOnPayment `protobuf:"bytes,8,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *Payment) Reset() {
	*x = Payment{}
	mi := &file_payment_v1_payment_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Payment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payment) ProtoMessage() {}

func (x *Payment) ProtoReflect() protoreflect.Message {
	mi := &file_payment_v1_payment_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payment.ProtoReflect.Descriptor instead.
func (*Payment) Descriptor() ([]byte, []int) {
	return file_payment_v1_payment_proto_rawDescGZIP(), []int{1}
}

func (x *Payment) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Payment) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Payment) GetMethod() PaymentMethod {
	if x != nil {
		return x.Method
	}
	return PaymentMethod_PAYMENT_METHOD_UNSPECIFIED
}

func (x *Payment) GetStatus() common.Status {
	if x != nil {
		return x.Status
	}
	return common.Status(0)
}

func (x *Payment) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Payment) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Payment) GetDateCreated() int64 {
	if x != nil {
		return x.DateCreated
	}
	return 0
}

func (x *Payment) GetProducts() []*ProductOnPayment {
	if x != nil {
		return x.Products
	}
	return nil
}

type GetPaymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPaymentRequest) Reset() {
	*x = GetPaymentRequest{}
	mi := &file_payment_v1_payment_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentRequest) ProtoMessage() {}

func (x *GetPaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_v1_payment_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentRequest.ProtoReflect.Descriptor instead.
func (*GetPaymentRequest) Descriptor() ([]byte, []int) {
	return file_payment_v1_payment_proto_rawDescGZIP(), []int{2}
}

func (x *GetPaymentRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetPaymentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Payment `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetPaymentResponse) Reset() {
	*x = GetPaymentResponse{}
	mi := &file_payment_v1_payment_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentResponse) ProtoMessage() {}

func (x *GetPaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_v1_payment_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentResponse.ProtoReflect.Descriptor instead.
func (*GetPaymentResponse) Descriptor() ([]byte, []int) {
	return file_payment_v1_payment_proto_rawDescGZIP(), []int{3}
}

func (x *GetPaymentResponse) GetData() *Payment {
	if x != nil {
		return x.Data
	}
	return nil
}

type ListPaymentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination      *common.PaginationRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	UserId          *int64                    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3,oneof" json:"user_id,omitempty"`
	Method          *PaymentMethod            `protobuf:"varint,3,opt,name=method,proto3,enum=payment.v1.PaymentMethod,oneof" json:"method,omitempty"`
	Status          *common.Status            `protobuf:"varint,4,opt,name=status,proto3,enum=common.Status,oneof" json:"status,omitempty"`
	Address         *string                   `protobuf:"bytes,5,opt,name=address,proto3,oneof" json:"address,omitempty"`
	TotalFrom       *int64                    `protobuf:"varint,6,opt,name=total_from,json=totalFrom,proto3,oneof" json:"total_from,omitempty"`
	TotalTo         *int64                    `protobuf:"varint,7,opt,name=total_to,json=totalTo,proto3,oneof" json:"total_to,omitempty"`
	DateCreatedFrom *int64                    `protobuf:"varint,8,opt,name=date_created_from,json=dateCreatedFrom,proto3,oneof" json:"date_created_from,omitempty"`
	DateCreatedTo   *int64                    `protobuf:"varint,9,opt,name=date_created_to,json=dateCreatedTo,proto3,oneof" json:"date_created_to,omitempty"`
}

func (x *ListPaymentsRequest) Reset() {
	*x = ListPaymentsRequest{}
	mi := &file_payment_v1_payment_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPaymentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPaymentsRequest) ProtoMessage() {}

func (x *ListPaymentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_v1_payment_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPaymentsRequest.ProtoReflect.Descriptor instead.
func (*ListPaymentsRequest) Descriptor() ([]byte, []int) {
	return file_payment_v1_payment_proto_rawDescGZIP(), []int{4}
}

func (x *ListPaymentsRequest) GetPagination() *common.PaginationRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ListPaymentsRequest) GetUserId() int64 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

func (x *ListPaymentsRequest) GetMethod() PaymentMethod {
	if x != nil && x.Method != nil {
		return *x.Method
	}
	return PaymentMethod_PAYMENT_METHOD_UNSPECIFIED
}

func (x *ListPaymentsRequest) GetStatus() common.Status {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return common.Status(0)
}

func (x *ListPaymentsRequest) GetAddress() string {
	if x != nil && x.Address != nil {
		return *x.Address
	}
	return ""
}

func (x *ListPaymentsRequest) GetTotalFrom() int64 {
	if x != nil && x.TotalFrom != nil {
		return *x.TotalFrom
	}
	return 0
}

func (x *ListPaymentsRequest) GetTotalTo() int64 {
	if x != nil && x.TotalTo != nil {
		return *x.TotalTo
	}
	return 0
}

func (x *ListPaymentsRequest) GetDateCreatedFrom() int64 {
	if x != nil && x.DateCreatedFrom != nil {
		return *x.DateCreatedFrom
	}
	return 0
}

func (x *ListPaymentsRequest) GetDateCreatedTo() int64 {
	if x != nil && x.DateCreatedTo != nil {
		return *x.DateCreatedTo
	}
	return 0
}

type ListPaymentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data       []*Payment                 `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Pagination *common.PaginationResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListPaymentsResponse) Reset() {
	*x = ListPaymentsResponse{}
	mi := &file_payment_v1_payment_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPaymentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPaymentsResponse) ProtoMessage() {}

func (x *ListPaymentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_v1_payment_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPaymentsResponse.ProtoReflect.Descriptor instead.
func (*ListPaymentsResponse) Descriptor() ([]byte, []int) {
	return file_payment_v1_payment_proto_rawDescGZIP(), []int{5}
}

func (x *ListPaymentsResponse) GetData() []*Payment {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ListPaymentsResponse) GetPagination() *common.PaginationResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type CreatePaymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId int64         `protobuf:"varint,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Method    PaymentMethod `protobuf:"varint,2,opt,name=method,proto3,enum=payment.v1.PaymentMethod" json:"method,omitempty"`
	Address   string        `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *CreatePaymentRequest) Reset() {
	*x = CreatePaymentRequest{}
	mi := &file_payment_v1_payment_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePaymentRequest) ProtoMessage() {}

func (x *CreatePaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_v1_payment_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePaymentRequest.ProtoReflect.Descriptor instead.
func (*CreatePaymentRequest) Descriptor() ([]byte, []int) {
	return file_payment_v1_payment_proto_rawDescGZIP(), []int{6}
}

func (x *CreatePaymentRequest) GetRequestId() int64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *CreatePaymentRequest) GetMethod() PaymentMethod {
	if x != nil {
		return x.Method
	}
	return PaymentMethod_PAYMENT_METHOD_UNSPECIFIED
}

func (x *CreatePaymentRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type CreatePaymentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId int64    `protobuf:"varint,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Url       string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Payment   *Payment `protobuf:"bytes,3,opt,name=payment,proto3" json:"payment,omitempty"`
}

func (x *CreatePaymentResponse) Reset() {
	*x = CreatePaymentResponse{}
	mi := &file_payment_v1_payment_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePaymentResponse) ProtoMessage() {}

func (x *CreatePaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_v1_payment_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePaymentResponse.ProtoReflect.Descriptor instead.
func (*CreatePaymentResponse) Descriptor() ([]byte, []int) {
	return file_payment_v1_payment_proto_rawDescGZIP(), []int{7}
}

func (x *CreatePaymentResponse) GetRequestId() int64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *CreatePaymentResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *CreatePaymentResponse) GetPayment() *Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

type UpdatePaymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Method  *PaymentMethod `protobuf:"varint,2,opt,name=method,proto3,enum=payment.v1.PaymentMethod,oneof" json:"method,omitempty"`
	Address *string        `protobuf:"bytes,3,opt,name=address,proto3,oneof" json:"address,omitempty"`
}

func (x *UpdatePaymentRequest) Reset() {
	*x = UpdatePaymentRequest{}
	mi := &file_payment_v1_payment_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatePaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePaymentRequest) ProtoMessage() {}

func (x *UpdatePaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_v1_payment_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePaymentRequest.ProtoReflect.Descriptor instead.
func (*UpdatePaymentRequest) Descriptor() ([]byte, []int) {
	return file_payment_v1_payment_proto_rawDescGZIP(), []int{8}
}

func (x *UpdatePaymentRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdatePaymentRequest) GetMethod() PaymentMethod {
	if x != nil && x.Method != nil {
		return *x.Method
	}
	return PaymentMethod_PAYMENT_METHOD_UNSPECIFIED
}

func (x *UpdatePaymentRequest) GetAddress() string {
	if x != nil && x.Address != nil {
		return *x.Address
	}
	return ""
}

type UpdatePaymentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdatePaymentResponse) Reset() {
	*x = UpdatePaymentResponse{}
	mi := &file_payment_v1_payment_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatePaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePaymentResponse) ProtoMessage() {}

func (x *UpdatePaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_v1_payment_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePaymentResponse.ProtoReflect.Descriptor instead.
func (*UpdatePaymentResponse) Descriptor() ([]byte, []int) {
	return file_payment_v1_payment_proto_rawDescGZIP(), []int{9}
}

type CancelPaymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CancelPaymentRequest) Reset() {
	*x = CancelPaymentRequest{}
	mi := &file_payment_v1_payment_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CancelPaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelPaymentRequest) ProtoMessage() {}

func (x *CancelPaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_v1_payment_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelPaymentRequest.ProtoReflect.Descriptor instead.
func (*CancelPaymentRequest) Descriptor() ([]byte, []int) {
	return file_payment_v1_payment_proto_rawDescGZIP(), []int{10}
}

func (x *CancelPaymentRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CancelPaymentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CancelPaymentResponse) Reset() {
	*x = CancelPaymentResponse{}
	mi := &file_payment_v1_payment_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CancelPaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelPaymentResponse) ProtoMessage() {}

func (x *CancelPaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_v1_payment_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelPaymentResponse.ProtoReflect.Descriptor instead.
func (*CancelPaymentResponse) Descriptor() ([]byte, []int) {
	return file_payment_v1_payment_proto_rawDescGZIP(), []int{11}
}

var File_payment_v1_payment_proto protoreflect.FileDescriptor

var file_payment_v1_payment_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x69,
	0x74, 0x65, 0x6d, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa8, 0x01, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x6e, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3e, 0x0a, 0x0d, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x52, 0x0c, 0x69, 0x74, 0x65, 0x6d, 0x51, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f,
	0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x49, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x9a, 0x02, 0x0a, 0x07,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x31, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x19, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x38,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4f, 0x6e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3d, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x88, 0x04, 0x0a,
	0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1c, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x48, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x36, 0x0a,
	0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x48, 0x01, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x02, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88,
	0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x88, 0x01,
	0x01, 0x12, 0x22, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x48, 0x04, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x72,
	0x6f, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x74,
	0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x48, 0x05, 0x52, 0x07, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x54, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x11, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x48, 0x06, 0x52, 0x0f, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x46,
	0x72, 0x6f, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x48,
	0x07, 0x52, 0x0d, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x6f,
	0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x42,
	0x09, 0x0a, 0x07, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x66, 0x72, 0x6f, 0x6d,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x74, 0x6f, 0x42, 0x14, 0x0a,
	0x12, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x66,
	0x72, 0x6f, 0x6d, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x22, 0x7b, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x27, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3a, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x82, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x06,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x77, 0x0a, 0x15, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x2d, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x22, 0x94, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x36, 0x0a, 0x06, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x48, 0x00, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x88,
	0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x0a, 0x0a,
	0x08, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x17, 0x0a, 0x15, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x26, 0x0a, 0x14, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2a, 0x7b, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x43, 0x41, 0x53, 0x48, 0x10, 0x01, 0x12, 0x17, 0x0a,
	0x13, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f,
	0x4d, 0x4f, 0x4d, 0x4f, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e,
	0x54, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x56, 0x4e, 0x50, 0x41, 0x59, 0x10, 0x03,
	0x42, 0xaf, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x42, 0x0c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x68, 0x6f, 0x70, 0x6e, 0x65, 0x78, 0x75, 0x73, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x6e, 0x65,
	0x78, 0x75, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76,
	0x31, 0x3b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x58,
	0x58, 0xaa, 0x02, 0x0a, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x0a, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payment_v1_payment_proto_rawDescOnce sync.Once
	file_payment_v1_payment_proto_rawDescData = file_payment_v1_payment_proto_rawDesc
)

func file_payment_v1_payment_proto_rawDescGZIP() []byte {
	file_payment_v1_payment_proto_rawDescOnce.Do(func() {
		file_payment_v1_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_payment_v1_payment_proto_rawDescData)
	})
	return file_payment_v1_payment_proto_rawDescData
}

var file_payment_v1_payment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_payment_v1_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_payment_v1_payment_proto_goTypes = []any{
	(PaymentMethod)(0),                // 0: payment.v1.PaymentMethod
	(*ProductOnPayment)(nil),          // 1: payment.v1.ProductOnPayment
	(*Payment)(nil),                   // 2: payment.v1.Payment
	(*GetPaymentRequest)(nil),         // 3: payment.v1.GetPaymentRequest
	(*GetPaymentResponse)(nil),        // 4: payment.v1.GetPaymentResponse
	(*ListPaymentsRequest)(nil),       // 5: payment.v1.ListPaymentsRequest
	(*ListPaymentsResponse)(nil),      // 6: payment.v1.ListPaymentsResponse
	(*CreatePaymentRequest)(nil),      // 7: payment.v1.CreatePaymentRequest
	(*CreatePaymentResponse)(nil),     // 8: payment.v1.CreatePaymentResponse
	(*UpdatePaymentRequest)(nil),      // 9: payment.v1.UpdatePaymentRequest
	(*UpdatePaymentResponse)(nil),     // 10: payment.v1.UpdatePaymentResponse
	(*CancelPaymentRequest)(nil),      // 11: payment.v1.CancelPaymentRequest
	(*CancelPaymentResponse)(nil),     // 12: payment.v1.CancelPaymentResponse
	(*common.ItemQuantityInt64)(nil),  // 13: common.ItemQuantityInt64
	(common.Status)(0),                // 14: common.Status
	(*common.PaginationRequest)(nil),  // 15: common.PaginationRequest
	(*common.PaginationResponse)(nil), // 16: common.PaginationResponse
}
var file_payment_v1_payment_proto_depIdxs = []int32{
	13, // 0: payment.v1.ProductOnPayment.item_quantity:type_name -> common.ItemQuantityInt64
	0,  // 1: payment.v1.Payment.method:type_name -> payment.v1.PaymentMethod
	14, // 2: payment.v1.Payment.status:type_name -> common.Status
	1,  // 3: payment.v1.Payment.products:type_name -> payment.v1.ProductOnPayment
	2,  // 4: payment.v1.GetPaymentResponse.data:type_name -> payment.v1.Payment
	15, // 5: payment.v1.ListPaymentsRequest.pagination:type_name -> common.PaginationRequest
	0,  // 6: payment.v1.ListPaymentsRequest.method:type_name -> payment.v1.PaymentMethod
	14, // 7: payment.v1.ListPaymentsRequest.status:type_name -> common.Status
	2,  // 8: payment.v1.ListPaymentsResponse.data:type_name -> payment.v1.Payment
	16, // 9: payment.v1.ListPaymentsResponse.pagination:type_name -> common.PaginationResponse
	0,  // 10: payment.v1.CreatePaymentRequest.method:type_name -> payment.v1.PaymentMethod
	2,  // 11: payment.v1.CreatePaymentResponse.payment:type_name -> payment.v1.Payment
	0,  // 12: payment.v1.UpdatePaymentRequest.method:type_name -> payment.v1.PaymentMethod
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_payment_v1_payment_proto_init() }
func file_payment_v1_payment_proto_init() {
	if File_payment_v1_payment_proto != nil {
		return
	}
	file_payment_v1_payment_proto_msgTypes[4].OneofWrappers = []any{}
	file_payment_v1_payment_proto_msgTypes[8].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_payment_v1_payment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_payment_v1_payment_proto_goTypes,
		DependencyIndexes: file_payment_v1_payment_proto_depIdxs,
		EnumInfos:         file_payment_v1_payment_proto_enumTypes,
		MessageInfos:      file_payment_v1_payment_proto_msgTypes,
	}.Build()
	File_payment_v1_payment_proto = out.File
	file_payment_v1_payment_proto_rawDesc = nil
	file_payment_v1_payment_proto_goTypes = nil
	file_payment_v1_payment_proto_depIdxs = nil
}
