package polecat

import (
	"os"
	"testing"
)

func TestCustomNamePoolCycling(t *testing.T) {
	dir, err := os.MkdirTemp("", "pooltest-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	customNames := []string{"ponder", "ridcully", "rincewind"}
	pool := NewNamePoolWithConfig(dir, "testrig", "", customNames, 0)
	pool.Reconcile(nil)

	tests := []struct {
		want string
	}{
		{"ponder"},
		{"ridcully"},
		{"rincewind"},
		// Pool exhausted - cycling begins
		{"ponder-2"},
		{"ridcully-2"},
		{"rincewind-2"},
		// Second cycle
		{"ponder-3"},
	}

	for i, tc := range tests {
		got, err := pool.Allocate()
		if err != nil {
			t.Fatalf("allocation %d: unexpected error: %v", i+1, err)
		}
		if got != tc.want {
			t.Errorf("allocation %d: got %q, want %q", i+1, got, tc.want)
		}
	}
}
