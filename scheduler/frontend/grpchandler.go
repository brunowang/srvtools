package frontend

// This is a compile-time assertion to ensure that this generated file
// is compatible with the protoc-gen-go-gframe package it is being compiled against.
import (
	"context"
	"github.com/brunowang/gframe/gflog"
	"github.com/brunowang/srvtools/scheduler/dto"
	scheduler "github.com/brunowang/srvtools/scheduler/pb"
	"github.com/brunowang/srvtools/scheduler/service"
	"go.uber.org/zap"
	"time"
)

type grpcHandler struct {
	scheduler.UnimplementedSchedulerServiceServer
	svc service.SchedulerServiceService
}

func NewGrpcHandler(svc service.SchedulerServiceService) *grpcHandler {
	return &grpcHandler{svc: svc}
}

func (g *grpcHandler) StartSchedule(ctx context.Context, req *scheduler.StartScheduleReq) (*scheduler.StartScheduleRsp, error) {
	params := &dto.StartScheduleReq{}
	params.Fill(req)

	gflog.Info(ctx, "grpcHandler StartSchedule processing")
	nowt := time.Now()

	result, err := g.svc.StartSchedule(ctx, params)
	if err != nil {
		gflog.Error(ctx, "grpcHandler StartSchedule error", zap.Error(err))
		return nil, err
	}
	gflog.Info(ctx, "grpcHandler logical processing finished", zap.Duration("latency", time.Since(nowt)))

	return result.ToPb(), nil
}

func (g *grpcHandler) CancelSchedule(ctx context.Context, req *scheduler.CancelScheduleReq) (*scheduler.Empty, error) {
	params := &dto.CancelScheduleReq{}
	params.Fill(req)

	gflog.Info(ctx, "grpcHandler CancelSchedule processing")
	nowt := time.Now()

	result, err := g.svc.CancelSchedule(ctx, params)
	if err != nil {
		gflog.Error(ctx, "grpcHandler CancelSchedule error", zap.Error(err))
		return nil, err
	}
	gflog.Info(ctx, "grpcHandler logical processing finished", zap.Duration("latency", time.Since(nowt)))

	return result.ToPb(), nil
}

func (g *grpcHandler) RefreshScheduleTime(ctx context.Context, req *scheduler.RefreshScheduleTimeReq) (*scheduler.Empty, error) {
	params := &dto.RefreshScheduleTimeReq{}
	params.Fill(req)

	gflog.Info(ctx, "grpcHandler RefreshScheduleTime processing")
	nowt := time.Now()

	result, err := g.svc.RefreshScheduleTime(ctx, params)
	if err != nil {
		gflog.Error(ctx, "grpcHandler RefreshScheduleTime error", zap.Error(err))
		return nil, err
	}
	gflog.Info(ctx, "grpcHandler logical processing finished", zap.Duration("latency", time.Since(nowt)))

	return result.ToPb(), nil
}

func (g *grpcHandler) GetReadySchedule(ctx context.Context, req *scheduler.GetReadyScheduleReq) (*scheduler.GetReadyScheduleRsp, error) {
	params := &dto.GetReadyScheduleReq{}
	params.Fill(req)

	gflog.Info(ctx, "grpcHandler GetReadySchedule processing")
	nowt := time.Now()

	result, err := g.svc.GetReadySchedule(ctx, params)
	if err != nil {
		gflog.Error(ctx, "grpcHandler GetReadySchedule error", zap.Error(err))
		return nil, err
	}
	gflog.Info(ctx, "grpcHandler logical processing finished", zap.Duration("latency", time.Since(nowt)))

	return result.ToPb(), nil
}

func (g *grpcHandler) GetScheduleTime(ctx context.Context, req *scheduler.GetScheduleTimeReq) (*scheduler.GetScheduleTimeRsp, error) {
	params := &dto.GetScheduleTimeReq{}
	params.Fill(req)

	gflog.Info(ctx, "grpcHandler GetScheduleTime processing")
	nowt := time.Now()

	result, err := g.svc.GetScheduleTime(ctx, params)
	if err != nil {
		gflog.Error(ctx, "grpcHandler GetScheduleTime error", zap.Error(err))
		return nil, err
	}
	gflog.Info(ctx, "grpcHandler logical processing finished", zap.Duration("latency", time.Since(nowt)))

	return result.ToPb(), nil
}

func (g *grpcHandler) CreateSchedulePlan(ctx context.Context, req *scheduler.CreateSchedulePlanReq) (*scheduler.CreateSchedulePlanRsp, error) {
	params := &dto.CreateSchedulePlanReq{}
	params.Fill(req)

	gflog.Info(ctx, "grpcHandler CreateSchedulePlan processing")
	nowt := time.Now()

	result, err := g.svc.CreateSchedulePlan(ctx, params)
	if err != nil {
		gflog.Error(ctx, "grpcHandler CreateSchedulePlan error", zap.Error(err))
		return nil, err
	}
	gflog.Info(ctx, "grpcHandler logical processing finished", zap.Duration("latency", time.Since(nowt)))

	return result.ToPb(), nil
}
