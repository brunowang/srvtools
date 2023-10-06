package service

// This is a compile-time assertion to ensure that this generated file
// is compatible with the protoc-gen-go-gframe package it is being compiled against.
import (
	"context"
	"github.com/brunowang/srvtools/scheduler/dto"
)

type SchedulerServiceService interface {
	StartSchedule(ctx context.Context, req *dto.StartScheduleReq) (*dto.StartScheduleRsp, error)

	CancelSchedule(ctx context.Context, req *dto.CancelScheduleReq) (*dto.Empty, error)

	RefreshScheduleTime(ctx context.Context, req *dto.RefreshScheduleTimeReq) (*dto.Empty, error)

	GetReadySchedule(ctx context.Context, req *dto.GetReadyScheduleReq) (*dto.GetReadyScheduleRsp, error)

	GetScheduleTime(ctx context.Context, req *dto.GetScheduleTimeReq) (*dto.GetScheduleTimeRsp, error)

	CreateSchedulePlan(ctx context.Context, req *dto.CreateSchedulePlanReq) (*dto.CreateSchedulePlanRsp, error)
}

type SchedulerServiceServiceImpl struct{}

func NewSchedulerService() SchedulerServiceService {
	return &SchedulerServiceServiceImpl{}
}

func (s *SchedulerServiceServiceImpl) StartSchedule(ctx context.Context, req *dto.StartScheduleReq) (*dto.StartScheduleRsp, error) {
	return new(dto.StartScheduleRsp), nil
}

func (s *SchedulerServiceServiceImpl) CancelSchedule(ctx context.Context, req *dto.CancelScheduleReq) (*dto.Empty, error) {
	return new(dto.Empty), nil
}

func (s *SchedulerServiceServiceImpl) RefreshScheduleTime(ctx context.Context, req *dto.RefreshScheduleTimeReq) (*dto.Empty, error) {
	return new(dto.Empty), nil
}

func (s *SchedulerServiceServiceImpl) GetReadySchedule(ctx context.Context, req *dto.GetReadyScheduleReq) (*dto.GetReadyScheduleRsp, error) {
	return new(dto.GetReadyScheduleRsp), nil
}

func (s *SchedulerServiceServiceImpl) GetScheduleTime(ctx context.Context, req *dto.GetScheduleTimeReq) (*dto.GetScheduleTimeRsp, error) {
	return new(dto.GetScheduleTimeRsp), nil
}

func (s *SchedulerServiceServiceImpl) CreateSchedulePlan(ctx context.Context, req *dto.CreateSchedulePlanReq) (*dto.CreateSchedulePlanRsp, error) {
	return new(dto.CreateSchedulePlanRsp), nil
}
