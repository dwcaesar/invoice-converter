package assert

import "testing"

func Equal(t *testing.T, got any, expected any) {
	t.Helper()

	if expected != got {
		t.Errorf(`assert.Equal
			got: %v
			expected: %v`, got, expected)
	}
}
