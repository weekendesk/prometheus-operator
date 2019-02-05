package v1

import (
	"reflect"
	"testing"
)

func TestCloneSelectorAndAddLabel(t *testing.T) {
	labels := map[string]string{
		"foo1":	"bar1",
		"foo2":	"bar2",
		"foo3":	"bar3",
	}

	cases := []struct {
		labels		map[string]string
		labelKey	string
		labelValue	string
		want		map[string]string
	}{
		{
			labels:	labels,
			want:	labels,
		},
		{
			labels:		labels,
			labelKey:	"foo4",
			labelValue:	"89",
			want: map[string]string{
				"foo1":	"bar1",
				"foo2":	"bar2",
				"foo3":	"bar3",
				"foo4":	"89",
			},
		},
		{
			labels:		nil,
			labelKey:	"foo4",
			labelValue:	"12",
			want: map[string]string{
				"foo4": "12",
			},
		},
	}

	for _, tc := range cases {
		ls_in := LabelSelector{MatchLabels: tc.labels}
		ls_out := LabelSelector{MatchLabels: tc.want}

		got := CloneSelectorAndAddLabel(&ls_in, tc.labelKey, tc.labelValue)
		if !reflect.DeepEqual(got, &ls_out) {
			t.Errorf("got %v, want %v", got, tc.want)
		}
	}
}

func TestAddLabelToSelector(t *testing.T) {
	labels := map[string]string{
		"foo1":	"bar1",
		"foo2":	"bar2",
		"foo3":	"bar3",
	}

	cases := []struct {
		labels		map[string]string
		labelKey	string
		labelValue	string
		want		map[string]string
	}{
		{
			labels:	labels,
			want:	labels,
		},
		{
			labels:		labels,
			labelKey:	"foo4",
			labelValue:	"89",
			want: map[string]string{
				"foo1":	"bar1",
				"foo2":	"bar2",
				"foo3":	"bar3",
				"foo4":	"89",
			},
		},
		{
			labels:		nil,
			labelKey:	"foo4",
			labelValue:	"12",
			want: map[string]string{
				"foo4": "12",
			},
		},
	}

	for _, tc := range cases {
		ls_in := LabelSelector{MatchLabels: tc.labels}
		ls_out := LabelSelector{MatchLabels: tc.want}

		got := AddLabelToSelector(&ls_in, tc.labelKey, tc.labelValue)
		if !reflect.DeepEqual(got, &ls_out) {
			t.Errorf("got %v, want %v", got, tc.want)
		}
	}
}
