package caseformat

import (
	"testing"
)

func TestLowerHyphen_ToLowerHyphen(t *testing.T) {
	tests := []struct {
		name string
		s    LowerHyphen
		want string
	}{
		{
			name: "正常系",
			s:    LowerHyphen("lower-hyphen"),
			want: "lower-hyphen",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToLowerHyphen(); got != tt.want {
				t.Errorf("LowerHyphen.ToLowerHyphen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerHyphen_ToLowerUnderscore(t *testing.T) {
	tests := []struct {
		name string
		s    LowerHyphen
		want string
	}{
		{
			name: "正常系",
			s:    LowerHyphen("lower-hyphen"),
			want: "lower_hyphen",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToLowerUnderscore(); got != tt.want {
				t.Errorf("LowerHyphen.ToLowerUnderscore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerHyphen_ToLowerCamel(t *testing.T) {
	tests := []struct {
		name string
		s    LowerHyphen
		want string
	}{
		{
			name: "正常系",
			s:    LowerHyphen("lower-hyphen"),
			want: "lowerHyphen",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToLowerCamel(); got != tt.want {
				t.Errorf("LowerHyphen.ToLowerCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerHyphen_ToUpperHyphen(t *testing.T) {
	tests := []struct {
		name string
		s    LowerHyphen
		want string
	}{
		{
			name: "正常系",
			s:    LowerHyphen("lower-hyphen"),
			want: "LOWER-HYPHEN",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToUpperHyphen(); got != tt.want {
				t.Errorf("LowerHyphen.ToUpperHyphen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerHyphen_ToUpperUnderscore(t *testing.T) {
	tests := []struct {
		name string
		s    LowerHyphen
		want string
	}{
		{
			name: "正常系",
			s:    LowerHyphen("lower-hyphen"),
			want: "LOWER_HYPHEN",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToUpperUnderscore(); got != tt.want {
				t.Errorf("LowerHyphen.ToUpperUnderscore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerHyphen_ToUpperCamel(t *testing.T) {
	tests := []struct {
		name string
		s    LowerHyphen
		want string
	}{
		{
			name: "正常系",
			s:    LowerHyphen("lower-hyphen"),
			want: "LowerHyphen",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToUpperCamel(); got != tt.want {
				t.Errorf("LowerHyphen.ToUpperCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}
