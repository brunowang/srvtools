package frontend

// This is a compile-time assertion to ensure that this generated file
// is compatible with the protoc-gen-go-gframe package it is being compiled against.
import (
	"github.com/brunowang/gframe/gflog"
	scheduler "github.com/brunowang/srvtools/scheduler/pb"
	"github.com/brunowang/srvtools/scheduler/service"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type grpcServer struct {
	h *grpcHandler
}

func NewGrpcServer(svc service.SchedulerServiceService) *grpcServer {
	return &grpcServer{
		h: NewGrpcHandler(svc),
	}
}

func (s *grpcServer) Serve(grpcl net.Listener) error {
	srv := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			gflog.UnaryEntryInterceptor,
			gflog.UnaryTraceInterceptor,
		),
		grpc_middleware.WithStreamServerChain(
			gflog.StreamEntryInterceptor,
			gflog.StreamTraceInterceptor,
		),
		grpc.MaxRecvMsgSize(1<<30), grpc.MaxSendMsgSize(1<<30),
	)

	scheduler.RegisterSchedulerServiceServer(srv, s.h)

	reflection.Register(srv)

	gflog.Info(nil, "grpc start listen", zap.String("address", grpcl.Addr().String()))

	return srv.Serve(grpcl)
}
