// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: payment/v1/service.proto

package paymentv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_payment_v1_service_proto protoreflect.FileDescriptor

var file_payment_v1_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x66,
	0x75, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe5, 0x06, 0x0a, 0x0e, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x2e, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x90, 0x02, 0x01, 0x12, 0x56,
	0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1f,
	0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x03, 0x90, 0x02, 0x01, 0x12, 0x56, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56,
	0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x20, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x0d, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x12, 0x1c, 0x2e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x66, 0x75,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x90, 0x02, 0x01, 0x12, 0x53, 0x0a,
	0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x73, 0x12, 0x1e, 0x2e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x66, 0x75, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x66, 0x75, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x90,
	0x02, 0x01, 0x12, 0x53, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x66, 0x75,
	0x6e, 0x64, 0x12, 0x1f, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x12, 0x1f, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x66, 0x75, 0x6e,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x66, 0x75,
	0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0c,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x12, 0x1f, 0x2e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0xaf, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x68, 0x6f, 0x70, 0x6e, 0x65, 0x78, 0x75, 0x73, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x6e,
	0x65, 0x78, 0x75, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x76, 0x31, 0x3b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50,
	0x58, 0x58, 0xaa, 0x02, 0x0a, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x0a, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_payment_v1_service_proto_goTypes = []any{
	(*GetPaymentRequest)(nil),     // 0: payment.v1.GetPaymentRequest
	(*ListPaymentsRequest)(nil),   // 1: payment.v1.ListPaymentsRequest
	(*CreatePaymentRequest)(nil),  // 2: payment.v1.CreatePaymentRequest
	(*UpdatePaymentRequest)(nil),  // 3: payment.v1.UpdatePaymentRequest
	(*CancelPaymentRequest)(nil),  // 4: payment.v1.CancelPaymentRequest
	(*GetRefundRequest)(nil),      // 5: payment.v1.GetRefundRequest
	(*ListRefundsRequest)(nil),    // 6: payment.v1.ListRefundsRequest
	(*CreateRefundRequest)(nil),   // 7: payment.v1.CreateRefundRequest
	(*UpdateRefundRequest)(nil),   // 8: payment.v1.UpdateRefundRequest
	(*CancelRefundRequest)(nil),   // 9: payment.v1.CancelRefundRequest
	(*GetPaymentResponse)(nil),    // 10: payment.v1.GetPaymentResponse
	(*ListPaymentsResponse)(nil),  // 11: payment.v1.ListPaymentsResponse
	(*CreatePaymentResponse)(nil), // 12: payment.v1.CreatePaymentResponse
	(*UpdatePaymentResponse)(nil), // 13: payment.v1.UpdatePaymentResponse
	(*CancelPaymentResponse)(nil), // 14: payment.v1.CancelPaymentResponse
	(*GetRefundResponse)(nil),     // 15: payment.v1.GetRefundResponse
	(*ListRefundsResponse)(nil),   // 16: payment.v1.ListRefundsResponse
	(*CreateRefundResponse)(nil),  // 17: payment.v1.CreateRefundResponse
	(*UpdateRefundResponse)(nil),  // 18: payment.v1.UpdateRefundResponse
	(*CancelRefundResponse)(nil),  // 19: payment.v1.CancelRefundResponse
}
var file_payment_v1_service_proto_depIdxs = []int32{
	0,  // 0: payment.v1.PaymentService.GetPayment:input_type -> payment.v1.GetPaymentRequest
	1,  // 1: payment.v1.PaymentService.ListPayments:input_type -> payment.v1.ListPaymentsRequest
	2,  // 2: payment.v1.PaymentService.CreatePayment:input_type -> payment.v1.CreatePaymentRequest
	3,  // 3: payment.v1.PaymentService.UpdatePayment:input_type -> payment.v1.UpdatePaymentRequest
	4,  // 4: payment.v1.PaymentService.CancelPayment:input_type -> payment.v1.CancelPaymentRequest
	5,  // 5: payment.v1.PaymentService.GetRefund:input_type -> payment.v1.GetRefundRequest
	6,  // 6: payment.v1.PaymentService.ListRefunds:input_type -> payment.v1.ListRefundsRequest
	7,  // 7: payment.v1.PaymentService.CreateRefund:input_type -> payment.v1.CreateRefundRequest
	8,  // 8: payment.v1.PaymentService.UpdateRefund:input_type -> payment.v1.UpdateRefundRequest
	9,  // 9: payment.v1.PaymentService.CancelRefund:input_type -> payment.v1.CancelRefundRequest
	10, // 10: payment.v1.PaymentService.GetPayment:output_type -> payment.v1.GetPaymentResponse
	11, // 11: payment.v1.PaymentService.ListPayments:output_type -> payment.v1.ListPaymentsResponse
	12, // 12: payment.v1.PaymentService.CreatePayment:output_type -> payment.v1.CreatePaymentResponse
	13, // 13: payment.v1.PaymentService.UpdatePayment:output_type -> payment.v1.UpdatePaymentResponse
	14, // 14: payment.v1.PaymentService.CancelPayment:output_type -> payment.v1.CancelPaymentResponse
	15, // 15: payment.v1.PaymentService.GetRefund:output_type -> payment.v1.GetRefundResponse
	16, // 16: payment.v1.PaymentService.ListRefunds:output_type -> payment.v1.ListRefundsResponse
	17, // 17: payment.v1.PaymentService.CreateRefund:output_type -> payment.v1.CreateRefundResponse
	18, // 18: payment.v1.PaymentService.UpdateRefund:output_type -> payment.v1.UpdateRefundResponse
	19, // 19: payment.v1.PaymentService.CancelRefund:output_type -> payment.v1.CancelRefundResponse
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_payment_v1_service_proto_init() }
func file_payment_v1_service_proto_init() {
	if File_payment_v1_service_proto != nil {
		return
	}
	file_payment_v1_payment_proto_init()
	file_payment_v1_refund_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_payment_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_payment_v1_service_proto_goTypes,
		DependencyIndexes: file_payment_v1_service_proto_depIdxs,
	}.Build()
	File_payment_v1_service_proto = out.File
	file_payment_v1_service_proto_rawDesc = nil
	file_payment_v1_service_proto_goTypes = nil
	file_payment_v1_service_proto_depIdxs = nil
}
