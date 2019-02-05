package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// GroupName is the group name for this API.
const GroupName = "meta.k8s.io"

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1"}

// WatchEventKind is name reserved for serializing watch events.
const WatchEventKind = "WatchEvent"

// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// AddToGroupVersion registers common meta types into schemas.
func AddToGroupVersion(scheme *runtime.Scheme, groupVersion schema.GroupVersion) {
	scheme.AddKnownTypeWithName(groupVersion.WithKind(WatchEventKind), &WatchEvent{})
	scheme.AddKnownTypeWithName(
		schema.GroupVersion{Group: groupVersion.Group, Version: runtime.APIVersionInternal}.WithKind(WatchEventKind),
		&InternalEvent{},
	)

	scheme.AddKnownTypes(groupVersion,
		&ListOptions{},
		&ExportOptions{},
		&GetOptions{},
		&DeleteOptions{},
	)
	scheme.AddConversionFuncs(
		Convert_versioned_Event_to_watch_Event,
		Convert_versioned_InternalEvent_to_versioned_Event,
		Convert_watch_Event_to_versioned_Event,
		Convert_versioned_Event_to_versioned_InternalEvent,
	)

	scheme.AddGeneratedDeepCopyFuncs(GetGeneratedDeepCopyFuncs()...)
	AddConversionFuncs(scheme)
	RegisterDefaults(scheme)
}

// scheme is the registry for the common types that adhere to the meta v1 API spec.
var scheme = runtime.NewScheme()

// ParameterCodec knows about query parameters used with the meta v1 API spec.
var ParameterCodec = runtime.NewParameterCodec(scheme)

func init() {
	scheme.AddUnversionedTypes(SchemeGroupVersion,
		&ListOptions{},
		&ExportOptions{},
		&GetOptions{},
		&DeleteOptions{},
	)

	scheme.AddGeneratedDeepCopyFuncs(GetGeneratedDeepCopyFuncs()...)
	RegisterDefaults(scheme)
}
