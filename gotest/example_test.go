package gotest

import (
	"sync"
	"testing"
)

type GVar struct {
	once *sync.Once
}

var g = &GVar{once: new(sync.Once)}

func loadResource() {
	// init local var
	println("...loading resrouce for test")
}

func setUp() {
	g.once.Do(loadResource)
}

type Test struct {
	a int
	b int
	c int
}

var tests = []Test{
	{1, 2, 3},
	{2, 3, 5},
	{2, 4, 6},
}

func TestAdd(t *testing.T) {
	setUp()
	for i, test := range tests {
		c := Add(test.a, test.b)
		if c != test.c {
			t.Errorf("#%d: Add(%d,%d)=%d; want %d", i, test.a, test.b, c, test.c)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(3<<5, 2<<6)
	}
}

var tests1 = []Test{
	{3, 2, 1},
	{3, 3, 0},
	{0, 4, 0},
}

func TestMod(t *testing.T) {
	setUp()
	for i, test := range tests1 {
		c := Mod(test.a, test.b)
		if c != test.c {
			t.Errorf("#%d: Mod(%d,%d)=%d; want %d", i, test.a, test.b, c, test.c)
		}
	}
}
