// @generated by protoc-gen-es v2.4.0 with parameter "target=ts"
// @generated from file product/v1/product_serial.proto (package product.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import type { PaginationRequest, PaginationResponse } from "../../common/pagination_pb";
import { file_common_pagination } from "../../common/pagination_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file product/v1/product_serial.proto.
 */
export const file_product_v1_product_serial: GenFile = /*@__PURE__*/
  fileDesc("Ch9wcm9kdWN0L3YxL3Byb2R1Y3Rfc2VyaWFsLnByb3RvEgpwcm9kdWN0LnYxIowBChNQcm9kdWN0U2VyaWFsRW50aXR5EhEKCXNlcmlhbF9pZBgBIAEoCRISCgpwcm9kdWN0X2lkGAIgASgDEg8KB2lzX3NvbGQYAyABKAgSEQoJaXNfYWN0aXZlGAQgASgIEhQKDGRhdGVfY3JlYXRlZBgFIAEoAxIUCgxkYXRlX3VwZGF0ZWQYBiABKAMiLAoXR2V0UHJvZHVjdFNlcmlhbFJlcXVlc3QSEQoJc2VyaWFsX2lkGAEgASgJIkkKGEdldFByb2R1Y3RTZXJpYWxSZXNwb25zZRItCgRkYXRhGAEgASgLMh8ucHJvZHVjdC52MS5Qcm9kdWN0U2VyaWFsRW50aXR5IsgCChlMaXN0UHJvZHVjdFNlcmlhbHNSZXF1ZXN0Ei0KCnBhZ2luYXRpb24YASABKAsyGS5jb21tb24uUGFnaW5hdGlvblJlcXVlc3QSFgoJc2VyaWFsX2lkGAIgASgJSACIAQESFwoKcHJvZHVjdF9pZBgDIAEoA0gBiAEBEhQKB2lzX3NvbGQYBCABKAhIAogBARIWCglpc19hY3RpdmUYBSABKAhIA4gBARIeChFkYXRlX2NyZWF0ZWRfZnJvbRgGIAEoA0gEiAEBEhwKD2RhdGVfY3JlYXRlZF90bxgHIAEoA0gFiAEBQgwKCl9zZXJpYWxfaWRCDQoLX3Byb2R1Y3RfaWRCCgoIX2lzX3NvbGRCDAoKX2lzX2FjdGl2ZUIUChJfZGF0ZV9jcmVhdGVkX2Zyb21CEgoQX2RhdGVfY3JlYXRlZF90byJ7ChpMaXN0UHJvZHVjdFNlcmlhbHNSZXNwb25zZRItCgRkYXRhGAEgAygLMh8ucHJvZHVjdC52MS5Qcm9kdWN0U2VyaWFsRW50aXR5Ei4KCnBhZ2luYXRpb24YAiABKAsyGi5jb21tb24uUGFnaW5hdGlvblJlc3BvbnNlImcKGkNyZWF0ZVByb2R1Y3RTZXJpYWxSZXF1ZXN0EhEKCXNlcmlhbF9pZBgBIAEoCRISCgpwcm9kdWN0X2lkGAIgASgDEg8KB2lzX3NvbGQYAyABKAgSEQoJaXNfYWN0aXZlGAQgASgIIkwKG0NyZWF0ZVByb2R1Y3RTZXJpYWxSZXNwb25zZRItCgRkYXRhGAEgASgLMh8ucHJvZHVjdC52MS5Qcm9kdWN0U2VyaWFsRW50aXR5IncKGlVwZGF0ZVByb2R1Y3RTZXJpYWxSZXF1ZXN0EhEKCXNlcmlhbF9pZBgBIAEoCRIUCgdpc19zb2xkGAIgASgISACIAQESFgoJaXNfYWN0aXZlGAMgASgISAGIAQFCCgoIX2lzX3NvbGRCDAoKX2lzX2FjdGl2ZSIdChtVcGRhdGVQcm9kdWN0U2VyaWFsUmVzcG9uc2UiLwoaRGVsZXRlUHJvZHVjdFNlcmlhbFJlcXVlc3QSEQoJc2VyaWFsX2lkGAEgASgJIh0KG0RlbGV0ZVByb2R1Y3RTZXJpYWxSZXNwb25zZUK1AQoOY29tLnByb2R1Y3QudjFCElByb2R1Y3RTZXJpYWxQcm90b1ABWkZnaXRodWIuY29tL3Nob3BuZXh1cy9zaG9wbmV4dXMtcHJvdG9idWYtZ2VuLWdvL3BiL3Byb2R1Y3QvdjE7cHJvZHVjdHYxogIDUFhYqgIKUHJvZHVjdC5WMcoCClByb2R1Y3RcVjHiAhZQcm9kdWN0XFYxXEdQQk1ldGFkYXRh6gILUHJvZHVjdDo6VjFiBnByb3RvMw", [file_common_pagination]);

/**
 * @generated from message product.v1.ProductSerialEntity
 */
export type ProductSerialEntity = Message<"product.v1.ProductSerialEntity"> & {
  /**
   * @generated from field: string serial_id = 1;
   */
  serialId: string;

  /**
   * @generated from field: int64 product_id = 2;
   */
  productId: bigint;

  /**
   * @generated from field: bool is_sold = 3;
   */
  isSold: boolean;

  /**
   * @generated from field: bool is_active = 4;
   */
  isActive: boolean;

  /**
   * @generated from field: int64 date_created = 5;
   */
  dateCreated: bigint;

  /**
   * @generated from field: int64 date_updated = 6;
   */
  dateUpdated: bigint;
};

/**
 * Describes the message product.v1.ProductSerialEntity.
 * Use `create(ProductSerialEntitySchema)` to create a new message.
 */
export const ProductSerialEntitySchema: GenMessage<ProductSerialEntity> = /*@__PURE__*/
  messageDesc(file_product_v1_product_serial, 0);

/**
 * @generated from message product.v1.GetProductSerialRequest
 */
export type GetProductSerialRequest = Message<"product.v1.GetProductSerialRequest"> & {
  /**
   * @generated from field: string serial_id = 1;
   */
  serialId: string;
};

/**
 * Describes the message product.v1.GetProductSerialRequest.
 * Use `create(GetProductSerialRequestSchema)` to create a new message.
 */
export const GetProductSerialRequestSchema: GenMessage<GetProductSerialRequest> = /*@__PURE__*/
  messageDesc(file_product_v1_product_serial, 1);

/**
 * @generated from message product.v1.GetProductSerialResponse
 */
export type GetProductSerialResponse = Message<"product.v1.GetProductSerialResponse"> & {
  /**
   * @generated from field: product.v1.ProductSerialEntity data = 1;
   */
  data?: ProductSerialEntity;
};

/**
 * Describes the message product.v1.GetProductSerialResponse.
 * Use `create(GetProductSerialResponseSchema)` to create a new message.
 */
export const GetProductSerialResponseSchema: GenMessage<GetProductSerialResponse> = /*@__PURE__*/
  messageDesc(file_product_v1_product_serial, 2);

/**
 * @generated from message product.v1.ListProductSerialsRequest
 */
export type ListProductSerialsRequest = Message<"product.v1.ListProductSerialsRequest"> & {
  /**
   * @generated from field: common.PaginationRequest pagination = 1;
   */
  pagination?: PaginationRequest;

  /**
   * @generated from field: optional string serial_id = 2;
   */
  serialId?: string;

  /**
   * @generated from field: optional int64 product_id = 3;
   */
  productId?: bigint;

  /**
   * @generated from field: optional bool is_sold = 4;
   */
  isSold?: boolean;

  /**
   * @generated from field: optional bool is_active = 5;
   */
  isActive?: boolean;

  /**
   * @generated from field: optional int64 date_created_from = 6;
   */
  dateCreatedFrom?: bigint;

  /**
   * @generated from field: optional int64 date_created_to = 7;
   */
  dateCreatedTo?: bigint;
};

/**
 * Describes the message product.v1.ListProductSerialsRequest.
 * Use `create(ListProductSerialsRequestSchema)` to create a new message.
 */
export const ListProductSerialsRequestSchema: GenMessage<ListProductSerialsRequest> = /*@__PURE__*/
  messageDesc(file_product_v1_product_serial, 3);

/**
 * @generated from message product.v1.ListProductSerialsResponse
 */
export type ListProductSerialsResponse = Message<"product.v1.ListProductSerialsResponse"> & {
  /**
   * @generated from field: repeated product.v1.ProductSerialEntity data = 1;
   */
  data: ProductSerialEntity[];

  /**
   * @generated from field: common.PaginationResponse pagination = 2;
   */
  pagination?: PaginationResponse;
};

/**
 * Describes the message product.v1.ListProductSerialsResponse.
 * Use `create(ListProductSerialsResponseSchema)` to create a new message.
 */
export const ListProductSerialsResponseSchema: GenMessage<ListProductSerialsResponse> = /*@__PURE__*/
  messageDesc(file_product_v1_product_serial, 4);

/**
 * @generated from message product.v1.CreateProductSerialRequest
 */
export type CreateProductSerialRequest = Message<"product.v1.CreateProductSerialRequest"> & {
  /**
   * @generated from field: string serial_id = 1;
   */
  serialId: string;

  /**
   * @generated from field: int64 product_id = 2;
   */
  productId: bigint;

  /**
   * @generated from field: bool is_sold = 3;
   */
  isSold: boolean;

  /**
   * @generated from field: bool is_active = 4;
   */
  isActive: boolean;
};

/**
 * Describes the message product.v1.CreateProductSerialRequest.
 * Use `create(CreateProductSerialRequestSchema)` to create a new message.
 */
export const CreateProductSerialRequestSchema: GenMessage<CreateProductSerialRequest> = /*@__PURE__*/
  messageDesc(file_product_v1_product_serial, 5);

/**
 * @generated from message product.v1.CreateProductSerialResponse
 */
export type CreateProductSerialResponse = Message<"product.v1.CreateProductSerialResponse"> & {
  /**
   * @generated from field: product.v1.ProductSerialEntity data = 1;
   */
  data?: ProductSerialEntity;
};

/**
 * Describes the message product.v1.CreateProductSerialResponse.
 * Use `create(CreateProductSerialResponseSchema)` to create a new message.
 */
export const CreateProductSerialResponseSchema: GenMessage<CreateProductSerialResponse> = /*@__PURE__*/
  messageDesc(file_product_v1_product_serial, 6);

/**
 * @generated from message product.v1.UpdateProductSerialRequest
 */
export type UpdateProductSerialRequest = Message<"product.v1.UpdateProductSerialRequest"> & {
  /**
   * @generated from field: string serial_id = 1;
   */
  serialId: string;

  /**
   * @generated from field: optional bool is_sold = 2;
   */
  isSold?: boolean;

  /**
   * @generated from field: optional bool is_active = 3;
   */
  isActive?: boolean;
};

/**
 * Describes the message product.v1.UpdateProductSerialRequest.
 * Use `create(UpdateProductSerialRequestSchema)` to create a new message.
 */
export const UpdateProductSerialRequestSchema: GenMessage<UpdateProductSerialRequest> = /*@__PURE__*/
  messageDesc(file_product_v1_product_serial, 7);

/**
 * @generated from message product.v1.UpdateProductSerialResponse
 */
export type UpdateProductSerialResponse = Message<"product.v1.UpdateProductSerialResponse"> & {
};

/**
 * Describes the message product.v1.UpdateProductSerialResponse.
 * Use `create(UpdateProductSerialResponseSchema)` to create a new message.
 */
export const UpdateProductSerialResponseSchema: GenMessage<UpdateProductSerialResponse> = /*@__PURE__*/
  messageDesc(file_product_v1_product_serial, 8);

/**
 * @generated from message product.v1.DeleteProductSerialRequest
 */
export type DeleteProductSerialRequest = Message<"product.v1.DeleteProductSerialRequest"> & {
  /**
   * @generated from field: string serial_id = 1;
   */
  serialId: string;
};

/**
 * Describes the message product.v1.DeleteProductSerialRequest.
 * Use `create(DeleteProductSerialRequestSchema)` to create a new message.
 */
export const DeleteProductSerialRequestSchema: GenMessage<DeleteProductSerialRequest> = /*@__PURE__*/
  messageDesc(file_product_v1_product_serial, 9);

/**
 * @generated from message product.v1.DeleteProductSerialResponse
 */
export type DeleteProductSerialResponse = Message<"product.v1.DeleteProductSerialResponse"> & {
};

/**
 * Describes the message product.v1.DeleteProductSerialResponse.
 * Use `create(DeleteProductSerialResponseSchema)` to create a new message.
 */
export const DeleteProductSerialResponseSchema: GenMessage<DeleteProductSerialResponse> = /*@__PURE__*/
  messageDesc(file_product_v1_product_serial, 10);

