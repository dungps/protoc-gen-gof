// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by generate-protos. DO NOT EDIT.

package genid

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
)

const File_google_protobuf_struct_proto = "google/protobuf/struct.proto"

// Full and short names for google.protobuf.NullValue.
const (
	NullValue_enum_fullname = "google.protobuf.NullValue"
	NullValue_enum_name     = "NullValue"
)

// Names for google.protobuf.Struct.
const (
	Struct_message_name     protoreflect.Name     = "Struct"
	Struct_message_fullname protoreflect.FullName = "google.protobuf.Struct"
)

// Field names for google.protobuf.Struct.
const (
	Struct_Fields_field_name protoreflect.Name = "fields"

	Struct_Fields_field_fullname protoreflect.FullName = "google.protobuf.Struct.fields"
)

// Field numbers for google.protobuf.Struct.
const (
	Struct_Fields_field_number protoreflect.FieldNumber = 1
)

// Names for google.protobuf.Struct.FieldsEntry.
const (
	Struct_FieldsEntry_message_name     protoreflect.Name     = "FieldsEntry"
	Struct_FieldsEntry_message_fullname protoreflect.FullName = "google.protobuf.Struct.FieldsEntry"
)

// Field names for google.protobuf.Struct.FieldsEntry.
const (
	Struct_FieldsEntry_Key_field_name   protoreflect.Name = "key"
	Struct_FieldsEntry_Value_field_name protoreflect.Name = "value"

	Struct_FieldsEntry_Key_field_fullname   protoreflect.FullName = "google.protobuf.Struct.FieldsEntry.key"
	Struct_FieldsEntry_Value_field_fullname protoreflect.FullName = "google.protobuf.Struct.FieldsEntry.value"
)

// Field numbers for google.protobuf.Struct.FieldsEntry.
const (
	Struct_FieldsEntry_Key_field_number   protoreflect.FieldNumber = 1
	Struct_FieldsEntry_Value_field_number protoreflect.FieldNumber = 2
)

// Names for google.protobuf.Value.
const (
	Value_message_name     protoreflect.Name     = "Value"
	Value_message_fullname protoreflect.FullName = "google.protobuf.Value"
)

// Field names for google.protobuf.Value.
const (
	Value_NullValue_field_name   protoreflect.Name = "null_value"
	Value_NumberValue_field_name protoreflect.Name = "number_value"
	Value_StringValue_field_name protoreflect.Name = "string_value"
	Value_BoolValue_field_name   protoreflect.Name = "bool_value"
	Value_StructValue_field_name protoreflect.Name = "struct_value"
	Value_ListValue_field_name   protoreflect.Name = "list_value"

	Value_NullValue_field_fullname   protoreflect.FullName = "google.protobuf.Value.null_value"
	Value_NumberValue_field_fullname protoreflect.FullName = "google.protobuf.Value.number_value"
	Value_StringValue_field_fullname protoreflect.FullName = "google.protobuf.Value.string_value"
	Value_BoolValue_field_fullname   protoreflect.FullName = "google.protobuf.Value.bool_value"
	Value_StructValue_field_fullname protoreflect.FullName = "google.protobuf.Value.struct_value"
	Value_ListValue_field_fullname   protoreflect.FullName = "google.protobuf.Value.list_value"
)

// Field numbers for google.protobuf.Value.
const (
	Value_NullValue_field_number   protoreflect.FieldNumber = 1
	Value_NumberValue_field_number protoreflect.FieldNumber = 2
	Value_StringValue_field_number protoreflect.FieldNumber = 3
	Value_BoolValue_field_number   protoreflect.FieldNumber = 4
	Value_StructValue_field_number protoreflect.FieldNumber = 5
	Value_ListValue_field_number   protoreflect.FieldNumber = 6
)

// Oneof names for google.protobuf.Value.
const (
	Value_Kind_oneof_name protoreflect.Name = "kind"

	Value_Kind_oneof_fullname protoreflect.FullName = "google.protobuf.Value.kind"
)

// Names for google.protobuf.ListValue.
const (
	ListValue_message_name     protoreflect.Name     = "ListValue"
	ListValue_message_fullname protoreflect.FullName = "google.protobuf.ListValue"
)

// Field names for google.protobuf.ListValue.
const (
	ListValue_Values_field_name protoreflect.Name = "values"

	ListValue_Values_field_fullname protoreflect.FullName = "google.protobuf.ListValue.values"
)

// Field numbers for google.protobuf.ListValue.
const (
	ListValue_Values_field_number protoreflect.FieldNumber = 1
)