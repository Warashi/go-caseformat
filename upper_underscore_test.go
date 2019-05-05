package caseformat

import "testing"

func TestUpperUnderscore_ToLowerHyphen(t *testing.T) {
	tests := []struct {
		name string
		s    UpperUnderscore
		want string
	}{
		{
			name: "normal_case",
			s:    UpperUnderscore("UPPER_UNDERSCORE"),
			want: "upper-underscore",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToLowerHyphen(); got != tt.want {
				t.Errorf("UpperUnderscore.ToLowerHyphen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperUnderscore_ToLowerUnderscore(t *testing.T) {
	tests := []struct {
		name string
		s    UpperUnderscore
		want string
	}{
		{
			name: "normal_case",
			s:    UpperUnderscore("UPPER_UNDERSCORE"),
			want: "upper_underscore",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToLowerUnderscore(); got != tt.want {
				t.Errorf("UpperUnderscore.ToLowerUnderscore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperUnderscore_ToLowerCamel(t *testing.T) {
	tests := []struct {
		name string
		s    UpperUnderscore
		want string
	}{
		{
			name: "normal_case",
			s:    UpperUnderscore("UPPER_UNDERSCORE"),
			want: "upperUnderscore",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToLowerCamel(); got != tt.want {
				t.Errorf("UpperUnderscore.ToLowerCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperUnderscore_ToUpperHyphen(t *testing.T) {
	tests := []struct {
		name string
		s    UpperUnderscore
		want string
	}{
		{
			name: "normal_case",
			s:    UpperUnderscore("UPPER_UNDERSCORE"),
			want: "UPPER-UNDERSCORE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToUpperHyphen(); got != tt.want {
				t.Errorf("UpperUnderscore.ToUpperHyphen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperUnderscore_ToUpperUnderscore(t *testing.T) {
	tests := []struct {
		name string
		s    UpperUnderscore
		want string
	}{
		{
			name: "normal_case",
			s:    UpperUnderscore("UPPER_UNDERSCORE"),
			want: "UPPER_UNDERSCORE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToUpperUnderscore(); got != tt.want {
				t.Errorf("UpperUnderscore.ToUpperUnderscore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperUnderscore_ToUpperCamel(t *testing.T) {
	tests := []struct {
		name string
		s    UpperUnderscore
		want string
	}{
		{
			name: "normal_case",
			s:    UpperUnderscore("UPPER_UNDERSCORE"),
			want: "UpperUnderscore",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToUpperCamel(); got != tt.want {
				t.Errorf("UpperUnderscore.ToUpperCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}
