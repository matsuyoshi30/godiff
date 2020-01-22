package godiff

import (
	"testing"
)

func TestDiff_LevenshteinDistance(t *testing.T) {
	testcases := []struct {
		inputstr1 string
		inputstr2 string
		expected  int
	}{
		{"", "", 0},
		{"abc", "abc", 0},
		{"kitten", "sitting", 5},
	}

	for _, tt := range testcases {
		d, err := NewDiff(tt.inputstr1, tt.inputstr2)
		if err != nil {
			t.Errorf("unexpected error %v\n", err)
		}
		if tt.expected != d.LevenshteinDistance() {
			t.Errorf("expected %v, but got %v\n", tt.expected, d.LevenshteinDistance())
		}
	}
}

func TestDiff_LongCommonSubSeq(t *testing.T) {
	testcases := []struct {
		inputstr1 string
		inputstr2 string
		expected  string
	}{
		{"", "", ""},
		{"abc", "abc", "abc"},
		{"kitten", "sitting", "ittn"},
	}

	for _, tt := range testcases {
		d, err := NewDiff(tt.inputstr1, tt.inputstr2)
		if err != nil {
			t.Errorf("unexpected error %v\n", err)
		}
		if tt.expected != d.LongCommonSubSeq() {
			t.Errorf("expected %v, but got %v\n", tt.expected, d.LongCommonSubSeq())
		}
	}
}

func TestDiff_ShowDiff(t *testing.T) {
	testcases := []struct {
		inputstr1 string
		inputstr2 string
		expected  []string
	}{
		{"ATGATCGGCAT", "CAATGTGAATC", []string{"ATGATCGGCAT_", "_CAATGTGAATC", "*++--++-+--*"}},
	}

	for _, tt := range testcases {
		d, err := NewDiff(tt.inputstr1, tt.inputstr2)
		if err != nil {
			t.Errorf("unexpected error %v\n", err)
		}
		diffs := d.ShowDiff()
		for idx, dd := range diffs {
			if tt.expected[idx] != dd {
				t.Errorf("expected %v, but got %v\n", tt.expected[idx], dd)
			}
		}
	}
}

func TestDiff_ShowFileDiff(t *testing.T) {
	testcases := []struct {
		inputfile1 string
		inputfile2 string
		expected   []string
	}{
		{"testdata/test1", "testdata/test2", []string{"= abcde", "+ a", "+ fffff"}},
	}

	for _, tt := range testcases {
		d, err := NewDiff(tt.inputfile1, tt.inputfile2)
		if err != nil {
			t.Errorf("unexpected error %v\n", err)
		}
		diffs, err := d.ShowFileDiff()
		if err != nil {
			t.Errorf("unexpected error %v\n", err)
		}
		for idx, dd := range diffs {
			if tt.expected[idx] != dd {
				t.Errorf("expected %v, but got %v\n", tt.expected[idx], dd)
			}
		}
	}
}
