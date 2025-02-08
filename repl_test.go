package main

import (
	"reflect"
	"testing"
)
func TestCleanInput(t *testing.T) {
	testCases := []struct {
        input    string
        expected []string
    }{
        {"", []string{}},
		{"hello", []string{"hello"}},
        {"hello world", []string{"hello", "world"}},
        {"  hello   world  ", []string{"hello", "world"}},
        {"hello,world", []string{"hello", "world"}},
        {"hello, world", []string{"hello", "world"}},
    }

    for _, tc := range testCases {
        result := cleanInput(tc.input)
        if !reflect.DeepEqual(result, tc.expected) {
            t.Errorf("cleanInput(%q) = %v; want %v", tc.input, result, tc.expected)
        }
    }
}