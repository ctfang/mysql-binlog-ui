package ctx

import (
	"errors"
	"testing"
)

func TestLogError(t *testing.T) {
	type args struct {
		msg string
		arr []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				msg: "test",
				arr: []interface{}{1, 2, errors.New("err = 3")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LogError(tt.args.msg, tt.args.arr...)
		})
	}
}
