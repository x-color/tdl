package cmd

import (
	"testing"
)

func TestCmdShowToCheckTextOverflow(t *testing.T) {
	testCases := []struct {
		num    int
		text   string
		output string
	}{
		{12, "test message", "test message   "},
		{10, "test message", "test messa..."},
		{-1, "test message", "test message"},
	}
	for _, tc := range testCases {
		optCmdShow.length = tc.num
		rowText := tc.text
		expected := tc.output
		actual := textOverflow(rowText)
		if actual != expected {
			msg := "Output text was not correctory."
			t.Fatalf("%s\nExpected: %v\nActual  : %v\n", msg, expected, actual)
		}
	}
}
