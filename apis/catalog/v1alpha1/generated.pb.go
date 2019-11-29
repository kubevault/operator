/*
Copyright The KubeVault Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kubevault.dev/operator/apis/catalog/v1alpha1/generated.proto

package v1alpha1

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func (m *VaultServerVersion) Reset()      { *m = VaultServerVersion{} }
func (*VaultServerVersion) ProtoMessage() {}
func (*VaultServerVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d85d87e86e417be, []int{0}
}
func (m *VaultServerVersion) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VaultServerVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *VaultServerVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VaultServerVersion.Merge(m, src)
}
func (m *VaultServerVersion) XXX_Size() int {
	return m.Size()
}
func (m *VaultServerVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_VaultServerVersion.DiscardUnknown(m)
}

var xxx_messageInfo_VaultServerVersion proto.InternalMessageInfo

func (m *VaultServerVersionExporter) Reset()      { *m = VaultServerVersionExporter{} }
func (*VaultServerVersionExporter) ProtoMessage() {}
func (*VaultServerVersionExporter) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d85d87e86e417be, []int{1}
}
func (m *VaultServerVersionExporter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VaultServerVersionExporter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *VaultServerVersionExporter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VaultServerVersionExporter.Merge(m, src)
}
func (m *VaultServerVersionExporter) XXX_Size() int {
	return m.Size()
}
func (m *VaultServerVersionExporter) XXX_DiscardUnknown() {
	xxx_messageInfo_VaultServerVersionExporter.DiscardUnknown(m)
}

var xxx_messageInfo_VaultServerVersionExporter proto.InternalMessageInfo

func (m *VaultServerVersionList) Reset()      { *m = VaultServerVersionList{} }
func (*VaultServerVersionList) ProtoMessage() {}
func (*VaultServerVersionList) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d85d87e86e417be, []int{2}
}
func (m *VaultServerVersionList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VaultServerVersionList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *VaultServerVersionList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VaultServerVersionList.Merge(m, src)
}
func (m *VaultServerVersionList) XXX_Size() int {
	return m.Size()
}
func (m *VaultServerVersionList) XXX_DiscardUnknown() {
	xxx_messageInfo_VaultServerVersionList.DiscardUnknown(m)
}

var xxx_messageInfo_VaultServerVersionList proto.InternalMessageInfo

func (m *VaultServerVersionSpec) Reset()      { *m = VaultServerVersionSpec{} }
func (*VaultServerVersionSpec) ProtoMessage() {}
func (*VaultServerVersionSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d85d87e86e417be, []int{3}
}
func (m *VaultServerVersionSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VaultServerVersionSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *VaultServerVersionSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VaultServerVersionSpec.Merge(m, src)
}
func (m *VaultServerVersionSpec) XXX_Size() int {
	return m.Size()
}
func (m *VaultServerVersionSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_VaultServerVersionSpec.DiscardUnknown(m)
}

var xxx_messageInfo_VaultServerVersionSpec proto.InternalMessageInfo

func (m *VaultServerVersionUnsealer) Reset()      { *m = VaultServerVersionUnsealer{} }
func (*VaultServerVersionUnsealer) ProtoMessage() {}
func (*VaultServerVersionUnsealer) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d85d87e86e417be, []int{4}
}
func (m *VaultServerVersionUnsealer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VaultServerVersionUnsealer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *VaultServerVersionUnsealer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VaultServerVersionUnsealer.Merge(m, src)
}
func (m *VaultServerVersionUnsealer) XXX_Size() int {
	return m.Size()
}
func (m *VaultServerVersionUnsealer) XXX_DiscardUnknown() {
	xxx_messageInfo_VaultServerVersionUnsealer.DiscardUnknown(m)
}

var xxx_messageInfo_VaultServerVersionUnsealer proto.InternalMessageInfo

func (m *VaultServerVersionVault) Reset()      { *m = VaultServerVersionVault{} }
func (*VaultServerVersionVault) ProtoMessage() {}
func (*VaultServerVersionVault) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d85d87e86e417be, []int{5}
}
func (m *VaultServerVersionVault) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VaultServerVersionVault) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *VaultServerVersionVault) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VaultServerVersionVault.Merge(m, src)
}
func (m *VaultServerVersionVault) XXX_Size() int {
	return m.Size()
}
func (m *VaultServerVersionVault) XXX_DiscardUnknown() {
	xxx_messageInfo_VaultServerVersionVault.DiscardUnknown(m)
}

var xxx_messageInfo_VaultServerVersionVault proto.InternalMessageInfo

func init() {
	proto.RegisterType((*VaultServerVersion)(nil), "kubevault.dev.operator.apis.catalog.v1alpha1.VaultServerVersion")
	proto.RegisterType((*VaultServerVersionExporter)(nil), "kubevault.dev.operator.apis.catalog.v1alpha1.VaultServerVersionExporter")
	proto.RegisterType((*VaultServerVersionList)(nil), "kubevault.dev.operator.apis.catalog.v1alpha1.VaultServerVersionList")
	proto.RegisterType((*VaultServerVersionSpec)(nil), "kubevault.dev.operator.apis.catalog.v1alpha1.VaultServerVersionSpec")
	proto.RegisterType((*VaultServerVersionUnsealer)(nil), "kubevault.dev.operator.apis.catalog.v1alpha1.VaultServerVersionUnsealer")
	proto.RegisterType((*VaultServerVersionVault)(nil), "kubevault.dev.operator.apis.catalog.v1alpha1.VaultServerVersionVault")
}

func init() {
	proto.RegisterFile("kubevault.dev/operator/apis/catalog/v1alpha1/generated.proto", fileDescriptor_0d85d87e86e417be)
}

var fileDescriptor_0d85d87e86e417be = []byte{
	// 554 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x4f, 0x8b, 0xd3, 0x40,
	0x14, 0x6f, 0x76, 0x5b, 0xb6, 0xce, 0x2a, 0xca, 0x1c, 0x34, 0xf4, 0x90, 0x2d, 0xf5, 0x52, 0x41,
	0x27, 0xb6, 0x88, 0x78, 0x10, 0xd1, 0xb2, 0x0b, 0x0a, 0x8a, 0x90, 0xc5, 0x3d, 0x88, 0x07, 0xa7,
	0xe9, 0xdb, 0x34, 0xdb, 0xa6, 0x33, 0x4c, 0x26, 0x41, 0x6f, 0x7e, 0x04, 0x3f, 0x56, 0x8f, 0x0b,
	0x5e, 0xf6, 0x62, 0xb1, 0xf1, 0x8b, 0xc8, 0x4c, 0x26, 0x4d, 0x35, 0xbb, 0x58, 0xa9, 0xb7, 0xbc,
	0x37, 0xf3, 0xfb, 0xf3, 0x7e, 0xf3, 0x82, 0x9e, 0x4e, 0x92, 0x21, 0xa4, 0x34, 0x99, 0x4a, 0x32,
	0x82, 0xd4, 0x65, 0x1c, 0x04, 0x95, 0x4c, 0xb8, 0x94, 0x87, 0xb1, 0xeb, 0x53, 0x49, 0xa7, 0x2c,
	0x70, 0xd3, 0x1e, 0x9d, 0xf2, 0x31, 0xed, 0xb9, 0x01, 0xcc, 0xd4, 0x39, 0x8c, 0x08, 0x17, 0x4c,
	0x32, 0x7c, 0xff, 0x37, 0x34, 0x29, 0xd0, 0x44, 0xa1, 0x89, 0x41, 0x93, 0x02, 0xdd, 0x7a, 0x10,
	0x84, 0x72, 0x9c, 0x0c, 0x89, 0xcf, 0x22, 0x37, 0x60, 0x01, 0x73, 0x35, 0xc9, 0x30, 0x39, 0xd5,
	0x95, 0x2e, 0xf4, 0x57, 0x4e, 0xde, 0x7a, 0x34, 0x79, 0x12, 0x93, 0x90, 0x29, 0x2b, 0x11, 0xf5,
	0xc7, 0xe1, 0x0c, 0xc4, 0x67, 0x97, 0x4f, 0x82, 0xdc, 0x5b, 0x04, 0x92, 0xba, 0x69, 0xc5, 0x52,
	0xcb, 0xbd, 0x0a, 0x25, 0x92, 0x99, 0x0c, 0x23, 0xa8, 0x00, 0x1e, 0xff, 0x0d, 0x10, 0xfb, 0x63,
	0x88, 0xe8, 0x9f, 0xb8, 0xce, 0xc2, 0x42, 0xf8, 0x44, 0x8d, 0x7e, 0x0c, 0x22, 0x05, 0x71, 0x02,
	0x22, 0x0e, 0xd9, 0x0c, 0x7f, 0x44, 0x4d, 0x65, 0x6d, 0x44, 0x25, 0xb5, 0xad, 0xb6, 0xd5, 0xdd,
	0xef, 0x3f, 0x24, 0xb9, 0x02, 0x59, 0x57, 0x20, 0x7c, 0x12, 0xe4, 0x31, 0xa9, 0xdb, 0x24, 0xed,
	0x91, 0xb7, 0xc3, 0x33, 0xf0, 0xe5, 0x1b, 0x90, 0x74, 0x80, 0xe7, 0x8b, 0x83, 0x5a, 0xb6, 0x38,
	0x40, 0x65, 0xcf, 0x5b, 0xb1, 0xe2, 0x53, 0x54, 0x8f, 0x39, 0xf8, 0xf6, 0x8e, 0x66, 0x3f, 0x24,
	0xff, 0xf2, 0x06, 0xa4, 0xea, 0xf8, 0x98, 0x83, 0x3f, 0xb8, 0x6e, 0x14, 0xeb, 0xaa, 0xf2, 0x34,
	0x7f, 0xe7, 0x05, 0x6a, 0x55, 0x6f, 0x1f, 0x7d, 0xe2, 0x4c, 0x48, 0x10, 0xf8, 0x2e, 0x6a, 0x84,
	0x11, 0x0d, 0x40, 0x0f, 0x79, 0x6d, 0x70, 0xc3, 0x10, 0x34, 0x5e, 0xa9, 0xa6, 0x97, 0x9f, 0x75,
	0xbe, 0x5b, 0xe8, 0x76, 0x95, 0xe3, 0x75, 0x18, 0x4b, 0xfc, 0xa1, 0x92, 0x13, 0xd9, 0x2c, 0x27,
	0x85, 0xd6, 0x29, 0xdd, 0x32, 0x92, 0xcd, 0xa2, 0xb3, 0x96, 0x11, 0xa0, 0x46, 0x28, 0x21, 0x8a,
	0xed, 0x9d, 0xf6, 0x6e, 0x77, 0xbf, 0xff, 0x7c, 0xdb, 0x90, 0xd6, 0xe6, 0x53, 0xb4, 0x5e, 0xce,
	0xde, 0xf9, 0xb6, 0x7b, 0xd9, 0x7c, 0x2a, 0x43, 0x7c, 0x0f, 0xed, 0xa5, 0x79, 0x69, 0x12, 0xba,
	0x69, 0x18, 0xf6, 0xcc, 0x2d, 0xaf, 0x38, 0xc7, 0x67, 0xa8, 0xa1, 0xad, 0x99, 0x17, 0x3d, 0xda,
	0xd6, 0xac, 0xee, 0x94, 0x8e, 0x75, 0xe9, 0xe5, 0x12, 0x38, 0x45, 0xcd, 0x64, 0x16, 0x03, 0x9d,
	0x82, 0xb0, 0x77, 0xb5, 0xdc, 0xcb, 0x6d, 0xe5, 0xde, 0x19, 0xbe, 0xf2, 0x41, 0x8a, 0x8e, 0xb7,
	0xd2, 0x52, 0xba, 0x60, 0x56, 0xc7, 0xae, 0xff, 0x1f, 0xdd, 0x62, 0x15, 0x4b, 0xdd, 0xa2, 0xe3,
	0xad, 0xb4, 0x70, 0x1f, 0xa1, 0x11, 0x70, 0x01, 0xbe, 0xfa, 0x73, 0xed, 0x46, 0xdb, 0xea, 0x36,
	0xcb, 0xdf, 0xeb, 0x70, 0x75, 0xe2, 0xad, 0xdd, 0xba, 0x7c, 0xf1, 0x8b, 0x99, 0x36, 0x5b, 0xfc,
	0x67, 0xe8, 0xce, 0x15, 0xef, 0xb2, 0x11, 0x7e, 0x40, 0xe6, 0x4b, 0xa7, 0x76, 0xbe, 0x74, 0x6a,
	0x17, 0x4b, 0xa7, 0xf6, 0x25, 0x73, 0xac, 0x79, 0xe6, 0x58, 0xe7, 0x99, 0x63, 0x5d, 0x64, 0x8e,
	0xf5, 0x23, 0x73, 0xac, 0xaf, 0x3f, 0x9d, 0xda, 0xfb, 0x66, 0x91, 0xce, 0xaf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x80, 0x10, 0x69, 0x19, 0xc7, 0x05, 0x00, 0x00,
}

func (m *VaultServerVersion) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VaultServerVersion) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VaultServerVersion) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Spec.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.ObjectMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *VaultServerVersionExporter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VaultServerVersionExporter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VaultServerVersionExporter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= len(m.Image)
	copy(dAtA[i:], m.Image)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Image)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *VaultServerVersionList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VaultServerVersionList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VaultServerVersionList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for iNdEx := len(m.Items) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Items[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenerated(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.ListMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *VaultServerVersionSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VaultServerVersionSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VaultServerVersionSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i--
	if m.Deprecated {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i--
	dAtA[i] = 0x28
	{
		size, err := m.Exporter.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.Unsealer.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Vault.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	i -= len(m.Version)
	copy(dAtA[i:], m.Version)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Version)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *VaultServerVersionUnsealer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VaultServerVersionUnsealer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VaultServerVersionUnsealer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= len(m.Image)
	copy(dAtA[i:], m.Image)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Image)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *VaultServerVersionVault) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VaultServerVersionVault) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VaultServerVersionVault) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= len(m.Image)
	copy(dAtA[i:], m.Image)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Image)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenerated(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VaultServerVersion) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *VaultServerVersionExporter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Image)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *VaultServerVersionList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *VaultServerVersionSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Version)
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Vault.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Unsealer.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Exporter.Size()
	n += 1 + l + sovGenerated(uint64(l))
	n += 2
	return n
}

func (m *VaultServerVersionUnsealer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Image)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *VaultServerVersionVault) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Image)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func sovGenerated(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *VaultServerVersion) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&VaultServerVersion{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ObjectMeta), "ObjectMeta", "v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Spec:` + strings.Replace(strings.Replace(this.Spec.String(), "VaultServerVersionSpec", "VaultServerVersionSpec", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *VaultServerVersionExporter) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&VaultServerVersionExporter{`,
		`Image:` + fmt.Sprintf("%v", this.Image) + `,`,
		`}`,
	}, "")
	return s
}
func (this *VaultServerVersionList) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForItems := "[]VaultServerVersion{"
	for _, f := range this.Items {
		repeatedStringForItems += strings.Replace(strings.Replace(f.String(), "VaultServerVersion", "VaultServerVersion", 1), `&`, ``, 1) + ","
	}
	repeatedStringForItems += "}"
	s := strings.Join([]string{`&VaultServerVersionList{`,
		`ListMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ListMeta), "ListMeta", "v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + repeatedStringForItems + `,`,
		`}`,
	}, "")
	return s
}
func (this *VaultServerVersionSpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&VaultServerVersionSpec{`,
		`Version:` + fmt.Sprintf("%v", this.Version) + `,`,
		`Vault:` + strings.Replace(strings.Replace(this.Vault.String(), "VaultServerVersionVault", "VaultServerVersionVault", 1), `&`, ``, 1) + `,`,
		`Unsealer:` + strings.Replace(strings.Replace(this.Unsealer.String(), "VaultServerVersionUnsealer", "VaultServerVersionUnsealer", 1), `&`, ``, 1) + `,`,
		`Exporter:` + strings.Replace(strings.Replace(this.Exporter.String(), "VaultServerVersionExporter", "VaultServerVersionExporter", 1), `&`, ``, 1) + `,`,
		`Deprecated:` + fmt.Sprintf("%v", this.Deprecated) + `,`,
		`}`,
	}, "")
	return s
}
func (this *VaultServerVersionUnsealer) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&VaultServerVersionUnsealer{`,
		`Image:` + fmt.Sprintf("%v", this.Image) + `,`,
		`}`,
	}, "")
	return s
}
func (this *VaultServerVersionVault) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&VaultServerVersionVault{`,
		`Image:` + fmt.Sprintf("%v", this.Image) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *VaultServerVersion) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: VaultServerVersion: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VaultServerVersion: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *VaultServerVersionExporter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: VaultServerVersionExporter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VaultServerVersionExporter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Image", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Image = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *VaultServerVersionList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: VaultServerVersionList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VaultServerVersionList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, VaultServerVersion{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *VaultServerVersionSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: VaultServerVersionSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VaultServerVersionSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vault", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Vault.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unsealer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Unsealer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exporter", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Exporter.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deprecated", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Deprecated = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *VaultServerVersionUnsealer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: VaultServerVersionUnsealer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VaultServerVersionUnsealer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Image", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Image = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *VaultServerVersionVault) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: VaultServerVersionVault: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VaultServerVersionVault: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Image", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Image = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGenerated(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthGenerated
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)
