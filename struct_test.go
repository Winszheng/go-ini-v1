package go_ini_v1

import (
	"reflect"
	"testing"
)

func TestFile_Section(t *testing.T) {
	type fields struct {
		sections map[string]*Section
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Section
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				sections: tt.fields.sections,
			}
			if got := f.Section(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Section() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSection_Key(t *testing.T) {
	type fields struct {
		value map[string]string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Section{
				value: tt.fields.value,
			}
			if got := s.Key(tt.args.key); got != tt.want {
				t.Errorf("Key() = %v, want %v", got, tt.want)
			}
		})
	}
}