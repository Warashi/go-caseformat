package caseformat

import "testing"

func TestLowerUnderscore_ToLowerHyphen(t *testing.T) {
	tests := []struct {
		name string
		s    LowerUnderscore
		want string
	}{
		{
			name: "normal_case",
			s:    LowerUnderscore("lower_underscore"),
			want: "lower-underscore",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToLowerHyphen(); got != tt.want {
				t.Errorf("LowerUnderscore.ToLowerHyphen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerUnderscore_ToLowerUnderscore(t *testing.T) {
	tests := []struct {
		name string
		s    LowerUnderscore
		want string
	}{
		{
			name: "normal_case",
			s:    LowerUnderscore("lower_underscore"),
			want: "lower_underscore",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToLowerUnderscore(); got != tt.want {
				t.Errorf("LowerUnderscore.ToLowerUnderscore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerUnderscore_ToLowerCamel(t *testing.T) {
	tests := []struct {
		name string
		s    LowerUnderscore
		want string
	}{
		{
			name: "normal_case",
			s:    LowerUnderscore("lower_underscore"),
			want: "lowerUnderscore",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToLowerCamel(); got != tt.want {
				t.Errorf("LowerUnderscore.ToLowerCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerUnderscore_ToUpperHyphen(t *testing.T) {
	tests := []struct {
		name string
		s    LowerUnderscore
		want string
	}{
		{
			name: "normal_case",
			s:    LowerUnderscore("lower_underscore"),
			want: "LOWER-UNDERSCORE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToUpperHyphen(); got != tt.want {
				t.Errorf("LowerUnderscore.ToUpperHyphen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerUnderscore_ToUpperUnderscore(t *testing.T) {
	tests := []struct {
		name string
		s    LowerUnderscore
		want string
	}{
		{
			name: "normal_case",
			s:    LowerUnderscore("lower_underscore"),
			want: "LOWER_UNDERSCORE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToUpperUnderscore(); got != tt.want {
				t.Errorf("LowerUnderscore.ToUpperUnderscore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerUnderscore_ToUpperCamel(t *testing.T) {
	tests := []struct {
		name string
		s    LowerUnderscore
		want string
	}{
		{
			name: "normal_case",
			s:    LowerUnderscore("lower_underscore"),
			want: "LowerUnderscore",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToUpperCamel(); got != tt.want {
				t.Errorf("LowerUnderscore.ToUpperCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}
