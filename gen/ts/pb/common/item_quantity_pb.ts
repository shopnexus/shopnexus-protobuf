// @generated by protoc-gen-es v2.4.0 with parameter "target=ts"
// @generated from file common/item_quantity.proto (package common, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file common/item_quantity.proto.
 */
export const file_common_item_quantity: GenFile = /*@__PURE__*/
  fileDesc("Chpjb21tb24vaXRlbV9xdWFudGl0eS5wcm90bxIGY29tbW9uIjYKEUl0ZW1RdWFudGl0eUludDY0Eg8KB2l0ZW1faWQYASABKAMSEAoIcXVhbnRpdHkYAiABKAMiNwoSSXRlbVF1YW50aXR5U3RyaW5nEg8KB2l0ZW1faWQYASABKAkSEAoIcXVhbnRpdHkYAiABKANCkQEKCmNvbS5jb21tb25CEUl0ZW1RdWFudGl0eVByb3RvUAFaOGdpdGh1Yi5jb20vc2hvcG5leHVzL3Nob3BuZXh1cy1wcm90b2J1Zi1nZW4tZ28vcGIvY29tbW9uogIDQ1hYqgIGQ29tbW9uygIGQ29tbW9u4gISQ29tbW9uXEdQQk1ldGFkYXRh6gIGQ29tbW9uYgZwcm90bzM");

/**
 * @generated from message common.ItemQuantityInt64
 */
export type ItemQuantityInt64 = Message<"common.ItemQuantityInt64"> & {
  /**
   * @generated from field: int64 item_id = 1;
   */
  itemId: bigint;

  /**
   * @generated from field: int64 quantity = 2;
   */
  quantity: bigint;
};

/**
 * Describes the message common.ItemQuantityInt64.
 * Use `create(ItemQuantityInt64Schema)` to create a new message.
 */
export const ItemQuantityInt64Schema: GenMessage<ItemQuantityInt64> = /*@__PURE__*/
  messageDesc(file_common_item_quantity, 0);

/**
 * @generated from message common.ItemQuantityString
 */
export type ItemQuantityString = Message<"common.ItemQuantityString"> & {
  /**
   * @generated from field: string item_id = 1;
   */
  itemId: string;

  /**
   * @generated from field: int64 quantity = 2;
   */
  quantity: bigint;
};

/**
 * Describes the message common.ItemQuantityString.
 * Use `create(ItemQuantityStringSchema)` to create a new message.
 */
export const ItemQuantityStringSchema: GenMessage<ItemQuantityString> = /*@__PURE__*/
  messageDesc(file_common_item_quantity, 1);

