// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: payment/v1/payment.proto

package paymentv1

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
	state         protoimpl.MessageState    `protogen:"open.v1"`
	Id            int64                     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ItemQuantity  *common.ItemQuantityInt64 `protobuf:"bytes,2,opt,name=item_quantity,json=itemQuantity,proto3" json:"item_quantity,omitempty"`
	SerialIds     []string                  `protobuf:"bytes,3,rep,name=serial_ids,json=serialIds,proto3" json:"serial_ids,omitempty"`
	Price         int64                     `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	TotalPrice    int64                     `protobuf:"varint,5,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

func (x *ProductOnPayment) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
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

type PaymentEntity struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        int64                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Method        PaymentMethod          `protobuf:"varint,3,opt,name=method,proto3,enum=payment.v1.PaymentMethod" json:"method,omitempty"`
	Status        common.Status          `protobuf:"varint,4,opt,name=status,proto3,enum=common.Status" json:"status,omitempty"`
	Address       string                 `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Total         int64                  `protobuf:"varint,6,opt,name=total,proto3" json:"total,omitempty"`
	DateCreated   int64                  `protobuf:"varint,7,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty"`
	Products      []*ProductOnPayment    `protobuf:"bytes,8,rep,name=products,proto3" json:"products,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PaymentEntity) Reset() {
	*x = PaymentEntity{}
	mi := &file_payment_v1_payment_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaymentEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentEntity) ProtoMessage() {}

func (x *PaymentEntity) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PaymentEntity.ProtoReflect.Descriptor instead.
func (*PaymentEntity) Descriptor() ([]byte, []int) {
	return file_payment_v1_payment_proto_rawDescGZIP(), []int{1}
}

func (x *PaymentEntity) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PaymentEntity) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PaymentEntity) GetMethod() PaymentMethod {
	if x != nil {
		return x.Method
	}
	return PaymentMethod_PAYMENT_METHOD_UNSPECIFIED
}

func (x *PaymentEntity) GetStatus() common.Status {
	if x != nil {
		return x.Status
	}
	return common.Status(0)
}

func (x *PaymentEntity) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *PaymentEntity) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PaymentEntity) GetDateCreated() int64 {
	if x != nil {
		return x.DateCreated
	}
	return 0
}

func (x *PaymentEntity) GetProducts() []*ProductOnPayment {
	if x != nil {
		return x.Products
	}
	return nil
}

type GetPaymentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          *PaymentEntity         `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

func (x *GetPaymentResponse) GetData() *PaymentEntity {
	if x != nil {
		return x.Data
	}
	return nil
}

type ListPaymentsRequest struct {
	state           protoimpl.MessageState    `protogen:"open.v1"`
	Pagination      *common.PaginationRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	UserId          *int64                    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3,oneof" json:"user_id,omitempty"`
	Method          *PaymentMethod            `protobuf:"varint,3,opt,name=method,proto3,enum=payment.v1.PaymentMethod,oneof" json:"method,omitempty"`
	Status          *common.Status            `protobuf:"varint,4,opt,name=status,proto3,enum=common.Status,oneof" json:"status,omitempty"`
	Address         *string                   `protobuf:"bytes,5,opt,name=address,proto3,oneof" json:"address,omitempty"`
	TotalFrom       *int64                    `protobuf:"varint,6,opt,name=total_from,json=totalFrom,proto3,oneof" json:"total_from,omitempty"`
	TotalTo         *int64                    `protobuf:"varint,7,opt,name=total_to,json=totalTo,proto3,oneof" json:"total_to,omitempty"`
	DateCreatedFrom *int64                    `protobuf:"varint,8,opt,name=date_created_from,json=dateCreatedFrom,proto3,oneof" json:"date_created_from,omitempty"`
	DateCreatedTo   *int64                    `protobuf:"varint,9,opt,name=date_created_to,json=dateCreatedTo,proto3,oneof" json:"date_created_to,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
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
	state         protoimpl.MessageState     `protogen:"open.v1"`
	Data          []*PaymentEntity           `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Pagination    *common.PaginationResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

func (x *ListPaymentsResponse) GetData() []*PaymentEntity {
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	RequestId     int64                  `protobuf:"varint,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Method        PaymentMethod          `protobuf:"varint,2,opt,name=method,proto3,enum=payment.v1.PaymentMethod" json:"method,omitempty"`
	Address       string                 `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	ProductIds    []int64                `protobuf:"varint,4,rep,packed,name=product_ids,json=productIds,proto3" json:"product_ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

func (x *CreatePaymentRequest) GetProductIds() []int64 {
	if x != nil {
		return x.ProductIds
	}
	return nil
}

type CreatePaymentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RequestId     int64                  `protobuf:"varint,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Url           string                 `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Payment       *PaymentEntity         `protobuf:"bytes,3,opt,name=payment,proto3" json:"payment,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

func (x *CreatePaymentResponse) GetPayment() *PaymentEntity {
	if x != nil {
		return x.Payment
	}
	return nil
}

type UpdatePaymentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Method        *PaymentMethod         `protobuf:"varint,2,opt,name=method,proto3,enum=payment.v1.PaymentMethod,oneof" json:"method,omitempty"`
	Address       *string                `protobuf:"bytes,3,opt,name=address,proto3,oneof" json:"address,omitempty"`
	Status        *common.Status         `protobuf:"varint,4,opt,name=status,proto3,enum=common.Status,oneof" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

func (x *UpdatePaymentRequest) GetStatus() common.Status {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return common.Status(0)
}

type UpdatePaymentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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

const file_payment_v1_payment_proto_rawDesc = "" +
	"\n" +
	"\x18payment/v1/payment.proto\x12\n" +
	"payment.v1\x1a\x1acommon/item_quantity.proto\x1a\x17common/pagination.proto\x1a\x13common/status.proto\"\xb8\x01\n" +
	"\x10ProductOnPayment\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12>\n" +
	"\ritem_quantity\x18\x02 \x01(\v2\x19.common.ItemQuantityInt64R\fitemQuantity\x12\x1d\n" +
	"\n" +
	"serial_ids\x18\x03 \x03(\tR\tserialIds\x12\x14\n" +
	"\x05price\x18\x04 \x01(\x03R\x05price\x12\x1f\n" +
	"\vtotal_price\x18\x05 \x01(\x03R\n" +
	"totalPrice\"\xa0\x02\n" +
	"\rPaymentEntity\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\x03R\x06userId\x121\n" +
	"\x06method\x18\x03 \x01(\x0e2\x19.payment.v1.PaymentMethodR\x06method\x12&\n" +
	"\x06status\x18\x04 \x01(\x0e2\x0e.common.StatusR\x06status\x12\x18\n" +
	"\aaddress\x18\x05 \x01(\tR\aaddress\x12\x14\n" +
	"\x05total\x18\x06 \x01(\x03R\x05total\x12!\n" +
	"\fdate_created\x18\a \x01(\x03R\vdateCreated\x128\n" +
	"\bproducts\x18\b \x03(\v2\x1c.payment.v1.ProductOnPaymentR\bproducts\"#\n" +
	"\x11GetPaymentRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\"C\n" +
	"\x12GetPaymentResponse\x12-\n" +
	"\x04data\x18\x01 \x01(\v2\x19.payment.v1.PaymentEntityR\x04data\"\x88\x04\n" +
	"\x13ListPaymentsRequest\x129\n" +
	"\n" +
	"pagination\x18\x01 \x01(\v2\x19.common.PaginationRequestR\n" +
	"pagination\x12\x1c\n" +
	"\auser_id\x18\x02 \x01(\x03H\x00R\x06userId\x88\x01\x01\x126\n" +
	"\x06method\x18\x03 \x01(\x0e2\x19.payment.v1.PaymentMethodH\x01R\x06method\x88\x01\x01\x12+\n" +
	"\x06status\x18\x04 \x01(\x0e2\x0e.common.StatusH\x02R\x06status\x88\x01\x01\x12\x1d\n" +
	"\aaddress\x18\x05 \x01(\tH\x03R\aaddress\x88\x01\x01\x12\"\n" +
	"\n" +
	"total_from\x18\x06 \x01(\x03H\x04R\ttotalFrom\x88\x01\x01\x12\x1e\n" +
	"\btotal_to\x18\a \x01(\x03H\x05R\atotalTo\x88\x01\x01\x12/\n" +
	"\x11date_created_from\x18\b \x01(\x03H\x06R\x0fdateCreatedFrom\x88\x01\x01\x12+\n" +
	"\x0fdate_created_to\x18\t \x01(\x03H\aR\rdateCreatedTo\x88\x01\x01B\n" +
	"\n" +
	"\b_user_idB\t\n" +
	"\a_methodB\t\n" +
	"\a_statusB\n" +
	"\n" +
	"\b_addressB\r\n" +
	"\v_total_fromB\v\n" +
	"\t_total_toB\x14\n" +
	"\x12_date_created_fromB\x12\n" +
	"\x10_date_created_to\"\x81\x01\n" +
	"\x14ListPaymentsResponse\x12-\n" +
	"\x04data\x18\x01 \x03(\v2\x19.payment.v1.PaymentEntityR\x04data\x12:\n" +
	"\n" +
	"pagination\x18\x02 \x01(\v2\x1a.common.PaginationResponseR\n" +
	"pagination\"\xa3\x01\n" +
	"\x14CreatePaymentRequest\x12\x1d\n" +
	"\n" +
	"request_id\x18\x01 \x01(\x03R\trequestId\x121\n" +
	"\x06method\x18\x02 \x01(\x0e2\x19.payment.v1.PaymentMethodR\x06method\x12\x18\n" +
	"\aaddress\x18\x03 \x01(\tR\aaddress\x12\x1f\n" +
	"\vproduct_ids\x18\x04 \x03(\x03R\n" +
	"productIds\"}\n" +
	"\x15CreatePaymentResponse\x12\x1d\n" +
	"\n" +
	"request_id\x18\x01 \x01(\x03R\trequestId\x12\x10\n" +
	"\x03url\x18\x02 \x01(\tR\x03url\x123\n" +
	"\apayment\x18\x03 \x01(\v2\x19.payment.v1.PaymentEntityR\apayment\"\xcc\x01\n" +
	"\x14UpdatePaymentRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x126\n" +
	"\x06method\x18\x02 \x01(\x0e2\x19.payment.v1.PaymentMethodH\x00R\x06method\x88\x01\x01\x12\x1d\n" +
	"\aaddress\x18\x03 \x01(\tH\x01R\aaddress\x88\x01\x01\x12+\n" +
	"\x06status\x18\x04 \x01(\x0e2\x0e.common.StatusH\x02R\x06status\x88\x01\x01B\t\n" +
	"\a_methodB\n" +
	"\n" +
	"\b_addressB\t\n" +
	"\a_status\"\x17\n" +
	"\x15UpdatePaymentResponse\"&\n" +
	"\x14CancelPaymentRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\"\x17\n" +
	"\x15CancelPaymentResponse*{\n" +
	"\rPaymentMethod\x12\x1e\n" +
	"\x1aPAYMENT_METHOD_UNSPECIFIED\x10\x00\x12\x17\n" +
	"\x13PAYMENT_METHOD_CASH\x10\x01\x12\x17\n" +
	"\x13PAYMENT_METHOD_MOMO\x10\x02\x12\x18\n" +
	"\x14PAYMENT_METHOD_VNPAY\x10\x03B\xaf\x01\n" +
	"\x0ecom.payment.v1B\fPaymentProtoP\x01ZFgithub.com/shopnexus/shopnexus-protobuf-gen-go/pb/payment/v1;paymentv1\xa2\x02\x03PXX\xaa\x02\n" +
	"Payment.V1\xca\x02\n" +
	"Payment\\V1\xe2\x02\x16Payment\\V1\\GPBMetadata\xea\x02\vPayment::V1b\x06proto3"

var (
	file_payment_v1_payment_proto_rawDescOnce sync.Once
	file_payment_v1_payment_proto_rawDescData []byte
)

func file_payment_v1_payment_proto_rawDescGZIP() []byte {
	file_payment_v1_payment_proto_rawDescOnce.Do(func() {
		file_payment_v1_payment_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_payment_v1_payment_proto_rawDesc), len(file_payment_v1_payment_proto_rawDesc)))
	})
	return file_payment_v1_payment_proto_rawDescData
}

var file_payment_v1_payment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_payment_v1_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_payment_v1_payment_proto_goTypes = []any{
	(PaymentMethod)(0),                // 0: payment.v1.PaymentMethod
	(*ProductOnPayment)(nil),          // 1: payment.v1.ProductOnPayment
	(*PaymentEntity)(nil),             // 2: payment.v1.PaymentEntity
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
	0,  // 1: payment.v1.PaymentEntity.method:type_name -> payment.v1.PaymentMethod
	14, // 2: payment.v1.PaymentEntity.status:type_name -> common.Status
	1,  // 3: payment.v1.PaymentEntity.products:type_name -> payment.v1.ProductOnPayment
	2,  // 4: payment.v1.GetPaymentResponse.data:type_name -> payment.v1.PaymentEntity
	15, // 5: payment.v1.ListPaymentsRequest.pagination:type_name -> common.PaginationRequest
	0,  // 6: payment.v1.ListPaymentsRequest.method:type_name -> payment.v1.PaymentMethod
	14, // 7: payment.v1.ListPaymentsRequest.status:type_name -> common.Status
	2,  // 8: payment.v1.ListPaymentsResponse.data:type_name -> payment.v1.PaymentEntity
	16, // 9: payment.v1.ListPaymentsResponse.pagination:type_name -> common.PaginationResponse
	0,  // 10: payment.v1.CreatePaymentRequest.method:type_name -> payment.v1.PaymentMethod
	2,  // 11: payment.v1.CreatePaymentResponse.payment:type_name -> payment.v1.PaymentEntity
	0,  // 12: payment.v1.UpdatePaymentRequest.method:type_name -> payment.v1.PaymentMethod
	14, // 13: payment.v1.UpdatePaymentRequest.status:type_name -> common.Status
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
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
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_payment_v1_payment_proto_rawDesc), len(file_payment_v1_payment_proto_rawDesc)),
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
	file_payment_v1_payment_proto_goTypes = nil
	file_payment_v1_payment_proto_depIdxs = nil
}
