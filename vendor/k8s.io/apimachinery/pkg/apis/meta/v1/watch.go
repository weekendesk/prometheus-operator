package v1

import (
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
)

// Event represents a single event to a watched resource.
//
// +protobuf=true
type WatchEvent struct {
	Type	string	`json:"type" protobuf:"bytes,1,opt,name=type"`

	// Object is:
	//  * If Type is Added or Modified: the new state of the object.
	//  * If Type is Deleted: the state of the object immediately before deletion.
	//  * If Type is Error: *Status is recommended; other types may make sense
	//    depending on context.
	Object	runtime.RawExtension	`json:"object" protobuf:"bytes,2,opt,name=object"`
}

func Convert_watch_Event_to_versioned_Event(in *watch.Event, out *WatchEvent, s conversion.Scope) error {
	out.Type = string(in.Type)
	switch t := in.Object.(type) {
	case *runtime.Unknown:

		out.Object.Raw = t.Raw
	case nil:
	default:
		out.Object.Object = in.Object
	}
	return nil
}

func Convert_versioned_InternalEvent_to_versioned_Event(in *InternalEvent, out *WatchEvent, s conversion.Scope) error {
	return Convert_watch_Event_to_versioned_Event((*watch.Event)(in), out, s)
}

func Convert_versioned_Event_to_watch_Event(in *WatchEvent, out *watch.Event, s conversion.Scope) error {
	out.Type = watch.EventType(in.Type)
	if in.Object.Object != nil {
		out.Object = in.Object.Object
	} else if in.Object.Raw != nil {

		out.Object = &runtime.Unknown{
			Raw:		in.Object.Raw,
			ContentType:	runtime.ContentTypeJSON,
		}
	}
	return nil
}

func Convert_versioned_Event_to_versioned_InternalEvent(in *WatchEvent, out *InternalEvent, s conversion.Scope) error {
	return Convert_versioned_Event_to_watch_Event(in, (*watch.Event)(out), s)
}

// InternalEvent makes watch.Event versioned
// +protobuf=false
type InternalEvent watch.Event

func (e *InternalEvent) GetObjectKind() schema.ObjectKind	{ return schema.EmptyObjectKind }
func (e *WatchEvent) GetObjectKind() schema.ObjectKind		{ return schema.EmptyObjectKind }
