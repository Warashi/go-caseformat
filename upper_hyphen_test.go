package caseformat

import (
	"testing"
)

func TestUpperHyphen_ToLowerHyphen(t *testing.T) {
	tests := []struct {
		name string
		s    UpperHyphen
		want string
	}{
		{
			name: "正常系",
			s:    UpperHyphen("UPPER-HYPHEN"),
			want: "upper-hyphen",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToLowerHyphen(); got != tt.want {
				t.Errorf("UpperHyphen.ToLowerHyphen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperHyphen_ToLowerUnderscore(t *testing.T) {
	tests := []struct {
		name string
		s    UpperHyphen
		want string
	}{
		{
			name: "正常系",
			s:    UpperHyphen("UPPER-HYPHEN"),
			want: "upper_hyphen",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToLowerUnderscore(); got != tt.want {
				t.Errorf("UpperHyphen.ToLowerUnderscore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperHyphen_ToLowerCamel(t *testing.T) {
	tests := []struct {
		name string
		s    UpperHyphen
		want string
	}{
		{
			name: "正常系",
			s:    UpperHyphen("UPPER-HYPHEN"),
			want: "upperHyphen",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToLowerCamel(); got != tt.want {
				t.Errorf("UpperHyphen.ToLowerCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperHyphen_ToUpperHyphen(t *testing.T) {
	tests := []struct {
		name string
		s    UpperHyphen
		want string
	}{
		{
			name: "正常系",
			s:    UpperHyphen("UPPER-HYPHEN"),
			want: "UPPER-HYPHEN",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToUpperHyphen(); got != tt.want {
				t.Errorf("UpperHyphen.ToUpperHyphen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperHyphen_ToUpperUnderscore(t *testing.T) {
	tests := []struct {
		name string
		s    UpperHyphen
		want string
	}{
		{
			name: "正常系",
			s:    UpperHyphen("UPPER-HYPHEN"),
			want: "UPPER_HYPHEN",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToUpperUnderscore(); got != tt.want {
				t.Errorf("UpperHyphen.ToUpperUnderscore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperHyphen_ToUpperCamel(t *testing.T) {
	tests := []struct {
		name string
		s    UpperHyphen
		want string
	}{
		{
			name: "正常系",
			s:    UpperHyphen("UPPER-HYPHEN"),
			want: "UpperHyphen",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToUpperCamel(); got != tt.want {
				t.Errorf("UpperHyphen.ToUpperCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}
