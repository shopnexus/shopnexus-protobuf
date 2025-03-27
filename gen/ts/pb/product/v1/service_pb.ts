// @generated by protoc-gen-es v2.2.3 with parameter "target=ts"
// @generated from file product/v1/service.proto (package product.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import type { CreateBrandRequestSchema, CreateBrandResponseSchema, DeleteBrandRequestSchema, DeleteBrandResponseSchema, GetBrandRequestSchema, GetBrandResponseSchema, ListBrandsRequestSchema, ListBrandsResponseSchema, UpdateBrandRequestSchema, UpdateBrandResponseSchema } from "./brand_pb";
import { file_product_v1_brand } from "./brand_pb";
import type { CreateProductRequestSchema, CreateProductResponseSchema, DeleteProductRequestSchema, DeleteProductResponseSchema, GetProductRequestSchema, GetProductResponseSchema, ListProductsRequestSchema, ListProductsResponseSchema, UpdateProductRequestSchema, UpdateProductResponseSchema } from "./product_pb";
import { file_product_v1_product } from "./product_pb";
import type { CreateProductModelRequestSchema, CreateProductModelResponseSchema, DeleteProductModelRequestSchema, DeleteProductModelResponseSchema, GetProductModelRequestSchema, GetProductModelResponseSchema, ListProductModelsRequestSchema, ListProductModelsResponseSchema, ListProductTypesRequestSchema, ListProductTypesResponseSchema, UpdateProductModelRequestSchema, UpdateProductModelResponseSchema } from "./product_model_pb";
import { file_product_v1_product_model } from "./product_model_pb";
import type { CreateSaleRequestSchema, CreateSaleResponseSchema, DeleteSaleRequestSchema, DeleteSaleResponseSchema, GetSaleRequestSchema, GetSaleResponseSchema, ListSalesRequestSchema, ListSalesResponseSchema, UpdateSaleRequestSchema, UpdateSaleResponseSchema } from "./sale_pb";
import { file_product_v1_sale } from "./sale_pb";
import type { CreateTagRequestSchema, CreateTagResponseSchema, DeleteTagRequestSchema, DeleteTagResponseSchema, GetTagRequestSchema, GetTagResponseSchema, ListTagsRequestSchema, ListTagsResponseSchema, UpdateTagRequestSchema, UpdateTagResponseSchema } from "./tag_pb";
import { file_product_v1_tag } from "./tag_pb";

/**
 * Describes the file product/v1/service.proto.
 */
export const file_product_v1_service: GenFile = /*@__PURE__*/
  fileDesc("Chhwcm9kdWN0L3YxL3NlcnZpY2UucHJvdG8SCnByb2R1Y3QudjEyshEKDlByb2R1Y3RTZXJ2aWNlEl8KD0dldFByb2R1Y3RNb2RlbBIiLnByb2R1Y3QudjEuR2V0UHJvZHVjdE1vZGVsUmVxdWVzdBojLnByb2R1Y3QudjEuR2V0UHJvZHVjdE1vZGVsUmVzcG9uc2UiA5ACARJlChFMaXN0UHJvZHVjdE1vZGVscxIkLnByb2R1Y3QudjEuTGlzdFByb2R1Y3RNb2RlbHNSZXF1ZXN0GiUucHJvZHVjdC52MS5MaXN0UHJvZHVjdE1vZGVsc1Jlc3BvbnNlIgOQAgESZQoSQ3JlYXRlUHJvZHVjdE1vZGVsEiUucHJvZHVjdC52MS5DcmVhdGVQcm9kdWN0TW9kZWxSZXF1ZXN0GiYucHJvZHVjdC52MS5DcmVhdGVQcm9kdWN0TW9kZWxSZXNwb25zZSIAEmUKElVwZGF0ZVByb2R1Y3RNb2RlbBIlLnByb2R1Y3QudjEuVXBkYXRlUHJvZHVjdE1vZGVsUmVxdWVzdBomLnByb2R1Y3QudjEuVXBkYXRlUHJvZHVjdE1vZGVsUmVzcG9uc2UiABJlChJEZWxldGVQcm9kdWN0TW9kZWwSJS5wcm9kdWN0LnYxLkRlbGV0ZVByb2R1Y3RNb2RlbFJlcXVlc3QaJi5wcm9kdWN0LnYxLkRlbGV0ZVByb2R1Y3RNb2RlbFJlc3BvbnNlIgASYgoQTGlzdFByb2R1Y3RUeXBlcxIjLnByb2R1Y3QudjEuTGlzdFByb2R1Y3RUeXBlc1JlcXVlc3QaJC5wcm9kdWN0LnYxLkxpc3RQcm9kdWN0VHlwZXNSZXNwb25zZSIDkAIBElAKCkdldFByb2R1Y3QSHS5wcm9kdWN0LnYxLkdldFByb2R1Y3RSZXF1ZXN0Gh4ucHJvZHVjdC52MS5HZXRQcm9kdWN0UmVzcG9uc2UiA5ACARJWCgxMaXN0UHJvZHVjdHMSHy5wcm9kdWN0LnYxLkxpc3RQcm9kdWN0c1JlcXVlc3QaIC5wcm9kdWN0LnYxLkxpc3RQcm9kdWN0c1Jlc3BvbnNlIgOQAgESVgoNQ3JlYXRlUHJvZHVjdBIgLnByb2R1Y3QudjEuQ3JlYXRlUHJvZHVjdFJlcXVlc3QaIS5wcm9kdWN0LnYxLkNyZWF0ZVByb2R1Y3RSZXNwb25zZSIAElYKDVVwZGF0ZVByb2R1Y3QSIC5wcm9kdWN0LnYxLlVwZGF0ZVByb2R1Y3RSZXF1ZXN0GiEucHJvZHVjdC52MS5VcGRhdGVQcm9kdWN0UmVzcG9uc2UiABJWCg1EZWxldGVQcm9kdWN0EiAucHJvZHVjdC52MS5EZWxldGVQcm9kdWN0UmVxdWVzdBohLnByb2R1Y3QudjEuRGVsZXRlUHJvZHVjdFJlc3BvbnNlIgASSgoIR2V0QnJhbmQSGy5wcm9kdWN0LnYxLkdldEJyYW5kUmVxdWVzdBocLnByb2R1Y3QudjEuR2V0QnJhbmRSZXNwb25zZSIDkAIBElAKCkxpc3RCcmFuZHMSHS5wcm9kdWN0LnYxLkxpc3RCcmFuZHNSZXF1ZXN0Gh4ucHJvZHVjdC52MS5MaXN0QnJhbmRzUmVzcG9uc2UiA5ACARJQCgtDcmVhdGVCcmFuZBIeLnByb2R1Y3QudjEuQ3JlYXRlQnJhbmRSZXF1ZXN0Gh8ucHJvZHVjdC52MS5DcmVhdGVCcmFuZFJlc3BvbnNlIgASUAoLVXBkYXRlQnJhbmQSHi5wcm9kdWN0LnYxLlVwZGF0ZUJyYW5kUmVxdWVzdBofLnByb2R1Y3QudjEuVXBkYXRlQnJhbmRSZXNwb25zZSIAElAKC0RlbGV0ZUJyYW5kEh4ucHJvZHVjdC52MS5EZWxldGVCcmFuZFJlcXVlc3QaHy5wcm9kdWN0LnYxLkRlbGV0ZUJyYW5kUmVzcG9uc2UiABJHCgdHZXRTYWxlEhoucHJvZHVjdC52MS5HZXRTYWxlUmVxdWVzdBobLnByb2R1Y3QudjEuR2V0U2FsZVJlc3BvbnNlIgOQAgESTQoJTGlzdFNhbGVzEhwucHJvZHVjdC52MS5MaXN0U2FsZXNSZXF1ZXN0Gh0ucHJvZHVjdC52MS5MaXN0U2FsZXNSZXNwb25zZSIDkAIBEk0KCkNyZWF0ZVNhbGUSHS5wcm9kdWN0LnYxLkNyZWF0ZVNhbGVSZXF1ZXN0Gh4ucHJvZHVjdC52MS5DcmVhdGVTYWxlUmVzcG9uc2UiABJNCgpVcGRhdGVTYWxlEh0ucHJvZHVjdC52MS5VcGRhdGVTYWxlUmVxdWVzdBoeLnByb2R1Y3QudjEuVXBkYXRlU2FsZVJlc3BvbnNlIgASTQoKRGVsZXRlU2FsZRIdLnByb2R1Y3QudjEuRGVsZXRlU2FsZVJlcXVlc3QaHi5wcm9kdWN0LnYxLkRlbGV0ZVNhbGVSZXNwb25zZSIAEkQKBkdldFRhZxIZLnByb2R1Y3QudjEuR2V0VGFnUmVxdWVzdBoaLnByb2R1Y3QudjEuR2V0VGFnUmVzcG9uc2UiA5ACARJKCghMaXN0VGFncxIbLnByb2R1Y3QudjEuTGlzdFRhZ3NSZXF1ZXN0GhwucHJvZHVjdC52MS5MaXN0VGFnc1Jlc3BvbnNlIgOQAgESSgoJQ3JlYXRlVGFnEhwucHJvZHVjdC52MS5DcmVhdGVUYWdSZXF1ZXN0Gh0ucHJvZHVjdC52MS5DcmVhdGVUYWdSZXNwb25zZSIAEkoKCVVwZGF0ZVRhZxIcLnByb2R1Y3QudjEuVXBkYXRlVGFnUmVxdWVzdBodLnByb2R1Y3QudjEuVXBkYXRlVGFnUmVzcG9uc2UiABJKCglEZWxldGVUYWcSHC5wcm9kdWN0LnYxLkRlbGV0ZVRhZ1JlcXVlc3QaHS5wcm9kdWN0LnYxLkRlbGV0ZVRhZ1Jlc3BvbnNlIgBCrwEKDmNvbS5wcm9kdWN0LnYxQgxTZXJ2aWNlUHJvdG9QAVpGZ2l0aHViLmNvbS9zaG9wbmV4dXMvc2hvcG5leHVzLXByb3RvYnVmLWdlbi1nby9wYi9wcm9kdWN0L3YxO3Byb2R1Y3R2MaICA1BYWKoCClByb2R1Y3QuVjHKAgpQcm9kdWN0XFYx4gIWUHJvZHVjdFxWMVxHUEJNZXRhZGF0YeoCC1Byb2R1Y3Q6OlYxYgZwcm90bzM", [file_product_v1_brand, file_product_v1_product, file_product_v1_product_model, file_product_v1_sale, file_product_v1_tag]);

/**
 * @generated from service product.v1.ProductService
 */
export const ProductService: GenService<{
  /**
   * PRODUCT MODEL
   *
   * @generated from rpc product.v1.ProductService.GetProductModel
   */
  getProductModel: {
    methodKind: "unary";
    input: typeof GetProductModelRequestSchema;
    output: typeof GetProductModelResponseSchema;
  },
  /**
   * @generated from rpc product.v1.ProductService.ListProductModels
   */
  listProductModels: {
    methodKind: "unary";
    input: typeof ListProductModelsRequestSchema;
    output: typeof ListProductModelsResponseSchema;
  },
  /**
   * @generated from rpc product.v1.ProductService.CreateProductModel
   */
  createProductModel: {
    methodKind: "unary";
    input: typeof CreateProductModelRequestSchema;
    output: typeof CreateProductModelResponseSchema;
  },
  /**
   * @generated from rpc product.v1.ProductService.UpdateProductModel
   */
  updateProductModel: {
    methodKind: "unary";
    input: typeof UpdateProductModelRequestSchema;
    output: typeof UpdateProductModelResponseSchema;
  },
  /**
   * @generated from rpc product.v1.ProductService.DeleteProductModel
   */
  deleteProductModel: {
    methodKind: "unary";
    input: typeof DeleteProductModelRequestSchema;
    output: typeof DeleteProductModelResponseSchema;
  },
  /**
   * PRODUCT TYPES
   *
   * @generated from rpc product.v1.ProductService.ListProductTypes
   */
  listProductTypes: {
    methodKind: "unary";
    input: typeof ListProductTypesRequestSchema;
    output: typeof ListProductTypesResponseSchema;
  },
  /**
   * PRODUCT
   *
   * @generated from rpc product.v1.ProductService.GetProduct
   */
  getProduct: {
    methodKind: "unary";
    input: typeof GetProductRequestSchema;
    output: typeof GetProductResponseSchema;
  },
  /**
   * @generated from rpc product.v1.ProductService.ListProducts
   */
  listProducts: {
    methodKind: "unary";
    input: typeof ListProductsRequestSchema;
    output: typeof ListProductsResponseSchema;
  },
  /**
   * @generated from rpc product.v1.ProductService.CreateProduct
   */
  createProduct: {
    methodKind: "unary";
    input: typeof CreateProductRequestSchema;
    output: typeof CreateProductResponseSchema;
  },
  /**
   * @generated from rpc product.v1.ProductService.UpdateProduct
   */
  updateProduct: {
    methodKind: "unary";
    input: typeof UpdateProductRequestSchema;
    output: typeof UpdateProductResponseSchema;
  },
  /**
   * @generated from rpc product.v1.ProductService.DeleteProduct
   */
  deleteProduct: {
    methodKind: "unary";
    input: typeof DeleteProductRequestSchema;
    output: typeof DeleteProductResponseSchema;
  },
  /**
   * BRAND
   *
   * @generated from rpc product.v1.ProductService.GetBrand
   */
  getBrand: {
    methodKind: "unary";
    input: typeof GetBrandRequestSchema;
    output: typeof GetBrandResponseSchema;
  },
  /**
   * @generated from rpc product.v1.ProductService.ListBrands
   */
  listBrands: {
    methodKind: "unary";
    input: typeof ListBrandsRequestSchema;
    output: typeof ListBrandsResponseSchema;
  },
  /**
   * @generated from rpc product.v1.ProductService.CreateBrand
   */
  createBrand: {
    methodKind: "unary";
    input: typeof CreateBrandRequestSchema;
    output: typeof CreateBrandResponseSchema;
  },
  /**
   * @generated from rpc product.v1.ProductService.UpdateBrand
   */
  updateBrand: {
    methodKind: "unary";
    input: typeof UpdateBrandRequestSchema;
    output: typeof UpdateBrandResponseSchema;
  },
  /**
   * @generated from rpc product.v1.ProductService.DeleteBrand
   */
  deleteBrand: {
    methodKind: "unary";
    input: typeof DeleteBrandRequestSchema;
    output: typeof DeleteBrandResponseSchema;
  },
  /**
   * SALE
   *
   * @generated from rpc product.v1.ProductService.GetSale
   */
  getSale: {
    methodKind: "unary";
    input: typeof GetSaleRequestSchema;
    output: typeof GetSaleResponseSchema;
  },
  /**
   * @generated from rpc product.v1.ProductService.ListSales
   */
  listSales: {
    methodKind: "unary";
    input: typeof ListSalesRequestSchema;
    output: typeof ListSalesResponseSchema;
  },
  /**
   * @generated from rpc product.v1.ProductService.CreateSale
   */
  createSale: {
    methodKind: "unary";
    input: typeof CreateSaleRequestSchema;
    output: typeof CreateSaleResponseSchema;
  },
  /**
   * @generated from rpc product.v1.ProductService.UpdateSale
   */
  updateSale: {
    methodKind: "unary";
    input: typeof UpdateSaleRequestSchema;
    output: typeof UpdateSaleResponseSchema;
  },
  /**
   * @generated from rpc product.v1.ProductService.DeleteSale
   */
  deleteSale: {
    methodKind: "unary";
    input: typeof DeleteSaleRequestSchema;
    output: typeof DeleteSaleResponseSchema;
  },
  /**
   * TAG
   *
   * @generated from rpc product.v1.ProductService.GetTag
   */
  getTag: {
    methodKind: "unary";
    input: typeof GetTagRequestSchema;
    output: typeof GetTagResponseSchema;
  },
  /**
   * @generated from rpc product.v1.ProductService.ListTags
   */
  listTags: {
    methodKind: "unary";
    input: typeof ListTagsRequestSchema;
    output: typeof ListTagsResponseSchema;
  },
  /**
   * @generated from rpc product.v1.ProductService.CreateTag
   */
  createTag: {
    methodKind: "unary";
    input: typeof CreateTagRequestSchema;
    output: typeof CreateTagResponseSchema;
  },
  /**
   * @generated from rpc product.v1.ProductService.UpdateTag
   */
  updateTag: {
    methodKind: "unary";
    input: typeof UpdateTagRequestSchema;
    output: typeof UpdateTagResponseSchema;
  },
  /**
   * @generated from rpc product.v1.ProductService.DeleteTag
   */
  deleteTag: {
    methodKind: "unary";
    input: typeof DeleteTagRequestSchema;
    output: typeof DeleteTagResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_product_v1_service, 0);

