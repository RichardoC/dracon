// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.7
// source: api/proto/v1/engine.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Scan related information, unique and immutable per scan run
// This message is copied from LaunchToolRequest to LaunchToolResponse
// by each producer wrapper
type ScanInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// scan unique identifier
	ScanUuid string `protobuf:"bytes,1,opt,name=scan_uuid,json=scanUuid,proto3" json:"scan_uuid,omitempty"`
	// timestamp of when the scan was triggered (passed to LaunchToolResponse)
	ScanStartTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=scan_start_time,json=scanStartTime,proto3" json:"scan_start_time,omitempty"`
}

func (x *ScanInfo) Reset() {
	*x = ScanInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_engine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanInfo) ProtoMessage() {}

func (x *ScanInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_engine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanInfo.ProtoReflect.Descriptor instead.
func (*ScanInfo) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_engine_proto_rawDescGZIP(), []int{0}
}

func (x *ScanInfo) GetScanUuid() string {
	if x != nil {
		return x.ScanUuid
	}
	return ""
}

func (x *ScanInfo) GetScanStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ScanStartTime
	}
	return nil
}

// LaunchToolReponse consists of a response built by a producer,
// to be interpreted by a consumer
type LaunchToolResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The scan information, see above for details
	ScanInfo *ScanInfo `protobuf:"bytes,1,opt,name=scan_info,json=scanInfo,proto3" json:"scan_info,omitempty"`
	// The name of the tool that ran the scan
	ToolName string `protobuf:"bytes,2,opt,name=tool_name,json=toolName,proto3" json:"tool_name,omitempty"`
	// Issues discovered during the scan
	Issues []*Issue `protobuf:"bytes,3,rep,name=issues,proto3" json:"issues,omitempty"`
}

func (x *LaunchToolResponse) Reset() {
	*x = LaunchToolResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_engine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LaunchToolResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LaunchToolResponse) ProtoMessage() {}

func (x *LaunchToolResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_engine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LaunchToolResponse.ProtoReflect.Descriptor instead.
func (*LaunchToolResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_engine_proto_rawDescGZIP(), []int{1}
}

func (x *LaunchToolResponse) GetScanInfo() *ScanInfo {
	if x != nil {
		return x.ScanInfo
	}
	return nil
}

func (x *LaunchToolResponse) GetToolName() string {
	if x != nil {
		return x.ToolName
	}
	return ""
}

func (x *LaunchToolResponse) GetIssues() []*Issue {
	if x != nil {
		return x.Issues
	}
	return nil
}

// An EnrichedLaunchToolResponse consists of deduplicated vulnerability
// information, with added metadata for consumers
type EnrichedLaunchToolResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The results of the original scan prior to enrichment
	OriginalResults *LaunchToolResponse `protobuf:"bytes,1,opt,name=original_results,json=originalResults,proto3" json:"original_results,omitempty"`
	// Enriched, deduplicated issues
	Issues []*EnrichedIssue `protobuf:"bytes,2,rep,name=issues,proto3" json:"issues,omitempty"`
}

func (x *EnrichedLaunchToolResponse) Reset() {
	*x = EnrichedLaunchToolResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_engine_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnrichedLaunchToolResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnrichedLaunchToolResponse) ProtoMessage() {}

func (x *EnrichedLaunchToolResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_engine_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnrichedLaunchToolResponse.ProtoReflect.Descriptor instead.
func (*EnrichedLaunchToolResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_engine_proto_rawDescGZIP(), []int{2}
}

func (x *EnrichedLaunchToolResponse) GetOriginalResults() *LaunchToolResponse {
	if x != nil {
		return x.OriginalResults
	}
	return nil
}

func (x *EnrichedLaunchToolResponse) GetIssues() []*EnrichedIssue {
	if x != nil {
		return x.Issues
	}
	return nil
}

var File_api_proto_v1_engine_proto protoreflect.FileDescriptor

var file_api_proto_v1_engine_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6f, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x2e, 0x64, 0x72, 0x61, 0x63, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x18,
	0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x08, 0x53, 0x63, 0x61,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x63, 0x61, 0x6e, 0x55, 0x75,
	0x69, 0x64, 0x12, 0x42, 0x0a, 0x0f, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x73, 0x63, 0x61, 0x6e, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x9d, 0x01, 0x0a, 0x12, 0x4c, 0x61, 0x75, 0x6e, 0x63,
	0x68, 0x54, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a,
	0x09, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x6f, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x64, 0x72, 0x61, 0x63, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x73,
	0x63, 0x61, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x6f, 0x6c, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6f, 0x6f, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6f, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x64,
	0x72, 0x61, 0x63, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x06,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x22, 0xa8, 0x01, 0x0a, 0x1a, 0x45, 0x6e, 0x72, 0x69, 0x63,
	0x68, 0x65, 0x64, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x54, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x10, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x6f, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x64, 0x72, 0x61, 0x63, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x54, 0x6f, 0x6f, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x38, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6f, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x2e, 0x64, 0x72, 0x61, 0x63, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x72, 0x69,
	0x63, 0x68, 0x65, 0x64, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65,
	0x73, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6f, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x64, 0x72, 0x61, 0x63, 0x6f, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_proto_v1_engine_proto_rawDescOnce sync.Once
	file_api_proto_v1_engine_proto_rawDescData = file_api_proto_v1_engine_proto_rawDesc
)

func file_api_proto_v1_engine_proto_rawDescGZIP() []byte {
	file_api_proto_v1_engine_proto_rawDescOnce.Do(func() {
		file_api_proto_v1_engine_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_v1_engine_proto_rawDescData)
	})
	return file_api_proto_v1_engine_proto_rawDescData
}

var file_api_proto_v1_engine_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_proto_v1_engine_proto_goTypes = []interface{}{
	(*ScanInfo)(nil),                   // 0: ocurity.dracon.v1.ScanInfo
	(*LaunchToolResponse)(nil),         // 1: ocurity.dracon.v1.LaunchToolResponse
	(*EnrichedLaunchToolResponse)(nil), // 2: ocurity.dracon.v1.EnrichedLaunchToolResponse
	(*timestamppb.Timestamp)(nil),      // 3: google.protobuf.Timestamp
	(*Issue)(nil),                      // 4: ocurity.dracon.v1.Issue
	(*EnrichedIssue)(nil),              // 5: ocurity.dracon.v1.EnrichedIssue
}
var file_api_proto_v1_engine_proto_depIdxs = []int32{
	3, // 0: ocurity.dracon.v1.ScanInfo.scan_start_time:type_name -> google.protobuf.Timestamp
	0, // 1: ocurity.dracon.v1.LaunchToolResponse.scan_info:type_name -> ocurity.dracon.v1.ScanInfo
	4, // 2: ocurity.dracon.v1.LaunchToolResponse.issues:type_name -> ocurity.dracon.v1.Issue
	1, // 3: ocurity.dracon.v1.EnrichedLaunchToolResponse.original_results:type_name -> ocurity.dracon.v1.LaunchToolResponse
	5, // 4: ocurity.dracon.v1.EnrichedLaunchToolResponse.issues:type_name -> ocurity.dracon.v1.EnrichedIssue
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_proto_v1_engine_proto_init() }
func file_api_proto_v1_engine_proto_init() {
	if File_api_proto_v1_engine_proto != nil {
		return
	}
	file_api_proto_v1_issue_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_proto_v1_engine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_v1_engine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LaunchToolResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_v1_engine_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnrichedLaunchToolResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_v1_engine_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_v1_engine_proto_goTypes,
		DependencyIndexes: file_api_proto_v1_engine_proto_depIdxs,
		MessageInfos:      file_api_proto_v1_engine_proto_msgTypes,
	}.Build()
	File_api_proto_v1_engine_proto = out.File
	file_api_proto_v1_engine_proto_rawDesc = nil
	file_api_proto_v1_engine_proto_goTypes = nil
	file_api_proto_v1_engine_proto_depIdxs = nil
}
