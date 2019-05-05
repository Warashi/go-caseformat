package caseformat_test

import (
	"testing"

	"github.com/Warashi/go-caseformat"
)

func TestUpperUnderscore_ToLowerHyphen(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal_case",
			s:    "UPPER_UNDERSCORE",
			want: "upper-underscore",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseformat.FromUpperUnderscore(tt.s).ToLowerHyphen(); got != tt.want {
				t.Errorf("upperUnderscore.ToLowerHyphen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperUnderscore_ToLowerUnderscore(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal_case",
			s:    "UPPER_UNDERSCORE",
			want: "upper_underscore",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseformat.FromUpperUnderscore(tt.s).ToLowerUnderscore(); got != tt.want {
				t.Errorf("upperUnderscore.ToLowerUnderscore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperUnderscore_ToLowerCamel(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal_case",
			s:    "UPPER_UNDERSCORE",
			want: "upperUnderscore",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseformat.FromUpperUnderscore(tt.s).ToLowerCamel(); got != tt.want {
				t.Errorf("upperUnderscore.ToLowerCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperUnderscore_ToUpperHyphen(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal_case",
			s:    "UPPER_UNDERSCORE",
			want: "UPPER-UNDERSCORE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseformat.FromUpperUnderscore(tt.s).ToUpperHyphen(); got != tt.want {
				t.Errorf("upperUnderscore.ToUpperHyphen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperUnderscore_ToUpperUnderscore(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal_case",
			s:    "UPPER_UNDERSCORE",
			want: "UPPER_UNDERSCORE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseformat.FromUpperUnderscore(tt.s).ToUpperUnderscore(); got != tt.want {
				t.Errorf("upperUnderscore.ToUpperUnderscore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperUnderscore_ToUpperCamel(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal_case",
			s:    "UPPER_UNDERSCORE",
			want: "UpperUnderscore",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseformat.FromUpperUnderscore(tt.s).ToUpperCamel(); got != tt.want {
				t.Errorf("upperUnderscore.ToUpperCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}
