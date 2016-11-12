package regex

import (
	"fmt"
	"testing"
)

func TestMatch(t *testing.T) {
	type args struct {
		regex string
		test  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"simple match", args{regex: "ab", test: "ab"}, true},
		{"star match 1", args{regex: "a*b", test: "ab"}, true},
		{"star match 2", args{regex: "a*b", test: "aab"}, true},
		{"multi star match", args{regex: "a*b*", test: "aaabb"}, true},
		{"star match empty 1", args{regex: "a*b*", test: "aaa"}, true},
		{"star match empty 2", args{regex: "a*b*", test: ""}, true},
		{"doesn't match", args{regex: "a*b*c", test: "aabb"}, false},
	}
	for _, tt := range tests {
		if got := Match(tt.args.regex, tt.args.test); got != tt.want {
			t.Errorf("%q. Match() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func ExampleMatch() {
	fmt.Println(Match("a*b*", "aaa"))
	fmt.Println(Match("a*b*", "aaabb"))
	fmt.Println(Match("a*b*", ""))
	fmt.Println(Match("aab*", "abb"))
	// Output:
	// true
	// true
	// true
	// false
}
