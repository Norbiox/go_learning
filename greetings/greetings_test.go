package greetings

import (
	"testing"
)

func TestGreet(t *testing.T) {
	tests := []struct {
		desc string // What we are testing
		name string // The name we will pass
		want string // What we expect to be returned
		expectErr bool // Do we expect an error?
	}{
		{
			desc: "Error: name is an empty string",
			expectErr: true,
			// name and expect are "", the zero value for string
		},
		{
			desc: "Success",
			name: "John",
			want: "Hello, John",
			// expectErr is set to the zero value, false
		},
	}

	// Executes each test
	for _, test := range tests {
		got, err := Greet(test.name)
		switch {
		case err == nil && test.expectErr:
			t.Errorf("Greet(%q) got no error but expected one", test.name)
			continue
		case err != nil && !test.expectErr:
			t.Errorf("Greet(%q) got an error: %v", test.name, err)
			continue
		case got != test.want:
			t.Errorf("Greet(%q) = %q, want %q", test.name, got, test.want)
		}
	}
}
