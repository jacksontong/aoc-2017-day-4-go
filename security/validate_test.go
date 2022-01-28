package security

import "testing"

func TestValidatePassword(t *testing.T) {
	type args struct {
		pass string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"aa bb cc dd ee",
			args{pass: "aa bb cc dd ee"},
			true,
		},
		{
			"aa bb cc dd aa",
			args{pass: "aa bb cc dd aa"},
			false,
		},
		{
			"iqqv iqqv neui iqqv",
			args{pass: "iqqv iqqv neui iqqv"},
			false,
		},
		{
			"aa bb cc dd aaa",
			args{pass: "aa bb cc dd aaa"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidatePassword(tt.args.pass); got != tt.want {
				t.Errorf("ValidatePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
