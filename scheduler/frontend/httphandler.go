package frontend

// This is a compile-time assertion to ensure that this generated file
// is compatible with the protoc-gen-go-gframe package it is being compiled against.
import (
	"encoding/json"
	"github.com/brunowang/gframe/gfhttp"
	"github.com/brunowang/gframe/gflog"
	"github.com/brunowang/srvtools/scheduler/dto"
	"github.com/brunowang/srvtools/scheduler/service"
	"github.com/gin-gonic/gin"
	ws "github.com/gorilla/websocket"
	"go.uber.org/zap"
	"net/http"
	"sync"
	"time"
)

type httpHandler struct {
	svc service.SchedulerServiceService
}

func NewHttpHandler(svc service.SchedulerServiceService) *httpHandler {
	return &httpHandler{svc: svc}
}

func (s *httpHandler) StartSchedule(ctx *gin.Context) {
	var req dto.StartScheduleReq
	if !gfhttp.BindJson(ctx, &req) {
		return
	}

	gflog.Info(ctx, "httpHandler StartSchedule processing")
	nowt := time.Now()

	rsp, err := s.svc.StartSchedule(ctx, &req)
	if err != nil {
		gflog.Error(ctx, "httpHandler StartSchedule error", zap.Error(err))
		gfhttp.NewResp(ctx).Err(err)
		return
	}
	gflog.Info(ctx, "httpHandler StartSchedule finished", zap.Duration("latency", time.Since(nowt)))

	gfhttp.NewResp(ctx).OK(rsp.ToPb())
}

func (s *httpHandler) CancelSchedule(ctx *gin.Context) {
	var req dto.CancelScheduleReq
	if !gfhttp.BindJson(ctx, &req) {
		return
	}

	gflog.Info(ctx, "httpHandler CancelSchedule processing")
	nowt := time.Now()

	rsp, err := s.svc.CancelSchedule(ctx, &req)
	if err != nil {
		gflog.Error(ctx, "httpHandler CancelSchedule error", zap.Error(err))
		gfhttp.NewResp(ctx).Err(err)
		return
	}
	gflog.Info(ctx, "httpHandler CancelSchedule finished", zap.Duration("latency", time.Since(nowt)))

	gfhttp.NewResp(ctx).OK(rsp.ToPb())
}

func (s *httpHandler) RefreshScheduleTime(ctx *gin.Context) {
	var req dto.RefreshScheduleTimeReq
	if !gfhttp.BindJson(ctx, &req) {
		return
	}

	gflog.Info(ctx, "httpHandler RefreshScheduleTime processing")
	nowt := time.Now()

	rsp, err := s.svc.RefreshScheduleTime(ctx, &req)
	if err != nil {
		gflog.Error(ctx, "httpHandler RefreshScheduleTime error", zap.Error(err))
		gfhttp.NewResp(ctx).Err(err)
		return
	}
	gflog.Info(ctx, "httpHandler RefreshScheduleTime finished", zap.Duration("latency", time.Since(nowt)))

	gfhttp.NewResp(ctx).OK(rsp.ToPb())
}

func (s *httpHandler) GetReadySchedule(ctx *gin.Context) {
	var req dto.GetReadyScheduleReq
	if !gfhttp.BindJson(ctx, &req) {
		return
	}

	gflog.Info(ctx, "httpHandler GetReadySchedule processing")
	nowt := time.Now()

	rsp, err := s.svc.GetReadySchedule(ctx, &req)
	if err != nil {
		gflog.Error(ctx, "httpHandler GetReadySchedule error", zap.Error(err))
		gfhttp.NewResp(ctx).Err(err)
		return
	}
	gflog.Info(ctx, "httpHandler GetReadySchedule finished", zap.Duration("latency", time.Since(nowt)))

	gfhttp.NewResp(ctx).OK(rsp.ToPb())
}

func (s *httpHandler) GetScheduleTime(ctx *gin.Context) {
	var req dto.GetScheduleTimeReq
	if !gfhttp.BindJson(ctx, &req) {
		return
	}

	gflog.Info(ctx, "httpHandler GetScheduleTime processing")
	nowt := time.Now()

	rsp, err := s.svc.GetScheduleTime(ctx, &req)
	if err != nil {
		gflog.Error(ctx, "httpHandler GetScheduleTime error", zap.Error(err))
		gfhttp.NewResp(ctx).Err(err)
		return
	}
	gflog.Info(ctx, "httpHandler GetScheduleTime finished", zap.Duration("latency", time.Since(nowt)))

	gfhttp.NewResp(ctx).OK(rsp.ToPb())
}

func (s *httpHandler) SetSchedulePlan(ctx *gin.Context) {
	var req dto.SetSchedulePlanReq
	if !gfhttp.BindJson(ctx, &req) {
		return
	}

	gflog.Info(ctx, "httpHandler SetSchedulePlan processing")
	nowt := time.Now()

	rsp, err := s.svc.SetSchedulePlan(ctx, &req)
	if err != nil {
		gflog.Error(ctx, "httpHandler SetSchedulePlan error", zap.Error(err))
		gfhttp.NewResp(ctx).Err(err)
		return
	}
	gflog.Info(ctx, "httpHandler SetSchedulePlan finished", zap.Duration("latency", time.Since(nowt)))

	gfhttp.NewResp(ctx).OK(rsp.ToPb())
}

func (s *httpHandler) GetSchedulePlan(ctx *gin.Context) {
	var req dto.GetSchedulePlanReq
	if !gfhttp.BindJson(ctx, &req) {
		return
	}

	gflog.Info(ctx, "httpHandler GetSchedulePlan processing")
	nowt := time.Now()

	rsp, err := s.svc.GetSchedulePlan(ctx, &req)
	if err != nil {
		gflog.Error(ctx, "httpHandler GetSchedulePlan error", zap.Error(err))
		gfhttp.NewResp(ctx).Err(err)
		return
	}
	gflog.Info(ctx, "httpHandler GetSchedulePlan finished", zap.Duration("latency", time.Since(nowt)))

	gfhttp.NewResp(ctx).OK(rsp.ToPb())
}
