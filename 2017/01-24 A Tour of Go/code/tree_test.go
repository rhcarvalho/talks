package main

import (
	"testing"

	"golang.org/x/tour/tree"
)

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool { return true }

func TestSame(t *testing.T) {
	tests := []struct {
		trees [2]*tree.Tree
		want  bool
	}{
		{
			trees: [2]*tree.Tree{tree.New(1), tree.New(1)},
			want:  true,
		},
		{
			trees: [2]*tree.Tree{tree.New(1), tree.New(2)},
			want:  false,
		},
	}
	for _, tt := range tests {
		if got := Same(tt.trees[0], tt.trees[1]); got != tt.want {
			t.Errorf("Same(%v, %v) = %v, want %v", tt.trees[0], tt.trees[1], got, tt.want)
		}
	}
}

func main() {
	tests := []testing.InternalTest{{"TestSame", TestSame}}
	matchAll := func(t string, pat string) (bool, error) { return true, nil }
	testing.Main(matchAll, tests, nil, nil)
}
