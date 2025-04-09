// @generated by protoc-gen-es v2.2.4 with parameter "target=ts"
// @generated from file product/v1/comment.proto (package product.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import type { PaginationRequest, PaginationResponse } from "../../common/pagination_pb";
import { file_common_pagination } from "../../common/pagination_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file product/v1/comment.proto.
 */
export const file_product_v1_comment: GenFile = /*@__PURE__*/
  fileDesc("Chhwcm9kdWN0L3YxL2NvbW1lbnQucHJvdG8SCnByb2R1Y3QudjEiuwEKDUNvbW1lbnRFbnRpdHkSCgoCaWQYASABKAMSDwoHdXNlcl9pZBgCIAEoAxIPCgdkZXN0X2lkGAMgASgDEgwKBGJvZHkYBCABKAkSDgoGdXB2b3RlGAUgASgDEhAKCGRvd252b3RlGAYgASgDEg0KBXNjb3JlGAcgASgFEhQKDGRhdGVfY3JlYXRlZBgIIAEoAxIUCgxkYXRlX3VwZGF0ZWQYCSABKAMSEQoJcmVzb3VyY2VzGAogAygJIlEKEUdldENvbW1lbnRSZXF1ZXN0Eg8KAmlkGAEgASgDSACIAQESFgoJc2VyaWFsX2lkGAIgASgJSAGIAQFCBQoDX2lkQgwKCl9zZXJpYWxfaWQiPQoSR2V0Q29tbWVudFJlc3BvbnNlEicKBGRhdGEYASABKAsyGS5wcm9kdWN0LnYxLkNvbW1lbnRFbnRpdHkigAQKE0xpc3RDb21tZW50c1JlcXVlc3QSLQoKcGFnaW5hdGlvbhgBIAEoCzIZLmNvbW1vbi5QYWdpbmF0aW9uUmVxdWVzdBIUCgd1c2VyX2lkGAIgASgDSACIAQESFAoHZGVzdF9pZBgDIAEoA0gBiAEBEhEKBGJvZHkYBCABKAlIAogBARIYCgt1cHZvdGVfZnJvbRgFIAEoA0gDiAEBEhYKCXVwdm90ZV90bxgGIAEoA0gEiAEBEhoKDWRvd252b3RlX2Zyb20YByABKANIBYgBARIYCgtkb3dudm90ZV90bxgIIAEoA0gGiAEBEhcKCnNjb3JlX2Zyb20YCSABKANIB4gBARIVCghzY29yZV90bxgKIAEoA0gIiAEBEh4KEWRhdGVfY3JlYXRlZF9mcm9tGAsgASgDSAmIAQESHAoPZGF0ZV9jcmVhdGVkX3RvGAwgASgDSAqIAQFCCgoIX3VzZXJfaWRCCgoIX2Rlc3RfaWRCBwoFX2JvZHlCDgoMX3Vwdm90ZV9mcm9tQgwKCl91cHZvdGVfdG9CEAoOX2Rvd252b3RlX2Zyb21CDgoMX2Rvd252b3RlX3RvQg0KC19zY29yZV9mcm9tQgsKCV9zY29yZV90b0IUChJfZGF0ZV9jcmVhdGVkX2Zyb21CEgoQX2RhdGVfY3JlYXRlZF90byJvChRMaXN0Q29tbWVudHNSZXNwb25zZRInCgRkYXRhGAEgAygLMhkucHJvZHVjdC52MS5Db21tZW50RW50aXR5Ei4KCnBhZ2luYXRpb24YAiABKAsyGi5jb21tb24uUGFnaW5hdGlvblJlc3BvbnNlIkgKFENyZWF0ZUNvbW1lbnRSZXF1ZXN0Eg8KB2Rlc3RfaWQYASABKAMSDAoEYm9keRgCIAEoCRIRCglyZXNvdXJjZXMYAyADKAkiFwoVQ3JlYXRlQ29tbWVudFJlc3BvbnNlIlEKFFVwZGF0ZUNvbW1lbnRSZXF1ZXN0EgoKAmlkGAEgASgDEhEKBGJvZHkYAiABKAlIAIgBARIRCglyZXNvdXJjZXMYAyADKAlCBwoFX2JvZHkiFwoVVXBkYXRlQ29tbWVudFJlc3BvbnNlIiIKFERlbGV0ZUNvbW1lbnRSZXF1ZXN0EgoKAmlkGAEgASgDIhcKFURlbGV0ZUNvbW1lbnRSZXNwb25zZUKvAQoOY29tLnByb2R1Y3QudjFCDENvbW1lbnRQcm90b1ABWkZnaXRodWIuY29tL3Nob3BuZXh1cy9zaG9wbmV4dXMtcHJvdG9idWYtZ2VuLWdvL3BiL3Byb2R1Y3QvdjE7cHJvZHVjdHYxogIDUFhYqgIKUHJvZHVjdC5WMcoCClByb2R1Y3RcVjHiAhZQcm9kdWN0XFYxXEdQQk1ldGFkYXRh6gILUHJvZHVjdDo6VjFiBnByb3RvMw", [file_common_pagination]);

/**
 * COMMENT ENTITY
 *
 * @generated from message product.v1.CommentEntity
 */
export type CommentEntity = Message<"product.v1.CommentEntity"> & {
  /**
   * @generated from field: int64 id = 1;
   */
  id: bigint;

  /**
   * @generated from field: int64 user_id = 2;
   */
  userId: bigint;

  /**
   * @generated from field: int64 dest_id = 3;
   */
  destId: bigint;

  /**
   * @generated from field: string body = 4;
   */
  body: string;

  /**
   * @generated from field: int64 upvote = 5;
   */
  upvote: bigint;

  /**
   * @generated from field: int64 downvote = 6;
   */
  downvote: bigint;

  /**
   * @generated from field: int32 score = 7;
   */
  score: number;

  /**
   * @generated from field: int64 date_created = 8;
   */
  dateCreated: bigint;

  /**
   * @generated from field: int64 date_updated = 9;
   */
  dateUpdated: bigint;

  /**
   * @generated from field: repeated string resources = 10;
   */
  resources: string[];
};

/**
 * Describes the message product.v1.CommentEntity.
 * Use `create(CommentEntitySchema)` to create a new message.
 */
export const CommentEntitySchema: GenMessage<CommentEntity> = /*@__PURE__*/
  messageDesc(file_product_v1_comment, 0);

/**
 * @generated from message product.v1.GetCommentRequest
 */
export type GetCommentRequest = Message<"product.v1.GetCommentRequest"> & {
  /**
   * @generated from field: optional int64 id = 1;
   */
  id?: bigint;

  /**
   * @generated from field: optional string serial_id = 2;
   */
  serialId?: string;
};

/**
 * Describes the message product.v1.GetCommentRequest.
 * Use `create(GetCommentRequestSchema)` to create a new message.
 */
export const GetCommentRequestSchema: GenMessage<GetCommentRequest> = /*@__PURE__*/
  messageDesc(file_product_v1_comment, 1);

/**
 * @generated from message product.v1.GetCommentResponse
 */
export type GetCommentResponse = Message<"product.v1.GetCommentResponse"> & {
  /**
   * @generated from field: product.v1.CommentEntity data = 1;
   */
  data?: CommentEntity;
};

/**
 * Describes the message product.v1.GetCommentResponse.
 * Use `create(GetCommentResponseSchema)` to create a new message.
 */
export const GetCommentResponseSchema: GenMessage<GetCommentResponse> = /*@__PURE__*/
  messageDesc(file_product_v1_comment, 2);

/**
 * @generated from message product.v1.ListCommentsRequest
 */
export type ListCommentsRequest = Message<"product.v1.ListCommentsRequest"> & {
  /**
   * @generated from field: common.PaginationRequest pagination = 1;
   */
  pagination?: PaginationRequest;

  /**
   * @generated from field: optional int64 user_id = 2;
   */
  userId?: bigint;

  /**
   * @generated from field: optional int64 dest_id = 3;
   */
  destId?: bigint;

  /**
   * @generated from field: optional string body = 4;
   */
  body?: string;

  /**
   * @generated from field: optional int64 upvote_from = 5;
   */
  upvoteFrom?: bigint;

  /**
   * @generated from field: optional int64 upvote_to = 6;
   */
  upvoteTo?: bigint;

  /**
   * @generated from field: optional int64 downvote_from = 7;
   */
  downvoteFrom?: bigint;

  /**
   * @generated from field: optional int64 downvote_to = 8;
   */
  downvoteTo?: bigint;

  /**
   * @generated from field: optional int64 score_from = 9;
   */
  scoreFrom?: bigint;

  /**
   * @generated from field: optional int64 score_to = 10;
   */
  scoreTo?: bigint;

  /**
   * @generated from field: optional int64 date_created_from = 11;
   */
  dateCreatedFrom?: bigint;

  /**
   * @generated from field: optional int64 date_created_to = 12;
   */
  dateCreatedTo?: bigint;
};

/**
 * Describes the message product.v1.ListCommentsRequest.
 * Use `create(ListCommentsRequestSchema)` to create a new message.
 */
export const ListCommentsRequestSchema: GenMessage<ListCommentsRequest> = /*@__PURE__*/
  messageDesc(file_product_v1_comment, 3);

/**
 * @generated from message product.v1.ListCommentsResponse
 */
export type ListCommentsResponse = Message<"product.v1.ListCommentsResponse"> & {
  /**
   * @generated from field: repeated product.v1.CommentEntity data = 1;
   */
  data: CommentEntity[];

  /**
   * @generated from field: common.PaginationResponse pagination = 2;
   */
  pagination?: PaginationResponse;
};

/**
 * Describes the message product.v1.ListCommentsResponse.
 * Use `create(ListCommentsResponseSchema)` to create a new message.
 */
export const ListCommentsResponseSchema: GenMessage<ListCommentsResponse> = /*@__PURE__*/
  messageDesc(file_product_v1_comment, 4);

/**
 * @generated from message product.v1.CreateCommentRequest
 */
export type CreateCommentRequest = Message<"product.v1.CreateCommentRequest"> & {
  /**
   * @generated from field: int64 dest_id = 1;
   */
  destId: bigint;

  /**
   * @generated from field: string body = 2;
   */
  body: string;

  /**
   * @generated from field: repeated string resources = 3;
   */
  resources: string[];
};

/**
 * Describes the message product.v1.CreateCommentRequest.
 * Use `create(CreateCommentRequestSchema)` to create a new message.
 */
export const CreateCommentRequestSchema: GenMessage<CreateCommentRequest> = /*@__PURE__*/
  messageDesc(file_product_v1_comment, 5);

/**
 * @generated from message product.v1.CreateCommentResponse
 */
export type CreateCommentResponse = Message<"product.v1.CreateCommentResponse"> & {
};

/**
 * Describes the message product.v1.CreateCommentResponse.
 * Use `create(CreateCommentResponseSchema)` to create a new message.
 */
export const CreateCommentResponseSchema: GenMessage<CreateCommentResponse> = /*@__PURE__*/
  messageDesc(file_product_v1_comment, 6);

/**
 * @generated from message product.v1.UpdateCommentRequest
 */
export type UpdateCommentRequest = Message<"product.v1.UpdateCommentRequest"> & {
  /**
   * @generated from field: int64 id = 1;
   */
  id: bigint;

  /**
   * @generated from field: optional string body = 2;
   */
  body?: string;

  /**
   * @generated from field: repeated string resources = 3;
   */
  resources: string[];
};

/**
 * Describes the message product.v1.UpdateCommentRequest.
 * Use `create(UpdateCommentRequestSchema)` to create a new message.
 */
export const UpdateCommentRequestSchema: GenMessage<UpdateCommentRequest> = /*@__PURE__*/
  messageDesc(file_product_v1_comment, 7);

/**
 * @generated from message product.v1.UpdateCommentResponse
 */
export type UpdateCommentResponse = Message<"product.v1.UpdateCommentResponse"> & {
};

/**
 * Describes the message product.v1.UpdateCommentResponse.
 * Use `create(UpdateCommentResponseSchema)` to create a new message.
 */
export const UpdateCommentResponseSchema: GenMessage<UpdateCommentResponse> = /*@__PURE__*/
  messageDesc(file_product_v1_comment, 8);

/**
 * @generated from message product.v1.DeleteCommentRequest
 */
export type DeleteCommentRequest = Message<"product.v1.DeleteCommentRequest"> & {
  /**
   * @generated from field: int64 id = 1;
   */
  id: bigint;
};

/**
 * Describes the message product.v1.DeleteCommentRequest.
 * Use `create(DeleteCommentRequestSchema)` to create a new message.
 */
export const DeleteCommentRequestSchema: GenMessage<DeleteCommentRequest> = /*@__PURE__*/
  messageDesc(file_product_v1_comment, 9);

/**
 * @generated from message product.v1.DeleteCommentResponse
 */
export type DeleteCommentResponse = Message<"product.v1.DeleteCommentResponse"> & {
};

/**
 * Describes the message product.v1.DeleteCommentResponse.
 * Use `create(DeleteCommentResponseSchema)` to create a new message.
 */
export const DeleteCommentResponseSchema: GenMessage<DeleteCommentResponse> = /*@__PURE__*/
  messageDesc(file_product_v1_comment, 10);

