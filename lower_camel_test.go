package caseformat

import "testing"

func TestLowerCamel_ToLowerHyphen(t *testing.T) {
	tests := []struct {
		name string
		s    LowerCamel
		want string
	}{
		{
			name: "normal_case",
			s:    LowerCamel("lowerCamel"),
			want: "lower-camel",
		},
		{
			name: "normal_JSON_case",
			s:    LowerCamel("lowerJSONCamel"),
			want: "lower-json-camel",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToLowerHyphen(); got != tt.want {
				t.Errorf("LowerCamel.ToLowerHyphen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerCamel_ToLowerUnderscore(t *testing.T) {
	tests := []struct {
		name string
		s    LowerCamel
		want string
	}{
		{
			name: "normal_case",
			s:    LowerCamel("lowerCamel"),
			want: "lower_camel",
		},
		{
			name: "normal_JSON_case",
			s:    LowerCamel("lowerJSONCamel"),
			want: "lower_json_camel",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToLowerUnderscore(); got != tt.want {
				t.Errorf("LowerCamel.ToLowerUnderscore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerCamel_ToLowerCamel(t *testing.T) {
	tests := []struct {
		name string
		s    LowerCamel
		want string
	}{
		{
			name: "normal_case",
			s:    LowerCamel("lowerCamel"),
			want: "lowerCamel",
		},
		{
			name: "normal_JSON_case",
			s:    LowerCamel("lowerJSONCamel"),
			want: "lowerJSONCamel",
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToLowerCamel(); got != tt.want {
				t.Errorf("LowerCamel.ToLowerCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerCamel_ToUpperHyphen(t *testing.T) {
	tests := []struct {
		name string
		s    LowerCamel
		want string
	}{
		{
			name: "normal_case",
			s:    LowerCamel("lowerCamel"),
			want: "LOWER-CAMEL",
		},
		{
			name: "normal_JSON_case",
			s:    LowerCamel("lowerJSONCamel"),
			want: "LOWER-JSON-CAMEL",
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToUpperHyphen(); got != tt.want {
				t.Errorf("LowerCamel.ToUpperHyphen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerCamel_ToUpperUnderscore(t *testing.T) {
	tests := []struct {
		name string
		s    LowerCamel
		want string
	}{
		{
			name: "normal_case",
			s:    LowerCamel("lowerCamel"),
			want: "LOWER_CAMEL",
		},
		{
			name: "normal_JSON_case",
			s:    LowerCamel("lowerJSONCamel"),
			want: "LOWER_JSON_CAMEL",
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToUpperUnderscore(); got != tt.want {
				t.Errorf("LowerCamel.ToUpperUnderscore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerCamel_ToUpperCamel(t *testing.T) {
	tests := []struct {
		name string
		s    LowerCamel
		want string
	}{
		{
			name: "normal_case",
			s:    LowerCamel("lowerCamel"),
			want: "LowerCamel",
		},
		{
			name: "normal_JSON_case",
			s:    LowerCamel("lowerJSONCamel"),
			want: "LowerJSONCamel",
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToUpperCamel(); got != tt.want {
				t.Errorf("LowerCamel.ToUpperCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}
