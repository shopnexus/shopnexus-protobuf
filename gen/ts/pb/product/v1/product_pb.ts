// @generated by protoc-gen-es v2.2.4 with parameter "target=ts"
// @generated from file product/v1/product.proto (package product.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import type { PaginationRequest, PaginationResponse } from "../../common/pagination_pb";
import { file_common_pagination } from "../../common/pagination_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file product/v1/product.proto.
 */
export const file_product_v1_product: GenFile = /*@__PURE__*/
  fileDesc("Chhwcm9kdWN0L3YxL3Byb2R1Y3QucHJvdG8SCnByb2R1Y3QudjEi3wEKDVByb2R1Y3RFbnRpdHkSCgoCaWQYASABKAMSEQoJc2VyaWFsX2lkGAIgASgJEhgKEHByb2R1Y3RfbW9kZWxfaWQYAyABKAMSEAoIcXVhbnRpdHkYBCABKAMSDAoEc29sZBgFIAEoAxIRCglhZGRfcHJpY2UYBiABKAMSEQoJaXNfYWN0aXZlGAcgASgIEhAKCG1ldGFkYXRhGAggASgMEhQKDGRhdGVfY3JlYXRlZBgJIAEoAxIUCgxkYXRlX3VwZGF0ZWQYCiABKAMSEQoJcmVzb3VyY2VzGAsgAygJIlEKEUdldFByb2R1Y3RSZXF1ZXN0Eg8KAmlkGAEgASgDSACIAQESFgoJc2VyaWFsX2lkGAIgASgJSAGIAQFCBQoDX2lkQgwKCl9zZXJpYWxfaWQiPQoSR2V0UHJvZHVjdFJlc3BvbnNlEicKBGRhdGEYASABKAsyGS5wcm9kdWN0LnYxLlByb2R1Y3RFbnRpdHkipgQKE0xpc3RQcm9kdWN0c1JlcXVlc3QSLQoKcGFnaW5hdGlvbhgBIAEoCzIZLmNvbW1vbi5QYWdpbmF0aW9uUmVxdWVzdBIdChBwcm9kdWN0X21vZGVsX2lkGAIgASgDSACIAQESGgoNcXVhbnRpdHlfZnJvbRgDIAEoA0gBiAEBEhgKC3F1YW50aXR5X3RvGAQgASgDSAKIAQESFgoJc29sZF9mcm9tGAUgASgDSAOIAQESFAoHc29sZF90bxgGIAEoA0gEiAEBEhsKDmFkZF9wcmljZV9mcm9tGAcgASgDSAWIAQESGQoMYWRkX3ByaWNlX3RvGAggASgDSAaIAQESFgoJaXNfYWN0aXZlGAkgASgISAeIAQESFQoIbWV0YWRhdGEYCiABKAxICIgBARIeChFkYXRlX2NyZWF0ZWRfZnJvbRgLIAEoA0gJiAEBEhwKD2RhdGVfY3JlYXRlZF90bxgMIAEoA0gKiAEBQhMKEV9wcm9kdWN0X21vZGVsX2lkQhAKDl9xdWFudGl0eV9mcm9tQg4KDF9xdWFudGl0eV90b0IMCgpfc29sZF9mcm9tQgoKCF9zb2xkX3RvQhEKD19hZGRfcHJpY2VfZnJvbUIPCg1fYWRkX3ByaWNlX3RvQgwKCl9pc19hY3RpdmVCCwoJX21ldGFkYXRhQhQKEl9kYXRlX2NyZWF0ZWRfZnJvbUISChBfZGF0ZV9jcmVhdGVkX3RvIm8KFExpc3RQcm9kdWN0c1Jlc3BvbnNlEicKBGRhdGEYASADKAsyGS5wcm9kdWN0LnYxLlByb2R1Y3RFbnRpdHkSLgoKcGFnaW5hdGlvbhgCIAEoCzIaLmNvbW1vbi5QYWdpbmF0aW9uUmVzcG9uc2UioAEKFENyZWF0ZVByb2R1Y3RSZXF1ZXN0EhEKCXNlcmlhbF9pZBgBIAEoCRIYChBwcm9kdWN0X21vZGVsX2lkGAIgASgDEhAKCHF1YW50aXR5GAMgASgDEhEKCWFkZF9wcmljZRgEIAEoAxIRCglpc19hY3RpdmUYBSABKAgSEAoIbWV0YWRhdGEYBiABKAwSEQoJcmVzb3VyY2VzGAcgAygJIkAKFUNyZWF0ZVByb2R1Y3RSZXNwb25zZRInCgRkYXRhGAEgASgLMhkucHJvZHVjdC52MS5Qcm9kdWN0RW50aXR5Ir8CChRVcGRhdGVQcm9kdWN0UmVxdWVzdBIKCgJpZBgBIAEoAxIWCglzZXJpYWxfaWQYAiABKAlIAIgBARIdChBwcm9kdWN0X21vZGVsX2lkGAMgASgDSAGIAQESFQoIcXVhbnRpdHkYBCABKANIAogBARIRCgRzb2xkGAUgASgDSAOIAQESFgoJYWRkX3ByaWNlGAYgASgDSASIAQESFgoJaXNfYWN0aXZlGAcgASgISAWIAQESFQoIbWV0YWRhdGEYCCABKAxIBogBARIRCglyZXNvdXJjZXMYCSADKAlCDAoKX3NlcmlhbF9pZEITChFfcHJvZHVjdF9tb2RlbF9pZEILCglfcXVhbnRpdHlCBwoFX3NvbGRCDAoKX2FkZF9wcmljZUIMCgpfaXNfYWN0aXZlQgsKCV9tZXRhZGF0YSIXChVVcGRhdGVQcm9kdWN0UmVzcG9uc2UiVAoURGVsZXRlUHJvZHVjdFJlcXVlc3QSDwoCaWQYASABKANIAIgBARIWCglzZXJpYWxfaWQYAiABKAlIAYgBAUIFCgNfaWRCDAoKX3NlcmlhbF9pZCIXChVEZWxldGVQcm9kdWN0UmVzcG9uc2VCrwEKDmNvbS5wcm9kdWN0LnYxQgxQcm9kdWN0UHJvdG9QAVpGZ2l0aHViLmNvbS9zaG9wbmV4dXMvc2hvcG5leHVzLXByb3RvYnVmLWdlbi1nby9wYi9wcm9kdWN0L3YxO3Byb2R1Y3R2MaICA1BYWKoCClByb2R1Y3QuVjHKAgpQcm9kdWN0XFYx4gIWUHJvZHVjdFxWMVxHUEJNZXRhZGF0YeoCC1Byb2R1Y3Q6OlYxYgZwcm90bzM", [file_common_pagination]);

/**
 * PRODUCT ENTITY
 *
 * @generated from message product.v1.ProductEntity
 */
export type ProductEntity = Message<"product.v1.ProductEntity"> & {
  /**
   * @generated from field: int64 id = 1;
   */
  id: bigint;

  /**
   * @generated from field: string serial_id = 2;
   */
  serialId: string;

  /**
   * @generated from field: int64 product_model_id = 3;
   */
  productModelId: bigint;

  /**
   * @generated from field: int64 quantity = 4;
   */
  quantity: bigint;

  /**
   * @generated from field: int64 sold = 5;
   */
  sold: bigint;

  /**
   * @generated from field: int64 add_price = 6;
   */
  addPrice: bigint;

  /**
   * @generated from field: bool is_active = 7;
   */
  isActive: boolean;

  /**
   * @generated from field: bytes metadata = 8;
   */
  metadata: Uint8Array;

  /**
   * @generated from field: int64 date_created = 9;
   */
  dateCreated: bigint;

  /**
   * @generated from field: int64 date_updated = 10;
   */
  dateUpdated: bigint;

  /**
   * @generated from field: repeated string resources = 11;
   */
  resources: string[];
};

/**
 * Describes the message product.v1.ProductEntity.
 * Use `create(ProductEntitySchema)` to create a new message.
 */
export const ProductEntitySchema: GenMessage<ProductEntity> = /*@__PURE__*/
  messageDesc(file_product_v1_product, 0);

/**
 * @generated from message product.v1.GetProductRequest
 */
export type GetProductRequest = Message<"product.v1.GetProductRequest"> & {
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
 * Describes the message product.v1.GetProductRequest.
 * Use `create(GetProductRequestSchema)` to create a new message.
 */
export const GetProductRequestSchema: GenMessage<GetProductRequest> = /*@__PURE__*/
  messageDesc(file_product_v1_product, 1);

/**
 * @generated from message product.v1.GetProductResponse
 */
export type GetProductResponse = Message<"product.v1.GetProductResponse"> & {
  /**
   * @generated from field: product.v1.ProductEntity data = 1;
   */
  data?: ProductEntity;
};

/**
 * Describes the message product.v1.GetProductResponse.
 * Use `create(GetProductResponseSchema)` to create a new message.
 */
export const GetProductResponseSchema: GenMessage<GetProductResponse> = /*@__PURE__*/
  messageDesc(file_product_v1_product, 2);

/**
 * @generated from message product.v1.ListProductsRequest
 */
export type ListProductsRequest = Message<"product.v1.ListProductsRequest"> & {
  /**
   * @generated from field: common.PaginationRequest pagination = 1;
   */
  pagination?: PaginationRequest;

  /**
   * @generated from field: optional int64 product_model_id = 2;
   */
  productModelId?: bigint;

  /**
   * @generated from field: optional int64 quantity_from = 3;
   */
  quantityFrom?: bigint;

  /**
   * @generated from field: optional int64 quantity_to = 4;
   */
  quantityTo?: bigint;

  /**
   * @generated from field: optional int64 sold_from = 5;
   */
  soldFrom?: bigint;

  /**
   * @generated from field: optional int64 sold_to = 6;
   */
  soldTo?: bigint;

  /**
   * @generated from field: optional int64 add_price_from = 7;
   */
  addPriceFrom?: bigint;

  /**
   * @generated from field: optional int64 add_price_to = 8;
   */
  addPriceTo?: bigint;

  /**
   * @generated from field: optional bool is_active = 9;
   */
  isActive?: boolean;

  /**
   * @generated from field: optional bytes metadata = 10;
   */
  metadata?: Uint8Array;

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
 * Describes the message product.v1.ListProductsRequest.
 * Use `create(ListProductsRequestSchema)` to create a new message.
 */
export const ListProductsRequestSchema: GenMessage<ListProductsRequest> = /*@__PURE__*/
  messageDesc(file_product_v1_product, 3);

/**
 * @generated from message product.v1.ListProductsResponse
 */
export type ListProductsResponse = Message<"product.v1.ListProductsResponse"> & {
  /**
   * @generated from field: repeated product.v1.ProductEntity data = 1;
   */
  data: ProductEntity[];

  /**
   * @generated from field: common.PaginationResponse pagination = 2;
   */
  pagination?: PaginationResponse;
};

/**
 * Describes the message product.v1.ListProductsResponse.
 * Use `create(ListProductsResponseSchema)` to create a new message.
 */
export const ListProductsResponseSchema: GenMessage<ListProductsResponse> = /*@__PURE__*/
  messageDesc(file_product_v1_product, 4);

/**
 * @generated from message product.v1.CreateProductRequest
 */
export type CreateProductRequest = Message<"product.v1.CreateProductRequest"> & {
  /**
   * @generated from field: string serial_id = 1;
   */
  serialId: string;

  /**
   * @generated from field: int64 product_model_id = 2;
   */
  productModelId: bigint;

  /**
   * @generated from field: int64 quantity = 3;
   */
  quantity: bigint;

  /**
   * @generated from field: int64 add_price = 4;
   */
  addPrice: bigint;

  /**
   * @generated from field: bool is_active = 5;
   */
  isActive: boolean;

  /**
   * @generated from field: bytes metadata = 6;
   */
  metadata: Uint8Array;

  /**
   * @generated from field: repeated string resources = 7;
   */
  resources: string[];
};

/**
 * Describes the message product.v1.CreateProductRequest.
 * Use `create(CreateProductRequestSchema)` to create a new message.
 */
export const CreateProductRequestSchema: GenMessage<CreateProductRequest> = /*@__PURE__*/
  messageDesc(file_product_v1_product, 5);

/**
 * @generated from message product.v1.CreateProductResponse
 */
export type CreateProductResponse = Message<"product.v1.CreateProductResponse"> & {
  /**
   * @generated from field: product.v1.ProductEntity data = 1;
   */
  data?: ProductEntity;
};

/**
 * Describes the message product.v1.CreateProductResponse.
 * Use `create(CreateProductResponseSchema)` to create a new message.
 */
export const CreateProductResponseSchema: GenMessage<CreateProductResponse> = /*@__PURE__*/
  messageDesc(file_product_v1_product, 6);

/**
 * @generated from message product.v1.UpdateProductRequest
 */
export type UpdateProductRequest = Message<"product.v1.UpdateProductRequest"> & {
  /**
   * @generated from field: int64 id = 1;
   */
  id: bigint;

  /**
   * @generated from field: optional string serial_id = 2;
   */
  serialId?: string;

  /**
   * @generated from field: optional int64 product_model_id = 3;
   */
  productModelId?: bigint;

  /**
   * @generated from field: optional int64 quantity = 4;
   */
  quantity?: bigint;

  /**
   * @generated from field: optional int64 sold = 5;
   */
  sold?: bigint;

  /**
   * @generated from field: optional int64 add_price = 6;
   */
  addPrice?: bigint;

  /**
   * @generated from field: optional bool is_active = 7;
   */
  isActive?: boolean;

  /**
   * @generated from field: optional bytes metadata = 8;
   */
  metadata?: Uint8Array;

  /**
   * @generated from field: repeated string resources = 9;
   */
  resources: string[];
};

/**
 * Describes the message product.v1.UpdateProductRequest.
 * Use `create(UpdateProductRequestSchema)` to create a new message.
 */
export const UpdateProductRequestSchema: GenMessage<UpdateProductRequest> = /*@__PURE__*/
  messageDesc(file_product_v1_product, 7);

/**
 * @generated from message product.v1.UpdateProductResponse
 */
export type UpdateProductResponse = Message<"product.v1.UpdateProductResponse"> & {
};

/**
 * Describes the message product.v1.UpdateProductResponse.
 * Use `create(UpdateProductResponseSchema)` to create a new message.
 */
export const UpdateProductResponseSchema: GenMessage<UpdateProductResponse> = /*@__PURE__*/
  messageDesc(file_product_v1_product, 8);

/**
 * @generated from message product.v1.DeleteProductRequest
 */
export type DeleteProductRequest = Message<"product.v1.DeleteProductRequest"> & {
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
 * Describes the message product.v1.DeleteProductRequest.
 * Use `create(DeleteProductRequestSchema)` to create a new message.
 */
export const DeleteProductRequestSchema: GenMessage<DeleteProductRequest> = /*@__PURE__*/
  messageDesc(file_product_v1_product, 9);

/**
 * @generated from message product.v1.DeleteProductResponse
 */
export type DeleteProductResponse = Message<"product.v1.DeleteProductResponse"> & {
};

/**
 * Describes the message product.v1.DeleteProductResponse.
 * Use `create(DeleteProductResponseSchema)` to create a new message.
 */
export const DeleteProductResponseSchema: GenMessage<DeleteProductResponse> = /*@__PURE__*/
  messageDesc(file_product_v1_product, 10);

