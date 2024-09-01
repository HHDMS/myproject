// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: course/course.proto

package course

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

type AddCourseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Teacher  string `protobuf:"bytes,2,opt,name=teacher,proto3" json:"teacher,omitempty"`
	Duration string `protobuf:"bytes,3,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *AddCourseResponse) Reset() {
	*x = AddCourseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_course_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCourseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCourseResponse) ProtoMessage() {}

func (x *AddCourseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_course_course_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCourseResponse.ProtoReflect.Descriptor instead.
func (*AddCourseResponse) Descriptor() ([]byte, []int) {
	return file_course_course_proto_rawDescGZIP(), []int{0}
}

func (x *AddCourseResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddCourseResponse) GetTeacher() string {
	if x != nil {
		return x.Teacher
	}
	return ""
}

func (x *AddCourseResponse) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

var File_course_course_proto protoreflect.FileDescriptor

var file_course_course_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x22, 0x5d, 0x0a,
	0x11, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0f, 0x5a, 0x0d,
	0x67, 0x69, 0x6e, 0x6a, 0x77, 0x74, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_course_course_proto_rawDescOnce sync.Once
	file_course_course_proto_rawDescData = file_course_course_proto_rawDesc
)

func file_course_course_proto_rawDescGZIP() []byte {
	file_course_course_proto_rawDescOnce.Do(func() {
		file_course_course_proto_rawDescData = protoimpl.X.CompressGZIP(file_course_course_proto_rawDescData)
	})
	return file_course_course_proto_rawDescData
}

var file_course_course_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_course_course_proto_goTypes = []any{
	(*AddCourseResponse)(nil), // 0: course.AddCourseResponse
}
var file_course_course_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_course_course_proto_init() }
func file_course_course_proto_init() {
	if File_course_course_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_course_course_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AddCourseResponse); i {
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
			RawDescriptor: file_course_course_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_course_course_proto_goTypes,
		DependencyIndexes: file_course_course_proto_depIdxs,
		MessageInfos:      file_course_course_proto_msgTypes,
	}.Build()
	File_course_course_proto = out.File
	file_course_course_proto_rawDesc = nil
	file_course_course_proto_goTypes = nil
	file_course_course_proto_depIdxs = nil
}
