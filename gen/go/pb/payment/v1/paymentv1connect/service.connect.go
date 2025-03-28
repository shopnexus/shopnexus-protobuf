// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: payment/v1/service.proto

package paymentv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/shopnexus/shopnexus-protobuf-gen-go/pb/payment/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// PaymentServiceName is the fully-qualified name of the PaymentService service.
	PaymentServiceName = "payment.v1.PaymentService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// PaymentServiceGetPaymentProcedure is the fully-qualified name of the PaymentService's GetPayment
	// RPC.
	PaymentServiceGetPaymentProcedure = "/payment.v1.PaymentService/GetPayment"
	// PaymentServiceListPaymentsProcedure is the fully-qualified name of the PaymentService's
	// ListPayments RPC.
	PaymentServiceListPaymentsProcedure = "/payment.v1.PaymentService/ListPayments"
	// PaymentServiceCreatePaymentProcedure is the fully-qualified name of the PaymentService's
	// CreatePayment RPC.
	PaymentServiceCreatePaymentProcedure = "/payment.v1.PaymentService/CreatePayment"
	// PaymentServiceUpdatePaymentProcedure is the fully-qualified name of the PaymentService's
	// UpdatePayment RPC.
	PaymentServiceUpdatePaymentProcedure = "/payment.v1.PaymentService/UpdatePayment"
	// PaymentServiceCancelPaymentProcedure is the fully-qualified name of the PaymentService's
	// CancelPayment RPC.
	PaymentServiceCancelPaymentProcedure = "/payment.v1.PaymentService/CancelPayment"
	// PaymentServiceGetRefundProcedure is the fully-qualified name of the PaymentService's GetRefund
	// RPC.
	PaymentServiceGetRefundProcedure = "/payment.v1.PaymentService/GetRefund"
	// PaymentServiceListRefundsProcedure is the fully-qualified name of the PaymentService's
	// ListRefunds RPC.
	PaymentServiceListRefundsProcedure = "/payment.v1.PaymentService/ListRefunds"
	// PaymentServiceCreateRefundProcedure is the fully-qualified name of the PaymentService's
	// CreateRefund RPC.
	PaymentServiceCreateRefundProcedure = "/payment.v1.PaymentService/CreateRefund"
	// PaymentServiceUpdateRefundProcedure is the fully-qualified name of the PaymentService's
	// UpdateRefund RPC.
	PaymentServiceUpdateRefundProcedure = "/payment.v1.PaymentService/UpdateRefund"
	// PaymentServiceCancelRefundProcedure is the fully-qualified name of the PaymentService's
	// CancelRefund RPC.
	PaymentServiceCancelRefundProcedure = "/payment.v1.PaymentService/CancelRefund"
)

// PaymentServiceClient is a client for the payment.v1.PaymentService service.
type PaymentServiceClient interface {
	GetPayment(context.Context, *connect.Request[v1.GetPaymentRequest]) (*connect.Response[v1.GetPaymentResponse], error)
	ListPayments(context.Context, *connect.Request[v1.ListPaymentsRequest]) (*connect.Response[v1.ListPaymentsResponse], error)
	CreatePayment(context.Context, *connect.Request[v1.CreatePaymentRequest]) (*connect.Response[v1.CreatePaymentResponse], error)
	UpdatePayment(context.Context, *connect.Request[v1.UpdatePaymentRequest]) (*connect.Response[v1.UpdatePaymentResponse], error)
	CancelPayment(context.Context, *connect.Request[v1.CancelPaymentRequest]) (*connect.Response[v1.CancelPaymentResponse], error)
	GetRefund(context.Context, *connect.Request[v1.GetRefundRequest]) (*connect.Response[v1.GetRefundResponse], error)
	ListRefunds(context.Context, *connect.Request[v1.ListRefundsRequest]) (*connect.Response[v1.ListRefundsResponse], error)
	CreateRefund(context.Context, *connect.Request[v1.CreateRefundRequest]) (*connect.Response[v1.CreateRefundResponse], error)
	UpdateRefund(context.Context, *connect.Request[v1.UpdateRefundRequest]) (*connect.Response[v1.UpdateRefundResponse], error)
	CancelRefund(context.Context, *connect.Request[v1.CancelRefundRequest]) (*connect.Response[v1.CancelRefundResponse], error)
}

// NewPaymentServiceClient constructs a client for the payment.v1.PaymentService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPaymentServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) PaymentServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	paymentServiceMethods := v1.File_payment_v1_service_proto.Services().ByName("PaymentService").Methods()
	return &paymentServiceClient{
		getPayment: connect.NewClient[v1.GetPaymentRequest, v1.GetPaymentResponse](
			httpClient,
			baseURL+PaymentServiceGetPaymentProcedure,
			connect.WithSchema(paymentServiceMethods.ByName("GetPayment")),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		listPayments: connect.NewClient[v1.ListPaymentsRequest, v1.ListPaymentsResponse](
			httpClient,
			baseURL+PaymentServiceListPaymentsProcedure,
			connect.WithSchema(paymentServiceMethods.ByName("ListPayments")),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		createPayment: connect.NewClient[v1.CreatePaymentRequest, v1.CreatePaymentResponse](
			httpClient,
			baseURL+PaymentServiceCreatePaymentProcedure,
			connect.WithSchema(paymentServiceMethods.ByName("CreatePayment")),
			connect.WithClientOptions(opts...),
		),
		updatePayment: connect.NewClient[v1.UpdatePaymentRequest, v1.UpdatePaymentResponse](
			httpClient,
			baseURL+PaymentServiceUpdatePaymentProcedure,
			connect.WithSchema(paymentServiceMethods.ByName("UpdatePayment")),
			connect.WithClientOptions(opts...),
		),
		cancelPayment: connect.NewClient[v1.CancelPaymentRequest, v1.CancelPaymentResponse](
			httpClient,
			baseURL+PaymentServiceCancelPaymentProcedure,
			connect.WithSchema(paymentServiceMethods.ByName("CancelPayment")),
			connect.WithClientOptions(opts...),
		),
		getRefund: connect.NewClient[v1.GetRefundRequest, v1.GetRefundResponse](
			httpClient,
			baseURL+PaymentServiceGetRefundProcedure,
			connect.WithSchema(paymentServiceMethods.ByName("GetRefund")),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		listRefunds: connect.NewClient[v1.ListRefundsRequest, v1.ListRefundsResponse](
			httpClient,
			baseURL+PaymentServiceListRefundsProcedure,
			connect.WithSchema(paymentServiceMethods.ByName("ListRefunds")),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		createRefund: connect.NewClient[v1.CreateRefundRequest, v1.CreateRefundResponse](
			httpClient,
			baseURL+PaymentServiceCreateRefundProcedure,
			connect.WithSchema(paymentServiceMethods.ByName("CreateRefund")),
			connect.WithClientOptions(opts...),
		),
		updateRefund: connect.NewClient[v1.UpdateRefundRequest, v1.UpdateRefundResponse](
			httpClient,
			baseURL+PaymentServiceUpdateRefundProcedure,
			connect.WithSchema(paymentServiceMethods.ByName("UpdateRefund")),
			connect.WithClientOptions(opts...),
		),
		cancelRefund: connect.NewClient[v1.CancelRefundRequest, v1.CancelRefundResponse](
			httpClient,
			baseURL+PaymentServiceCancelRefundProcedure,
			connect.WithSchema(paymentServiceMethods.ByName("CancelRefund")),
			connect.WithClientOptions(opts...),
		),
	}
}

// paymentServiceClient implements PaymentServiceClient.
type paymentServiceClient struct {
	getPayment    *connect.Client[v1.GetPaymentRequest, v1.GetPaymentResponse]
	listPayments  *connect.Client[v1.ListPaymentsRequest, v1.ListPaymentsResponse]
	createPayment *connect.Client[v1.CreatePaymentRequest, v1.CreatePaymentResponse]
	updatePayment *connect.Client[v1.UpdatePaymentRequest, v1.UpdatePaymentResponse]
	cancelPayment *connect.Client[v1.CancelPaymentRequest, v1.CancelPaymentResponse]
	getRefund     *connect.Client[v1.GetRefundRequest, v1.GetRefundResponse]
	listRefunds   *connect.Client[v1.ListRefundsRequest, v1.ListRefundsResponse]
	createRefund  *connect.Client[v1.CreateRefundRequest, v1.CreateRefundResponse]
	updateRefund  *connect.Client[v1.UpdateRefundRequest, v1.UpdateRefundResponse]
	cancelRefund  *connect.Client[v1.CancelRefundRequest, v1.CancelRefundResponse]
}

// GetPayment calls payment.v1.PaymentService.GetPayment.
func (c *paymentServiceClient) GetPayment(ctx context.Context, req *connect.Request[v1.GetPaymentRequest]) (*connect.Response[v1.GetPaymentResponse], error) {
	return c.getPayment.CallUnary(ctx, req)
}

// ListPayments calls payment.v1.PaymentService.ListPayments.
func (c *paymentServiceClient) ListPayments(ctx context.Context, req *connect.Request[v1.ListPaymentsRequest]) (*connect.Response[v1.ListPaymentsResponse], error) {
	return c.listPayments.CallUnary(ctx, req)
}

// CreatePayment calls payment.v1.PaymentService.CreatePayment.
func (c *paymentServiceClient) CreatePayment(ctx context.Context, req *connect.Request[v1.CreatePaymentRequest]) (*connect.Response[v1.CreatePaymentResponse], error) {
	return c.createPayment.CallUnary(ctx, req)
}

// UpdatePayment calls payment.v1.PaymentService.UpdatePayment.
func (c *paymentServiceClient) UpdatePayment(ctx context.Context, req *connect.Request[v1.UpdatePaymentRequest]) (*connect.Response[v1.UpdatePaymentResponse], error) {
	return c.updatePayment.CallUnary(ctx, req)
}

// CancelPayment calls payment.v1.PaymentService.CancelPayment.
func (c *paymentServiceClient) CancelPayment(ctx context.Context, req *connect.Request[v1.CancelPaymentRequest]) (*connect.Response[v1.CancelPaymentResponse], error) {
	return c.cancelPayment.CallUnary(ctx, req)
}

// GetRefund calls payment.v1.PaymentService.GetRefund.
func (c *paymentServiceClient) GetRefund(ctx context.Context, req *connect.Request[v1.GetRefundRequest]) (*connect.Response[v1.GetRefundResponse], error) {
	return c.getRefund.CallUnary(ctx, req)
}

// ListRefunds calls payment.v1.PaymentService.ListRefunds.
func (c *paymentServiceClient) ListRefunds(ctx context.Context, req *connect.Request[v1.ListRefundsRequest]) (*connect.Response[v1.ListRefundsResponse], error) {
	return c.listRefunds.CallUnary(ctx, req)
}

// CreateRefund calls payment.v1.PaymentService.CreateRefund.
func (c *paymentServiceClient) CreateRefund(ctx context.Context, req *connect.Request[v1.CreateRefundRequest]) (*connect.Response[v1.CreateRefundResponse], error) {
	return c.createRefund.CallUnary(ctx, req)
}

// UpdateRefund calls payment.v1.PaymentService.UpdateRefund.
func (c *paymentServiceClient) UpdateRefund(ctx context.Context, req *connect.Request[v1.UpdateRefundRequest]) (*connect.Response[v1.UpdateRefundResponse], error) {
	return c.updateRefund.CallUnary(ctx, req)
}

// CancelRefund calls payment.v1.PaymentService.CancelRefund.
func (c *paymentServiceClient) CancelRefund(ctx context.Context, req *connect.Request[v1.CancelRefundRequest]) (*connect.Response[v1.CancelRefundResponse], error) {
	return c.cancelRefund.CallUnary(ctx, req)
}

// PaymentServiceHandler is an implementation of the payment.v1.PaymentService service.
type PaymentServiceHandler interface {
	GetPayment(context.Context, *connect.Request[v1.GetPaymentRequest]) (*connect.Response[v1.GetPaymentResponse], error)
	ListPayments(context.Context, *connect.Request[v1.ListPaymentsRequest]) (*connect.Response[v1.ListPaymentsResponse], error)
	CreatePayment(context.Context, *connect.Request[v1.CreatePaymentRequest]) (*connect.Response[v1.CreatePaymentResponse], error)
	UpdatePayment(context.Context, *connect.Request[v1.UpdatePaymentRequest]) (*connect.Response[v1.UpdatePaymentResponse], error)
	CancelPayment(context.Context, *connect.Request[v1.CancelPaymentRequest]) (*connect.Response[v1.CancelPaymentResponse], error)
	GetRefund(context.Context, *connect.Request[v1.GetRefundRequest]) (*connect.Response[v1.GetRefundResponse], error)
	ListRefunds(context.Context, *connect.Request[v1.ListRefundsRequest]) (*connect.Response[v1.ListRefundsResponse], error)
	CreateRefund(context.Context, *connect.Request[v1.CreateRefundRequest]) (*connect.Response[v1.CreateRefundResponse], error)
	UpdateRefund(context.Context, *connect.Request[v1.UpdateRefundRequest]) (*connect.Response[v1.UpdateRefundResponse], error)
	CancelRefund(context.Context, *connect.Request[v1.CancelRefundRequest]) (*connect.Response[v1.CancelRefundResponse], error)
}

// NewPaymentServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPaymentServiceHandler(svc PaymentServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	paymentServiceMethods := v1.File_payment_v1_service_proto.Services().ByName("PaymentService").Methods()
	paymentServiceGetPaymentHandler := connect.NewUnaryHandler(
		PaymentServiceGetPaymentProcedure,
		svc.GetPayment,
		connect.WithSchema(paymentServiceMethods.ByName("GetPayment")),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	paymentServiceListPaymentsHandler := connect.NewUnaryHandler(
		PaymentServiceListPaymentsProcedure,
		svc.ListPayments,
		connect.WithSchema(paymentServiceMethods.ByName("ListPayments")),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	paymentServiceCreatePaymentHandler := connect.NewUnaryHandler(
		PaymentServiceCreatePaymentProcedure,
		svc.CreatePayment,
		connect.WithSchema(paymentServiceMethods.ByName("CreatePayment")),
		connect.WithHandlerOptions(opts...),
	)
	paymentServiceUpdatePaymentHandler := connect.NewUnaryHandler(
		PaymentServiceUpdatePaymentProcedure,
		svc.UpdatePayment,
		connect.WithSchema(paymentServiceMethods.ByName("UpdatePayment")),
		connect.WithHandlerOptions(opts...),
	)
	paymentServiceCancelPaymentHandler := connect.NewUnaryHandler(
		PaymentServiceCancelPaymentProcedure,
		svc.CancelPayment,
		connect.WithSchema(paymentServiceMethods.ByName("CancelPayment")),
		connect.WithHandlerOptions(opts...),
	)
	paymentServiceGetRefundHandler := connect.NewUnaryHandler(
		PaymentServiceGetRefundProcedure,
		svc.GetRefund,
		connect.WithSchema(paymentServiceMethods.ByName("GetRefund")),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	paymentServiceListRefundsHandler := connect.NewUnaryHandler(
		PaymentServiceListRefundsProcedure,
		svc.ListRefunds,
		connect.WithSchema(paymentServiceMethods.ByName("ListRefunds")),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	paymentServiceCreateRefundHandler := connect.NewUnaryHandler(
		PaymentServiceCreateRefundProcedure,
		svc.CreateRefund,
		connect.WithSchema(paymentServiceMethods.ByName("CreateRefund")),
		connect.WithHandlerOptions(opts...),
	)
	paymentServiceUpdateRefundHandler := connect.NewUnaryHandler(
		PaymentServiceUpdateRefundProcedure,
		svc.UpdateRefund,
		connect.WithSchema(paymentServiceMethods.ByName("UpdateRefund")),
		connect.WithHandlerOptions(opts...),
	)
	paymentServiceCancelRefundHandler := connect.NewUnaryHandler(
		PaymentServiceCancelRefundProcedure,
		svc.CancelRefund,
		connect.WithSchema(paymentServiceMethods.ByName("CancelRefund")),
		connect.WithHandlerOptions(opts...),
	)
	return "/payment.v1.PaymentService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case PaymentServiceGetPaymentProcedure:
			paymentServiceGetPaymentHandler.ServeHTTP(w, r)
		case PaymentServiceListPaymentsProcedure:
			paymentServiceListPaymentsHandler.ServeHTTP(w, r)
		case PaymentServiceCreatePaymentProcedure:
			paymentServiceCreatePaymentHandler.ServeHTTP(w, r)
		case PaymentServiceUpdatePaymentProcedure:
			paymentServiceUpdatePaymentHandler.ServeHTTP(w, r)
		case PaymentServiceCancelPaymentProcedure:
			paymentServiceCancelPaymentHandler.ServeHTTP(w, r)
		case PaymentServiceGetRefundProcedure:
			paymentServiceGetRefundHandler.ServeHTTP(w, r)
		case PaymentServiceListRefundsProcedure:
			paymentServiceListRefundsHandler.ServeHTTP(w, r)
		case PaymentServiceCreateRefundProcedure:
			paymentServiceCreateRefundHandler.ServeHTTP(w, r)
		case PaymentServiceUpdateRefundProcedure:
			paymentServiceUpdateRefundHandler.ServeHTTP(w, r)
		case PaymentServiceCancelRefundProcedure:
			paymentServiceCancelRefundHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedPaymentServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPaymentServiceHandler struct{}

func (UnimplementedPaymentServiceHandler) GetPayment(context.Context, *connect.Request[v1.GetPaymentRequest]) (*connect.Response[v1.GetPaymentResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("payment.v1.PaymentService.GetPayment is not implemented"))
}

func (UnimplementedPaymentServiceHandler) ListPayments(context.Context, *connect.Request[v1.ListPaymentsRequest]) (*connect.Response[v1.ListPaymentsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("payment.v1.PaymentService.ListPayments is not implemented"))
}

func (UnimplementedPaymentServiceHandler) CreatePayment(context.Context, *connect.Request[v1.CreatePaymentRequest]) (*connect.Response[v1.CreatePaymentResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("payment.v1.PaymentService.CreatePayment is not implemented"))
}

func (UnimplementedPaymentServiceHandler) UpdatePayment(context.Context, *connect.Request[v1.UpdatePaymentRequest]) (*connect.Response[v1.UpdatePaymentResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("payment.v1.PaymentService.UpdatePayment is not implemented"))
}

func (UnimplementedPaymentServiceHandler) CancelPayment(context.Context, *connect.Request[v1.CancelPaymentRequest]) (*connect.Response[v1.CancelPaymentResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("payment.v1.PaymentService.CancelPayment is not implemented"))
}

func (UnimplementedPaymentServiceHandler) GetRefund(context.Context, *connect.Request[v1.GetRefundRequest]) (*connect.Response[v1.GetRefundResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("payment.v1.PaymentService.GetRefund is not implemented"))
}

func (UnimplementedPaymentServiceHandler) ListRefunds(context.Context, *connect.Request[v1.ListRefundsRequest]) (*connect.Response[v1.ListRefundsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("payment.v1.PaymentService.ListRefunds is not implemented"))
}

func (UnimplementedPaymentServiceHandler) CreateRefund(context.Context, *connect.Request[v1.CreateRefundRequest]) (*connect.Response[v1.CreateRefundResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("payment.v1.PaymentService.CreateRefund is not implemented"))
}

func (UnimplementedPaymentServiceHandler) UpdateRefund(context.Context, *connect.Request[v1.UpdateRefundRequest]) (*connect.Response[v1.UpdateRefundResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("payment.v1.PaymentService.UpdateRefund is not implemented"))
}

func (UnimplementedPaymentServiceHandler) CancelRefund(context.Context, *connect.Request[v1.CancelRefundRequest]) (*connect.Response[v1.CancelRefundResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("payment.v1.PaymentService.CancelRefund is not implemented"))
}
