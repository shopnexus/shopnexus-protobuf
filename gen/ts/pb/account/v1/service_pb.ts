// @generated by protoc-gen-es v2.2.3 with parameter "target=ts"
// @generated from file account/v1/service.proto (package account.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import type { GetUserPublicRequestSchema, GetUserPublicResponseSchema, GetUserRequestSchema, GetUserResponseSchema, LoginAdminRequestSchema, LoginAdminResponseSchema, LoginUserRequestSchema, LoginUserResponseSchema, RegisterAdminRequestSchema, RegisterAdminResponseSchema, RegisterUserRequestSchema, RegisterUserResponseSchema, UpdateAccountRequestSchema, UpdateAccountResponseSchema, UpdateUserRequestSchema, UpdateUserResponseSchema } from "./account_pb";
import { file_account_v1_account } from "./account_pb";
import type { AddCartItemRequestSchema, AddCartItemResponseSchema, ClearCartRequestSchema, ClearCartResponseSchema, GetCartRequestSchema, GetCartResponseSchema, UpdateCartItemRequestSchema, UpdateCartItemResponseSchema } from "./cart_pb";
import { file_account_v1_cart } from "./cart_pb";
import type { CreateAddressRequestSchema, CreateAddressResponseSchema, DeleteAddressRequestSchema, DeleteAddressResponseSchema, GetAddressRequestSchema, GetAddressResponseSchema, ListAddressesRequestSchema, ListAddressesResponseSchema, UpdateAddressRequestSchema, UpdateAddressResponseSchema } from "./address_pb";
import { file_account_v1_address } from "./address_pb";

/**
 * Describes the file account/v1/service.proto.
 */
export const file_account_v1_service: GenFile = /*@__PURE__*/
  fileDesc("ChhhY2NvdW50L3YxL3NlcnZpY2UucHJvdG8SCmFjY291bnQudjEymgsKDkFjY291bnRTZXJ2aWNlEkcKB0dldFVzZXISGi5hY2NvdW50LnYxLkdldFVzZXJSZXF1ZXN0GhsuYWNjb3VudC52MS5HZXRVc2VyUmVzcG9uc2UiA5ACARJZCg1HZXRVc2VyUHVibGljEiAuYWNjb3VudC52MS5HZXRVc2VyUHVibGljUmVxdWVzdBohLmFjY291bnQudjEuR2V0VXNlclB1YmxpY1Jlc3BvbnNlIgOQAgESVgoNVXBkYXRlQWNjb3VudBIgLmFjY291bnQudjEuVXBkYXRlQWNjb3VudFJlcXVlc3QaIS5hY2NvdW50LnYxLlVwZGF0ZUFjY291bnRSZXNwb25zZSIAEk0KClVwZGF0ZVVzZXISHS5hY2NvdW50LnYxLlVwZGF0ZVVzZXJSZXF1ZXN0Gh4uYWNjb3VudC52MS5VcGRhdGVVc2VyUmVzcG9uc2UiABJKCglMb2dpblVzZXISHC5hY2NvdW50LnYxLkxvZ2luVXNlclJlcXVlc3QaHS5hY2NvdW50LnYxLkxvZ2luVXNlclJlc3BvbnNlIgASTQoKTG9naW5BZG1pbhIdLmFjY291bnQudjEuTG9naW5BZG1pblJlcXVlc3QaHi5hY2NvdW50LnYxLkxvZ2luQWRtaW5SZXNwb25zZSIAElMKDFJlZ2lzdGVyVXNlchIfLmFjY291bnQudjEuUmVnaXN0ZXJVc2VyUmVxdWVzdBogLmFjY291bnQudjEuUmVnaXN0ZXJVc2VyUmVzcG9uc2UiABJWCg1SZWdpc3RlckFkbWluEiAuYWNjb3VudC52MS5SZWdpc3RlckFkbWluUmVxdWVzdBohLmFjY291bnQudjEuUmVnaXN0ZXJBZG1pblJlc3BvbnNlIgASRwoHR2V0Q2FydBIaLmFjY291bnQudjEuR2V0Q2FydFJlcXVlc3QaGy5hY2NvdW50LnYxLkdldENhcnRSZXNwb25zZSIDkAIBElAKC0FkZENhcnRJdGVtEh4uYWNjb3VudC52MS5BZGRDYXJ0SXRlbVJlcXVlc3QaHy5hY2NvdW50LnYxLkFkZENhcnRJdGVtUmVzcG9uc2UiABJZCg5VcGRhdGVDYXJ0SXRlbRIhLmFjY291bnQudjEuVXBkYXRlQ2FydEl0ZW1SZXF1ZXN0GiIuYWNjb3VudC52MS5VcGRhdGVDYXJ0SXRlbVJlc3BvbnNlIgASSgoJQ2xlYXJDYXJ0EhwuYWNjb3VudC52MS5DbGVhckNhcnRSZXF1ZXN0Gh0uYWNjb3VudC52MS5DbGVhckNhcnRSZXNwb25zZSIAElAKCkdldEFkZHJlc3MSHS5hY2NvdW50LnYxLkdldEFkZHJlc3NSZXF1ZXN0Gh4uYWNjb3VudC52MS5HZXRBZGRyZXNzUmVzcG9uc2UiA5ACARJZCg1MaXN0QWRkcmVzc2VzEiAuYWNjb3VudC52MS5MaXN0QWRkcmVzc2VzUmVxdWVzdBohLmFjY291bnQudjEuTGlzdEFkZHJlc3Nlc1Jlc3BvbnNlIgOQAgESVgoNQ3JlYXRlQWRkcmVzcxIgLmFjY291bnQudjEuQ3JlYXRlQWRkcmVzc1JlcXVlc3QaIS5hY2NvdW50LnYxLkNyZWF0ZUFkZHJlc3NSZXNwb25zZSIAElYKDVVwZGF0ZUFkZHJlc3MSIC5hY2NvdW50LnYxLlVwZGF0ZUFkZHJlc3NSZXF1ZXN0GiEuYWNjb3VudC52MS5VcGRhdGVBZGRyZXNzUmVzcG9uc2UiABJWCg1EZWxldGVBZGRyZXNzEiAuYWNjb3VudC52MS5EZWxldGVBZGRyZXNzUmVxdWVzdBohLmFjY291bnQudjEuRGVsZXRlQWRkcmVzc1Jlc3BvbnNlIgBCrwEKDmNvbS5hY2NvdW50LnYxQgxTZXJ2aWNlUHJvdG9QAVpGZ2l0aHViLmNvbS9zaG9wbmV4dXMvc2hvcG5leHVzLXByb3RvYnVmLWdlbi1nby9wYi9hY2NvdW50L3YxO2FjY291bnR2MaICA0FYWKoCCkFjY291bnQuVjHKAgpBY2NvdW50XFYx4gIWQWNjb3VudFxWMVxHUEJNZXRhZGF0YeoCC0FjY291bnQ6OlYxYgZwcm90bzM", [file_account_v1_account, file_account_v1_cart, file_account_v1_address]);

/**
 * @generated from service account.v1.AccountService
 */
export const AccountService: GenService<{
  /**
   * @generated from rpc account.v1.AccountService.GetUser
   */
  getUser: {
    methodKind: "unary";
    input: typeof GetUserRequestSchema;
    output: typeof GetUserResponseSchema;
  },
  /**
   * @generated from rpc account.v1.AccountService.GetUserPublic
   */
  getUserPublic: {
    methodKind: "unary";
    input: typeof GetUserPublicRequestSchema;
    output: typeof GetUserPublicResponseSchema;
  },
  /**
   * @generated from rpc account.v1.AccountService.UpdateAccount
   */
  updateAccount: {
    methodKind: "unary";
    input: typeof UpdateAccountRequestSchema;
    output: typeof UpdateAccountResponseSchema;
  },
  /**
   * @generated from rpc account.v1.AccountService.UpdateUser
   */
  updateUser: {
    methodKind: "unary";
    input: typeof UpdateUserRequestSchema;
    output: typeof UpdateUserResponseSchema;
  },
  /**
   * Login, register
   *
   * @generated from rpc account.v1.AccountService.LoginUser
   */
  loginUser: {
    methodKind: "unary";
    input: typeof LoginUserRequestSchema;
    output: typeof LoginUserResponseSchema;
  },
  /**
   * @generated from rpc account.v1.AccountService.LoginAdmin
   */
  loginAdmin: {
    methodKind: "unary";
    input: typeof LoginAdminRequestSchema;
    output: typeof LoginAdminResponseSchema;
  },
  /**
   * @generated from rpc account.v1.AccountService.RegisterUser
   */
  registerUser: {
    methodKind: "unary";
    input: typeof RegisterUserRequestSchema;
    output: typeof RegisterUserResponseSchema;
  },
  /**
   * @generated from rpc account.v1.AccountService.RegisterAdmin
   */
  registerAdmin: {
    methodKind: "unary";
    input: typeof RegisterAdminRequestSchema;
    output: typeof RegisterAdminResponseSchema;
  },
  /**
   * Cart
   *
   * @generated from rpc account.v1.AccountService.GetCart
   */
  getCart: {
    methodKind: "unary";
    input: typeof GetCartRequestSchema;
    output: typeof GetCartResponseSchema;
  },
  /**
   * @generated from rpc account.v1.AccountService.AddCartItem
   */
  addCartItem: {
    methodKind: "unary";
    input: typeof AddCartItemRequestSchema;
    output: typeof AddCartItemResponseSchema;
  },
  /**
   * @generated from rpc account.v1.AccountService.UpdateCartItem
   */
  updateCartItem: {
    methodKind: "unary";
    input: typeof UpdateCartItemRequestSchema;
    output: typeof UpdateCartItemResponseSchema;
  },
  /**
   * @generated from rpc account.v1.AccountService.ClearCart
   */
  clearCart: {
    methodKind: "unary";
    input: typeof ClearCartRequestSchema;
    output: typeof ClearCartResponseSchema;
  },
  /**
   * Address
   *
   * @generated from rpc account.v1.AccountService.GetAddress
   */
  getAddress: {
    methodKind: "unary";
    input: typeof GetAddressRequestSchema;
    output: typeof GetAddressResponseSchema;
  },
  /**
   * @generated from rpc account.v1.AccountService.ListAddresses
   */
  listAddresses: {
    methodKind: "unary";
    input: typeof ListAddressesRequestSchema;
    output: typeof ListAddressesResponseSchema;
  },
  /**
   * @generated from rpc account.v1.AccountService.CreateAddress
   */
  createAddress: {
    methodKind: "unary";
    input: typeof CreateAddressRequestSchema;
    output: typeof CreateAddressResponseSchema;
  },
  /**
   * @generated from rpc account.v1.AccountService.UpdateAddress
   */
  updateAddress: {
    methodKind: "unary";
    input: typeof UpdateAddressRequestSchema;
    output: typeof UpdateAddressResponseSchema;
  },
  /**
   * @generated from rpc account.v1.AccountService.DeleteAddress
   */
  deleteAddress: {
    methodKind: "unary";
    input: typeof DeleteAddressRequestSchema;
    output: typeof DeleteAddressResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_account_v1_service, 0);

