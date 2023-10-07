package frontend

// This is a compile-time assertion to ensure that this generated file
// is compatible with the protoc-gen-go-gframe package it is being compiled against.
import (
	"github.com/brunowang/gframe/gflog"
	"github.com/brunowang/srvtools/scheduler/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net"
)

type httpServer struct {
	h *httpHandler
}

func NewHttpServer(svc service.SchedulerServiceService) *httpServer {
	return &httpServer{
		h: NewHttpHandler(svc),
	}
}

func (s *httpServer) Serve(httpl net.Listener) error {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gflog.TraceMiddleware())

	router.POST("/api/scheduler/start_schedule", s.h.StartSchedule)
	router.POST("/api/scheduler/cancel_schedule", s.h.CancelSchedule)
	router.POST("/api/scheduler/refresh_schedule_time", s.h.RefreshScheduleTime)
	router.GET("/api/scheduler/get_ready_schedule", s.h.GetReadySchedule)
	router.GET("/api/scheduler/get_schedule_time", s.h.GetScheduleTime)
	router.POST("/api/scheduler/set_schedule_plan", s.h.SetSchedulePlan)
	router.GET("/api/scheduler/get_schedule_plan", s.h.GetSchedulePlan)

	gflog.Info(nil, "Http start listen", zap.String("address", httpl.Addr().String()))

	return router.RunListener(httpl)
}
