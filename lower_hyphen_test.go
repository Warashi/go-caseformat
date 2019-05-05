package caseformat_test

import (
	"testing"

	"github.com/Warashi/go-caseformat"
)

func TestLowerHyphen_ToLowerHyphen(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal_case",
			s:    "lower-hyphen",
			want: "lower-hyphen",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseformat.FromLowerHyphen(tt.s).ToLowerHyphen(); got != tt.want {
				t.Errorf("lowerHyphen.ToLowerHyphen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerHyphen_ToLowerUnderscore(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal_case",
			s:    "lower-hyphen",
			want: "lower_hyphen",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseformat.FromLowerHyphen(tt.s).ToLowerUnderscore(); got != tt.want {
				t.Errorf("lowerHyphen.ToLowerUnderscore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerHyphen_ToLowerCamel(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal_case",
			s:    "lower-hyphen",
			want: "lowerHyphen",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseformat.FromLowerHyphen(tt.s).ToLowerCamel(); got != tt.want {
				t.Errorf("lowerHyphen.ToLowerCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerHyphen_ToUpperHyphen(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal_case",
			s:    "lower-hyphen",
			want: "LOWER-HYPHEN",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseformat.FromLowerHyphen(tt.s).ToUpperHyphen(); got != tt.want {
				t.Errorf("lowerHyphen.ToUpperHyphen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerHyphen_ToUpperUnderscore(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal_case",
			s:    "lower-hyphen",
			want: "LOWER_HYPHEN",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseformat.FromLowerHyphen(tt.s).ToUpperUnderscore(); got != tt.want {
				t.Errorf("lowerHyphen.ToUpperUnderscore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerHyphen_ToUpperCamel(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal_case",
			s:    "lower-hyphen",
			want: "LowerHyphen",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseformat.FromLowerHyphen(tt.s).ToUpperCamel(); got != tt.want {
				t.Errorf("lowerHyphen.ToUpperCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}
