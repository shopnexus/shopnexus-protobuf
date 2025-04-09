// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: product/v1/comment.proto

package productv1

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

// COMMENT ENTITY
type CommentEntity struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        int64                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DestId        int64                  `protobuf:"varint,3,opt,name=dest_id,json=destId,proto3" json:"dest_id,omitempty"`
	Body          string                 `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	Upvote        int64                  `protobuf:"varint,5,opt,name=upvote,proto3" json:"upvote,omitempty"`
	Downvote      int64                  `protobuf:"varint,6,opt,name=downvote,proto3" json:"downvote,omitempty"`
	Score         int32                  `protobuf:"varint,7,opt,name=score,proto3" json:"score,omitempty"`
	DateCreated   int64                  `protobuf:"varint,8,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty"`
	DateUpdated   int64                  `protobuf:"varint,9,opt,name=date_updated,json=dateUpdated,proto3" json:"date_updated,omitempty"`
	Resources     []string               `protobuf:"bytes,10,rep,name=resources,proto3" json:"resources,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CommentEntity) Reset() {
	*x = CommentEntity{}
	mi := &file_product_v1_comment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CommentEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentEntity) ProtoMessage() {}

func (x *CommentEntity) ProtoReflect() protoreflect.Message {
	mi := &file_product_v1_comment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentEntity.ProtoReflect.Descriptor instead.
func (*CommentEntity) Descriptor() ([]byte, []int) {
	return file_product_v1_comment_proto_rawDescGZIP(), []int{0}
}

func (x *CommentEntity) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CommentEntity) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CommentEntity) GetDestId() int64 {
	if x != nil {
		return x.DestId
	}
	return 0
}

func (x *CommentEntity) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *CommentEntity) GetUpvote() int64 {
	if x != nil {
		return x.Upvote
	}
	return 0
}

func (x *CommentEntity) GetDownvote() int64 {
	if x != nil {
		return x.Downvote
	}
	return 0
}

func (x *CommentEntity) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *CommentEntity) GetDateCreated() int64 {
	if x != nil {
		return x.DateCreated
	}
	return 0
}

func (x *CommentEntity) GetDateUpdated() int64 {
	if x != nil {
		return x.DateUpdated
	}
	return 0
}

func (x *CommentEntity) GetResources() []string {
	if x != nil {
		return x.Resources
	}
	return nil
}

type GetCommentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            *int64                 `protobuf:"varint,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	SerialId      *string                `protobuf:"bytes,2,opt,name=serial_id,json=serialId,proto3,oneof" json:"serial_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCommentRequest) Reset() {
	*x = GetCommentRequest{}
	mi := &file_product_v1_comment_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentRequest) ProtoMessage() {}

func (x *GetCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_v1_comment_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentRequest.ProtoReflect.Descriptor instead.
func (*GetCommentRequest) Descriptor() ([]byte, []int) {
	return file_product_v1_comment_proto_rawDescGZIP(), []int{1}
}

func (x *GetCommentRequest) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *GetCommentRequest) GetSerialId() string {
	if x != nil && x.SerialId != nil {
		return *x.SerialId
	}
	return ""
}

type GetCommentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          *CommentEntity         `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCommentResponse) Reset() {
	*x = GetCommentResponse{}
	mi := &file_product_v1_comment_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentResponse) ProtoMessage() {}

func (x *GetCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_v1_comment_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentResponse.ProtoReflect.Descriptor instead.
func (*GetCommentResponse) Descriptor() ([]byte, []int) {
	return file_product_v1_comment_proto_rawDescGZIP(), []int{2}
}

func (x *GetCommentResponse) GetData() *CommentEntity {
	if x != nil {
		return x.Data
	}
	return nil
}

type ListCommentsRequest struct {
	state           protoimpl.MessageState    `protogen:"open.v1"`
	Pagination      *common.PaginationRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	UserId          *int64                    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3,oneof" json:"user_id,omitempty"`
	DestId          *int64                    `protobuf:"varint,3,opt,name=dest_id,json=destId,proto3,oneof" json:"dest_id,omitempty"`
	Body            *string                   `protobuf:"bytes,4,opt,name=body,proto3,oneof" json:"body,omitempty"`
	UpvoteFrom      *int64                    `protobuf:"varint,5,opt,name=upvote_from,json=upvoteFrom,proto3,oneof" json:"upvote_from,omitempty"`
	UpvoteTo        *int64                    `protobuf:"varint,6,opt,name=upvote_to,json=upvoteTo,proto3,oneof" json:"upvote_to,omitempty"`
	DownvoteFrom    *int64                    `protobuf:"varint,7,opt,name=downvote_from,json=downvoteFrom,proto3,oneof" json:"downvote_from,omitempty"`
	DownvoteTo      *int64                    `protobuf:"varint,8,opt,name=downvote_to,json=downvoteTo,proto3,oneof" json:"downvote_to,omitempty"`
	ScoreFrom       *int64                    `protobuf:"varint,9,opt,name=score_from,json=scoreFrom,proto3,oneof" json:"score_from,omitempty"`
	ScoreTo         *int64                    `protobuf:"varint,10,opt,name=score_to,json=scoreTo,proto3,oneof" json:"score_to,omitempty"`
	DateCreatedFrom *int64                    `protobuf:"varint,11,opt,name=date_created_from,json=dateCreatedFrom,proto3,oneof" json:"date_created_from,omitempty"`
	DateCreatedTo   *int64                    `protobuf:"varint,12,opt,name=date_created_to,json=dateCreatedTo,proto3,oneof" json:"date_created_to,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ListCommentsRequest) Reset() {
	*x = ListCommentsRequest{}
	mi := &file_product_v1_comment_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCommentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCommentsRequest) ProtoMessage() {}

func (x *ListCommentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_v1_comment_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCommentsRequest.ProtoReflect.Descriptor instead.
func (*ListCommentsRequest) Descriptor() ([]byte, []int) {
	return file_product_v1_comment_proto_rawDescGZIP(), []int{3}
}

func (x *ListCommentsRequest) GetPagination() *common.PaginationRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ListCommentsRequest) GetUserId() int64 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

func (x *ListCommentsRequest) GetDestId() int64 {
	if x != nil && x.DestId != nil {
		return *x.DestId
	}
	return 0
}

func (x *ListCommentsRequest) GetBody() string {
	if x != nil && x.Body != nil {
		return *x.Body
	}
	return ""
}

func (x *ListCommentsRequest) GetUpvoteFrom() int64 {
	if x != nil && x.UpvoteFrom != nil {
		return *x.UpvoteFrom
	}
	return 0
}

func (x *ListCommentsRequest) GetUpvoteTo() int64 {
	if x != nil && x.UpvoteTo != nil {
		return *x.UpvoteTo
	}
	return 0
}

func (x *ListCommentsRequest) GetDownvoteFrom() int64 {
	if x != nil && x.DownvoteFrom != nil {
		return *x.DownvoteFrom
	}
	return 0
}

func (x *ListCommentsRequest) GetDownvoteTo() int64 {
	if x != nil && x.DownvoteTo != nil {
		return *x.DownvoteTo
	}
	return 0
}

func (x *ListCommentsRequest) GetScoreFrom() int64 {
	if x != nil && x.ScoreFrom != nil {
		return *x.ScoreFrom
	}
	return 0
}

func (x *ListCommentsRequest) GetScoreTo() int64 {
	if x != nil && x.ScoreTo != nil {
		return *x.ScoreTo
	}
	return 0
}

func (x *ListCommentsRequest) GetDateCreatedFrom() int64 {
	if x != nil && x.DateCreatedFrom != nil {
		return *x.DateCreatedFrom
	}
	return 0
}

func (x *ListCommentsRequest) GetDateCreatedTo() int64 {
	if x != nil && x.DateCreatedTo != nil {
		return *x.DateCreatedTo
	}
	return 0
}

type ListCommentsResponse struct {
	state         protoimpl.MessageState     `protogen:"open.v1"`
	Data          []*CommentEntity           `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Pagination    *common.PaginationResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCommentsResponse) Reset() {
	*x = ListCommentsResponse{}
	mi := &file_product_v1_comment_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCommentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCommentsResponse) ProtoMessage() {}

func (x *ListCommentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_v1_comment_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCommentsResponse.ProtoReflect.Descriptor instead.
func (*ListCommentsResponse) Descriptor() ([]byte, []int) {
	return file_product_v1_comment_proto_rawDescGZIP(), []int{4}
}

func (x *ListCommentsResponse) GetData() []*CommentEntity {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ListCommentsResponse) GetPagination() *common.PaginationResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type CreateCommentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DestId        int64                  `protobuf:"varint,1,opt,name=dest_id,json=destId,proto3" json:"dest_id,omitempty"`
	Body          string                 `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Resources     []string               `protobuf:"bytes,3,rep,name=resources,proto3" json:"resources,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCommentRequest) Reset() {
	*x = CreateCommentRequest{}
	mi := &file_product_v1_comment_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommentRequest) ProtoMessage() {}

func (x *CreateCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_v1_comment_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommentRequest.ProtoReflect.Descriptor instead.
func (*CreateCommentRequest) Descriptor() ([]byte, []int) {
	return file_product_v1_comment_proto_rawDescGZIP(), []int{5}
}

func (x *CreateCommentRequest) GetDestId() int64 {
	if x != nil {
		return x.DestId
	}
	return 0
}

func (x *CreateCommentRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *CreateCommentRequest) GetResources() []string {
	if x != nil {
		return x.Resources
	}
	return nil
}

type CreateCommentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCommentResponse) Reset() {
	*x = CreateCommentResponse{}
	mi := &file_product_v1_comment_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommentResponse) ProtoMessage() {}

func (x *CreateCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_v1_comment_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommentResponse.ProtoReflect.Descriptor instead.
func (*CreateCommentResponse) Descriptor() ([]byte, []int) {
	return file_product_v1_comment_proto_rawDescGZIP(), []int{6}
}

type UpdateCommentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Body          *string                `protobuf:"bytes,2,opt,name=body,proto3,oneof" json:"body,omitempty"`
	Resources     []string               `protobuf:"bytes,3,rep,name=resources,proto3" json:"resources,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateCommentRequest) Reset() {
	*x = UpdateCommentRequest{}
	mi := &file_product_v1_comment_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCommentRequest) ProtoMessage() {}

func (x *UpdateCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_v1_comment_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCommentRequest.ProtoReflect.Descriptor instead.
func (*UpdateCommentRequest) Descriptor() ([]byte, []int) {
	return file_product_v1_comment_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateCommentRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateCommentRequest) GetBody() string {
	if x != nil && x.Body != nil {
		return *x.Body
	}
	return ""
}

func (x *UpdateCommentRequest) GetResources() []string {
	if x != nil {
		return x.Resources
	}
	return nil
}

type UpdateCommentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateCommentResponse) Reset() {
	*x = UpdateCommentResponse{}
	mi := &file_product_v1_comment_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCommentResponse) ProtoMessage() {}

func (x *UpdateCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_v1_comment_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCommentResponse.ProtoReflect.Descriptor instead.
func (*UpdateCommentResponse) Descriptor() ([]byte, []int) {
	return file_product_v1_comment_proto_rawDescGZIP(), []int{8}
}

type DeleteCommentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteCommentRequest) Reset() {
	*x = DeleteCommentRequest{}
	mi := &file_product_v1_comment_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCommentRequest) ProtoMessage() {}

func (x *DeleteCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_v1_comment_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCommentRequest.ProtoReflect.Descriptor instead.
func (*DeleteCommentRequest) Descriptor() ([]byte, []int) {
	return file_product_v1_comment_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteCommentRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteCommentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteCommentResponse) Reset() {
	*x = DeleteCommentResponse{}
	mi := &file_product_v1_comment_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCommentResponse) ProtoMessage() {}

func (x *DeleteCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_v1_comment_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCommentResponse.ProtoReflect.Descriptor instead.
func (*DeleteCommentResponse) Descriptor() ([]byte, []int) {
	return file_product_v1_comment_proto_rawDescGZIP(), []int{10}
}

var File_product_v1_comment_proto protoreflect.FileDescriptor

var file_product_v1_comment_proto_rawDesc = string([]byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x93, 0x02, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x65,
	0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x64, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x70, 0x76, 0x6f, 0x74,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x6f, 0x77, 0x6e, 0x76, 0x6f, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x64, 0x6f, 0x77, 0x6e, 0x76, 0x6f, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0x5f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x20, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x88, 0x01,
	0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x73, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x22, 0x43, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x86, 0x05, 0x0a, 0x13,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48,
	0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x07,
	0x64, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52,
	0x06, 0x64, 0x65, 0x73, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x66, 0x72,
	0x6f, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x48, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x76, 0x6f,
	0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x75, 0x70, 0x76,
	0x6f, 0x74, 0x65, 0x5f, 0x74, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x48, 0x04, 0x52, 0x08,
	0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x54, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x64,
	0x6f, 0x77, 0x6e, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x05, 0x52, 0x0c, 0x64, 0x6f, 0x77, 0x6e, 0x76, 0x6f, 0x74, 0x65, 0x46, 0x72,
	0x6f, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x64, 0x6f, 0x77, 0x6e, 0x76, 0x6f, 0x74,
	0x65, 0x5f, 0x74, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x48, 0x06, 0x52, 0x0a, 0x64, 0x6f,
	0x77, 0x6e, 0x76, 0x6f, 0x74, 0x65, 0x54, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x48,
	0x07, 0x52, 0x09, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x88, 0x01, 0x01, 0x12,
	0x1e, 0x0a, 0x08, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x74, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x03, 0x48, 0x08, 0x52, 0x07, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x54, 0x6f, 0x88, 0x01, 0x01, 0x12,
	0x2f, 0x0a, 0x11, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x66, 0x72, 0x6f, 0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x48, 0x09, 0x52, 0x0f, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x88, 0x01, 0x01,
	0x12, 0x2b, 0x0a, 0x0f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x74, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x48, 0x0a, 0x52, 0x0d, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x6f, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a,
	0x08, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x64, 0x65,
	0x73, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x42, 0x0e,
	0x0a, 0x0c, 0x5f, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x75, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x74, 0x6f, 0x42, 0x10, 0x0a, 0x0e,
	0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x42, 0x0e,
	0x0a, 0x0c, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x76, 0x6f, 0x74, 0x65, 0x5f, 0x74, 0x6f, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x42, 0x0b, 0x0a,
	0x09, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x74, 0x6f, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x72, 0x6f, 0x6d,
	0x42, 0x12, 0x0a, 0x10, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3a, 0x0a, 0x0a,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x61, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x64, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x64, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x1c, 0x0a,
	0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0x17, 0x0a, 0x15, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x66, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x17, 0x0a, 0x15,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x17, 0x0a,
	0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xaf, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x6e, 0x65, 0x78, 0x75, 0x73, 0x2f,
	0x73, 0x68, 0x6f, 0x70, 0x6e, 0x65, 0x78, 0x75, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x16, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_product_v1_comment_proto_rawDescOnce sync.Once
	file_product_v1_comment_proto_rawDescData []byte
)

func file_product_v1_comment_proto_rawDescGZIP() []byte {
	file_product_v1_comment_proto_rawDescOnce.Do(func() {
		file_product_v1_comment_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_product_v1_comment_proto_rawDesc), len(file_product_v1_comment_proto_rawDesc)))
	})
	return file_product_v1_comment_proto_rawDescData
}

var file_product_v1_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_product_v1_comment_proto_goTypes = []any{
	(*CommentEntity)(nil),             // 0: product.v1.CommentEntity
	(*GetCommentRequest)(nil),         // 1: product.v1.GetCommentRequest
	(*GetCommentResponse)(nil),        // 2: product.v1.GetCommentResponse
	(*ListCommentsRequest)(nil),       // 3: product.v1.ListCommentsRequest
	(*ListCommentsResponse)(nil),      // 4: product.v1.ListCommentsResponse
	(*CreateCommentRequest)(nil),      // 5: product.v1.CreateCommentRequest
	(*CreateCommentResponse)(nil),     // 6: product.v1.CreateCommentResponse
	(*UpdateCommentRequest)(nil),      // 7: product.v1.UpdateCommentRequest
	(*UpdateCommentResponse)(nil),     // 8: product.v1.UpdateCommentResponse
	(*DeleteCommentRequest)(nil),      // 9: product.v1.DeleteCommentRequest
	(*DeleteCommentResponse)(nil),     // 10: product.v1.DeleteCommentResponse
	(*common.PaginationRequest)(nil),  // 11: common.PaginationRequest
	(*common.PaginationResponse)(nil), // 12: common.PaginationResponse
}
var file_product_v1_comment_proto_depIdxs = []int32{
	0,  // 0: product.v1.GetCommentResponse.data:type_name -> product.v1.CommentEntity
	11, // 1: product.v1.ListCommentsRequest.pagination:type_name -> common.PaginationRequest
	0,  // 2: product.v1.ListCommentsResponse.data:type_name -> product.v1.CommentEntity
	12, // 3: product.v1.ListCommentsResponse.pagination:type_name -> common.PaginationResponse
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_product_v1_comment_proto_init() }
func file_product_v1_comment_proto_init() {
	if File_product_v1_comment_proto != nil {
		return
	}
	file_product_v1_comment_proto_msgTypes[1].OneofWrappers = []any{}
	file_product_v1_comment_proto_msgTypes[3].OneofWrappers = []any{}
	file_product_v1_comment_proto_msgTypes[7].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_product_v1_comment_proto_rawDesc), len(file_product_v1_comment_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_product_v1_comment_proto_goTypes,
		DependencyIndexes: file_product_v1_comment_proto_depIdxs,
		MessageInfos:      file_product_v1_comment_proto_msgTypes,
	}.Build()
	File_product_v1_comment_proto = out.File
	file_product_v1_comment_proto_goTypes = nil
	file_product_v1_comment_proto_depIdxs = nil
}
