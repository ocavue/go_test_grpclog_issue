package util

import "testing"

func TestAdd(t *testing.T) {
	t.Run("1+2", func(t *testing.T) {
		if Add(1, 2) != 3 {
			t.Errorf("1+2!=3")
		}
	})
}
