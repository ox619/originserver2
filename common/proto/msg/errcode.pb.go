// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.0
// source: msgproto/errcode.proto

package msg

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

type ErrCode int32

const (
	ErrCode_OK                  ErrCode = 0
	ErrCode_InterNalError       ErrCode = 1  //内部错误
	ErrCode_TokenError          ErrCode = 2  //Token验证错误
	ErrCode_ConnExceeded        ErrCode = 3  //超过最大连接数
	ErrCode_RepeatLoginReq      ErrCode = 4  //已经是登陆请求,不允许重复请求
	ErrCode_PlatTypeError       ErrCode = 5  //平台类型错误
	ErrCode_PlatIdError         ErrCode = 6  //平台Id错误
	ErrCode_ServerFull          ErrCode = 7  //服务器已经满载
	ErrCode_CloseServerError    ErrCode = 8  //正在关服
	ErrCode_SessionInvalid      ErrCode = 9  //Session失效
	ErrCode_LockLoginPleaseWait ErrCode = 10 //等会儿重试登陆
	ErrCode_LoginCDError        ErrCode = 11 //登陆CD错误
	ErrCode_DBReturnError       ErrCode = 12 //数据库错误返回
	ErrCode_LoginParamError     ErrCode = 13 //登录参数错误
)

// Enum value maps for ErrCode.
var (
	ErrCode_name = map[int32]string{
		0:  "OK",
		1:  "InterNalError",
		2:  "TokenError",
		3:  "ConnExceeded",
		4:  "RepeatLoginReq",
		5:  "PlatTypeError",
		6:  "PlatIdError",
		7:  "ServerFull",
		8:  "CloseServerError",
		9:  "SessionInvalid",
		10: "LockLoginPleaseWait",
		11: "LoginCDError",
		12: "DBReturnError",
		13: "LoginParamError",
	}
	ErrCode_value = map[string]int32{
		"OK":                  0,
		"InterNalError":       1,
		"TokenError":          2,
		"ConnExceeded":        3,
		"RepeatLoginReq":      4,
		"PlatTypeError":       5,
		"PlatIdError":         6,
		"ServerFull":          7,
		"CloseServerError":    8,
		"SessionInvalid":      9,
		"LockLoginPleaseWait": 10,
		"LoginCDError":        11,
		"DBReturnError":       12,
		"LoginParamError":     13,
	}
)

func (x ErrCode) Enum() *ErrCode {
	p := new(ErrCode)
	*p = x
	return p
}

func (x ErrCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrCode) Descriptor() protoreflect.EnumDescriptor {
	return file_msgproto_errcode_proto_enumTypes[0].Descriptor()
}

func (ErrCode) Type() protoreflect.EnumType {
	return &file_msgproto_errcode_proto_enumTypes[0]
}

func (x ErrCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrCode.Descriptor instead.
func (ErrCode) EnumDescriptor() ([]byte, []int) {
	return file_msgproto_errcode_proto_rawDescGZIP(), []int{0}
}

var File_msgproto_errcode_proto protoreflect.FileDescriptor

var file_msgproto_errcode_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6d, 0x73, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x72, 0x72, 0x63, 0x6f,
	0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d, 0x73, 0x67, 0x2a, 0x8b, 0x02,
	0x0a, 0x07, 0x45, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10,
	0x00, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6c, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x6e, 0x45, 0x78, 0x63, 0x65,
	0x65, 0x64, 0x65, 0x64, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x6c,
	0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x05, 0x12, 0x0f, 0x0a,
	0x0b, 0x50, 0x6c, 0x61, 0x74, 0x49, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x06, 0x12, 0x0e,
	0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x46, 0x75, 0x6c, 0x6c, 0x10, 0x07, 0x12, 0x14,
	0x0a, 0x10, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x10, 0x08, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x10, 0x09, 0x12, 0x17, 0x0a, 0x13, 0x4c, 0x6f, 0x63, 0x6b,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x57, 0x61, 0x69, 0x74, 0x10,
	0x0a, 0x12, 0x10, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x44, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x10, 0x0b, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x42, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x10, 0x0c, 0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x0d, 0x42, 0x07, 0x5a, 0x05, 0x2e,
	0x3b, 0x6d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msgproto_errcode_proto_rawDescOnce sync.Once
	file_msgproto_errcode_proto_rawDescData = file_msgproto_errcode_proto_rawDesc
)

func file_msgproto_errcode_proto_rawDescGZIP() []byte {
	file_msgproto_errcode_proto_rawDescOnce.Do(func() {
		file_msgproto_errcode_proto_rawDescData = protoimpl.X.CompressGZIP(file_msgproto_errcode_proto_rawDescData)
	})
	return file_msgproto_errcode_proto_rawDescData
}

var file_msgproto_errcode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_msgproto_errcode_proto_goTypes = []interface{}{
	(ErrCode)(0), // 0: msg.ErrCode
}
var file_msgproto_errcode_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_msgproto_errcode_proto_init() }
func file_msgproto_errcode_proto_init() {
	if File_msgproto_errcode_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_msgproto_errcode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_msgproto_errcode_proto_goTypes,
		DependencyIndexes: file_msgproto_errcode_proto_depIdxs,
		EnumInfos:         file_msgproto_errcode_proto_enumTypes,
	}.Build()
	File_msgproto_errcode_proto = out.File
	file_msgproto_errcode_proto_rawDesc = nil
	file_msgproto_errcode_proto_goTypes = nil
	file_msgproto_errcode_proto_depIdxs = nil
}