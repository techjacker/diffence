package diffence

import "testing"

func Test_noop(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no op test",
			args: args{msg: "test message"},
			want: "test message",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := noop(tt.args.msg); got != tt.want {
				t.Errorf("noop() = %v, want %v", got, tt.want)
			}
		})
	}
}
