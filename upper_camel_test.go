package caseformat

import "testing"

func TestUpperCamel_ToLowerHyphen(t *testing.T) {
	tests := []struct {
		name string
		s    UpperCamel
		want string
	}{
		{
			name: "normal_case",
			s:    UpperCamel("UpperCamel"),
			want: "upper-camel",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToLowerHyphen(); got != tt.want {
				t.Errorf("UpperCamel.ToLowerHyphen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperCamel_ToLowerUnderscore(t *testing.T) {
	tests := []struct {
		name string
		s    UpperCamel
		want string
	}{
		{
			name: "normal_case",
			s:    UpperCamel("UpperCamel"),
			want: "upper_camel",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToLowerUnderscore(); got != tt.want {
				t.Errorf("UpperCamel.ToLowerUnderscore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperCamel_ToLowerCamel(t *testing.T) {
	tests := []struct {
		name string
		s    UpperCamel
		want string
	}{
		{
			name: "normal_case",
			s:    UpperCamel("UpperCamel"),
			want: "upperCamel",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToLowerCamel(); got != tt.want {
				t.Errorf("UpperCamel.ToLowerCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperCamel_ToUpperHyphen(t *testing.T) {
	tests := []struct {
		name string
		s    UpperCamel
		want string
	}{
		{
			name: "normal_case",
			s:    UpperCamel("UpperCamel"),
			want: "UPPER-CAMEL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToUpperHyphen(); got != tt.want {
				t.Errorf("UpperCamel.ToUpperHyphen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperCamel_ToUpperUnderscore(t *testing.T) {
	tests := []struct {
		name string
		s    UpperCamel
		want string
	}{
		{
			name: "normal_case",
			s:    UpperCamel("UpperCamel"),
			want: "UPPER_CAMEL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToUpperUnderscore(); got != tt.want {
				t.Errorf("UpperCamel.ToUpperUnderscore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperCamel_ToUpperCamel(t *testing.T) {
	tests := []struct {
		name string
		s    UpperCamel
		want string
	}{
		{
			name: "normal_case",
			s:    UpperCamel("UpperCamel"),
			want: "UpperCamel",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToUpperCamel(); got != tt.want {
				t.Errorf("UpperCamel.ToUpperCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}
