// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: time.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetRawTimeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRawTimeReq) Reset() {
	*x = GetRawTimeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_time_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRawTimeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRawTimeReq) ProtoMessage() {}

func (x *GetRawTimeReq) ProtoReflect() protoreflect.Message {
	mi := &file_time_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRawTimeReq.ProtoReflect.Descriptor instead.
func (*GetRawTimeReq) Descriptor() ([]byte, []int) {
	return file_time_proto_rawDescGZIP(), []int{0}
}

type GetRawTimeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *GetRawTimeResp) Reset() {
	*x = GetRawTimeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_time_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRawTimeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRawTimeResp) ProtoMessage() {}

func (x *GetRawTimeResp) ProtoReflect() protoreflect.Message {
	mi := &file_time_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRawTimeResp.ProtoReflect.Descriptor instead.
func (*GetRawTimeResp) Descriptor() ([]byte, []int) {
	return file_time_proto_rawDescGZIP(), []int{1}
}

func (x *GetRawTimeResp) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type GetFmtTimeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetFmtTimeReq) Reset() {
	*x = GetFmtTimeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_time_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFmtTimeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFmtTimeReq) ProtoMessage() {}

func (x *GetFmtTimeReq) ProtoReflect() protoreflect.Message {
	mi := &file_time_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFmtTimeReq.ProtoReflect.Descriptor instead.
func (*GetFmtTimeReq) Descriptor() ([]byte, []int) {
	return file_time_proto_rawDescGZIP(), []int{2}
}

type GetFmtTimeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FmtTime string `protobuf:"bytes,1,opt,name=fmtTime,proto3" json:"fmtTime,omitempty"`
}

func (x *GetFmtTimeResp) Reset() {
	*x = GetFmtTimeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_time_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFmtTimeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFmtTimeResp) ProtoMessage() {}

func (x *GetFmtTimeResp) ProtoReflect() protoreflect.Message {
	mi := &file_time_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFmtTimeResp.ProtoReflect.Descriptor instead.
func (*GetFmtTimeResp) Descriptor() ([]byte, []int) {
	return file_time_proto_rawDescGZIP(), []int{3}
}

func (x *GetFmtTimeResp) GetFmtTime() string {
	if x != nil {
		return x.FmtTime
	}
	return ""
}

type GetFmtTimeFromRawReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *GetFmtTimeFromRawReq) Reset() {
	*x = GetFmtTimeFromRawReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_time_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFmtTimeFromRawReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFmtTimeFromRawReq) ProtoMessage() {}

func (x *GetFmtTimeFromRawReq) ProtoReflect() protoreflect.Message {
	mi := &file_time_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFmtTimeFromRawReq.ProtoReflect.Descriptor instead.
func (*GetFmtTimeFromRawReq) Descriptor() ([]byte, []int) {
	return file_time_proto_rawDescGZIP(), []int{4}
}

func (x *GetFmtTimeFromRawReq) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type GetFmtTimeFromRawResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FmtTime string `protobuf:"bytes,1,opt,name=fmtTime,proto3" json:"fmtTime,omitempty"`
}

func (x *GetFmtTimeFromRawResp) Reset() {
	*x = GetFmtTimeFromRawResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_time_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFmtTimeFromRawResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFmtTimeFromRawResp) ProtoMessage() {}

func (x *GetFmtTimeFromRawResp) ProtoReflect() protoreflect.Message {
	mi := &file_time_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFmtTimeFromRawResp.ProtoReflect.Descriptor instead.
func (*GetFmtTimeFromRawResp) Descriptor() ([]byte, []int) {
	return file_time_proto_rawDescGZIP(), []int{5}
}

func (x *GetFmtTimeFromRawResp) GetFmtTime() string {
	if x != nil {
		return x.FmtTime
	}
	return ""
}

var File_time_proto protoreflect.FileDescriptor

var file_time_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x52, 0x61, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x22, 0x2e, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x52, 0x61, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x0f, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x46, 0x6d, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x22, 0x2a, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x46, 0x6d, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18,
	0x0a, 0x07, 0x66, 0x6d, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x66, 0x6d, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x34, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x46,
	0x6d, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x52, 0x61, 0x77, 0x52, 0x65, 0x71,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x31,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x46, 0x6d, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x72, 0x6f, 0x6d,
	0x52, 0x61, 0x77, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x6d, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x6d, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x32, 0x82, 0x02, 0x0a, 0x0a, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x49, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x61, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b,
	0x2e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x61, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61,
	0x77, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x46, 0x6d, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x2e, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6d, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6d, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x46, 0x6d, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x52, 0x61, 0x77, 0x12, 0x22, 0x2e, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6d,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x52, 0x61, 0x77, 0x52, 0x65, 0x71, 0x1a,
	0x23, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x46, 0x6d, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x52, 0x61, 0x77,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_time_proto_rawDescOnce sync.Once
	file_time_proto_rawDescData = file_time_proto_rawDesc
)

func file_time_proto_rawDescGZIP() []byte {
	file_time_proto_rawDescOnce.Do(func() {
		file_time_proto_rawDescData = protoimpl.X.CompressGZIP(file_time_proto_rawDescData)
	})
	return file_time_proto_rawDescData
}

var file_time_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_time_proto_goTypes = []interface{}{
	(*GetRawTimeReq)(nil),         // 0: time_service.GetRawTimeReq
	(*GetRawTimeResp)(nil),        // 1: time_service.GetRawTimeResp
	(*GetFmtTimeReq)(nil),         // 2: time_service.GetFmtTimeReq
	(*GetFmtTimeResp)(nil),        // 3: time_service.GetFmtTimeResp
	(*GetFmtTimeFromRawReq)(nil),  // 4: time_service.GetFmtTimeFromRawReq
	(*GetFmtTimeFromRawResp)(nil), // 5: time_service.GetFmtTimeFromRawResp
}
var file_time_proto_depIdxs = []int32{
	0, // 0: time_service.TimeServer.GetRawTime:input_type -> time_service.GetRawTimeReq
	2, // 1: time_service.TimeServer.GetFmtTime:input_type -> time_service.GetFmtTimeReq
	4, // 2: time_service.TimeServer.GetFmtTimeFromRaw:input_type -> time_service.GetFmtTimeFromRawReq
	1, // 3: time_service.TimeServer.GetRawTime:output_type -> time_service.GetRawTimeResp
	3, // 4: time_service.TimeServer.GetFmtTime:output_type -> time_service.GetFmtTimeResp
	5, // 5: time_service.TimeServer.GetFmtTimeFromRaw:output_type -> time_service.GetFmtTimeFromRawResp
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_time_proto_init() }
func file_time_proto_init() {
	if File_time_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_time_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRawTimeReq); i {
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
		file_time_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRawTimeResp); i {
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
		file_time_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFmtTimeReq); i {
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
		file_time_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFmtTimeResp); i {
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
		file_time_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFmtTimeFromRawReq); i {
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
		file_time_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFmtTimeFromRawResp); i {
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
			RawDescriptor: file_time_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_time_proto_goTypes,
		DependencyIndexes: file_time_proto_depIdxs,
		MessageInfos:      file_time_proto_msgTypes,
	}.Build()
	File_time_proto = out.File
	file_time_proto_rawDesc = nil
	file_time_proto_goTypes = nil
	file_time_proto_depIdxs = nil
}