package dto

// This is a compile-time assertion to ensure that this generated file
// is compatible with the protoc-gen-go-gframe package it is being compiled against.
import (
	"github.com/brunowang/srvtools/scheduler"
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

type SetSchedulePlanReq struct {
	scheduler.SetSchedulePlanReq `json:",inline"`
}

func (m *SetSchedulePlanReq) IsValid() bool {

	if m.GetPlan() == nil {
		return false
	}

	return true
}

func (m *SetSchedulePlanReq) Fill(pb *scheduler.SetSchedulePlanReq) {
	if pb == nil {
		return
	}
	m.SetSchedulePlanReq = *pb
	return
}

type SetSchedulePlanRsp struct {
	scheduler.SetSchedulePlanRsp `json:",inline"`

	Err error `json:"err"`
}

func (m *SetSchedulePlanRsp) ToPb() *scheduler.SetSchedulePlanRsp {
	return &m.SetSchedulePlanRsp
}

func (m *SetSchedulePlanRsp) ToJson() []byte {
	js, _ := jspb.MarshalToString(m)
	return []byte(js)
}

type GetSchedulePlanReq struct {
	scheduler.GetSchedulePlanReq `json:",inline"`
}

func (m *GetSchedulePlanReq) IsValid() bool {

	if m.GetPlanId() == "" {
		return false
	}

	return true
}

func (m *GetSchedulePlanReq) Fill(pb *scheduler.GetSchedulePlanReq) {
	if pb == nil {
		return
	}
	m.GetSchedulePlanReq = *pb
	return
}

type GetSchedulePlanRsp struct {
	scheduler.GetSchedulePlanRsp `json:",inline"`

	Err error `json:"err"`
}

func (m *GetSchedulePlanRsp) ToPb() *scheduler.GetSchedulePlanRsp {
	return &m.GetSchedulePlanRsp
}

func (m *GetSchedulePlanRsp) ToJson() []byte {
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
