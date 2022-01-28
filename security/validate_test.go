package security

import (
	"testing"
)

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
		{
			"abcde fghij",
			args{pass: "abcde fghij"},
			true,
		},
		{
			"abcde xyz ecdab",
			args{pass: "abcde xyz ecdab"},
			false,
		},
		{
			"a ab abc abd abf abj",
			args{pass: "a ab abc abd abf abj"},
			true,
		},
		{
			"iiii oiii ooii oooi oooo",
			args{pass: "iiii oiii ooii oooi oooo"},
			true,
		},
		{
			"oiii ioii iioi iiio",
			args{pass: "oiii ioii iioi iiio"},
			false,
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

func Test_isAnagram(t *testing.T) {
	type args struct {
		w1 string
		w2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"abcde fghij",
			args{"abcde", "fghij"},
			false,
		},
		{
			"abcde fghij",
			args{"abcde", "fghij"},
			false,
		},
		{
			"ioii iioi",
			args{"ioii", "iioi"},
			true,
		},
		{
			"abcde ecdab",
			args{"abcde", "ecdab"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.w1, tt.args.w2); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
