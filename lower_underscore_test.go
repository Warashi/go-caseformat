package caseformat_test

import (
	"testing"

	"github.com/Warashi/go-caseformat"
)

func TestLowerUnderscore_ToLowerHyphen(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal_case",
			s:    "lower_underscore",
			want: "lower-underscore",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseformat.FromLowerUnderscore(tt.s).ToLowerHyphen(); got != tt.want {
				t.Errorf("lowerUnderscore.ToLowerHyphen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerUnderscore_ToLowerUnderscore(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal_case",
			s:    "lower_underscore",
			want: "lower_underscore",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseformat.FromLowerUnderscore(tt.s).ToLowerUnderscore(); got != tt.want {
				t.Errorf("lowerUnderscore.ToLowerUnderscore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerUnderscore_ToLowerCamel(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal_case",
			s:    "lower_underscore",
			want: "lowerUnderscore",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseformat.FromLowerUnderscore(tt.s).ToLowerCamel(); got != tt.want {
				t.Errorf("lowerUnderscore.ToLowerCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerUnderscore_ToUpperHyphen(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal_case",
			s:    "lower_underscore",
			want: "LOWER-UNDERSCORE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseformat.FromLowerUnderscore(tt.s).ToUpperHyphen(); got != tt.want {
				t.Errorf("lowerUnderscore.ToUpperHyphen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerUnderscore_ToUpperUnderscore(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal_case",
			s:    "lower_underscore",
			want: "LOWER_UNDERSCORE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseformat.FromLowerUnderscore(tt.s).ToUpperUnderscore(); got != tt.want {
				t.Errorf("lowerUnderscore.ToUpperUnderscore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerUnderscore_ToUpperCamel(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal_case",
			s:    "lower_underscore",
			want: "LowerUnderscore",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseformat.FromLowerUnderscore(tt.s).ToUpperCamel(); got != tt.want {
				t.Errorf("lowerUnderscore.ToUpperCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}
