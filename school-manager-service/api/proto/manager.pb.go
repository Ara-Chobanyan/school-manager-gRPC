// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: api/proto/manager.proto

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

type Student struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName    string  `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName     string  `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Grade        float32 `protobuf:"fixed32,4,opt,name=grade,proto3" json:"grade,omitempty"`
	Behavior     float32 `protobuf:"fixed32,5,opt,name=behavior,proto3" json:"behavior,omitempty"`
	Satisfaction float32 `protobuf:"fixed32,6,opt,name=satisfaction,proto3" json:"satisfaction,omitempty"`
	Critical     bool    `protobuf:"varint,7,opt,name=critical,proto3" json:"critical,omitempty"`
	Exceptional  bool    `protobuf:"varint,8,opt,name=exceptional,proto3" json:"exceptional,omitempty"`
	Tutelary     bool    `protobuf:"varint,9,opt,name=tutelary,proto3" json:"tutelary,omitempty"`
	Staff        bool    `protobuf:"varint,10,opt,name=staff,proto3" json:"staff,omitempty"`
}

func (x *Student) Reset() {
	*x = Student{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Student) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Student) ProtoMessage() {}

func (x *Student) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Student.ProtoReflect.Descriptor instead.
func (*Student) Descriptor() ([]byte, []int) {
	return file_api_proto_manager_proto_rawDescGZIP(), []int{0}
}

func (x *Student) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Student) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Student) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Student) GetGrade() float32 {
	if x != nil {
		return x.Grade
	}
	return 0
}

func (x *Student) GetBehavior() float32 {
	if x != nil {
		return x.Behavior
	}
	return 0
}

func (x *Student) GetSatisfaction() float32 {
	if x != nil {
		return x.Satisfaction
	}
	return 0
}

func (x *Student) GetCritical() bool {
	if x != nil {
		return x.Critical
	}
	return false
}

func (x *Student) GetExceptional() bool {
	if x != nil {
		return x.Exceptional
	}
	return false
}

func (x *Student) GetTutelary() bool {
	if x != nil {
		return x.Tutelary
	}
	return false
}

func (x *Student) GetStaff() bool {
	if x != nil {
		return x.Staff
	}
	return false
}

type Helper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Helper) Reset() {
	*x = Helper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Helper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Helper) ProtoMessage() {}

func (x *Helper) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Helper.ProtoReflect.Descriptor instead.
func (*Helper) Descriptor() ([]byte, []int) {
	return file_api_proto_manager_proto_rawDescGZIP(), []int{1}
}

func (x *Helper) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type StudentList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Students []*Student `protobuf:"bytes,1,rep,name=students,proto3" json:"students,omitempty"`
}

func (x *StudentList) Reset() {
	*x = StudentList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentList) ProtoMessage() {}

func (x *StudentList) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentList.ProtoReflect.Descriptor instead.
func (*StudentList) Descriptor() ([]byte, []int) {
	return file_api_proto_manager_proto_rawDescGZIP(), []int{2}
}

func (x *StudentList) GetStudents() []*Student {
	if x != nil {
		return x.Students
	}
	return nil
}

type StudentParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StudentParam) Reset() {
	*x = StudentParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentParam) ProtoMessage() {}

func (x *StudentParam) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentParam.ProtoReflect.Descriptor instead.
func (*StudentParam) Descriptor() ([]byte, []int) {
	return file_api_proto_manager_proto_rawDescGZIP(), []int{3}
}

var File_api_proto_manager_proto protoreflect.FileDescriptor

var file_api_proto_manager_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x22, 0x9b, 0x02, 0x0a, 0x07, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72,
	0x61, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x08, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x12, 0x22, 0x0a, 0x0c,
	0x73, 0x61, 0x74, 0x69, 0x73, 0x66, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0c, 0x73, 0x61, 0x74, 0x69, 0x73, 0x66, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x12, 0x20, 0x0a, 0x0b,
	0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x1a,
	0x0a, 0x08, 0x74, 0x75, 0x74, 0x65, 0x6c, 0x61, 0x72, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x74, 0x75, 0x74, 0x65, 0x6c, 0x61, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x66, 0x66, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x73, 0x74, 0x61, 0x66, 0x66,
	0x22, 0x20, 0x0a, 0x06, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x3b, 0x0a, 0x0b, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x2c, 0x0a, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x22,
	0x0e, 0x0a, 0x0c, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x32,
	0x7c, 0x0a, 0x0d, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x12, 0x33, 0x0a, 0x08, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x12, 0x0f, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x1a, 0x14, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0b, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x0f, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x48,
	0x65, 0x6c, 0x70, 0x65, 0x72, 0x1a, 0x14, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x43, 0x5a,
	0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x72, 0x61, 0x2d,
	0x43, 0x68, 0x6f, 0x62, 0x61, 0x6e, 0x79, 0x61, 0x6e, 0x2f, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c,
	0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2d,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_manager_proto_rawDescOnce sync.Once
	file_api_proto_manager_proto_rawDescData = file_api_proto_manager_proto_rawDesc
)

func file_api_proto_manager_proto_rawDescGZIP() []byte {
	file_api_proto_manager_proto_rawDescOnce.Do(func() {
		file_api_proto_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_manager_proto_rawDescData)
	})
	return file_api_proto_manager_proto_rawDescData
}

var file_api_proto_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_proto_manager_proto_goTypes = []interface{}{
	(*Student)(nil),      // 0: manager.Student
	(*Helper)(nil),       // 1: manager.Helper
	(*StudentList)(nil),  // 2: manager.StudentList
	(*StudentParam)(nil), // 3: manager.StudentParam
}
var file_api_proto_manager_proto_depIdxs = []int32{
	0, // 0: manager.StudentList.students:type_name -> manager.Student
	1, // 1: manager.SchoolManager.Behavior:input_type -> manager.Helper
	1, // 2: manager.SchoolManager.Exceptional:input_type -> manager.Helper
	2, // 3: manager.SchoolManager.Behavior:output_type -> manager.StudentList
	2, // 4: manager.SchoolManager.Exceptional:output_type -> manager.StudentList
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_proto_manager_proto_init() }
func file_api_proto_manager_proto_init() {
	if File_api_proto_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Student); i {
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
		file_api_proto_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Helper); i {
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
		file_api_proto_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudentList); i {
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
		file_api_proto_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudentParam); i {
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
			RawDescriptor: file_api_proto_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_manager_proto_goTypes,
		DependencyIndexes: file_api_proto_manager_proto_depIdxs,
		MessageInfos:      file_api_proto_manager_proto_msgTypes,
	}.Build()
	File_api_proto_manager_proto = out.File
	file_api_proto_manager_proto_rawDesc = nil
	file_api_proto_manager_proto_goTypes = nil
	file_api_proto_manager_proto_depIdxs = nil
}
