// Copyright 2021 Security Scorecard Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: cron/data/metadata.proto

package data

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

type ShardMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShardLoc  *string `protobuf:"bytes,1,opt,name=shard_loc,json=shardLoc,proto3,oneof" json:"shard_loc,omitempty"`
	NumShard  *int32  `protobuf:"varint,2,opt,name=num_shard,json=numShard,proto3,oneof" json:"num_shard,omitempty"`
	CommitSha *string `protobuf:"bytes,3,opt,name=commit_sha,json=commitSha,proto3,oneof" json:"commit_sha,omitempty"`
}

func (x *ShardMetadata) Reset() {
	*x = ShardMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cron_data_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShardMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShardMetadata) ProtoMessage() {}

func (x *ShardMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_cron_data_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShardMetadata.ProtoReflect.Descriptor instead.
func (*ShardMetadata) Descriptor() ([]byte, []int) {
	return file_cron_data_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *ShardMetadata) GetShardLoc() string {
	if x != nil && x.ShardLoc != nil {
		return *x.ShardLoc
	}
	return ""
}

func (x *ShardMetadata) GetNumShard() int32 {
	if x != nil && x.NumShard != nil {
		return *x.NumShard
	}
	return 0
}

func (x *ShardMetadata) GetCommitSha() string {
	if x != nil && x.CommitSha != nil {
		return *x.CommitSha
	}
	return ""
}

var File_cron_data_metadata_proto protoreflect.FileDescriptor

var file_cron_data_metadata_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x72, 0x6f, 0x6e, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x6f, 0x73, 0x73, 0x66,
	0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x63, 0x72, 0x6f, 0x6e, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x22, 0xa2, 0x01,
	0x0a, 0x0d, 0x53, 0x68, 0x61, 0x72, 0x64, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x20, 0x0a, 0x09, 0x73, 0x68, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x6f, 0x63, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x73, 0x68, 0x61, 0x72, 0x64, 0x4c, 0x6f, 0x63, 0x88, 0x01,
	0x01, 0x12, 0x20, 0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x08, 0x6e, 0x75, 0x6d, 0x53, 0x68, 0x61, 0x72, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x73, 0x68,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x53, 0x68, 0x61, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x73, 0x68, 0x61, 0x72,
	0x64, 0x5f, 0x6c, 0x6f, 0x63, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6e, 0x75, 0x6d, 0x5f, 0x73, 0x68,
	0x61, 0x72, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x73,
	0x68, 0x61, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6f, 0x73, 0x73, 0x66, 0x2f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x2f,
	0x63, 0x72, 0x6f, 0x6e, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_cron_data_metadata_proto_rawDescOnce sync.Once
	file_cron_data_metadata_proto_rawDescData = file_cron_data_metadata_proto_rawDesc
)

func file_cron_data_metadata_proto_rawDescGZIP() []byte {
	file_cron_data_metadata_proto_rawDescOnce.Do(func() {
		file_cron_data_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_cron_data_metadata_proto_rawDescData)
	})
	return file_cron_data_metadata_proto_rawDescData
}

var file_cron_data_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cron_data_metadata_proto_goTypes = []interface{}{
	(*ShardMetadata)(nil), // 0: ossf.scorecard.cron.internal.data.ShardMetadata
}
var file_cron_data_metadata_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cron_data_metadata_proto_init() }
func file_cron_data_metadata_proto_init() {
	if File_cron_data_metadata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cron_data_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShardMetadata); i {
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
	file_cron_data_metadata_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cron_data_metadata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cron_data_metadata_proto_goTypes,
		DependencyIndexes: file_cron_data_metadata_proto_depIdxs,
		MessageInfos:      file_cron_data_metadata_proto_msgTypes,
	}.Build()
	File_cron_data_metadata_proto = out.File
	file_cron_data_metadata_proto_rawDesc = nil
	file_cron_data_metadata_proto_goTypes = nil
	file_cron_data_metadata_proto_depIdxs = nil
}
