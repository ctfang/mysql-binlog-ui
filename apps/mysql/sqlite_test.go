package mysql

import (
	"reflect"
	"testing"
)

func Test_getSqliteDBName(t *testing.T) {
	type args struct {
		binlogFilePath string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want2 string
	}{
		{
			name:  "腾讯云格式",
			args:  args{"test/Q-sh-mhhy-prod-farm_binlog_mysqlbin.014384"},
			want:  "014384",
			want2: "Q-sh-mhhy-prod-farm-binlog-mysqlbin",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, table := getSqliteDBName(tt.args.binlogFilePath)
			if !reflect.DeepEqual(table, tt.want) {
				t.Errorf("getSqliteDBName() got = %v, want %v", table, tt.want)
			}
		})
	}
}
