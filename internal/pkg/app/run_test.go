package app

import "testing"

func TestMake(t *testing.T) {
	testCases := []struct {
		name          string
		expectedName  string
		isErrExpected bool
	}{
		{
			name:          "success",
			expectedName:  templatesName,
			isErrExpected: false,
		},
	}
	for _, tc := range testCases {
		tc := tc
		name, err := Make().Name()
		if (err != nil) != tc.isErrExpected {
			t.Fatalf("TestMake: expecting an error")
		}

		if want, got := tc.expectedName, name; want != got {
			t.Fatalf("TestMake: want: %v, got: %v", want, got)
		}
	}
}
