package v1

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/ugorji/go/codec"
)

type GroupVersionHolder struct {
	GV GroupVersion `json:"val"`
}

func TestGroupVersionUnmarshalJSON(t *testing.T) {
	cases := []struct {
		input	[]byte
		expect	GroupVersion
	}{
		{[]byte(`{"val": "v1"}`), GroupVersion{"", "v1"}},
		{[]byte(`{"val": "extensions/v1beta1"}`), GroupVersion{"extensions", "v1beta1"}},
	}

	for _, c := range cases {
		var result GroupVersionHolder

		if err := json.Unmarshal([]byte(c.input), &result); err != nil {
			t.Errorf("JSON codec failed to unmarshal input '%v': %v", c.input, err)
		}
		if !reflect.DeepEqual(result.GV, c.expect) {
			t.Errorf("JSON codec failed to unmarshal input '%s': expected %+v, got %+v", c.input, c.expect, result.GV)
		}

		if err := codec.NewDecoderBytes(c.input, new(codec.JsonHandle)).Decode(&result); err != nil {
			t.Errorf("Ugorji codec failed to unmarshal input '%v': %v", c.input, err)
		}
		if !reflect.DeepEqual(result.GV, c.expect) {
			t.Errorf("Ugorji codec failed to unmarshal input '%s': expected %+v, got %+v", c.input, c.expect, result.GV)
		}
	}
}

func TestGroupVersionMarshalJSON(t *testing.T) {
	cases := []struct {
		input	GroupVersion
		expect	[]byte
	}{
		{GroupVersion{"", "v1"}, []byte(`{"val":"v1"}`)},
		{GroupVersion{"extensions", "v1beta1"}, []byte(`{"val":"extensions/v1beta1"}`)},
	}

	for _, c := range cases {
		input := GroupVersionHolder{c.input}
		result, err := json.Marshal(&input)
		if err != nil {
			t.Errorf("Failed to marshal input '%v': %v", input, err)
		}
		if !reflect.DeepEqual(result, c.expect) {
			t.Errorf("Failed to marshal input '%+v': expected: %s, got: %s", input, c.expect, result)
		}
	}
}
