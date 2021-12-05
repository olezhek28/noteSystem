// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: note.proto

/*
Package note_v1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package note_v1

import (
	"context"
	"io"
	"net/http"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = descriptor.ForMessage
var _ = metadata.Join

func request_NoteV1_GetNotesList_0(ctx context.Context, marshaler runtime.Marshaler, client NoteV1Client, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq emptypb.Empty
	var metadata runtime.ServerMetadata

	msg, err := client.GetNotesList(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_NoteV1_GetNotesList_0(ctx context.Context, marshaler runtime.Marshaler, server NoteV1Server, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq emptypb.Empty
	var metadata runtime.ServerMetadata

	msg, err := server.GetNotesList(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterNoteV1HandlerServer registers the http handlers for service NoteV1 to "mux".
// UnaryRPC     :call NoteV1Server directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterNoteV1HandlerFromEndpoint instead.
func RegisterNoteV1HandlerServer(ctx context.Context, mux *runtime.ServeMux, server NoteV1Server) error {

	mux.Handle("GET", pattern_NoteV1_GetNotesList_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_NoteV1_GetNotesList_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_NoteV1_GetNotesList_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterNoteV1HandlerFromEndpoint is same as RegisterNoteV1Handler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterNoteV1HandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterNoteV1Handler(ctx, mux, conn)
}

// RegisterNoteV1Handler registers the http handlers for service NoteV1 to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterNoteV1Handler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterNoteV1HandlerClient(ctx, mux, NewNoteV1Client(conn))
}

// RegisterNoteV1HandlerClient registers the http handlers for service NoteV1
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "NoteV1Client".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "NoteV1Client"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "NoteV1Client" to call the correct interceptors.
func RegisterNoteV1HandlerClient(ctx context.Context, mux *runtime.ServeMux, client NoteV1Client) error {

	mux.Handle("GET", pattern_NoteV1_GetNotesList_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_NoteV1_GetNotesList_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_NoteV1_GetNotesList_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_NoteV1_GetNotesList_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"note", "v1", "list"}, "", runtime.AssumeColonVerbOpt(true)))
)

var (
	forward_NoteV1_GetNotesList_0 = runtime.ForwardResponseMessage
)
