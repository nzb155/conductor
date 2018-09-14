// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/workflowtask.proto

package model // import "github.com/netflix/conductor/client/gogrpc/conductor/model"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _struct "github.com/golang/protobuf/ptypes/struct"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type WorkflowTask struct {
	Name                           string                                    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TaskReferenceName              string                                    `protobuf:"bytes,2,opt,name=task_reference_name,json=taskReferenceName,proto3" json:"task_reference_name,omitempty"`
	Description                    string                                    `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	InputParameters                map[string]*_struct.Value                 `protobuf:"bytes,4,rep,name=input_parameters,json=inputParameters,proto3" json:"input_parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Type                           string                                    `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	DynamicTaskNameParam           string                                    `protobuf:"bytes,6,opt,name=dynamic_task_name_param,json=dynamicTaskNameParam,proto3" json:"dynamic_task_name_param,omitempty"`
	CaseValueParam                 string                                    `protobuf:"bytes,7,opt,name=case_value_param,json=caseValueParam,proto3" json:"case_value_param,omitempty"`
	CaseExpression                 string                                    `protobuf:"bytes,8,opt,name=case_expression,json=caseExpression,proto3" json:"case_expression,omitempty"`
	DecisionCases                  map[string]*WorkflowTask_WorkflowTaskList `protobuf:"bytes,9,rep,name=decision_cases,json=decisionCases,proto3" json:"decision_cases,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	DynamicForkTasksParam          string                                    `protobuf:"bytes,10,opt,name=dynamic_fork_tasks_param,json=dynamicForkTasksParam,proto3" json:"dynamic_fork_tasks_param,omitempty"`
	DynamicForkTasksInputParamName string                                    `protobuf:"bytes,11,opt,name=dynamic_fork_tasks_input_param_name,json=dynamicForkTasksInputParamName,proto3" json:"dynamic_fork_tasks_input_param_name,omitempty"`
	DefaultCase                    []*WorkflowTask                           `protobuf:"bytes,12,rep,name=default_case,json=defaultCase,proto3" json:"default_case,omitempty"`
	ForkTasks                      []*WorkflowTask_WorkflowTaskList          `protobuf:"bytes,13,rep,name=fork_tasks,json=forkTasks,proto3" json:"fork_tasks,omitempty"`
	StartDelay                     int32                                     `protobuf:"varint,14,opt,name=start_delay,json=startDelay,proto3" json:"start_delay,omitempty"`
	SubWorkflowParam               *SubWorkflowParams                        `protobuf:"bytes,15,opt,name=sub_workflow_param,json=subWorkflowParam,proto3" json:"sub_workflow_param,omitempty"`
	JoinOn                         []string                                  `protobuf:"bytes,16,rep,name=join_on,json=joinOn,proto3" json:"join_on,omitempty"`
	Sink                           string                                    `protobuf:"bytes,17,opt,name=sink,proto3" json:"sink,omitempty"`
	Optional                       bool                                      `protobuf:"varint,18,opt,name=optional,proto3" json:"optional,omitempty"`
	TaskDefinition                 *TaskDef                                  `protobuf:"bytes,19,opt,name=task_definition,json=taskDefinition,proto3" json:"task_definition,omitempty"`
	XXX_NoUnkeyedLiteral           struct{}                                  `json:"-"`
	XXX_unrecognized               []byte                                    `json:"-"`
	XXX_sizecache                  int32                                     `json:"-"`
}

func (m *WorkflowTask) Reset()         { *m = WorkflowTask{} }
func (m *WorkflowTask) String() string { return proto.CompactTextString(m) }
func (*WorkflowTask) ProtoMessage()    {}
func (*WorkflowTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_workflowtask_db933374415937fc, []int{0}
}
func (m *WorkflowTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowTask.Unmarshal(m, b)
}
func (m *WorkflowTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowTask.Marshal(b, m, deterministic)
}
func (dst *WorkflowTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowTask.Merge(dst, src)
}
func (m *WorkflowTask) XXX_Size() int {
	return xxx_messageInfo_WorkflowTask.Size(m)
}
func (m *WorkflowTask) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowTask.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowTask proto.InternalMessageInfo

func (m *WorkflowTask) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WorkflowTask) GetTaskReferenceName() string {
	if m != nil {
		return m.TaskReferenceName
	}
	return ""
}

func (m *WorkflowTask) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *WorkflowTask) GetInputParameters() map[string]*_struct.Value {
	if m != nil {
		return m.InputParameters
	}
	return nil
}

func (m *WorkflowTask) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *WorkflowTask) GetDynamicTaskNameParam() string {
	if m != nil {
		return m.DynamicTaskNameParam
	}
	return ""
}

func (m *WorkflowTask) GetCaseValueParam() string {
	if m != nil {
		return m.CaseValueParam
	}
	return ""
}

func (m *WorkflowTask) GetCaseExpression() string {
	if m != nil {
		return m.CaseExpression
	}
	return ""
}

func (m *WorkflowTask) GetDecisionCases() map[string]*WorkflowTask_WorkflowTaskList {
	if m != nil {
		return m.DecisionCases
	}
	return nil
}

func (m *WorkflowTask) GetDynamicForkTasksParam() string {
	if m != nil {
		return m.DynamicForkTasksParam
	}
	return ""
}

func (m *WorkflowTask) GetDynamicForkTasksInputParamName() string {
	if m != nil {
		return m.DynamicForkTasksInputParamName
	}
	return ""
}

func (m *WorkflowTask) GetDefaultCase() []*WorkflowTask {
	if m != nil {
		return m.DefaultCase
	}
	return nil
}

func (m *WorkflowTask) GetForkTasks() []*WorkflowTask_WorkflowTaskList {
	if m != nil {
		return m.ForkTasks
	}
	return nil
}

func (m *WorkflowTask) GetStartDelay() int32 {
	if m != nil {
		return m.StartDelay
	}
	return 0
}

func (m *WorkflowTask) GetSubWorkflowParam() *SubWorkflowParams {
	if m != nil {
		return m.SubWorkflowParam
	}
	return nil
}

func (m *WorkflowTask) GetJoinOn() []string {
	if m != nil {
		return m.JoinOn
	}
	return nil
}

func (m *WorkflowTask) GetSink() string {
	if m != nil {
		return m.Sink
	}
	return ""
}

func (m *WorkflowTask) GetOptional() bool {
	if m != nil {
		return m.Optional
	}
	return false
}

func (m *WorkflowTask) GetTaskDefinition() *TaskDef {
	if m != nil {
		return m.TaskDefinition
	}
	return nil
}

type WorkflowTask_WorkflowTaskList struct {
	Tasks                []*WorkflowTask `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *WorkflowTask_WorkflowTaskList) Reset()         { *m = WorkflowTask_WorkflowTaskList{} }
func (m *WorkflowTask_WorkflowTaskList) String() string { return proto.CompactTextString(m) }
func (*WorkflowTask_WorkflowTaskList) ProtoMessage()    {}
func (*WorkflowTask_WorkflowTaskList) Descriptor() ([]byte, []int) {
	return fileDescriptor_workflowtask_db933374415937fc, []int{0, 0}
}
func (m *WorkflowTask_WorkflowTaskList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowTask_WorkflowTaskList.Unmarshal(m, b)
}
func (m *WorkflowTask_WorkflowTaskList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowTask_WorkflowTaskList.Marshal(b, m, deterministic)
}
func (dst *WorkflowTask_WorkflowTaskList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowTask_WorkflowTaskList.Merge(dst, src)
}
func (m *WorkflowTask_WorkflowTaskList) XXX_Size() int {
	return xxx_messageInfo_WorkflowTask_WorkflowTaskList.Size(m)
}
func (m *WorkflowTask_WorkflowTaskList) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowTask_WorkflowTaskList.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowTask_WorkflowTaskList proto.InternalMessageInfo

func (m *WorkflowTask_WorkflowTaskList) GetTasks() []*WorkflowTask {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func init() {
	proto.RegisterType((*WorkflowTask)(nil), "conductor.proto.WorkflowTask")
	proto.RegisterMapType((map[string]*WorkflowTask_WorkflowTaskList)(nil), "conductor.proto.WorkflowTask.DecisionCasesEntry")
	proto.RegisterMapType((map[string]*_struct.Value)(nil), "conductor.proto.WorkflowTask.InputParametersEntry")
	proto.RegisterType((*WorkflowTask_WorkflowTaskList)(nil), "conductor.proto.WorkflowTask.WorkflowTaskList")
}

func init() {
	proto.RegisterFile("model/workflowtask.proto", fileDescriptor_workflowtask_db933374415937fc)
}

var fileDescriptor_workflowtask_db933374415937fc = []byte{
	// 688 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdf, 0x4f, 0x1b, 0x39,
	0x10, 0x56, 0x08, 0x01, 0x32, 0x81, 0x24, 0x18, 0xee, 0xb0, 0x72, 0xc7, 0x5d, 0x44, 0x1f, 0x9a,
	0x87, 0x6a, 0x53, 0x05, 0x55, 0xad, 0x78, 0x6a, 0x69, 0x68, 0x55, 0xf5, 0x17, 0xda, 0x56, 0x45,
	0x42, 0xaa, 0x56, 0x9b, 0x5d, 0x6f, 0xea, 0x66, 0x63, 0xaf, 0x6c, 0x6f, 0x21, 0x7f, 0x79, 0x5f,
	0x2b, 0x8f, 0x77, 0x93, 0x10, 0x10, 0x6a, 0xdf, 0xc6, 0xdf, 0x7c, 0xf3, 0xe3, 0xb3, 0x67, 0x0c,
	0x74, 0x2a, 0x63, 0x96, 0xf6, 0xaf, 0xa4, 0x9a, 0x24, 0xa9, 0xbc, 0x32, 0xa1, 0x9e, 0x78, 0x99,
	0x92, 0x46, 0x92, 0x56, 0x24, 0x45, 0x9c, 0x47, 0x46, 0x2a, 0x07, 0x74, 0xf6, 0x1c, 0xd5, 0x52,
	0x62, 0x96, 0x14, 0xe0, 0xa1, 0x03, 0x75, 0x3e, 0x2a, 0x53, 0x64, 0xa1, 0x0a, 0xa7, 0xba, 0x70,
	0xff, 0x3b, 0x96, 0x72, 0x9c, 0xb2, 0x3e, 0x9e, 0x46, 0x79, 0xd2, 0xd7, 0x46, 0xe5, 0x91, 0x71,
	0xde, 0xa3, 0x9f, 0x75, 0xd8, 0xbe, 0x28, 0xc2, 0x3e, 0x87, 0x7a, 0x42, 0x08, 0xac, 0x8b, 0x70,
	0xca, 0x68, 0xa5, 0x5b, 0xe9, 0xd5, 0x7d, 0xb4, 0x89, 0x07, 0x7b, 0xb6, 0x64, 0xa0, 0x58, 0xc2,
	0x14, 0x13, 0x11, 0x0b, 0x90, 0xb2, 0x86, 0x94, 0x5d, 0xeb, 0xf2, 0x4b, 0xcf, 0x07, 0xcb, 0xef,
	0x42, 0x23, 0x66, 0x3a, 0x52, 0x3c, 0x33, 0x5c, 0x0a, 0x5a, 0x45, 0xde, 0x32, 0x44, 0xbe, 0x42,
	0x9b, 0x8b, 0x2c, 0x37, 0x01, 0xb6, 0xca, 0x0c, 0x53, 0x9a, 0xae, 0x77, 0xab, 0xbd, 0xc6, 0x60,
	0xe0, 0xad, 0x88, 0xf6, 0x96, 0xdb, 0xf3, 0xde, 0xd8, 0xa8, 0xf3, 0x79, 0xd0, 0x99, 0x30, 0x6a,
	0xe6, 0xb7, 0xf8, 0x4d, 0xd4, 0x8a, 0x30, 0xb3, 0x8c, 0xd1, 0x9a, 0x13, 0x61, 0x6d, 0xf2, 0x04,
	0x0e, 0xe2, 0x99, 0x08, 0xa7, 0x3c, 0x0a, 0x50, 0x8c, 0x95, 0xe0, 0xca, 0xd3, 0x0d, 0xa4, 0xed,
	0x17, 0x6e, 0x5b, 0xc7, 0xca, 0xc0, 0x7c, 0xa4, 0x07, 0xed, 0x28, 0xd4, 0x2c, 0xf8, 0x11, 0xa6,
	0x79, 0xc9, 0xdf, 0x44, 0x7e, 0xd3, 0xe2, 0x5f, 0x2c, 0xec, 0x98, 0x0f, 0xa1, 0x85, 0x4c, 0x76,
	0x9d, 0x29, 0xa6, 0xb5, 0x55, 0xbe, 0xb5, 0x20, 0x9e, 0xcd, 0x51, 0x72, 0x01, 0xcd, 0x98, 0x45,
	0xdc, 0xda, 0x81, 0x75, 0x69, 0x5a, 0x47, 0xe9, 0x8f, 0xef, 0x97, 0x3e, 0x2c, 0x62, 0x5e, 0xda,
	0x10, 0x27, 0x7c, 0x27, 0x5e, 0xc6, 0xc8, 0x53, 0xa0, 0xa5, 0xc4, 0x44, 0xaa, 0x09, 0xea, 0xd4,
	0x45, 0xcf, 0x80, 0xad, 0xfc, 0x55, 0xf8, 0x5f, 0x49, 0x35, 0xb1, 0x49, 0xb5, 0x6b, 0xfd, 0x2d,
	0x3c, 0xb8, 0x23, 0x70, 0xe9, 0x85, 0xdc, 0x83, 0x37, 0x30, 0xc7, 0x7f, 0xab, 0x39, 0x16, 0x6f,
	0x82, 0xaf, 0xff, 0x1c, 0xb6, 0x63, 0x96, 0x84, 0x79, 0x6a, 0x50, 0x1d, 0xdd, 0x46, 0x71, 0x87,
	0xf7, 0x8a, 0xb3, 0xd3, 0x81, 0x21, 0x56, 0x08, 0x79, 0x0f, 0xb0, 0x68, 0x83, 0xee, 0x60, 0xbc,
	0x77, 0xff, 0xe5, 0x2c, 0x1f, 0xde, 0x71, 0x6d, 0xfc, 0x7a, 0x52, 0xb6, 0x47, 0xfe, 0x87, 0x86,
	0x36, 0xa1, 0x32, 0x41, 0xcc, 0xd2, 0x70, 0x46, 0x9b, 0xdd, 0x4a, 0xaf, 0xe6, 0x03, 0x42, 0x43,
	0x8b, 0x90, 0x73, 0x20, 0x3a, 0x1f, 0x05, 0xe5, 0xfa, 0x14, 0x37, 0xd6, 0xea, 0x56, 0x7a, 0x8d,
	0xc1, 0xd1, 0xad, 0xba, 0x9f, 0xf2, 0x51, 0x59, 0x0d, 0x45, 0x6b, 0xbf, 0xad, 0x57, 0x20, 0x72,
	0x00, 0x9b, 0xdf, 0x25, 0x17, 0x81, 0x14, 0xb4, 0xdd, 0xad, 0xf6, 0xea, 0xfe, 0x86, 0x3d, 0x7e,
	0x14, 0x76, 0x32, 0x35, 0x17, 0x13, 0xba, 0xeb, 0x26, 0xd3, 0xda, 0xa4, 0x03, 0x5b, 0x12, 0xd7,
	0x22, 0x4c, 0x29, 0xe9, 0x56, 0x7a, 0x5b, 0xfe, 0xfc, 0x4c, 0x5e, 0x40, 0x0b, 0xa7, 0x35, 0x66,
	0x09, 0x17, 0x1c, 0xd7, 0x69, 0x0f, 0xfb, 0xa2, 0xb7, 0xfa, 0xb2, 0x62, 0x87, 0x2c, 0xf1, 0x9b,
	0xc6, 0x19, 0x05, 0xbf, 0xf3, 0x1a, 0xda, 0xab, 0xb7, 0x43, 0x8e, 0xa1, 0xe6, 0x2e, 0xb7, 0xf2,
	0x3b, 0x8f, 0xe3, 0xb8, 0x9d, 0x4b, 0xd8, 0xbf, 0x6b, 0xfd, 0x48, 0x1b, 0xaa, 0x13, 0x36, 0x2b,
	0x7e, 0x0c, 0x6b, 0x92, 0x47, 0x50, 0xc3, 0x7d, 0xc1, 0x2f, 0xa2, 0x31, 0xf8, 0xdb, 0x73, 0x7f,
	0x90, 0x57, 0xfe, 0x41, 0x1e, 0xae, 0x8d, 0xef, 0x48, 0x27, 0x6b, 0xcf, 0x2a, 0x9d, 0x0c, 0xc8,
	0xed, 0xf9, 0xbe, 0x23, 0xf3, 0xf0, 0x66, 0xe6, 0x3f, 0x9d, 0x8a, 0x45, 0xc5, 0x53, 0x0e, 0xff,
	0x44, 0x72, 0xea, 0x09, 0x66, 0x92, 0x94, 0x5f, 0xaf, 0xe6, 0x39, 0x6d, 0x2e, 0xc7, 0x9e, 0x8f,
	0x2e, 0x4f, 0xc6, 0xdc, 0x7c, 0xcb, 0x47, 0x5e, 0x24, 0xa7, 0xfd, 0x22, 0xa6, 0x3f, 0x8f, 0xe9,
	0x47, 0x29, 0x67, 0xc2, 0xf4, 0xc7, 0x72, 0xac, 0xb2, 0x68, 0x09, 0xc7, 0xaf, 0x79, 0xb4, 0x81,
	0x29, 0x8f, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0xc5, 0xea, 0x4c, 0x04, 0xea, 0x05, 0x00, 0x00,
}
