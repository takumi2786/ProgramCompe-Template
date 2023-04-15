package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// *************** practice ******************
func TestSlice(t *testing.T) {
	{
		fmt.Println("キャパシティが足りる時のappend: スライスが指し示す配列が同じ")
		var list []int = make([]int, 1, 10)
		list[0] = 0
		fmt.Printf("&list: %p\n", &list)
		fmt.Printf("&list[0]: %p\n", &list[0])

		list = append(list, 1)
		fmt.Printf("&list: %p\n", &list)
		fmt.Printf("&list[0]: %p\n", &list[0])
	}
	{
		fmt.Println("キャパシティが足りない時のappend: スライスが指し示す配列が変わる")
		var list []int = make([]int, 1)
		list[0] = 0
		fmt.Printf("&list: %p\n", &list)
		fmt.Printf("&list[0]: %p\n", &list[0])

		list = append(list, 1, 2, 3, 4)
		fmt.Printf("&list: %p\n", &list)
		fmt.Printf("&list[0]: %p\n", &list[0])
	}
	{
		fmt.Println("関数内では、スライスは別物、指す先が同じ")
		var list []int = []int{1, 2, 3, 4, 5, 6, 7}
		f := func(l []int) {
			l[0] = 10
			fmt.Printf("&l: %p\n", &l)
			fmt.Printf("&l[0]: %p\n", &l[0])
		}
		fmt.Printf("&list: %p\n", &list)
		fmt.Printf("&list[0]: %p\n", &list[0])
		f(list)
		fmt.Printf("&list: %p\n", &list)
		fmt.Printf("&list[0]: %p\n", &list[0])
	}
	{
		fmt.Println("キャパシティが足りる時のappend: 指す先は変わらない。呼び出し元には影響がない。")
		var list []int = make([]int, 1, 10)
		f := func(l []int) {
			l = append(l, 1, 2, 3)
			fmt.Printf("&l: %p\n", &l)
			fmt.Printf("&l[0]: %p\n", &l[0])
		}
		f(list)
		fmt.Printf("&list: %p\n", &list)
		fmt.Printf("&list[0]: %p\n", &list[0])
		fmt.Println(list)
	}
	{
		fmt.Println("キャパシティが足りない時のappend: 指す先が変更されるが、呼び出し元には影響がない。")
		var list []int = []int{1}
		f := func(l []int) {
			fmt.Printf("&l: %p\n", &l)
			fmt.Printf("&l[0]: %p\n", &l[0])
			l = append(l, 1, 2, 3)
			fmt.Printf("&l: %p\n", &l)
			fmt.Printf("&l[0]: %p\n", &l[0])
		}
		fmt.Printf("&list: %p\n", &list)
		fmt.Printf("&list[0]: %p\n", &list[0])
		f(list)
		fmt.Printf("&list: %p\n", &list)
		fmt.Printf("&list[0]: %p\n", &list[0])
		fmt.Println(list)
	}
	{
		fmt.Println("グローバル変数の場合、呼び出し元に影響する")
		var list []int = []int{1}
		f := func() {
			fmt.Printf("&l: %p\n", &list)
			fmt.Printf("&l[0]: %p\n", &list[0])
			list = append(list, 1, 2, 3)
			fmt.Printf("&l: %p\n", &list)
			fmt.Printf("&l[0]: %p\n", &list[0])
		}
		// fmt.Printf("&list: %p\n", &list)
		// fmt.Printf("&list[0]: %p\n", &list[0])
		f()
		// fmt.Printf("&list: %p\n", &list)
		// fmt.Printf("&list[0]: %p\n", &list[0])
		fmt.Println(list)
	}
	{
		fmt.Println("ポインタ渡しの場合、呼び出し元に影響する")
		var list []int = []int{1}
		f := func(l *[]int) {
			fmt.Printf("&l: %p\n", l)
			fmt.Printf("&l[0]: %p\n", *l)
			*l = append(*l, 1, 2, 3)
			fmt.Printf("&l: %p\n", l)
			fmt.Printf("&l[0]: %p\n", *l)
		}
		// fmt.Printf("&list: %p\n", &list)
		// fmt.Printf("&list[0]: %p\n", &list[0])
		f(&list)
		// fmt.Printf("&list: %p\n", &list)
		// fmt.Printf("&list[0]: %p\n", &list[0])
		fmt.Println(list)
	}
	{
		fmt.Println("Mapの場合1: グローバル変数の場合、呼び出し元に影響する")
		var Map [][]int = make([][]int, 1)
		var first []int = []int{1, 2, 3}
		Map[0] = first
		f := func() {
			var second []int = []int{1, 2, 3}
			fmt.Printf("&Map: %p\n", &Map)
			fmt.Printf("&Map[0]: %p\n", &Map[0])
			fmt.Printf("&Map[0][0]: %p\n", &Map[0][0])
			Map = append(Map, second)
			fmt.Printf("&Map: %p\n", &Map)
			fmt.Printf("&Map[0]: %p\n", &Map[0])
			fmt.Printf("&Map[0][0]: %p\n", &Map[0][0])
		}
		// fmt.Printf("&list: %p\n", &list)
		// fmt.Printf("&list[0]: %p\n", &list[0])
		f()
		// fmt.Printf("&list: %p\n", &list)
		// fmt.Printf("&list[0]: %p\n", &list[0])
		fmt.Println(Map)
	}
	{
		fmt.Println("Mapの場合2: 引数として渡した場合、呼び出し元に影響しない")
		var Map [][]int = make([][]int, 1)
		var first []int = []int{1, 2, 3}
		Map[0] = first
		f := func(Map [][]int) {
			var second []int = []int{1, 2, 3}
			fmt.Printf("&Map: %p\n", &Map)
			fmt.Printf("&Map[0]: %p\n", &Map[0])
			fmt.Printf("&Map[0][0]: %p\n", &Map[0][0])
			Map = append(Map, second)
			fmt.Printf("&Map: %p\n", &Map)
			fmt.Printf("&Map[0]: %p\n", &Map[0])
			fmt.Printf("&Map[0][0]: %p\n", &Map[0][0])
		}
		// fmt.Printf("&list: %p\n", &list)
		// fmt.Printf("&list[0]: %p\n", &list[0])
		f(Map)
		// fmt.Printf("&list: %p\n", &list)
		// fmt.Printf("&list[0]: %p\n", &list[0])
		fmt.Println(Map)
	}
}

// *************** testing ******************

func TestSortAndReturnIdxs_1(t *testing.T) {
	// noChange=trueの場合、元の値を保持する
	slice := []int{3, 1, 4}
	sortedIdx := sortAndReturnIdxs(slice, true)
	assert.Equal(t, []int{3, 1, 4}, slice)
	assert.Equal(t, []int{1, 0, 2}, sortedIdx)
}

func TestSortAndReturnIdxs_2(t *testing.T) {
	// noChange=falseの場合、元の値を変更する
	slice := []int{3, 1, 4}
	sortedIdx := sortAndReturnIdxs(slice, false)
	assert.Equal(t, []int{1, 3, 4}, slice)
	assert.Equal(t, []int{1, 0, 2}, sortedIdx)
}

func TestBinarySearch_1(t *testing.T) {
	var input []int = []int{1, 2, 3}
	ans := binarySearch(input, 4)
	assert.Equal(t, 3, input[ans])
}

func TestBinarySearch_2(t *testing.T) {
	var input []int = []int{1, 2, 3, 10, 98}
	ans := binarySearch(input, 5)
	assert.Equal(t, 3, input[ans])
}

func TestBinarySearch_3(t *testing.T) {
	var input []int = []int{1, 2, 3, 10, 98}
	ans := binarySearch(input, 0)
	assert.Equal(t, 1, input[ans])
}

func TestBinarySearch_4(t *testing.T) {
	var input []int = []int{1, 2, 3, 10, 98}
	ans := binarySearch(input, -1)
	assert.Equal(t, 1, input[ans])
}

func TestBinarySearch_5(t *testing.T) {
	var input []int = []int{10, 9, 5, 0}
	ans := binarySearch(input, 4)
	assert.Equal(t, 5, input[ans])
}

func TestContains(t *testing.T) {
	var r []int = []int{1, 2, 3}
	if contains(r, 1) != true {
		t.Errorf("has_bit() = %v, want %v", contains(r, 1), true)
	}
}

func TestHasBit(t *testing.T) {
	if hasBit(0b1001, 0) != true {
		t.Errorf("has_bit() = %v, want %v", hasBit(0b1001, 0), true)
	}
}

func TestMinInt(t *testing.T) {
	if minInt(1, 2) != 1 {
		t.Errorf("add() = %v, want %v", minInt(1, 2), 3)
	}
}

// go test -timeout 30s -run ^TestPopInt$  -v
func TestPopInt(t *testing.T) {
	{ // index = first
		// arrange
		var r []int = []int{0, 1, 2, 3, 4, 5}

		// act
		resp_r, resp_v, error := popInt(r, 0)

		// assert
		var ans_r []int = []int{1, 2, 3, 4, 5}
		assert.Equal(t, error, nil)
		assert.Equal(t, 0, resp_v)
		assert.Equal(t, ans_r, resp_r)
	}
	{ // index = 3
		// arrange
		var r []int = []int{0, 1, 2, 3, 4, 5}

		// act
		resp_r, resp_v, error := popInt(r, 3)

		// assert
		var ans_r []int = []int{0, 1, 2, 4, 5}
		assert.Equal(t, error, nil)
		assert.Equal(t, 3, resp_v)
		assert.Equal(t, ans_r, resp_r)
	}
	{ // index = last
		// arrange
		var r []int = []int{0, 1, 2, 3, 4, 5}

		// act
		index := 5
		resp_r, resp_v, error := popInt(r, index)

		// assert
		var ans_r []int = []int{0, 1, 2, 3, 4}
		assert.Equal(t, nil, error)
		assert.Equal(t, 5, resp_v)
		assert.Equal(t, ans_r, resp_r)
	}
	{ // index = out of range: error
		// arrange
		var r []int = []int{0, 1, 2, 3, 4, 5}

		// act
		index := 10
		_, _, error := popInt(r, index)

		// assert
		assert.Equal(t, fmt.Errorf("index is out of range"), error)
	}
	{ // index = out of range: panic
		// arrange
		var r []int = []int{0, 1, 2, 3, 4, 5}

		// act & assert
		index := 10
		assert.Panics(t, func() { popInt_(r, index) }, "param error was not handled")
		assert.PanicsWithError(t, "index is out of range", func() { popInt_(r, index) })
	}
}

func TestPopInt2(t *testing.T) {
	// index: first
	{
		// arrange
		var input_r [][]int
		{
			var r0 []int = []int{1, 1, 1, 1}
			input_r = append(input_r, r0)

			var r1 []int = []int{1, 2, 3, 4}
			input_r = append(input_r, r1)

			var r2 []int = []int{5, 6, 7, 8}
			input_r = append(input_r, r2)

			var r3 []int = []int{9, 10, 11, 12}
			input_r = append(input_r, r3)
		}

		// act
		resp_r, resp_v := popInt2_(input_r, 1)

		// assert
		var ans_v []int = []int{1, 2, 3, 4}
		var ans_r [][]int
		{
			var r0 []int = []int{1, 1, 1, 1}
			ans_r = append(ans_r, r0)

			var r2 []int = []int{5, 6, 7, 8}
			ans_r = append(ans_r, r2)

			var r3 []int = []int{9, 10, 11, 12}
			ans_r = append(ans_r, r3)
		}

		assert.Equal(t, ans_r, resp_r)
		assert.Equal(t, ans_v, resp_v)
	}
	{ // empty
		// arrange
		var input_r [][]int

		// act
		resp_r, resp_v := popInt2_(input_r, 1)
		assert.Nil(t, resp_v)
		var ans_r [][]int
		assert.Equal(t, resp_r, ans_r)
	}
}

func TestListCombination(t *testing.T) {
	var list []int = []int{1, 2, 3}
	selectionMap := ListCombination(list, 2)

	assert.Contains(t, selectionMap, []int{1, 2})
	assert.Contains(t, selectionMap, []int{1, 3})
	assert.Contains(t, selectionMap, []int{2, 3})
}
