package util

import (
	"testing"
)

func TestRandString(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "TestRandString",
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomString(tt.want)
			t.Logf("RandString() = %v", got)
			if len(got) != tt.want {
				t.Errorf("RandString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateSalt(t *testing.T) {
	t.Log(GenerateSalt())
	t.Log(GenerateSalt())
	t.Log(GenerateSalt())
	t.Log(GenerateSalt())
	t.Log(GenerateSalt())
}
