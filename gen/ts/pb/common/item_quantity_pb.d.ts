// @generated by protoc-gen-es v2.2.3
// @generated from file common/item_quantity.proto (package common, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file common/item_quantity.proto.
 */
export declare const file_common_item_quantity: GenFile;

/**
 * @generated from message common.ItemQuantityInt64
 */
export declare type ItemQuantityInt64 = Message<"common.ItemQuantityInt64"> & {
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
export declare const ItemQuantityInt64Schema: GenMessage<ItemQuantityInt64>;

/**
 * @generated from message common.ItemQuantityString
 */
export declare type ItemQuantityString = Message<"common.ItemQuantityString"> & {
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
export declare const ItemQuantityStringSchema: GenMessage<ItemQuantityString>;

