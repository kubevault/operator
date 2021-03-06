/*
Copyright AppsCode Inc. and Contributors

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
// source: kubevault.dev/apimachinery/apis/config/v1alpha1/generated.proto

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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func (m *AWSAuthConfig) Reset()      { *m = AWSAuthConfig{} }
func (*AWSAuthConfig) ProtoMessage() {}
func (*AWSAuthConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b590075f4b07411, []int{0}
}
func (m *AWSAuthConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AWSAuthConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *AWSAuthConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AWSAuthConfig.Merge(m, src)
}
func (m *AWSAuthConfig) XXX_Size() int {
	return m.Size()
}
func (m *AWSAuthConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AWSAuthConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AWSAuthConfig proto.InternalMessageInfo

func (m *AzureAuthConfig) Reset()      { *m = AzureAuthConfig{} }
func (*AzureAuthConfig) ProtoMessage() {}
func (*AzureAuthConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b590075f4b07411, []int{1}
}
func (m *AzureAuthConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AzureAuthConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *AzureAuthConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AzureAuthConfig.Merge(m, src)
}
func (m *AzureAuthConfig) XXX_Size() int {
	return m.Size()
}
func (m *AzureAuthConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AzureAuthConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AzureAuthConfig proto.InternalMessageInfo

func (m *KubernetesAuthConfig) Reset()      { *m = KubernetesAuthConfig{} }
func (*KubernetesAuthConfig) ProtoMessage() {}
func (*KubernetesAuthConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b590075f4b07411, []int{2}
}
func (m *KubernetesAuthConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KubernetesAuthConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *KubernetesAuthConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KubernetesAuthConfig.Merge(m, src)
}
func (m *KubernetesAuthConfig) XXX_Size() int {
	return m.Size()
}
func (m *KubernetesAuthConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_KubernetesAuthConfig.DiscardUnknown(m)
}

var xxx_messageInfo_KubernetesAuthConfig proto.InternalMessageInfo

func (m *VaultServerConfiguration) Reset()      { *m = VaultServerConfiguration{} }
func (*VaultServerConfiguration) ProtoMessage() {}
func (*VaultServerConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b590075f4b07411, []int{3}
}
func (m *VaultServerConfiguration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VaultServerConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *VaultServerConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VaultServerConfiguration.Merge(m, src)
}
func (m *VaultServerConfiguration) XXX_Size() int {
	return m.Size()
}
func (m *VaultServerConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_VaultServerConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_VaultServerConfiguration proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AWSAuthConfig)(nil), "kubevault.dev.apimachinery.apis.config.v1alpha1.AWSAuthConfig")
	proto.RegisterType((*AzureAuthConfig)(nil), "kubevault.dev.apimachinery.apis.config.v1alpha1.AzureAuthConfig")
	proto.RegisterType((*KubernetesAuthConfig)(nil), "kubevault.dev.apimachinery.apis.config.v1alpha1.KubernetesAuthConfig")
	proto.RegisterType((*VaultServerConfiguration)(nil), "kubevault.dev.apimachinery.apis.config.v1alpha1.VaultServerConfiguration")
}

func init() {
	proto.RegisterFile("kubevault.dev/apimachinery/apis/config/v1alpha1/generated.proto", fileDescriptor_5b590075f4b07411)
}

var fileDescriptor_5b590075f4b07411 = []byte{
	// 629 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xc7, 0xe3, 0xa6, 0xed, 0x2f, 0xd9, 0xfc, 0x08, 0x74, 0x41, 0x28, 0xe4, 0x60, 0x47, 0x39,
	0x40, 0x0f, 0x60, 0xab, 0x15, 0x20, 0x4e, 0x85, 0xa4, 0xa5, 0xa5, 0x20, 0xa1, 0x6a, 0x03, 0xa9,
	0xe0, 0xb6, 0x71, 0xa6, 0xb1, 0x95, 0xd8, 0x6b, 0xd6, 0xbb, 0xae, 0xe0, 0xc4, 0x23, 0xf0, 0x0c,
	0x3c, 0x4d, 0x25, 0x2e, 0x3d, 0xf6, 0x14, 0x51, 0xf3, 0x02, 0x3c, 0x02, 0xf2, 0x66, 0xd3, 0xe6,
	0x4f, 0x21, 0xea, 0xcd, 0xb3, 0xf3, 0xfd, 0x7e, 0x66, 0x67, 0xec, 0x31, 0x7a, 0xde, 0x97, 0x1d,
	0x48, 0xa8, 0x1c, 0x08, 0xbb, 0x0b, 0x89, 0x43, 0x23, 0x3f, 0xa0, 0xae, 0xe7, 0x87, 0xc0, 0x3f,
	0x67, 0x41, 0xec, 0xb8, 0x2c, 0x3c, 0xf2, 0x7b, 0x4e, 0xb2, 0x41, 0x07, 0x91, 0x47, 0x37, 0x9c,
	0x1e, 0x84, 0xc0, 0xa9, 0x80, 0xae, 0x1d, 0x71, 0x26, 0x18, 0x76, 0xa6, 0x00, 0xf6, 0x24, 0x20,
	0x0b, 0x62, 0x7b, 0x04, 0xb0, 0xc7, 0x80, 0xea, 0xa3, 0x9e, 0x2f, 0x3c, 0xd9, 0xb1, 0x5d, 0x16,
	0x38, 0x3d, 0xd6, 0x63, 0x8e, 0xe2, 0x74, 0xe4, 0x91, 0x8a, 0x54, 0xa0, 0x9e, 0x46, 0xfc, 0xea,
	0xe3, 0xfe, 0xb3, 0xd8, 0xf6, 0xd9, 0xf4, 0xcd, 0xa2, 0x7e, 0x6f, 0x74, 0xbb, 0x00, 0x04, 0x75,
	0x92, 0xb9, 0x5b, 0x55, 0x9d, 0xbf, 0xb9, 0xb8, 0x0c, 0x85, 0x1f, 0xc0, 0x9c, 0xe1, 0xe9, 0x22,
	0x43, 0xec, 0x7a, 0x10, 0xd0, 0x59, 0x5f, 0x7d, 0x17, 0xdd, 0x68, 0x1c, 0xb6, 0x1a, 0x52, 0x78,
	0xdb, 0xaa, 0x4f, 0xfc, 0x04, 0x95, 0x3c, 0xa0, 0x5d, 0xe0, 0x6d, 0x3a, 0x90, 0x50, 0x31, 0x6a,
	0xc6, 0x7a, 0xb1, 0x79, 0xfb, 0x64, 0x68, 0xe5, 0xd2, 0xa1, 0x55, 0x7a, 0x75, 0x99, 0x22, 0x93,
	0xba, 0xfa, 0x6f, 0x03, 0xdd, 0x6c, 0x7c, 0x91, 0x1c, 0x26, 0x50, 0x5b, 0xa8, 0x1c, 0xcb, 0x4e,
	0xec, 0x72, 0x3f, 0x12, 0x3e, 0x0b, 0xf7, 0x77, 0x34, 0xed, 0xae, 0xa6, 0x95, 0x5b, 0x53, 0x59,
	0x32, 0xa3, 0xc6, 0x7b, 0x68, 0x8d, 0x43, 0xcc, 0x24, 0x77, 0x61, 0x8f, 0x33, 0x19, 0xbd, 0xa5,
	0x01, 0x54, 0x96, 0x14, 0xe2, 0x9e, 0x46, 0xac, 0x91, 0x59, 0x01, 0x99, 0xf7, 0xe0, 0xfb, 0x68,
	0x35, 0x09, 0x94, 0x3b, 0xaf, 0xdc, 0x65, 0xed, 0x5e, 0x6d, 0xab, 0x53, 0xa2, 0xb3, 0xf8, 0x21,
	0x2a, 0x24, 0x41, 0x1c, 0x2b, 0xe5, 0xb2, 0x52, 0xde, 0xd2, 0xca, 0x42, 0x5b, 0x9f, 0x93, 0x0b,
	0x45, 0xfd, 0xc7, 0x12, 0xba, 0xf3, 0x46, 0x76, 0x80, 0x87, 0x20, 0x20, 0x9e, 0xe8, 0xfb, 0x35,
	0xc2, 0x31, 0xf0, 0xc4, 0x77, 0xa1, 0xe1, 0xba, 0x4c, 0x86, 0x42, 0x01, 0x47, 0xbd, 0x57, 0x35,
	0x10, 0xb7, 0xe6, 0x14, 0xe4, 0x0a, 0x17, 0xfe, 0x84, 0x2c, 0xc1, 0xfa, 0x10, 0x12, 0x48, 0x7c,
	0x38, 0x06, 0x3e, 0x6f, 0xd3, 0x13, 0x79, 0xa0, 0xc1, 0xd6, 0xbb, 0x7f, 0xcb, 0xc9, 0x22, 0x1e,
	0x16, 0xa8, 0x26, 0x63, 0x38, 0x60, 0xdd, 0xe9, 0xdc, 0x2e, 0xe3, 0xdb, 0xad, 0xfd, 0x1d, 0xee,
	0x27, 0xc0, 0xd5, 0x1c, 0x0b, 0xcd, 0x75, 0x5d, 0xb3, 0xf6, 0x7e, 0x81, 0x9e, 0x2c, 0x24, 0xd6,
	0xbf, 0xe7, 0x51, 0xa5, 0x9d, 0xad, 0x61, 0xa6, 0x01, 0x3e, 0x1a, 0xa5, 0xe4, 0x34, 0xfb, 0x16,
	0x70, 0x0d, 0x2d, 0x47, 0x54, 0x78, 0x7a, 0x86, 0xff, 0xeb, 0xb2, 0xcb, 0x07, 0x54, 0x78, 0x44,
	0x65, 0xb0, 0x83, 0x8a, 0x6a, 0x89, 0x09, 0x1b, 0x8c, 0x27, 0xb2, 0xa6, 0x65, 0xc5, 0xf6, 0x38,
	0x41, 0x2e, 0x35, 0x58, 0x22, 0xd4, 0xbf, 0x78, 0x79, 0xaa, 0x9f, 0xd2, 0xe6, 0x4b, 0xfb, 0x9a,
	0x3f, 0x03, 0xfb, 0xaa, 0xf7, 0xdf, 0x2c, 0xa7, 0x43, 0x0b, 0x5d, 0x66, 0xc8, 0x44, 0x21, 0x4c,
	0xd1, 0x0a, 0xcd, 0xd6, 0x44, 0x7d, 0x5f, 0xa5, 0xcd, 0x17, 0xd7, 0xae, 0x38, 0xb3, 0x64, 0xcd,
	0x62, 0x3a, 0xb4, 0x56, 0xd4, 0x21, 0x19, 0x91, 0xf1, 0x07, 0x94, 0xa7, 0xc7, 0x71, 0x65, 0x45,
	0x15, 0xd8, 0xba, 0x7e, 0x81, 0xc9, 0xdf, 0x41, 0xf3, 0xbf, 0x74, 0x68, 0xe5, 0x1b, 0x87, 0x2d,
	0x92, 0x31, 0x9b, 0xf6, 0xc9, 0xb9, 0x99, 0x3b, 0x3d, 0x37, 0x73, 0x67, 0xe7, 0x66, 0xee, 0x6b,
	0x6a, 0x1a, 0x27, 0xa9, 0x69, 0x9c, 0xa6, 0xa6, 0x71, 0x96, 0x9a, 0xc6, 0xcf, 0xd4, 0x34, 0xbe,
	0xfd, 0x32, 0x73, 0x1f, 0x0b, 0x63, 0xd6, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x92, 0x87, 0xb8,
	0x3a, 0x9e, 0x05, 0x00, 0x00,
}

func (m *AWSAuthConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AWSAuthConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AWSAuthConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= len(m.HeaderValue)
	copy(dAtA[i:], m.HeaderValue)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.HeaderValue)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *AzureAuthConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AzureAuthConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AzureAuthConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= len(m.VmssName)
	copy(dAtA[i:], m.VmssName)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.VmssName)))
	i--
	dAtA[i] = 0x22
	i -= len(m.VmName)
	copy(dAtA[i:], m.VmName)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.VmName)))
	i--
	dAtA[i] = 0x1a
	i -= len(m.ResourceGroupName)
	copy(dAtA[i:], m.ResourceGroupName)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.ResourceGroupName)))
	i--
	dAtA[i] = 0x12
	i -= len(m.SubscriptionID)
	copy(dAtA[i:], m.SubscriptionID)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.SubscriptionID)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *KubernetesAuthConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KubernetesAuthConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KubernetesAuthConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i--
	if m.UsePodServiceAccountForCSIDriver {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i--
	dAtA[i] = 0x18
	i -= len(m.TokenReviewerServiceAccountName)
	copy(dAtA[i:], m.TokenReviewerServiceAccountName)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.TokenReviewerServiceAccountName)))
	i--
	dAtA[i] = 0x12
	i -= len(m.ServiceAccountName)
	copy(dAtA[i:], m.ServiceAccountName)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.ServiceAccountName)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *VaultServerConfiguration) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VaultServerConfiguration) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VaultServerConfiguration) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AWS != nil {
		{
			size, err := m.AWS.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.Azure != nil {
		{
			size, err := m.Azure.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Kubernetes != nil {
		{
			size, err := m.Kubernetes.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	i -= len(m.VaultRole)
	copy(dAtA[i:], m.VaultRole)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.VaultRole)))
	i--
	dAtA[i] = 0x12
	i -= len(m.Path)
	copy(dAtA[i:], m.Path)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Path)))
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
func (m *AWSAuthConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.HeaderValue)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *AzureAuthConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SubscriptionID)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.ResourceGroupName)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.VmName)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.VmssName)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *KubernetesAuthConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ServiceAccountName)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.TokenReviewerServiceAccountName)
	n += 1 + l + sovGenerated(uint64(l))
	n += 2
	return n
}

func (m *VaultServerConfiguration) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Path)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.VaultRole)
	n += 1 + l + sovGenerated(uint64(l))
	if m.Kubernetes != nil {
		l = m.Kubernetes.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.Azure != nil {
		l = m.Azure.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.AWS != nil {
		l = m.AWS.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *AWSAuthConfig) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AWSAuthConfig{`,
		`HeaderValue:` + fmt.Sprintf("%v", this.HeaderValue) + `,`,
		`}`,
	}, "")
	return s
}
func (this *AzureAuthConfig) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AzureAuthConfig{`,
		`SubscriptionID:` + fmt.Sprintf("%v", this.SubscriptionID) + `,`,
		`ResourceGroupName:` + fmt.Sprintf("%v", this.ResourceGroupName) + `,`,
		`VmName:` + fmt.Sprintf("%v", this.VmName) + `,`,
		`VmssName:` + fmt.Sprintf("%v", this.VmssName) + `,`,
		`}`,
	}, "")
	return s
}
func (this *KubernetesAuthConfig) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&KubernetesAuthConfig{`,
		`ServiceAccountName:` + fmt.Sprintf("%v", this.ServiceAccountName) + `,`,
		`TokenReviewerServiceAccountName:` + fmt.Sprintf("%v", this.TokenReviewerServiceAccountName) + `,`,
		`UsePodServiceAccountForCSIDriver:` + fmt.Sprintf("%v", this.UsePodServiceAccountForCSIDriver) + `,`,
		`}`,
	}, "")
	return s
}
func (this *VaultServerConfiguration) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&VaultServerConfiguration{`,
		`Path:` + fmt.Sprintf("%v", this.Path) + `,`,
		`VaultRole:` + fmt.Sprintf("%v", this.VaultRole) + `,`,
		`Kubernetes:` + strings.Replace(this.Kubernetes.String(), "KubernetesAuthConfig", "KubernetesAuthConfig", 1) + `,`,
		`Azure:` + strings.Replace(this.Azure.String(), "AzureAuthConfig", "AzureAuthConfig", 1) + `,`,
		`AWS:` + strings.Replace(this.AWS.String(), "AWSAuthConfig", "AWSAuthConfig", 1) + `,`,
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
func (m *AWSAuthConfig) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: AWSAuthConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AWSAuthConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeaderValue", wireType)
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
			m.HeaderValue = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
func (m *AzureAuthConfig) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: AzureAuthConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AzureAuthConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubscriptionID", wireType)
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
			m.SubscriptionID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceGroupName", wireType)
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
			m.ResourceGroupName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VmName", wireType)
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
			m.VmName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VmssName", wireType)
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
			m.VmssName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
func (m *KubernetesAuthConfig) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: KubernetesAuthConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KubernetesAuthConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceAccountName", wireType)
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
			m.ServiceAccountName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenReviewerServiceAccountName", wireType)
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
			m.TokenReviewerServiceAccountName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UsePodServiceAccountForCSIDriver", wireType)
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
			m.UsePodServiceAccountForCSIDriver = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
func (m *VaultServerConfiguration) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: VaultServerConfiguration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VaultServerConfiguration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
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
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VaultRole", wireType)
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
			m.VaultRole = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kubernetes", wireType)
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
			if m.Kubernetes == nil {
				m.Kubernetes = &KubernetesAuthConfig{}
			}
			if err := m.Kubernetes.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Azure", wireType)
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
			if m.Azure == nil {
				m.Azure = &AzureAuthConfig{}
			}
			if err := m.Azure.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AWS", wireType)
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
			if m.AWS == nil {
				m.AWS = &AWSAuthConfig{}
			}
			if err := m.AWS.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
	depth := 0
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
		case 1:
			iNdEx += 8
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
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenerated
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenerated
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenerated        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenerated = fmt.Errorf("proto: unexpected end of group")
)
