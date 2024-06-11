package idm

import "testing"

func TestTrim(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{"", []string{""}},
		{"Austin Zhou", []string{"Austin", "Zhou"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Trim(tt.name)
			if len(got) != len(tt.want) {
				t.Errorf("(%v): got %v, want %v", tt.name, len(got), len(tt.want))
			}
		})
	}
}
