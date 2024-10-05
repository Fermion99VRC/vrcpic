package cmd

import (
	"fmt"
	"testing"
	"time"
)

func TestParseFilename(t *testing.T) {
	tests := []struct {
		filename string
		expected time.Time
	}{
		{
			"VRChat_2024-01-01_01-02-30.622_1920x1080.png",
			time.Date(2024, 1, 1, 1, 2, 30, 622000000, time.Local),
		},
		{
			"VRChat_2023-02-04_11-30-23.113_3840x2160.png",
			time.Date(2023, 2, 4, 11, 30, 23, 113000000, time.Local),
		},
		{
			"VRChat_2022-12-31_23-59-00.047_1920x1080.png",
			time.Date(2022, 12, 31, 23, 59, 0, 47000000, time.Local),
		},
	}
	for _, test := range tests {
		t.Run(
			test.filename,
			func(t *testing.T) {
				actual, err := ParseFilename(test.filename)
				if err != nil {
					t.Errorf("cannot parse filename: %v", err)
				}
				if !actual.Equal(test.expected) {
					t.Errorf("ParseFilename(%q) = %q, want %q", test.filename, actual, test.expected)
				}
			},
		)
	}
}

func ExampleParseFilename() {
	datetime, _ := ParseFilename("VRChat_2024-01-01_01-02-30.622_1920x1080.png")
	fmt.Println(datetime)
	// Output:
	// 2024-01-01 01:02:30.622 +0000 UTC
}

func TestGetDirectoryToMove(t *testing.T) {
	tests := []struct {
		name     string
		datetime time.Time
		expected string
	}{
		{
			"2024-01-01 morning",
			time.Date(2024, 1, 1, 1, 2, 0, 0, time.Local),
			"2024-01/0101",
		},
		{
			"2023-02-04 morning",
			time.Date(2023, 2, 4, 11, 30, 0, 0, time.Local),
			"2023-02/0204",
		},
		{
			"2022-12-31 afternoon",
			time.Date(2022, 12, 31, 23, 59, 0, 0, time.Local),
			"2023-01/0101",
		},
	}
	for _, test := range tests {
		t.Run(
			test.name,
			func(t *testing.T) {
				actual := GetDirectoryToMove(test.datetime)
				if actual != test.expected {
					t.Errorf("GetDirectoryToMove(%q) = %q, want %q", test.datetime, actual, test.expected)
				}
			},
		)
	}
}

func ExampleGetDirectoryToMove_morning() {
	datetime := time.Date(2024, 4, 1, 6, 0, 0, 0, time.Local)
	directory := GetDirectoryToMove(datetime)
	fmt.Println(directory)
	// Output:
	// 2024-04/0401
}

func ExampleGetDirectoryToMove_afternoon() {
	datetime := time.Date(2024, 4, 1, 15, 0, 0, 0, time.Local)
	directory := GetDirectoryToMove(datetime)
	fmt.Println(directory)
	// Output:
	// 2024-04/0402
}

func TestCopyFile(t *testing.T) {
	// TODO: Add unitests
}
