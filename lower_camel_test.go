package caseformat_test

import (
	"testing"

	"github.com/Warashi/go-caseformat"
)

func TestLowerCamel_ToLowerHyphen(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal_case",
			s:    "lowerCamel",
			want: "lower-camel",
		},
		{
			name: "normal_JSON_case",
			s:    "lowerJSONCamel",
			want: "lower-json-camel",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseformat.FromLowerCamel(tt.s).ToLowerHyphen(); got != tt.want {
				t.Errorf("lowerCamel.ToLowerHyphen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerCamel_ToLowerUnderscore(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal_case",
			s:    "lowerCamel",
			want: "lower_camel",
		},
		{
			name: "normal_JSON_case",
			s:    "lowerJSONCamel",
			want: "lower_json_camel",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseformat.FromLowerCamel(tt.s).ToLowerUnderscore(); got != tt.want {
				t.Errorf("lowerCamel.ToLowerUnderscore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerCamel_ToLowerCamel(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal_case",
			s:    "lowerCamel",
			want: "lowerCamel",
		},
		{
			name: "normal_JSON_case",
			s:    "lowerJSONCamel",
			want: "lowerJSONCamel",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseformat.FromLowerCamel(tt.s).ToLowerCamel(); got != tt.want {
				t.Errorf("lowerCamel.ToLowerCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerCamel_ToUpperHyphen(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal_case",
			s:    "lowerCamel",
			want: "LOWER-CAMEL",
		},
		{
			name: "normal_JSON_case",
			s:    "lowerJSONCamel",
			want: "LOWER-JSON-CAMEL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseformat.FromLowerCamel(tt.s).ToUpperHyphen(); got != tt.want {
				t.Errorf("lowerCamel.ToUpperHyphen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerCamel_ToUpperUnderscore(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal_case",
			s:    "lowerCamel",
			want: "LOWER_CAMEL",
		},
		{
			name: "normal_JSON_case",
			s:    "lowerJSONCamel",
			want: "LOWER_JSON_CAMEL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseformat.FromLowerCamel(tt.s).ToUpperUnderscore(); got != tt.want {
				t.Errorf("lowerCamel.ToUpperUnderscore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerCamel_ToUpperCamel(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "normal_case",
			s:    "lowerCamel",
			want: "LowerCamel",
		},
		{
			name: "normal_JSON_case",
			s:    "lowerJSONCamel",
			want: "LowerJSONCamel",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := caseformat.FromLowerCamel(tt.s).ToUpperCamel(); got != tt.want {
				t.Errorf("lowerCamel.ToUpperCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}
