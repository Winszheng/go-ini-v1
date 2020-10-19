package go_ini_v1

import (
	"reflect"
	"testing"
)

func TestList_Listen(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{}
			if err := l.Listen(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("Listen() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWatch(t *testing.T) {
	type args struct {
		filename string
		listen   List
	}
	tests := []struct {
		name    string
		args    args
		want    *File
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Watch(tt.args.filename, tt.args.listen)
			if (err != nil) != tt.wantErr {
				t.Errorf("Watch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Watch() got = %v, want %v", got, tt.want)
			}
		})
	}
}