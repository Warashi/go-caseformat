package caseformat_test

import (
	"testing"

	"github.com/Warashi/go-caseformat"
)

func TestUpperCamel_ToLowerHyphen(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal_case",
			s:    "UpperCamel",
			want: "upper-camel",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseformat.FromUpperCamel(tt.s).ToLowerHyphen(); got != tt.want {
				t.Errorf("upperCamel.ToLowerHyphen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperCamel_ToLowerUnderscore(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal_case",
			s:    "UpperCamel",
			want: "upper_camel",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseformat.FromUpperCamel(tt.s).ToLowerUnderscore(); got != tt.want {
				t.Errorf("upperCamel.ToLowerUnderscore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperCamel_ToLowerCamel(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal_case",
			s:    "UpperCamel",
			want: "upperCamel",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseformat.FromUpperCamel(tt.s).ToLowerCamel(); got != tt.want {
				t.Errorf("upperCamel.ToLowerCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperCamel_ToUpperHyphen(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal_case",
			s:    "UpperCamel",
			want: "UPPER-CAMEL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseformat.FromUpperCamel(tt.s).ToUpperHyphen(); got != tt.want {
				t.Errorf("upperCamel.ToUpperHyphen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperCamel_ToUpperUnderscore(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal_case",
			s:    "UpperCamel",
			want: "UPPER_CAMEL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseformat.FromUpperCamel(tt.s).ToUpperUnderscore(); got != tt.want {
				t.Errorf("upperCamel.ToUpperUnderscore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperCamel_ToUpperCamel(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal_case",
			s:    "UpperCamel",
			want: "UpperCamel",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseformat.FromUpperCamel(tt.s).ToUpperCamel(); got != tt.want {
				t.Errorf("upperCamel.ToUpperCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}
