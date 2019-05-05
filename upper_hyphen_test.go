package caseformat_test

import (
	"testing"

	"github.com/Warashi/go-caseformat"
)

func TestUpperHyphen_ToLowerHyphen(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal_case",
			s:    "UPPER-HYPHEN",
			want: "upper-hyphen",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseformat.FromUpperHyphen(tt.s).ToLowerHyphen(); got != tt.want {
				t.Errorf("upperHyphen.ToLowerHyphen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperHyphen_ToLowerUnderscore(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal_case",
			s:    "UPPER-HYPHEN",
			want: "upper_hyphen",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseformat.FromUpperHyphen(tt.s).ToLowerUnderscore(); got != tt.want {
				t.Errorf("upperHyphen.ToLowerUnderscore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperHyphen_ToLowerCamel(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal_case",
			s:    "UPPER-HYPHEN",
			want: "upperHyphen",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseformat.FromUpperHyphen(tt.s).ToLowerCamel(); got != tt.want {
				t.Errorf("upperHyphen.ToLowerCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperHyphen_ToUpperHyphen(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal_case",
			s:    "UPPER-HYPHEN",
			want: "UPPER-HYPHEN",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseformat.FromUpperHyphen(tt.s).ToUpperHyphen(); got != tt.want {
				t.Errorf("upperHyphen.ToUpperHyphen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperHyphen_ToUpperUnderscore(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal_case",
			s:    "UPPER-HYPHEN",
			want: "UPPER_HYPHEN",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseformat.FromUpperHyphen(tt.s).ToUpperUnderscore(); got != tt.want {
				t.Errorf("upperHyphen.ToUpperUnderscore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperHyphen_ToUpperCamel(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal_case",
			s:    "UPPER-HYPHEN",
			want: "UpperHyphen",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseformat.FromUpperHyphen(tt.s).ToUpperCamel(); got != tt.want {
				t.Errorf("upperHyphen.ToUpperCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}
