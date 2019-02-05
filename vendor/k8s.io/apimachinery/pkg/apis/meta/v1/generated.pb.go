/*
	Package v1 is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/meta/v1/generated.proto

	It has these top-level messages:
		APIGroup
		APIGroupList
		APIResource
		APIResourceList
		APIVersions
		DeleteOptions
		Duration
		ExportOptions
		GetOptions
		GroupKind
		GroupResource
		GroupVersion
		GroupVersionForDiscovery
		GroupVersionKind
		GroupVersionResource
		LabelSelector
		LabelSelectorRequirement
		ListMeta
		ListOptions
		ObjectMeta
		OwnerReference
		Preconditions
		RootPaths
		ServerAddressByClientCIDR
		Status
		StatusCause
		StatusDetails
		Time
		Timestamp
		TypeMeta
		Verbs
		WatchEvent
*/
package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import time "time"
import k8s_io_apimachinery_pkg_types "k8s.io/apimachinery/pkg/types"

import strings "strings"
import reflect "reflect"
import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

func (m *APIGroup) Reset()			{ *m = APIGroup{} }
func (*APIGroup) ProtoMessage()			{}
func (*APIGroup) Descriptor() ([]byte, []int)	{ return fileDescriptorGenerated, []int{0} }

func (m *APIGroupList) Reset()				{ *m = APIGroupList{} }
func (*APIGroupList) ProtoMessage()			{}
func (*APIGroupList) Descriptor() ([]byte, []int)	{ return fileDescriptorGenerated, []int{1} }

func (m *APIResource) Reset()				{ *m = APIResource{} }
func (*APIResource) ProtoMessage()			{}
func (*APIResource) Descriptor() ([]byte, []int)	{ return fileDescriptorGenerated, []int{2} }

func (m *APIResourceList) Reset()			{ *m = APIResourceList{} }
func (*APIResourceList) ProtoMessage()			{}
func (*APIResourceList) Descriptor() ([]byte, []int)	{ return fileDescriptorGenerated, []int{3} }

func (m *APIVersions) Reset()				{ *m = APIVersions{} }
func (*APIVersions) ProtoMessage()			{}
func (*APIVersions) Descriptor() ([]byte, []int)	{ return fileDescriptorGenerated, []int{4} }

func (m *DeleteOptions) Reset()				{ *m = DeleteOptions{} }
func (*DeleteOptions) ProtoMessage()			{}
func (*DeleteOptions) Descriptor() ([]byte, []int)	{ return fileDescriptorGenerated, []int{5} }

func (m *Duration) Reset()			{ *m = Duration{} }
func (*Duration) ProtoMessage()			{}
func (*Duration) Descriptor() ([]byte, []int)	{ return fileDescriptorGenerated, []int{6} }

func (m *ExportOptions) Reset()				{ *m = ExportOptions{} }
func (*ExportOptions) ProtoMessage()			{}
func (*ExportOptions) Descriptor() ([]byte, []int)	{ return fileDescriptorGenerated, []int{7} }

func (m *GetOptions) Reset()			{ *m = GetOptions{} }
func (*GetOptions) ProtoMessage()		{}
func (*GetOptions) Descriptor() ([]byte, []int)	{ return fileDescriptorGenerated, []int{8} }

func (m *GroupKind) Reset()			{ *m = GroupKind{} }
func (*GroupKind) ProtoMessage()		{}
func (*GroupKind) Descriptor() ([]byte, []int)	{ return fileDescriptorGenerated, []int{9} }

func (m *GroupResource) Reset()				{ *m = GroupResource{} }
func (*GroupResource) ProtoMessage()			{}
func (*GroupResource) Descriptor() ([]byte, []int)	{ return fileDescriptorGenerated, []int{10} }

func (m *GroupVersion) Reset()				{ *m = GroupVersion{} }
func (*GroupVersion) ProtoMessage()			{}
func (*GroupVersion) Descriptor() ([]byte, []int)	{ return fileDescriptorGenerated, []int{11} }

func (m *GroupVersionForDiscovery) Reset()	{ *m = GroupVersionForDiscovery{} }
func (*GroupVersionForDiscovery) ProtoMessage()	{}
func (*GroupVersionForDiscovery) Descriptor() ([]byte, []int) {
	return fileDescriptorGenerated, []int{12}
}

func (m *GroupVersionKind) Reset()			{ *m = GroupVersionKind{} }
func (*GroupVersionKind) ProtoMessage()			{}
func (*GroupVersionKind) Descriptor() ([]byte, []int)	{ return fileDescriptorGenerated, []int{13} }

func (m *GroupVersionResource) Reset()				{ *m = GroupVersionResource{} }
func (*GroupVersionResource) ProtoMessage()			{}
func (*GroupVersionResource) Descriptor() ([]byte, []int)	{ return fileDescriptorGenerated, []int{14} }

func (m *LabelSelector) Reset()				{ *m = LabelSelector{} }
func (*LabelSelector) ProtoMessage()			{}
func (*LabelSelector) Descriptor() ([]byte, []int)	{ return fileDescriptorGenerated, []int{15} }

func (m *LabelSelectorRequirement) Reset()	{ *m = LabelSelectorRequirement{} }
func (*LabelSelectorRequirement) ProtoMessage()	{}
func (*LabelSelectorRequirement) Descriptor() ([]byte, []int) {
	return fileDescriptorGenerated, []int{16}
}

func (m *ListMeta) Reset()			{ *m = ListMeta{} }
func (*ListMeta) ProtoMessage()			{}
func (*ListMeta) Descriptor() ([]byte, []int)	{ return fileDescriptorGenerated, []int{17} }

func (m *ListOptions) Reset()				{ *m = ListOptions{} }
func (*ListOptions) ProtoMessage()			{}
func (*ListOptions) Descriptor() ([]byte, []int)	{ return fileDescriptorGenerated, []int{18} }

func (m *ObjectMeta) Reset()			{ *m = ObjectMeta{} }
func (*ObjectMeta) ProtoMessage()		{}
func (*ObjectMeta) Descriptor() ([]byte, []int)	{ return fileDescriptorGenerated, []int{19} }

func (m *OwnerReference) Reset()			{ *m = OwnerReference{} }
func (*OwnerReference) ProtoMessage()			{}
func (*OwnerReference) Descriptor() ([]byte, []int)	{ return fileDescriptorGenerated, []int{20} }

func (m *Preconditions) Reset()				{ *m = Preconditions{} }
func (*Preconditions) ProtoMessage()			{}
func (*Preconditions) Descriptor() ([]byte, []int)	{ return fileDescriptorGenerated, []int{21} }

func (m *RootPaths) Reset()			{ *m = RootPaths{} }
func (*RootPaths) ProtoMessage()		{}
func (*RootPaths) Descriptor() ([]byte, []int)	{ return fileDescriptorGenerated, []int{22} }

func (m *ServerAddressByClientCIDR) Reset()		{ *m = ServerAddressByClientCIDR{} }
func (*ServerAddressByClientCIDR) ProtoMessage()	{}
func (*ServerAddressByClientCIDR) Descriptor() ([]byte, []int) {
	return fileDescriptorGenerated, []int{23}
}

func (m *Status) Reset()			{ *m = Status{} }
func (*Status) ProtoMessage()			{}
func (*Status) Descriptor() ([]byte, []int)	{ return fileDescriptorGenerated, []int{24} }

func (m *StatusCause) Reset()				{ *m = StatusCause{} }
func (*StatusCause) ProtoMessage()			{}
func (*StatusCause) Descriptor() ([]byte, []int)	{ return fileDescriptorGenerated, []int{25} }

func (m *StatusDetails) Reset()				{ *m = StatusDetails{} }
func (*StatusDetails) ProtoMessage()			{}
func (*StatusDetails) Descriptor() ([]byte, []int)	{ return fileDescriptorGenerated, []int{26} }

func (m *Time) Reset()				{ *m = Time{} }
func (*Time) ProtoMessage()			{}
func (*Time) Descriptor() ([]byte, []int)	{ return fileDescriptorGenerated, []int{27} }

func (m *Timestamp) Reset()			{ *m = Timestamp{} }
func (*Timestamp) ProtoMessage()		{}
func (*Timestamp) Descriptor() ([]byte, []int)	{ return fileDescriptorGenerated, []int{28} }

func (m *TypeMeta) Reset()			{ *m = TypeMeta{} }
func (*TypeMeta) ProtoMessage()			{}
func (*TypeMeta) Descriptor() ([]byte, []int)	{ return fileDescriptorGenerated, []int{29} }

func (m *Verbs) Reset()				{ *m = Verbs{} }
func (*Verbs) ProtoMessage()			{}
func (*Verbs) Descriptor() ([]byte, []int)	{ return fileDescriptorGenerated, []int{30} }

func (m *WatchEvent) Reset()			{ *m = WatchEvent{} }
func (*WatchEvent) ProtoMessage()		{}
func (*WatchEvent) Descriptor() ([]byte, []int)	{ return fileDescriptorGenerated, []int{31} }

func init() {
	proto.RegisterType((*APIGroup)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1.APIGroup")
	proto.RegisterType((*APIGroupList)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1.APIGroupList")
	proto.RegisterType((*APIResource)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1.APIResource")
	proto.RegisterType((*APIResourceList)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1.APIResourceList")
	proto.RegisterType((*APIVersions)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1.APIVersions")
	proto.RegisterType((*DeleteOptions)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1.DeleteOptions")
	proto.RegisterType((*Duration)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1.Duration")
	proto.RegisterType((*ExportOptions)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1.ExportOptions")
	proto.RegisterType((*GetOptions)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1.GetOptions")
	proto.RegisterType((*GroupKind)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1.GroupKind")
	proto.RegisterType((*GroupResource)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1.GroupResource")
	proto.RegisterType((*GroupVersion)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1.GroupVersion")
	proto.RegisterType((*GroupVersionForDiscovery)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1.GroupVersionForDiscovery")
	proto.RegisterType((*GroupVersionKind)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1.GroupVersionKind")
	proto.RegisterType((*GroupVersionResource)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1.GroupVersionResource")
	proto.RegisterType((*LabelSelector)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1.LabelSelector")
	proto.RegisterType((*LabelSelectorRequirement)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1.LabelSelectorRequirement")
	proto.RegisterType((*ListMeta)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1.ListMeta")
	proto.RegisterType((*ListOptions)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1.ListOptions")
	proto.RegisterType((*ObjectMeta)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1.ObjectMeta")
	proto.RegisterType((*OwnerReference)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1.OwnerReference")
	proto.RegisterType((*Preconditions)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1.Preconditions")
	proto.RegisterType((*RootPaths)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1.RootPaths")
	proto.RegisterType((*ServerAddressByClientCIDR)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1.ServerAddressByClientCIDR")
	proto.RegisterType((*Status)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1.Status")
	proto.RegisterType((*StatusCause)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1.StatusCause")
	proto.RegisterType((*StatusDetails)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1.StatusDetails")
	proto.RegisterType((*Time)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1.Time")
	proto.RegisterType((*Timestamp)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1.Timestamp")
	proto.RegisterType((*TypeMeta)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1.TypeMeta")
	proto.RegisterType((*Verbs)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1.Verbs")
	proto.RegisterType((*WatchEvent)(nil), "k8s.io.apimachinery.pkg.apis.meta.v1.WatchEvent")
}
func (m *APIGroup) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *APIGroup) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Name)))
	i += copy(data[i:], m.Name)
	if len(m.Versions) > 0 {
		for _, msg := range m.Versions {
			data[i] = 0x12
			i++
			i = encodeVarintGenerated(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	data[i] = 0x1a
	i++
	i = encodeVarintGenerated(data, i, uint64(m.PreferredVersion.Size()))
	n1, err := m.PreferredVersion.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if len(m.ServerAddressByClientCIDRs) > 0 {
		for _, msg := range m.ServerAddressByClientCIDRs {
			data[i] = 0x22
			i++
			i = encodeVarintGenerated(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *APIGroupList) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *APIGroupList) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Groups) > 0 {
		for _, msg := range m.Groups {
			data[i] = 0xa
			i++
			i = encodeVarintGenerated(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *APIResource) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *APIResource) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Name)))
	i += copy(data[i:], m.Name)
	data[i] = 0x10
	i++
	if m.Namespaced {
		data[i] = 1
	} else {
		data[i] = 0
	}
	i++
	data[i] = 0x1a
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Kind)))
	i += copy(data[i:], m.Kind)
	if m.Verbs != nil {
		data[i] = 0x22
		i++
		i = encodeVarintGenerated(data, i, uint64(m.Verbs.Size()))
		n2, err := m.Verbs.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.ShortNames) > 0 {
		for _, s := range m.ShortNames {
			data[i] = 0x2a
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	return i, nil
}

func (m *APIResourceList) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *APIResourceList) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.GroupVersion)))
	i += copy(data[i:], m.GroupVersion)
	if len(m.APIResources) > 0 {
		for _, msg := range m.APIResources {
			data[i] = 0x12
			i++
			i = encodeVarintGenerated(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *APIVersions) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *APIVersions) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Versions) > 0 {
		for _, s := range m.Versions {
			data[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	if len(m.ServerAddressByClientCIDRs) > 0 {
		for _, msg := range m.ServerAddressByClientCIDRs {
			data[i] = 0x12
			i++
			i = encodeVarintGenerated(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *DeleteOptions) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *DeleteOptions) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.GracePeriodSeconds != nil {
		data[i] = 0x8
		i++
		i = encodeVarintGenerated(data, i, uint64(*m.GracePeriodSeconds))
	}
	if m.Preconditions != nil {
		data[i] = 0x12
		i++
		i = encodeVarintGenerated(data, i, uint64(m.Preconditions.Size()))
		n3, err := m.Preconditions.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.OrphanDependents != nil {
		data[i] = 0x18
		i++
		if *m.OrphanDependents {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.PropagationPolicy != nil {
		data[i] = 0x22
		i++
		i = encodeVarintGenerated(data, i, uint64(len(*m.PropagationPolicy)))
		i += copy(data[i:], *m.PropagationPolicy)
	}
	return i, nil
}

func (m *Duration) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Duration) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintGenerated(data, i, uint64(m.Duration))
	return i, nil
}

func (m *ExportOptions) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ExportOptions) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	if m.Export {
		data[i] = 1
	} else {
		data[i] = 0
	}
	i++
	data[i] = 0x10
	i++
	if m.Exact {
		data[i] = 1
	} else {
		data[i] = 0
	}
	i++
	return i, nil
}

func (m *GetOptions) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *GetOptions) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.ResourceVersion)))
	i += copy(data[i:], m.ResourceVersion)
	return i, nil
}

func (m *GroupKind) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *GroupKind) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Group)))
	i += copy(data[i:], m.Group)
	data[i] = 0x12
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Kind)))
	i += copy(data[i:], m.Kind)
	return i, nil
}

func (m *GroupResource) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *GroupResource) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Group)))
	i += copy(data[i:], m.Group)
	data[i] = 0x12
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Resource)))
	i += copy(data[i:], m.Resource)
	return i, nil
}

func (m *GroupVersion) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *GroupVersion) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Group)))
	i += copy(data[i:], m.Group)
	data[i] = 0x12
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Version)))
	i += copy(data[i:], m.Version)
	return i, nil
}

func (m *GroupVersionForDiscovery) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *GroupVersionForDiscovery) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.GroupVersion)))
	i += copy(data[i:], m.GroupVersion)
	data[i] = 0x12
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Version)))
	i += copy(data[i:], m.Version)
	return i, nil
}

func (m *GroupVersionKind) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *GroupVersionKind) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Group)))
	i += copy(data[i:], m.Group)
	data[i] = 0x12
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Version)))
	i += copy(data[i:], m.Version)
	data[i] = 0x1a
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Kind)))
	i += copy(data[i:], m.Kind)
	return i, nil
}

func (m *GroupVersionResource) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *GroupVersionResource) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Group)))
	i += copy(data[i:], m.Group)
	data[i] = 0x12
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Version)))
	i += copy(data[i:], m.Version)
	data[i] = 0x1a
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Resource)))
	i += copy(data[i:], m.Resource)
	return i, nil
}

func (m *LabelSelector) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *LabelSelector) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.MatchLabels) > 0 {
		for k := range m.MatchLabels {
			data[i] = 0xa
			i++
			v := m.MatchLabels[k]
			mapSize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + len(v) + sovGenerated(uint64(len(v)))
			i = encodeVarintGenerated(data, i, uint64(mapSize))
			data[i] = 0xa
			i++
			i = encodeVarintGenerated(data, i, uint64(len(k)))
			i += copy(data[i:], k)
			data[i] = 0x12
			i++
			i = encodeVarintGenerated(data, i, uint64(len(v)))
			i += copy(data[i:], v)
		}
	}
	if len(m.MatchExpressions) > 0 {
		for _, msg := range m.MatchExpressions {
			data[i] = 0x12
			i++
			i = encodeVarintGenerated(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *LabelSelectorRequirement) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *LabelSelectorRequirement) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Key)))
	i += copy(data[i:], m.Key)
	data[i] = 0x12
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Operator)))
	i += copy(data[i:], m.Operator)
	if len(m.Values) > 0 {
		for _, s := range m.Values {
			data[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	return i, nil
}

func (m *ListMeta) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ListMeta) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.SelfLink)))
	i += copy(data[i:], m.SelfLink)
	data[i] = 0x12
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.ResourceVersion)))
	i += copy(data[i:], m.ResourceVersion)
	return i, nil
}

func (m *ListOptions) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ListOptions) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.LabelSelector)))
	i += copy(data[i:], m.LabelSelector)
	data[i] = 0x12
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.FieldSelector)))
	i += copy(data[i:], m.FieldSelector)
	data[i] = 0x18
	i++
	if m.Watch {
		data[i] = 1
	} else {
		data[i] = 0
	}
	i++
	data[i] = 0x22
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.ResourceVersion)))
	i += copy(data[i:], m.ResourceVersion)
	if m.TimeoutSeconds != nil {
		data[i] = 0x28
		i++
		i = encodeVarintGenerated(data, i, uint64(*m.TimeoutSeconds))
	}
	return i, nil
}

func (m *ObjectMeta) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ObjectMeta) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Name)))
	i += copy(data[i:], m.Name)
	data[i] = 0x12
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.GenerateName)))
	i += copy(data[i:], m.GenerateName)
	data[i] = 0x1a
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Namespace)))
	i += copy(data[i:], m.Namespace)
	data[i] = 0x22
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.SelfLink)))
	i += copy(data[i:], m.SelfLink)
	data[i] = 0x2a
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.UID)))
	i += copy(data[i:], m.UID)
	data[i] = 0x32
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.ResourceVersion)))
	i += copy(data[i:], m.ResourceVersion)
	data[i] = 0x38
	i++
	i = encodeVarintGenerated(data, i, uint64(m.Generation))
	data[i] = 0x42
	i++
	i = encodeVarintGenerated(data, i, uint64(m.CreationTimestamp.Size()))
	n4, err := m.CreationTimestamp.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	if m.DeletionTimestamp != nil {
		data[i] = 0x4a
		i++
		i = encodeVarintGenerated(data, i, uint64(m.DeletionTimestamp.Size()))
		n5, err := m.DeletionTimestamp.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.DeletionGracePeriodSeconds != nil {
		data[i] = 0x50
		i++
		i = encodeVarintGenerated(data, i, uint64(*m.DeletionGracePeriodSeconds))
	}
	if len(m.Labels) > 0 {
		for k := range m.Labels {
			data[i] = 0x5a
			i++
			v := m.Labels[k]
			mapSize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + len(v) + sovGenerated(uint64(len(v)))
			i = encodeVarintGenerated(data, i, uint64(mapSize))
			data[i] = 0xa
			i++
			i = encodeVarintGenerated(data, i, uint64(len(k)))
			i += copy(data[i:], k)
			data[i] = 0x12
			i++
			i = encodeVarintGenerated(data, i, uint64(len(v)))
			i += copy(data[i:], v)
		}
	}
	if len(m.Annotations) > 0 {
		for k := range m.Annotations {
			data[i] = 0x62
			i++
			v := m.Annotations[k]
			mapSize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + len(v) + sovGenerated(uint64(len(v)))
			i = encodeVarintGenerated(data, i, uint64(mapSize))
			data[i] = 0xa
			i++
			i = encodeVarintGenerated(data, i, uint64(len(k)))
			i += copy(data[i:], k)
			data[i] = 0x12
			i++
			i = encodeVarintGenerated(data, i, uint64(len(v)))
			i += copy(data[i:], v)
		}
	}
	if len(m.OwnerReferences) > 0 {
		for _, msg := range m.OwnerReferences {
			data[i] = 0x6a
			i++
			i = encodeVarintGenerated(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Finalizers) > 0 {
		for _, s := range m.Finalizers {
			data[i] = 0x72
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	data[i] = 0x7a
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.ClusterName)))
	i += copy(data[i:], m.ClusterName)
	return i, nil
}

func (m *OwnerReference) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *OwnerReference) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Kind)))
	i += copy(data[i:], m.Kind)
	data[i] = 0x1a
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Name)))
	i += copy(data[i:], m.Name)
	data[i] = 0x22
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.UID)))
	i += copy(data[i:], m.UID)
	data[i] = 0x2a
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.APIVersion)))
	i += copy(data[i:], m.APIVersion)
	if m.Controller != nil {
		data[i] = 0x30
		i++
		if *m.Controller {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.BlockOwnerDeletion != nil {
		data[i] = 0x38
		i++
		if *m.BlockOwnerDeletion {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *Preconditions) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Preconditions) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.UID != nil {
		data[i] = 0xa
		i++
		i = encodeVarintGenerated(data, i, uint64(len(*m.UID)))
		i += copy(data[i:], *m.UID)
	}
	return i, nil
}

func (m *RootPaths) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RootPaths) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Paths) > 0 {
		for _, s := range m.Paths {
			data[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	return i, nil
}

func (m *ServerAddressByClientCIDR) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ServerAddressByClientCIDR) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.ClientCIDR)))
	i += copy(data[i:], m.ClientCIDR)
	data[i] = 0x12
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.ServerAddress)))
	i += copy(data[i:], m.ServerAddress)
	return i, nil
}

func (m *Status) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Status) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(m.ListMeta.Size()))
	n6, err := m.ListMeta.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	data[i] = 0x12
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Status)))
	i += copy(data[i:], m.Status)
	data[i] = 0x1a
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Message)))
	i += copy(data[i:], m.Message)
	data[i] = 0x22
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Reason)))
	i += copy(data[i:], m.Reason)
	if m.Details != nil {
		data[i] = 0x2a
		i++
		i = encodeVarintGenerated(data, i, uint64(m.Details.Size()))
		n7, err := m.Details.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	data[i] = 0x30
	i++
	i = encodeVarintGenerated(data, i, uint64(m.Code))
	return i, nil
}

func (m *StatusCause) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *StatusCause) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Type)))
	i += copy(data[i:], m.Type)
	data[i] = 0x12
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Message)))
	i += copy(data[i:], m.Message)
	data[i] = 0x1a
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Field)))
	i += copy(data[i:], m.Field)
	return i, nil
}

func (m *StatusDetails) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *StatusDetails) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Name)))
	i += copy(data[i:], m.Name)
	data[i] = 0x12
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Group)))
	i += copy(data[i:], m.Group)
	data[i] = 0x1a
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Kind)))
	i += copy(data[i:], m.Kind)
	if len(m.Causes) > 0 {
		for _, msg := range m.Causes {
			data[i] = 0x22
			i++
			i = encodeVarintGenerated(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	data[i] = 0x28
	i++
	i = encodeVarintGenerated(data, i, uint64(m.RetryAfterSeconds))
	return i, nil
}

func (m *Timestamp) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Timestamp) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintGenerated(data, i, uint64(m.Seconds))
	data[i] = 0x10
	i++
	i = encodeVarintGenerated(data, i, uint64(m.Nanos))
	return i, nil
}

func (m *TypeMeta) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *TypeMeta) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Kind)))
	i += copy(data[i:], m.Kind)
	data[i] = 0x12
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.APIVersion)))
	i += copy(data[i:], m.APIVersion)
	return i, nil
}

func (m Verbs) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m Verbs) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m) > 0 {
		for _, s := range m {
			data[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	return i, nil
}

func (m *WatchEvent) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *WatchEvent) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.Type)))
	i += copy(data[i:], m.Type)
	data[i] = 0x12
	i++
	i = encodeVarintGenerated(data, i, uint64(m.Object.Size()))
	n8, err := m.Object.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n8
	return i, nil
}

func encodeFixed64Generated(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Generated(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGenerated(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *APIGroup) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Versions) > 0 {
		for _, e := range m.Versions {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	l = m.PreferredVersion.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.ServerAddressByClientCIDRs) > 0 {
		for _, e := range m.ServerAddressByClientCIDRs {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *APIGroupList) Size() (n int) {
	var l int
	_ = l
	if len(m.Groups) > 0 {
		for _, e := range m.Groups {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *APIResource) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	n += 1 + l + sovGenerated(uint64(l))
	n += 2
	l = len(m.Kind)
	n += 1 + l + sovGenerated(uint64(l))
	if m.Verbs != nil {
		l = m.Verbs.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if len(m.ShortNames) > 0 {
		for _, s := range m.ShortNames {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *APIResourceList) Size() (n int) {
	var l int
	_ = l
	l = len(m.GroupVersion)
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.APIResources) > 0 {
		for _, e := range m.APIResources {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *APIVersions) Size() (n int) {
	var l int
	_ = l
	if len(m.Versions) > 0 {
		for _, s := range m.Versions {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if len(m.ServerAddressByClientCIDRs) > 0 {
		for _, e := range m.ServerAddressByClientCIDRs {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *DeleteOptions) Size() (n int) {
	var l int
	_ = l
	if m.GracePeriodSeconds != nil {
		n += 1 + sovGenerated(uint64(*m.GracePeriodSeconds))
	}
	if m.Preconditions != nil {
		l = m.Preconditions.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.OrphanDependents != nil {
		n += 2
	}
	if m.PropagationPolicy != nil {
		l = len(*m.PropagationPolicy)
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *Duration) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovGenerated(uint64(m.Duration))
	return n
}

func (m *ExportOptions) Size() (n int) {
	var l int
	_ = l
	n += 2
	n += 2
	return n
}

func (m *GetOptions) Size() (n int) {
	var l int
	_ = l
	l = len(m.ResourceVersion)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *GroupKind) Size() (n int) {
	var l int
	_ = l
	l = len(m.Group)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Kind)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *GroupResource) Size() (n int) {
	var l int
	_ = l
	l = len(m.Group)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Resource)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *GroupVersion) Size() (n int) {
	var l int
	_ = l
	l = len(m.Group)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Version)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *GroupVersionForDiscovery) Size() (n int) {
	var l int
	_ = l
	l = len(m.GroupVersion)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Version)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *GroupVersionKind) Size() (n int) {
	var l int
	_ = l
	l = len(m.Group)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Version)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Kind)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *GroupVersionResource) Size() (n int) {
	var l int
	_ = l
	l = len(m.Group)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Version)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Resource)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *LabelSelector) Size() (n int) {
	var l int
	_ = l
	if len(m.MatchLabels) > 0 {
		for k, v := range m.MatchLabels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + len(v) + sovGenerated(uint64(len(v)))
			n += mapEntrySize + 1 + sovGenerated(uint64(mapEntrySize))
		}
	}
	if len(m.MatchExpressions) > 0 {
		for _, e := range m.MatchExpressions {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *LabelSelectorRequirement) Size() (n int) {
	var l int
	_ = l
	l = len(m.Key)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Operator)
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Values) > 0 {
		for _, s := range m.Values {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *ListMeta) Size() (n int) {
	var l int
	_ = l
	l = len(m.SelfLink)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.ResourceVersion)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *ListOptions) Size() (n int) {
	var l int
	_ = l
	l = len(m.LabelSelector)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.FieldSelector)
	n += 1 + l + sovGenerated(uint64(l))
	n += 2
	l = len(m.ResourceVersion)
	n += 1 + l + sovGenerated(uint64(l))
	if m.TimeoutSeconds != nil {
		n += 1 + sovGenerated(uint64(*m.TimeoutSeconds))
	}
	return n
}

func (m *ObjectMeta) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.GenerateName)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Namespace)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.SelfLink)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.UID)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.ResourceVersion)
	n += 1 + l + sovGenerated(uint64(l))
	n += 1 + sovGenerated(uint64(m.Generation))
	l = m.CreationTimestamp.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if m.DeletionTimestamp != nil {
		l = m.DeletionTimestamp.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.DeletionGracePeriodSeconds != nil {
		n += 1 + sovGenerated(uint64(*m.DeletionGracePeriodSeconds))
	}
	if len(m.Labels) > 0 {
		for k, v := range m.Labels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + len(v) + sovGenerated(uint64(len(v)))
			n += mapEntrySize + 1 + sovGenerated(uint64(mapEntrySize))
		}
	}
	if len(m.Annotations) > 0 {
		for k, v := range m.Annotations {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovGenerated(uint64(len(k))) + 1 + len(v) + sovGenerated(uint64(len(v)))
			n += mapEntrySize + 1 + sovGenerated(uint64(mapEntrySize))
		}
	}
	if len(m.OwnerReferences) > 0 {
		for _, e := range m.OwnerReferences {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	if len(m.Finalizers) > 0 {
		for _, s := range m.Finalizers {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	l = len(m.ClusterName)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *OwnerReference) Size() (n int) {
	var l int
	_ = l
	l = len(m.Kind)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Name)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.UID)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.APIVersion)
	n += 1 + l + sovGenerated(uint64(l))
	if m.Controller != nil {
		n += 2
	}
	if m.BlockOwnerDeletion != nil {
		n += 2
	}
	return n
}

func (m *Preconditions) Size() (n int) {
	var l int
	_ = l
	if m.UID != nil {
		l = len(*m.UID)
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *RootPaths) Size() (n int) {
	var l int
	_ = l
	if len(m.Paths) > 0 {
		for _, s := range m.Paths {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *ServerAddressByClientCIDR) Size() (n int) {
	var l int
	_ = l
	l = len(m.ClientCIDR)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.ServerAddress)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *Status) Size() (n int) {
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Status)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Message)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Reason)
	n += 1 + l + sovGenerated(uint64(l))
	if m.Details != nil {
		l = m.Details.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	n += 1 + sovGenerated(uint64(m.Code))
	return n
}

func (m *StatusCause) Size() (n int) {
	var l int
	_ = l
	l = len(m.Type)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Message)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Field)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *StatusDetails) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Group)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.Kind)
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Causes) > 0 {
		for _, e := range m.Causes {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	n += 1 + sovGenerated(uint64(m.RetryAfterSeconds))
	return n
}

func (m *Timestamp) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovGenerated(uint64(m.Seconds))
	n += 1 + sovGenerated(uint64(m.Nanos))
	return n
}

func (m *TypeMeta) Size() (n int) {
	var l int
	_ = l
	l = len(m.Kind)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.APIVersion)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m Verbs) Size() (n int) {
	var l int
	_ = l
	if len(m) > 0 {
		for _, s := range m {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *WatchEvent) Size() (n int) {
	var l int
	_ = l
	l = len(m.Type)
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Object.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *APIGroup) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&APIGroup{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Versions:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Versions), "GroupVersionForDiscovery", "GroupVersionForDiscovery", 1), `&`, ``, 1) + `,`,
		`PreferredVersion:` + strings.Replace(strings.Replace(this.PreferredVersion.String(), "GroupVersionForDiscovery", "GroupVersionForDiscovery", 1), `&`, ``, 1) + `,`,
		`ServerAddressByClientCIDRs:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ServerAddressByClientCIDRs), "ServerAddressByClientCIDR", "ServerAddressByClientCIDR", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *APIGroupList) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&APIGroupList{`,
		`Groups:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Groups), "APIGroup", "APIGroup", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *APIResource) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&APIResource{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Namespaced:` + fmt.Sprintf("%v", this.Namespaced) + `,`,
		`Kind:` + fmt.Sprintf("%v", this.Kind) + `,`,
		`Verbs:` + strings.Replace(fmt.Sprintf("%v", this.Verbs), "Verbs", "Verbs", 1) + `,`,
		`ShortNames:` + fmt.Sprintf("%v", this.ShortNames) + `,`,
		`}`,
	}, "")
	return s
}
func (this *APIResourceList) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&APIResourceList{`,
		`GroupVersion:` + fmt.Sprintf("%v", this.GroupVersion) + `,`,
		`APIResources:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.APIResources), "APIResource", "APIResource", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *DeleteOptions) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DeleteOptions{`,
		`GracePeriodSeconds:` + valueToStringGenerated(this.GracePeriodSeconds) + `,`,
		`Preconditions:` + strings.Replace(fmt.Sprintf("%v", this.Preconditions), "Preconditions", "Preconditions", 1) + `,`,
		`OrphanDependents:` + valueToStringGenerated(this.OrphanDependents) + `,`,
		`PropagationPolicy:` + valueToStringGenerated(this.PropagationPolicy) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Duration) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Duration{`,
		`Duration:` + fmt.Sprintf("%v", this.Duration) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ExportOptions) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ExportOptions{`,
		`Export:` + fmt.Sprintf("%v", this.Export) + `,`,
		`Exact:` + fmt.Sprintf("%v", this.Exact) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetOptions) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetOptions{`,
		`ResourceVersion:` + fmt.Sprintf("%v", this.ResourceVersion) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GroupVersionForDiscovery) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GroupVersionForDiscovery{`,
		`GroupVersion:` + fmt.Sprintf("%v", this.GroupVersion) + `,`,
		`Version:` + fmt.Sprintf("%v", this.Version) + `,`,
		`}`,
	}, "")
	return s
}
func (this *LabelSelector) String() string {
	if this == nil {
		return "nil"
	}
	keysForMatchLabels := make([]string, 0, len(this.MatchLabels))
	for k := range this.MatchLabels {
		keysForMatchLabels = append(keysForMatchLabels, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForMatchLabels)
	mapStringForMatchLabels := "map[string]string{"
	for _, k := range keysForMatchLabels {
		mapStringForMatchLabels += fmt.Sprintf("%v: %v,", k, this.MatchLabels[k])
	}
	mapStringForMatchLabels += "}"
	s := strings.Join([]string{`&LabelSelector{`,
		`MatchLabels:` + mapStringForMatchLabels + `,`,
		`MatchExpressions:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.MatchExpressions), "LabelSelectorRequirement", "LabelSelectorRequirement", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *LabelSelectorRequirement) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&LabelSelectorRequirement{`,
		`Key:` + fmt.Sprintf("%v", this.Key) + `,`,
		`Operator:` + fmt.Sprintf("%v", this.Operator) + `,`,
		`Values:` + fmt.Sprintf("%v", this.Values) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ListMeta) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ListMeta{`,
		`SelfLink:` + fmt.Sprintf("%v", this.SelfLink) + `,`,
		`ResourceVersion:` + fmt.Sprintf("%v", this.ResourceVersion) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ListOptions) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ListOptions{`,
		`LabelSelector:` + fmt.Sprintf("%v", this.LabelSelector) + `,`,
		`FieldSelector:` + fmt.Sprintf("%v", this.FieldSelector) + `,`,
		`Watch:` + fmt.Sprintf("%v", this.Watch) + `,`,
		`ResourceVersion:` + fmt.Sprintf("%v", this.ResourceVersion) + `,`,
		`TimeoutSeconds:` + valueToStringGenerated(this.TimeoutSeconds) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ObjectMeta) String() string {
	if this == nil {
		return "nil"
	}
	keysForLabels := make([]string, 0, len(this.Labels))
	for k := range this.Labels {
		keysForLabels = append(keysForLabels, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForLabels)
	mapStringForLabels := "map[string]string{"
	for _, k := range keysForLabels {
		mapStringForLabels += fmt.Sprintf("%v: %v,", k, this.Labels[k])
	}
	mapStringForLabels += "}"
	keysForAnnotations := make([]string, 0, len(this.Annotations))
	for k := range this.Annotations {
		keysForAnnotations = append(keysForAnnotations, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForAnnotations)
	mapStringForAnnotations := "map[string]string{"
	for _, k := range keysForAnnotations {
		mapStringForAnnotations += fmt.Sprintf("%v: %v,", k, this.Annotations[k])
	}
	mapStringForAnnotations += "}"
	s := strings.Join([]string{`&ObjectMeta{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`GenerateName:` + fmt.Sprintf("%v", this.GenerateName) + `,`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`SelfLink:` + fmt.Sprintf("%v", this.SelfLink) + `,`,
		`UID:` + fmt.Sprintf("%v", this.UID) + `,`,
		`ResourceVersion:` + fmt.Sprintf("%v", this.ResourceVersion) + `,`,
		`Generation:` + fmt.Sprintf("%v", this.Generation) + `,`,
		`CreationTimestamp:` + strings.Replace(strings.Replace(this.CreationTimestamp.String(), "Time", "Time", 1), `&`, ``, 1) + `,`,
		`DeletionTimestamp:` + strings.Replace(fmt.Sprintf("%v", this.DeletionTimestamp), "Time", "Time", 1) + `,`,
		`DeletionGracePeriodSeconds:` + valueToStringGenerated(this.DeletionGracePeriodSeconds) + `,`,
		`Labels:` + mapStringForLabels + `,`,
		`Annotations:` + mapStringForAnnotations + `,`,
		`OwnerReferences:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.OwnerReferences), "OwnerReference", "OwnerReference", 1), `&`, ``, 1) + `,`,
		`Finalizers:` + fmt.Sprintf("%v", this.Finalizers) + `,`,
		`ClusterName:` + fmt.Sprintf("%v", this.ClusterName) + `,`,
		`}`,
	}, "")
	return s
}
func (this *OwnerReference) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&OwnerReference{`,
		`Kind:` + fmt.Sprintf("%v", this.Kind) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`UID:` + fmt.Sprintf("%v", this.UID) + `,`,
		`APIVersion:` + fmt.Sprintf("%v", this.APIVersion) + `,`,
		`Controller:` + valueToStringGenerated(this.Controller) + `,`,
		`BlockOwnerDeletion:` + valueToStringGenerated(this.BlockOwnerDeletion) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Preconditions) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Preconditions{`,
		`UID:` + valueToStringGenerated(this.UID) + `,`,
		`}`,
	}, "")
	return s
}
func (this *RootPaths) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RootPaths{`,
		`Paths:` + fmt.Sprintf("%v", this.Paths) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ServerAddressByClientCIDR) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ServerAddressByClientCIDR{`,
		`ClientCIDR:` + fmt.Sprintf("%v", this.ClientCIDR) + `,`,
		`ServerAddress:` + fmt.Sprintf("%v", this.ServerAddress) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Status) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Status{`,
		`ListMeta:` + strings.Replace(strings.Replace(this.ListMeta.String(), "ListMeta", "ListMeta", 1), `&`, ``, 1) + `,`,
		`Status:` + fmt.Sprintf("%v", this.Status) + `,`,
		`Message:` + fmt.Sprintf("%v", this.Message) + `,`,
		`Reason:` + fmt.Sprintf("%v", this.Reason) + `,`,
		`Details:` + strings.Replace(fmt.Sprintf("%v", this.Details), "StatusDetails", "StatusDetails", 1) + `,`,
		`Code:` + fmt.Sprintf("%v", this.Code) + `,`,
		`}`,
	}, "")
	return s
}
func (this *StatusCause) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StatusCause{`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Message:` + fmt.Sprintf("%v", this.Message) + `,`,
		`Field:` + fmt.Sprintf("%v", this.Field) + `,`,
		`}`,
	}, "")
	return s
}
func (this *StatusDetails) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StatusDetails{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Group:` + fmt.Sprintf("%v", this.Group) + `,`,
		`Kind:` + fmt.Sprintf("%v", this.Kind) + `,`,
		`Causes:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Causes), "StatusCause", "StatusCause", 1), `&`, ``, 1) + `,`,
		`RetryAfterSeconds:` + fmt.Sprintf("%v", this.RetryAfterSeconds) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Timestamp) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Timestamp{`,
		`Seconds:` + fmt.Sprintf("%v", this.Seconds) + `,`,
		`Nanos:` + fmt.Sprintf("%v", this.Nanos) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TypeMeta) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TypeMeta{`,
		`Kind:` + fmt.Sprintf("%v", this.Kind) + `,`,
		`APIVersion:` + fmt.Sprintf("%v", this.APIVersion) + `,`,
		`}`,
	}, "")
	return s
}
func (this *WatchEvent) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&WatchEvent{`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Object:` + strings.Replace(strings.Replace(this.Object.String(), "RawExtension", "k8s_io_apimachinery_pkg_runtime.RawExtension", 1), `&`, ``, 1) + `,`,
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
func (m *APIGroup) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: APIGroup: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: APIGroup: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Versions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Versions = append(m.Versions, GroupVersionForDiscovery{})
			if err := m.Versions[len(m.Versions)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreferredVersion", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PreferredVersion.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerAddressByClientCIDRs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServerAddressByClientCIDRs = append(m.ServerAddressByClientCIDRs, ServerAddressByClientCIDR{})
			if err := m.ServerAddressByClientCIDRs[len(m.ServerAddressByClientCIDRs)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *APIGroupList) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: APIGroupList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: APIGroupList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Groups", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Groups = append(m.Groups, APIGroup{})
			if err := m.Groups[len(m.Groups)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *APIResource) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: APIResource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: APIResource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespaced", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Namespaced = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Kind = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Verbs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Verbs == nil {
				m.Verbs = Verbs{}
			}
			if err := m.Verbs.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShortNames", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ShortNames = append(m.ShortNames, string(data[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *APIResourceList) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: APIResourceList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: APIResourceList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GroupVersion = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field APIResources", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.APIResources = append(m.APIResources, APIResource{})
			if err := m.APIResources[len(m.APIResources)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *APIVersions) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: APIVersions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: APIVersions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Versions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Versions = append(m.Versions, string(data[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerAddressByClientCIDRs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServerAddressByClientCIDRs = append(m.ServerAddressByClientCIDRs, ServerAddressByClientCIDR{})
			if err := m.ServerAddressByClientCIDRs[len(m.ServerAddressByClientCIDRs)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *DeleteOptions) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DeleteOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeleteOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GracePeriodSeconds", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.GracePeriodSeconds = &v
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Preconditions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Preconditions == nil {
				m.Preconditions = &Preconditions{}
			}
			if err := m.Preconditions.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrphanDependents", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.OrphanDependents = &b
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PropagationPolicy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := DeletionPropagation(data[iNdEx:postIndex])
			m.PropagationPolicy = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *Duration) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Duration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Duration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			m.Duration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Duration |= (time.Duration(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *ExportOptions) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExportOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExportOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Export", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Export = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exact", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Exact = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *GetOptions) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResourceVersion = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *GroupKind) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GroupKind: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GroupKind: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Group", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Group = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Kind = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *GroupResource) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GroupResource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GroupResource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Group", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Group = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resource", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Resource = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *GroupVersion) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GroupVersion: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GroupVersion: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Group", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Group = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
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
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *GroupVersionForDiscovery) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GroupVersionForDiscovery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GroupVersionForDiscovery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GroupVersion = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
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
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *GroupVersionKind) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GroupVersionKind: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GroupVersionKind: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Group", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Group = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
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
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Kind = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *GroupVersionResource) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GroupVersionResource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GroupVersionResource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Group", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Group = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
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
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resource", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Resource = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *LabelSelector) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LabelSelector: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LabelSelector: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MatchLabels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthGenerated
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(data[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			var valuekey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				valuekey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapvalue uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLenmapvalue |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapvalue := int(stringLenmapvalue)
			if intStringLenmapvalue < 0 {
				return ErrInvalidLengthGenerated
			}
			postStringIndexmapvalue := iNdEx + intStringLenmapvalue
			if postStringIndexmapvalue > l {
				return io.ErrUnexpectedEOF
			}
			mapvalue := string(data[iNdEx:postStringIndexmapvalue])
			iNdEx = postStringIndexmapvalue
			if m.MatchLabels == nil {
				m.MatchLabels = make(map[string]string)
			}
			m.MatchLabels[mapkey] = mapvalue
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MatchExpressions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MatchExpressions = append(m.MatchExpressions, LabelSelectorRequirement{})
			if err := m.MatchExpressions[len(m.MatchExpressions)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *LabelSelectorRequirement) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LabelSelectorRequirement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LabelSelectorRequirement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Operator = LabelSelectorOperator(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Values", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Values = append(m.Values, string(data[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *ListMeta) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SelfLink", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SelfLink = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResourceVersion = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *ListOptions) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LabelSelector", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LabelSelector = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FieldSelector", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FieldSelector = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Watch", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Watch = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResourceVersion = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutSeconds", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.TimeoutSeconds = &v
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *ObjectMeta) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ObjectMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ObjectMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenerateName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GenerateName = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SelfLink", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SelfLink = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = k8s_io_apimachinery_pkg_types.UID(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResourceVersion = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Generation", wireType)
			}
			m.Generation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Generation |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreationTimestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CreationTimestamp.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeletionTimestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DeletionTimestamp == nil {
				m.DeletionTimestamp = &Time{}
			}
			if err := m.DeletionTimestamp.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeletionGracePeriodSeconds", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.DeletionGracePeriodSeconds = &v
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthGenerated
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(data[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			var valuekey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				valuekey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapvalue uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLenmapvalue |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapvalue := int(stringLenmapvalue)
			if intStringLenmapvalue < 0 {
				return ErrInvalidLengthGenerated
			}
			postStringIndexmapvalue := iNdEx + intStringLenmapvalue
			if postStringIndexmapvalue > l {
				return io.ErrUnexpectedEOF
			}
			mapvalue := string(data[iNdEx:postStringIndexmapvalue])
			iNdEx = postStringIndexmapvalue
			if m.Labels == nil {
				m.Labels = make(map[string]string)
			}
			m.Labels[mapkey] = mapvalue
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Annotations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthGenerated
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(data[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			var valuekey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				valuekey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapvalue uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLenmapvalue |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapvalue := int(stringLenmapvalue)
			if intStringLenmapvalue < 0 {
				return ErrInvalidLengthGenerated
			}
			postStringIndexmapvalue := iNdEx + intStringLenmapvalue
			if postStringIndexmapvalue > l {
				return io.ErrUnexpectedEOF
			}
			mapvalue := string(data[iNdEx:postStringIndexmapvalue])
			iNdEx = postStringIndexmapvalue
			if m.Annotations == nil {
				m.Annotations = make(map[string]string)
			}
			m.Annotations[mapkey] = mapvalue
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerReferences", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnerReferences = append(m.OwnerReferences, OwnerReference{})
			if err := m.OwnerReferences[len(m.OwnerReferences)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Finalizers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Finalizers = append(m.Finalizers, string(data[iNdEx:postIndex]))
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClusterName = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *OwnerReference) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OwnerReference: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OwnerReference: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Kind = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = k8s_io_apimachinery_pkg_types.UID(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field APIVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.APIVersion = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Controller", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Controller = &b
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockOwnerDeletion", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.BlockOwnerDeletion = &b
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *Preconditions) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Preconditions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Preconditions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := k8s_io_apimachinery_pkg_types.UID(data[iNdEx:postIndex])
			m.UID = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *RootPaths) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RootPaths: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RootPaths: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Paths", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Paths = append(m.Paths, string(data[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *ServerAddressByClientCIDR) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ServerAddressByClientCIDR: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServerAddressByClientCIDR: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientCIDR", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientCIDR = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServerAddress = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *Status) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Status: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Status: illegal tag %d (wire type %d)", fieldNum, wire)
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
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reason = StatusReason(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Details", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Details == nil {
				m.Details = &StatusDetails{}
			}
			if err := m.Details.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Code |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *StatusCause) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StatusCause: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatusCause: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = CauseType(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Field = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *StatusDetails) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StatusDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatusDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Group", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Group = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Kind = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Causes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Causes = append(m.Causes, StatusCause{})
			if err := m.Causes[len(m.Causes)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RetryAfterSeconds", wireType)
			}
			m.RetryAfterSeconds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.RetryAfterSeconds |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *Timestamp) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Timestamp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Timestamp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seconds", wireType)
			}
			m.Seconds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Seconds |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nanos", wireType)
			}
			m.Nanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Nanos |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *TypeMeta) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TypeMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TypeMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Kind = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field APIVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.APIVersion = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *Verbs) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Verbs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Verbs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			*m = append(*m, string(data[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func (m *WatchEvent) Unmarshal(data []byte) error {
	l := len(data)
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
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WatchEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WatchEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Object", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Object.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
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
func skipGenerated(data []byte) (n int, err error) {
	l := len(data)
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
			b := data[iNdEx]
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
				if data[iNdEx-1] < 0x80 {
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
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
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
					b := data[iNdEx]
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
				next, err := skipGenerated(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthGenerated	= fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated		= fmt.Errorf("proto: integer overflow")
)

var fileDescriptorGenerated = []byte{

	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x19, 0xcd, 0x6f, 0x23, 0x57,
	0x3d, 0x63, 0xc7, 0x5e, 0xfb, 0xe7, 0x38, 0x1f, 0xaf, 0x59, 0x70, 0x23, 0x61, 0xa7, 0xd3, 0x0a,
	0xa5, 0xb0, 0xb5, 0x49, 0x0a, 0xd5, 0xb2, 0xc0, 0x42, 0x26, 0xce, 0x46, 0x51, 0x37, 0x9b, 0xe8,
	0xa5, 0xbb, 0x88, 0x65, 0x85, 0x98, 0xcc, 0xbc, 0x38, 0x43, 0xc6, 0x33, 0xc3, 0x7b, 0x63, 0x6f,
	0x4c, 0x0f, 0x54, 0x2a, 0x48, 0x1c, 0x10, 0xda, 0x23, 0x07, 0x84, 0xba, 0x82, 0x1b, 0x37, 0xfe,
	0x06, 0x24, 0xf6, 0x58, 0x89, 0x0b, 0x07, 0x64, 0xb1, 0xee, 0x81, 0x23, 0xf7, 0x9c, 0xd0, 0x7b,
	0xf3, 0xe6, 0xcb, 0x8e, 0x9b, 0x31, 0xed, 0xa1, 0xa7, 0x78, 0x7e, 0xdf, 0xef, 0xf7, 0xfd, 0x5e,
	0xe0, 0xe0, 0xfc, 0x36, 0x6b, 0x5a, 0x6e, 0xeb, 0xbc, 0x77, 0x42, 0xa8, 0x43, 0x7c, 0xc2, 0x5a,
	0x7d, 0xe2, 0x98, 0x2e, 0x6d, 0x49, 0x84, 0xee, 0x59, 0x5d, 0xdd, 0x38, 0xb3, 0x1c, 0x42, 0x07,
	0x2d, 0xef, 0xbc, 0xc3, 0x01, 0xac, 0xd5, 0x25, 0xbe, 0xde, 0xea, 0x6f, 0xb6, 0x3a, 0xc4, 0x21,
	0x54, 0xf7, 0x89, 0xd9, 0xf4, 0xa8, 0xeb, 0xbb, 0xe8, 0x8d, 0x80, 0xab, 0x99, 0xe4, 0x6a, 0x7a,
	0xe7, 0x1d, 0x0e, 0x60, 0x4d, 0xce, 0xd5, 0xec, 0x6f, 0xae, 0xbd, 0xd5, 0xb1, 0xfc, 0xb3, 0xde,
	0x49, 0xd3, 0x70, 0xbb, 0xad, 0x8e, 0xdb, 0x71, 0x5b, 0x82, 0xf9, 0xa4, 0x77, 0x2a, 0xbe, 0xc4,
	0x87, 0xf8, 0x15, 0x08, 0x5d, 0x9b, 0x6a, 0x0a, 0xed, 0x39, 0xbe, 0xd5, 0x25, 0xe3, 0x56, 0xac,
	0xbd, 0x73, 0x1d, 0x03, 0x33, 0xce, 0x48, 0x57, 0x9f, 0xe0, 0x7b, 0x7b, 0x1a, 0x5f, 0xcf, 0xb7,
	0xec, 0x96, 0xe5, 0xf8, 0xcc, 0xa7, 0xe3, 0x4c, 0xea, 0xdf, 0xf3, 0x50, 0xda, 0x3e, 0xda, 0xdf,
	0xa3, 0x6e, 0xcf, 0x43, 0xeb, 0x30, 0xef, 0xe8, 0x5d, 0x52, 0x53, 0xd6, 0x95, 0x8d, 0xb2, 0xb6,
	0xf0, 0x62, 0xd8, 0x98, 0x1b, 0x0d, 0x1b, 0xf3, 0x0f, 0xf4, 0x2e, 0xc1, 0x02, 0x83, 0x6c, 0x28,
	0xf5, 0x09, 0x65, 0x96, 0xeb, 0xb0, 0x5a, 0x6e, 0x3d, 0xbf, 0x51, 0xd9, 0xba, 0xdb, 0xcc, 0xe2,
	0xb4, 0xa6, 0x50, 0xf0, 0x28, 0x60, 0xbd, 0xe7, 0xd2, 0xb6, 0xc5, 0x0c, 0xb7, 0x4f, 0xe8, 0x40,
	0x5b, 0x96, 0x5a, 0x4a, 0x12, 0xc9, 0x70, 0xa4, 0x01, 0xfd, 0x4a, 0x81, 0x65, 0x8f, 0x92, 0x53,
	0x42, 0x29, 0x31, 0x25, 0xbe, 0x96, 0x5f, 0x57, 0x3e, 0x07, 0xb5, 0x35, 0xa9, 0x76, 0xf9, 0x68,
	0x4c, 0x3e, 0x9e, 0xd0, 0x88, 0xfe, 0xa4, 0xc0, 0x1a, 0x23, 0xb4, 0x4f, 0xe8, 0xb6, 0x69, 0x52,
	0xc2, 0x98, 0x36, 0xd8, 0xb1, 0x2d, 0xe2, 0xf8, 0x3b, 0xfb, 0x6d, 0xcc, 0x6a, 0xf3, 0xc2, 0x0f,
	0xdf, 0xcf, 0x66, 0xd0, 0xf1, 0x34, 0x39, 0x9a, 0x2a, 0x2d, 0x5a, 0x9b, 0x4a, 0xc2, 0xf0, 0xa7,
	0x98, 0xa1, 0x9e, 0xc2, 0x42, 0x18, 0xc8, 0xfb, 0x16, 0xf3, 0xd1, 0x23, 0x28, 0x76, 0xf8, 0x07,
	0xab, 0x29, 0xc2, 0xc0, 0x66, 0x36, 0x03, 0x43, 0x19, 0xda, 0xa2, 0xb4, 0xa7, 0x28, 0x3e, 0x19,
	0x96, 0xd2, 0xd4, 0x0f, 0x73, 0x50, 0xd9, 0x3e, 0xda, 0xc7, 0x84, 0xb9, 0x3d, 0x6a, 0x90, 0x0c,
	0x49, 0xb3, 0x05, 0xc0, 0xff, 0x32, 0x4f, 0x37, 0x88, 0x59, 0xcb, 0xad, 0x2b, 0x1b, 0x25, 0x0d,
	0x49, 0x3a, 0x78, 0x10, 0x61, 0x70, 0x82, 0x8a, 0x4b, 0x3d, 0xb7, 0x1c, 0x53, 0x44, 0x3b, 0x21,
	0xf5, 0x5d, 0xcb, 0x31, 0xb1, 0xc0, 0xa0, 0xfb, 0x50, 0xe8, 0x13, 0x7a, 0xc2, 0xfd, 0xcf, 0x13,
	0xe2, 0xeb, 0xd9, 0x8e, 0xf7, 0x88, 0xb3, 0x68, 0xe5, 0xd1, 0xb0, 0x51, 0x10, 0x3f, 0x71, 0x20,
	0x04, 0x35, 0x01, 0xd8, 0x99, 0x4b, 0x7d, 0x61, 0x4e, 0xad, 0xb0, 0x9e, 0xdf, 0x28, 0x6b, 0x8b,
	0xdc, 0xbe, 0xe3, 0x08, 0x8a, 0x13, 0x14, 0xea, 0x5f, 0x15, 0x58, 0x4a, 0x78, 0x41, 0x78, 0xfc,
	0x36, 0x2c, 0x74, 0x12, 0xf9, 0x26, 0x3d, 0xb2, 0x2a, 0x6d, 0x5f, 0x48, 0xe6, 0x22, 0x4e, 0x51,
	0x22, 0x02, 0x65, 0x2a, 0x25, 0x85, 0x75, 0xb5, 0x99, 0x39, 0x5c, 0xa1, 0x0d, 0xb1, 0xa6, 0x04,
	0x90, 0xe1, 0x58, 0xb2, 0xfa, 0x1f, 0x45, 0x84, 0x2e, 0xac, 0x34, 0xb4, 0x91, 0xa8, 0x66, 0x45,
	0x1c, 0x79, 0x61, 0x4a, 0x25, 0x5e, 0x53, 0x02, 0xb9, 0x2f, 0x44, 0x09, 0xdc, 0x29, 0xfd, 0xfe,
	0xa3, 0xc6, 0xdc, 0x07, 0xff, 0x5a, 0x9f, 0x53, 0x3f, 0xc9, 0x41, 0xb5, 0x4d, 0x6c, 0xe2, 0x93,
	0x43, 0xcf, 0x17, 0x27, 0xb8, 0x07, 0xa8, 0x43, 0x75, 0x83, 0x1c, 0x11, 0x6a, 0xb9, 0xe6, 0x31,
	0x31, 0x5c, 0xc7, 0x64, 0x22, 0x44, 0x79, 0xed, 0x4b, 0xa3, 0x61, 0x03, 0xed, 0x4d, 0x60, 0xf1,
	0x15, 0x1c, 0xc8, 0x86, 0xaa, 0x47, 0xc5, 0x6f, 0xcb, 0x97, 0x6d, 0x90, 0xa7, 0xdf, 0xdb, 0xd9,
	0xce, 0x7e, 0x94, 0x64, 0xd5, 0x56, 0x46, 0xc3, 0x46, 0x35, 0x05, 0xc2, 0x69, 0xe1, 0xe8, 0x07,
	0xb0, 0xec, 0x52, 0xef, 0x4c, 0x77, 0xda, 0xc4, 0x23, 0x8e, 0x49, 0x1c, 0x9f, 0x89, 0x92, 0x28,
	0x69, 0xab, 0xbc, 0x79, 0x1d, 0x8e, 0xe1, 0xf0, 0x04, 0x35, 0x7a, 0x0c, 0x2b, 0x1e, 0x75, 0x3d,
	0xbd, 0xa3, 0x73, 0x89, 0x47, 0xae, 0x6d, 0x19, 0x03, 0x51, 0x32, 0x65, 0xed, 0xd6, 0x68, 0xd8,
	0x58, 0x39, 0x1a, 0x47, 0x5e, 0x0e, 0x1b, 0xaf, 0x08, 0xd7, 0x71, 0x48, 0x8c, 0xc4, 0x93, 0x62,
	0xd4, 0x7d, 0x28, 0xb5, 0x7b, 0x54, 0x40, 0xd0, 0xf7, 0xa0, 0x64, 0xca, 0xdf, 0xd2, 0xab, 0xaf,
	0x85, 0x9d, 0x3d, 0xa4, 0xb9, 0x1c, 0x36, 0xaa, 0x7c, 0x80, 0x35, 0x43, 0x00, 0x8e, 0x58, 0xd4,
	0x27, 0x50, 0xdd, 0xbd, 0xf0, 0x5c, 0xea, 0x87, 0xf1, 0xfa, 0x2a, 0x14, 0x89, 0x00, 0x08, 0x69,
	0xa5, 0xb8, 0x1d, 0x05, 0x64, 0x58, 0x62, 0xd1, 0xeb, 0x50, 0x20, 0x17, 0xba, 0xe1, 0xcb, 0xbe,
	0x52, 0x95, 0x64, 0x85, 0x5d, 0x0e, 0xc4, 0x01, 0x4e, 0x3d, 0x04, 0xd8, 0x23, 0x91, 0xe8, 0x6d,
	0x58, 0x0a, 0x6b, 0x22, 0x5d, 0xaa, 0x5f, 0x96, 0xcc, 0x4b, 0x38, 0x8d, 0xc6, 0xe3, 0xf4, 0xea,
	0x13, 0x28, 0x8b, 0x72, 0xe6, 0xfd, 0x88, 0x9b, 0x20, 0xaa, 0x59, 0x4a, 0x89, 0x4c, 0x10, 0x14,
	0x38, 0xc0, 0x45, 0x0d, 0x2d, 0x37, 0xad, 0xa1, 0x25, 0xb2, 0xd7, 0x86, 0x6a, 0xc0, 0x1b, 0xf6,
	0xd8, 0x4c, 0x1a, 0x6e, 0x41, 0x29, 0x34, 0x53, 0x6a, 0x89, 0x66, 0x6b, 0x28, 0x08, 0x47, 0x14,
	0x09, 0x6d, 0x67, 0x90, 0x6a, 0x4d, 0xd9, 0x94, 0xbd, 0x09, 0x37, 0x64, 0x73, 0x90, 0xba, 0x96,
	0x24, 0xd9, 0x8d, 0xd0, 0x67, 0x21, 0x3e, 0xa1, 0xe9, 0x97, 0x50, 0x9b, 0x36, 0x90, 0x3f, 0x43,
	0xf3, 0xcc, 0x6e, 0x8a, 0xfa, 0x3b, 0x05, 0x96, 0x93, 0x92, 0xb2, 0x87, 0x2f, 0xbb, 0x92, 0xeb,
	0x47, 0x57, 0xc2, 0x23, 0x7f, 0x54, 0x60, 0x35, 0x75, 0xb4, 0x99, 0x22, 0x3e, 0x83, 0x51, 0xc9,
	0xe4, 0xc8, 0xcf, 0x90, 0x1c, 0xff, 0xc8, 0x41, 0xf5, 0xbe, 0x7e, 0x42, 0xec, 0x63, 0x62, 0x13,
	0xc3, 0x77, 0x29, 0x7a, 0x1f, 0x2a, 0x5d, 0xdd, 0x37, 0xce, 0x04, 0x34, 0x5c, 0x2e, 0xda, 0xd9,
	0xda, 0x5f, 0x4a, 0x52, 0xf3, 0x20, 0x16, 0xb3, 0xeb, 0xf8, 0x74, 0xa0, 0xbd, 0x22, 0x4d, 0xaa,
	0x24, 0x30, 0x38, 0xa9, 0x4d, 0x6c, 0x84, 0xe2, 0x7b, 0xf7, 0xc2, 0xe3, 0xfd, 0x7f, 0xf6, 0x45,
	0x34, 0x65, 0x02, 0x26, 0x3f, 0xef, 0x59, 0x94, 0x74, 0x89, 0xe3, 0xc7, 0x1b, 0xe1, 0xc1, 0x98,
	0x7c, 0x3c, 0xa1, 0x71, 0xed, 0x2e, 0x2c, 0x8f, 0x1b, 0x8f, 0x96, 0x21, 0x7f, 0x4e, 0x06, 0x41,
	0xbc, 0x30, 0xff, 0x89, 0x56, 0xa1, 0xd0, 0xd7, 0xed, 0x9e, 0xac, 0x46, 0x1c, 0x7c, 0xdc, 0xc9,
	0xdd, 0x56, 0xd4, 0x3f, 0x2b, 0x50, 0x9b, 0x66, 0x08, 0xfa, 0x4a, 0x42, 0x90, 0x56, 0x91, 0x56,
	0xe5, 0xdf, 0x25, 0x83, 0x40, 0xea, 0x2e, 0x94, 0x5c, 0x8f, 0xef, 0xf0, 0x2e, 0x95, 0x51, 0x7f,
	0x33, 0x8c, 0xe4, 0xa1, 0x84, 0x5f, 0x0e, 0x1b, 0x37, 0x53, 0xe2, 0x43, 0x04, 0x8e, 0x58, 0x91,
	0x0a, 0x45, 0x61, 0x0f, 0x9f, 0x27, 0x7c, 0xf2, 0x03, 0xef, 0xad, 0x8f, 0x04, 0x04, 0x4b, 0x8c,
	0xfa, 0x3e, 0x94, 0xf8, 0x62, 0x73, 0x40, 0x7c, 0x9d, 0x27, 0x10, 0x23, 0xf6, 0xe9, 0x7d, 0xcb,
	0x39, 0x97, 0xa6, 0x45, 0x09, 0x74, 0x2c, 0xe1, 0x38, 0xa2, 0xb8, 0xaa, 0xc5, 0xe6, 0x66, 0x6c,
	0xb1, 0x7f, 0xc9, 0x41, 0x85, 0x6b, 0x0f, 0xbb, 0xf6, 0x77, 0xa0, 0x6a, 0x27, 0xcf, 0x24, 0xad,
	0xb8, 0x29, 0x05, 0xa6, 0xb3, 0x14, 0xa7, 0x69, 0x39, 0xf3, 0xa9, 0x45, 0x6c, 0x33, 0x62, 0xce,
	0xa5, 0x99, 0xef, 0x25, 0x91, 0x38, 0x4d, 0xcb, 0x6b, 0xf1, 0x29, 0x8f, 0xb6, 0x9c, 0xbc, 0x51,
	0x2d, 0xfe, 0x90, 0x03, 0x71, 0x80, 0xbb, 0xea, 0xc4, 0xf3, 0xb3, 0x9d, 0x18, 0xdd, 0x81, 0x45,
	0x3e, 0x1e, 0xdd, 0x9e, 0x1f, 0xae, 0x27, 0x05, 0x31, 0x48, 0xd1, 0x68, 0xd8, 0x58, 0x7c, 0x2f,
	0x85, 0xc1, 0x63, 0x94, 0xea, 0x87, 0x00, 0x70, 0x78, 0xf2, 0x33, 0x62, 0x04, 0xd1, 0xba, 0x7e,
	0x29, 0xe7, 0xfd, 0x56, 0xde, 0x05, 0x39, 0x54, 0x3a, 0x24, 0xee, 0xb7, 0x09, 0x1c, 0x4e, 0x51,
	0xa2, 0x16, 0x94, 0xa3, 0x45, 0x5d, 0xf6, 0x92, 0x15, 0xc9, 0x56, 0x8e, 0xb6, 0x79, 0x1c, 0xd3,
	0xa4, 0x52, 0x67, 0xfe, 0xda, 0xd4, 0xd1, 0x20, 0xdf, 0xb3, 0x4c, 0x71, 0xf4, 0xb2, 0xf6, 0x8d,
	0x30, 0xfd, 0x1f, 0xee, 0xb7, 0x2f, 0x87, 0x8d, 0xd7, 0xa6, 0x5d, 0x71, 0xfd, 0x81, 0x47, 0x58,
	0xf3, 0xe1, 0x7e, 0x1b, 0x73, 0xe6, 0xab, 0x82, 0x51, 0x9c, 0x31, 0x18, 0x5b, 0x00, 0xf2, 0xd4,
	0x9c, 0xfb, 0x46, 0x10, 0x88, 0xf0, 0xd2, 0xb2, 0x17, 0x61, 0x70, 0x82, 0x0a, 0x31, 0x58, 0x31,
	0x28, 0x11, 0xbf, 0x79, 0xb8, 0x98, 0xaf, 0x77, 0xbd, 0x5a, 0x49, 0xec, 0x87, 0x5f, 0xcb, 0xd6,
	0x9d, 0x38, 0x9b, 0xf6, 0xaa, 0x54, 0xb3, 0xb2, 0x33, 0x2e, 0x0c, 0x4f, 0xca, 0x47, 0x2e, 0xac,
	0x98, 0x72, 0x5d, 0x8b, 0x95, 0x96, 0x67, 0x56, 0x7a, 0x93, 0x2b, 0x6c, 0x8f, 0x0b, 0xc2, 0x93,
	0xb2, 0xd1, 0x4f, 0x60, 0x2d, 0x04, 0x4e, 0xee, 0xcc, 0x35, 0x10, 0x9e, 0xaa, 0xf3, 0x2d, 0xbe,
	0x3d, 0x95, 0x0a, 0x7f, 0x8a, 0x04, 0x64, 0x42, 0xd1, 0x0e, 0x66, 0x4b, 0x45, 0x34, 0xf6, 0xef,
	0x66, 0x3b, 0x45, 0x9c, 0xfd, 0xcd, 0xe4, 0x4c, 0x89, 0xf6, 0x46, 0x39, 0x4e, 0xa4, 0x6c, 0x74,
	0x01, 0x15, 0xdd, 0x71, 0x5c, 0x5f, 0x0f, 0xb6, 0xf8, 0x05, 0xa1, 0x6a, 0x7b, 0x66, 0x55, 0xdb,
	0xb1, 0x8c, 0xb1, 0x19, 0x96, 0xc0, 0xe0, 0xa4, 0x2a, 0xf4, 0x14, 0x96, 0xdc, 0xa7, 0x0e, 0xa1,
	0x98, 0x9c, 0x12, 0x4a, 0x1c, 0x7e, 0xe5, 0xab, 0x0a, 0xed, 0xdf, 0xcc, 0xa8, 0x3d, 0xc5, 0x1c,
	0xa7, 0x74, 0x1a, 0xce, 0xf0, 0xb8, 0x16, 0x7e, 0xc7, 0x3d, 0xb5, 0x1c, 0xdd, 0xb6, 0x7e, 0x41,
	0x28, 0xab, 0x2d, 0xc6, 0x77, 0xdc, 0x7b, 0x11, 0x14, 0x27, 0x28, 0xd0, 0xb7, 0xa0, 0x62, 0xd8,
	0x3d, 0xe6, 0x13, 0x2a, 0x3a, 0xc4, 0x92, 0xa8, 0xa0, 0xe8, 0x7c, 0x3b, 0x31, 0x0a, 0x27, 0xe9,
	0xd6, 0xbe, 0x0d, 0x95, 0xff, 0x73, 0x2e, 0xf2, 0xb9, 0x3a, 0xee, 0xd0, 0x99, 0xe6, 0xea, 0xdf,
	0x72, 0xb0, 0x98, 0x76, 0x43, 0xb4, 0x8d, 0x29, 0x53, 0x1f, 0x12, 0xc2, 0x5e, 0x99, 0x9f, 0xda,
	0x2b, 0x65, 0x4b, 0x9a, 0xff, 0x2c, 0x2d, 0x69, 0x0b, 0x40, 0xf7, 0xac, 0xb0, 0x1b, 0x05, 0xdd,
	0x2d, 0xea, 0x27, 0xf1, 0xa5, 0x1c, 0x27, 0xa8, 0x78, 0xc0, 0x0c, 0xd7, 0xf1, 0xa9, 0x6b, 0xdb,
	0x84, 0x8a, 0x0e, 0x56, 0x0a, 0x02, 0xb6, 0x13, 0x41, 0x71, 0x82, 0x82, 0xdf, 0x71, 0x4f, 0x6c,
	0xd7, 0x38, 0x17, 0x2e, 0x08, 0xab, 0x4f, 0xf4, 0xae, 0x52, 0x70, 0xc7, 0xd5, 0x26, 0xb0, 0xf8,
	0x0a, 0x0e, 0xf5, 0x10, 0xd2, 0xb7, 0x52, 0x74, 0x37, 0x70, 0x80, 0x12, 0x5d, 0x1b, 0x67, 0x3b,
	0xbc, 0x7a, 0x0b, 0xca, 0xd8, 0x75, 0xfd, 0x23, 0xdd, 0x3f, 0x63, 0xa8, 0x01, 0x05, 0x8f, 0xff,
	0x90, 0x4f, 0x0e, 0xe2, 0x2d, 0x46, 0x60, 0x70, 0x00, 0x57, 0x7f, 0xab, 0xc0, 0xab, 0x53, 0x5f,
	0x00, 0xb8, 0x23, 0x8d, 0xe8, 0x4b, 0x9a, 0x14, 0x39, 0x32, 0xa6, 0xc3, 0x09, 0x2a, 0x3e, 0xfe,
	0x53, 0xcf, 0x06, 0xe3, 0xe3, 0x3f, 0xa5, 0x0d, 0xa7, 0x69, 0xd5, 0xff, 0xe6, 0xa0, 0x78, 0xec,
	0xeb, 0x7e, 0x8f, 0xa1, 0x27, 0x50, 0xe2, 0x55, 0x68, 0xea, 0xbe, 0x2e, 0x34, 0x67, 0x7e, 0x55,
	0x0b, 0xd7, 0xa8, 0x78, 0xf2, 0x85, 0x10, 0x1c, 0x49, 0xe4, 0x57, 0x5e, 0x26, 0xf4, 0x48, 0xf3,
	0xa2, 0xd6, 0x15, 0x68, 0xc7, 0x12, 0xcb, 0xd7, 0xfe, 0x2e, 0x61, 0x4c, 0xef, 0x84, 0x39, 0x1b,
	0xad, 0xfd, 0x07, 0x01, 0x18, 0x87, 0x78, 0xf4, 0x0e, 0x14, 0x29, 0xd1, 0x59, 0xb4, 0x8c, 0xd4,
	0x43, 0x91, 0x58, 0x40, 0x2f, 0x87, 0x8d, 0x05, 0x29, 0x5c, 0x7c, 0x63, 0x49, 0x8d, 0x1e, 0xc3,
	0x0d, 0x93, 0xf8, 0xba, 0x65, 0x07, 0x3b, 0x48, 0xe6, 0xf7, 0x8d, 0x40, 0x58, 0x3b, 0x60, 0xd5,
	0x2a, 0xdc, 0x26, 0xf9, 0x81, 0x43, 0x81, 0xbc, 0xde, 0x0c, 0xd7, 0x24, 0x22, 0x9f, 0x0b, 0x71,
	0xbd, 0xed, 0xb8, 0x26, 0xc1, 0x02, 0xa3, 0x3e, 0x53, 0xa0, 0x12, 0x48, 0xda, 0xd1, 0x7b, 0x8c,
	0xa0, 0xcd, 0xe8, 0x14, 0x41, 0xb8, 0xc3, 0x01, 0x39, 0xff, 0xde, 0xc0, 0x23, 0x97, 0xc3, 0x46,
	0x59, 0x90, 0xf1, 0x8f, 0xe8, 0x00, 0x09, 0x1f, 0xe5, 0xae, 0xf1, 0xd1, 0xeb, 0x50, 0x10, 0xfb,
	0x9e, 0x74, 0x66, 0xb4, 0xde, 0x89, 0x9d, 0x10, 0x07, 0x38, 0xf5, 0x0f, 0x39, 0xa8, 0xa6, 0x0e,
	0x97, 0x61, 0xc5, 0x8a, 0xee, 0x70, 0xb9, 0x0c, 0xef, 0x02, 0xd3, 0x1f, 0x3a, 0x7f, 0x04, 0x45,
	0x83, 0x9f, 0x2f, 0x7c, 0x69, 0xde, 0x9c, 0x25, 0x14, 0xc2, 0x33, 0x71, 0x26, 0x89, 0x4f, 0x86,
	0xa5, 0x40, 0xb4, 0x07, 0x2b, 0x94, 0xf8, 0x74, 0xb0, 0x7d, 0xea, 0x13, 0x9a, 0x5c, 0x3a, 0x0b,
	0xf1, 0x12, 0x82, 0xc7, 0x09, 0xf0, 0x24, 0x8f, 0x6a, 0xc3, 0x3c, 0x5f, 0x10, 0xb8, 0xdb, 0x59,
	0xea, 0x69, 0x2d, 0x72, 0x7b, 0xc8, 0x1c, 0xe2, 0xb9, 0x77, 0x1c, 0xdd, 0x71, 0x83, 0x64, 0x2f,
	0xc4, 0xde, 0x79, 0xc0, 0x81, 0x38, 0xc0, 0xdd, 0x59, 0xe5, 0x17, 0xd1, 0xdf, 0x3c, 0x6f, 0xcc,
	0x3d, 0x7b, 0xde, 0x98, 0xfb, 0xe8, 0xb9, 0xbc, 0x94, 0xfe, 0x18, 0xca, 0xf1, 0x3a, 0xf2, 0x39,
	0xab, 0x54, 0x7f, 0x0a, 0x25, 0x9e, 0x49, 0xe1, 0x1a, 0x7d, 0xcd, 0xf0, 0x48, 0xb7, 0xf5, 0x5c,
	0x96, 0xb6, 0xae, 0x6e, 0x41, 0xf0, 0xf6, 0xcc, 0x3b, 0xa1, 0xe5, 0x93, 0x6e, 0xaa, 0x13, 0xee,
	0x73, 0x00, 0x0e, 0xe0, 0x89, 0x7b, 0xf8, 0xaf, 0x15, 0x00, 0x71, 0xdf, 0xd8, 0xed, 0xf3, 0x3b,
	0xe2, 0x3a, 0xcc, 0xf3, 0x16, 0x3b, 0x6e, 0x98, 0x28, 0x01, 0x81, 0x41, 0x0f, 0xa1, 0xe8, 0x8a,
	0x35, 0x45, 0x3e, 0x50, 0xbe, 0x35, 0x35, 0x6b, 0xe4, 0xbf, 0x95, 0x9a, 0x58, 0x7f, 0xba, 0x7b,
	0xe1, 0x13, 0x87, 0xdb, 0x18, 0x67, 0x4c, 0xb0, 0xeb, 0x60, 0x29, 0x4c, 0x7b, 0xe3, 0xc5, 0xcb,
	0xfa, 0xdc, 0xc7, 0x2f, 0xeb, 0x73, 0xff, 0x7c, 0x59, 0x9f, 0xfb, 0x60, 0x54, 0x57, 0x5e, 0x8c,
	0xea, 0xca, 0xc7, 0xa3, 0xba, 0xf2, 0xef, 0x51, 0x5d, 0x79, 0xf6, 0x49, 0x7d, 0xee, 0x71, 0xae,
	0xbf, 0xf9, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb1, 0xe8, 0xc1, 0x2f, 0x98, 0x1b, 0x00, 0x00,
}
