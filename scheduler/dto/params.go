package dto

// This is a compile-time assertion to ensure that this generated file
// is compatible with the protoc-gen-go-gframe package it is being compiled against.
import (
	scheduler "github.com/brunowang/srvtools/scheduler/pb"
	"github.com/golang/protobuf/jsonpb"
)

var jspb = jsonpb.Marshaler{OrigName: true, EmitDefaults: true}

type StartScheduleReq struct {
	scheduler.StartScheduleReq `json:",inline"`
}

func (m *StartScheduleReq) IsValid() bool {

	if m.GetItem() == nil {
		return false
	}

	return true
}

func (m *StartScheduleReq) Fill(pb *scheduler.StartScheduleReq) {
	if pb == nil {
		return
	}
	m.StartScheduleReq = *pb
	return
}

type StartScheduleRsp struct {
	scheduler.StartScheduleRsp `json:",inline"`

	Err error `json:"err"`
}

func (m *StartScheduleRsp) ToPb() *scheduler.StartScheduleRsp {
	return &m.StartScheduleRsp
}

func (m *StartScheduleRsp) ToJson() []byte {
	js, _ := jspb.MarshalToString(m)
	return []byte(js)
}

type CancelScheduleReq struct {
	scheduler.CancelScheduleReq `json:",inline"`
}

func (m *CancelScheduleReq) IsValid() bool {

	if m.GetItemId() == 0 {
		return false
	}

	return true
}

func (m *CancelScheduleReq) Fill(pb *scheduler.CancelScheduleReq) {
	if pb == nil {
		return
	}
	m.CancelScheduleReq = *pb
	return
}

type RefreshScheduleTimeReq struct {
	scheduler.RefreshScheduleTimeReq `json:",inline"`
}

func (m *RefreshScheduleTimeReq) IsValid() bool {

	if m.GetItemId() == 0 {
		return false
	}

	return true
}

func (m *RefreshScheduleTimeReq) Fill(pb *scheduler.RefreshScheduleTimeReq) {
	if pb == nil {
		return
	}
	m.RefreshScheduleTimeReq = *pb
	return
}

type GetReadyScheduleReq struct {
	scheduler.GetReadyScheduleReq `json:",inline"`
}

func (m *GetReadyScheduleReq) IsValid() bool {

	if m.GetBucket() == "" {
		return false
	}

	if m.GetPlanId() == "" {
		return false
	}

	if m.GetTime() == 0 {
		return false
	}

	if m.GetCount() == 0 {
		return false
	}

	return true
}

func (m *GetReadyScheduleReq) Fill(pb *scheduler.GetReadyScheduleReq) {
	if pb == nil {
		return
	}
	m.GetReadyScheduleReq = *pb
	return
}

type GetReadyScheduleRsp struct {
	scheduler.GetReadyScheduleRsp `json:",inline"`

	Err error `json:"err"`
}

func (m *GetReadyScheduleRsp) ToPb() *scheduler.GetReadyScheduleRsp {
	return &m.GetReadyScheduleRsp
}

func (m *GetReadyScheduleRsp) ToJson() []byte {
	js, _ := jspb.MarshalToString(m)
	return []byte(js)
}

type GetScheduleTimeReq struct {
	scheduler.GetScheduleTimeReq `json:",inline"`
}

func (m *GetScheduleTimeReq) IsValid() bool {

	if m.GetItemId() == 0 {
		return false
	}

	if m.GetPlanId() == "" {
		return false
	}

	return true
}

func (m *GetScheduleTimeReq) Fill(pb *scheduler.GetScheduleTimeReq) {
	if pb == nil {
		return
	}
	m.GetScheduleTimeReq = *pb
	return
}

type GetScheduleTimeRsp struct {
	scheduler.GetScheduleTimeRsp `json:",inline"`

	Err error `json:"err"`
}

func (m *GetScheduleTimeRsp) ToPb() *scheduler.GetScheduleTimeRsp {
	return &m.GetScheduleTimeRsp
}

func (m *GetScheduleTimeRsp) ToJson() []byte {
	js, _ := jspb.MarshalToString(m)
	return []byte(js)
}

type CreateSchedulePlanReq struct {
	scheduler.CreateSchedulePlanReq `json:",inline"`
}

func (m *CreateSchedulePlanReq) IsValid() bool {

	if m.GetPlan() == nil {
		return false
	}

	return true
}

func (m *CreateSchedulePlanReq) Fill(pb *scheduler.CreateSchedulePlanReq) {
	if pb == nil {
		return
	}
	m.CreateSchedulePlanReq = *pb
	return
}

type CreateSchedulePlanRsp struct {
	scheduler.CreateSchedulePlanRsp `json:",inline"`

	Err error `json:"err"`
}

func (m *CreateSchedulePlanRsp) ToPb() *scheduler.CreateSchedulePlanRsp {
	return &m.CreateSchedulePlanRsp
}

func (m *CreateSchedulePlanRsp) ToJson() []byte {
	js, _ := jspb.MarshalToString(m)
	return []byte(js)
}

type ScheduleItem struct {
	scheduler.ScheduleItem `json:",inline"`
}

type SchedulePlan struct {
	scheduler.SchedulePlan `json:",inline"`
}

type Empty struct {
	scheduler.Empty `json:",inline"`

	Err error `json:"err"`
}

func (m *Empty) ToPb() *scheduler.Empty {
	return &m.Empty
}

func (m *Empty) ToJson() []byte {
	js, _ := jspb.MarshalToString(m)
	return []byte(js)
}
