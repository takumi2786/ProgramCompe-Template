// https://atcoder.jp/contests/typical-algorithm/tasks/typical_algorithm_c

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

const DEBUG bool = true
const INPUT_PATH = "./input.txt"
const OUTPUT_PATH = "./output.txt"

func main() {
	// initialize
	in := getReader()
	out := getWriter()
	defer out.Flush()

	// 入力値を取得
	var N, M int
	fmt.Fscan(in, &N, &M)

}

func binarySearch(input []int, val int) int {
	// 入力した配列の中の最も近い値を取得する
	reverse := false
	if input[0] > input[len(input)-1] {
		reverse = true
	}
	left := 0
	right := len(input) - 1
	for {
		if abs(right-left) <= 1 {
			if abs(input[left]-val) < abs(input[right]-val) {
				return left
			} else {
				return right
			}
		}
		mid := left + (right-left)/2
		if input[mid] == val {
			return mid
		} else if !reverse {
			if val < input[mid] {
				right = mid
			} else {
				left = mid
			}
		} else {
			if val < input[mid] {
				left = mid
			} else {
				right = mid
			}
		}
	}
}

// util 関数
func getReader() *bufio.Reader {
	if DEBUG {
		f, e := os.Open(INPUT_PATH)
		if e != nil {
			log.Fatal(e)
		}
		return bufio.NewReader(f)
	} else {
		return bufio.NewReader(os.Stdin)
	}
}

func getWriter() *bufio.Writer {
	if DEBUG {
		f, e := os.Create(OUTPUT_PATH)
		if e != nil {
			log.Fatal(e)
		}
		return bufio.NewWriter(f)
	} else {
		return bufio.NewWriter(os.Stdout)
	}
}

// sortAndReturnIdxsで利用するsort.Interfaceの実装
// https://cs.opensource.google/go/go/+/go1.20.3:src/sort/sort.go;l=14
type sortable struct {
	idxs []int
	vals []int
}

func (s sortable) Len() int           { return len(s.vals) }
func (s sortable) Less(i, j int) bool { return s.vals[i] < s.vals[j] }
func (s sortable) Swap(i, j int) {
	// 値と同時にインデックスをswapする
	s.vals[i], s.vals[j] = s.vals[j], s.vals[i]
	s.idxs[i], s.idxs[j] = s.idxs[j], s.idxs[i]
}

// ソートを実行し、結果のインデックスを取得する
// noChange=trueの場合、valsを書き換える
// ref: https://stackoverflow.com/questions/31141202/get-the-indices-of-the-array-after-sorting-in-golang
func sortAndReturnIdxs(vals []int, noChange bool) []int {
	vals_ := make([]int, len(vals))
	if noChange {
		// 値コピー
		copy(vals_, vals)
	} else {
		// 参照渡し
		vals_ = vals
	}

	var idxs = make([]int, len(vals_))
	for i := 0; i < len(vals_); i++ {
		idxs[i] = i
	}
	sort.Sort(sortable{idxs, vals_})
	return idxs
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

func contains(s []int, e int) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

// ビット列nの位置pのビットが立っているかどうか
func hasBit(n int, p int) bool {
	return n&(1<<p) > 0
}

// 2つの数値のうち、小さい方を返す
func minInt(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}

// 配列の指定されたインデックスの要素を取得・削除する
// 失敗するとerror
func popInt(r []int, index int) ([]int, int, error) {
	if index < 0 || index > len(r)-1 {
		return nil, 0, fmt.Errorf("index is out of range")
	}
	// 指定した位置の要素
	resp_v := r[index]

	// 指定した位置の要素を削除したスライスを作成
	resp_r := make([]int, len(r))
	copy(resp_r, r) // 元のスライスを書き換えないようにコピーする

	resp_r = append(r[0:index], r[index+1:]...) // ...を利用すると、複数の値をメソッドに渡せる
	return resp_r, resp_v, nil
}

// 配列の指定されたインデックスの要素を取得・削除する
// 失敗するとpanic
func popInt_(r []int, index int) ([]int, int) {
	resp_r, resp_v, error := popInt(r, index)

	if error != nil {
		panic(error)
	}
	return resp_r, resp_v
}

// 配列の指定されたインデックスの要素を取得・削除する: 二次元配列ver
// 失敗するとerror
func popInt2(r [][]int, index int) ([][]int, []int, error) {
	if len(r) == 0 {
		return r, nil, nil
	}

	if index < 0 || index > len(r)-1 {
		return nil, nil, fmt.Errorf("index is out of range")
	}
	// 指定した位置の要素
	resp_v := r[index]

	// 指定した位置の要素を削除したスライスを作成
	resp_r := make([][]int, len(r))
	copy(resp_r, r) // 元のスライスを書き換えないようにコピーする

	resp_r = append(r[0:index], r[index+1:]...) // ...を利用すると、複数の値をメソッドに渡せる
	return resp_r, resp_v, nil
}

// 配列の指定されたインデックスの要素を取得・削除する: 二次元配列ver
// 失敗するとpanic
func popInt2_(r [][]int, index int) ([][]int, []int) {
	resp_r, resp_v, error := popInt2(r, index)

	if error != nil {
		panic(error)
	}
	return resp_r, resp_v
}

// 順列計算
func Permutation(n int, k int) int {
	v := 1
	if 0 < k && k <= n {
		for i := 0; i < k; i++ {
			v *= (n - i)
		}
	} else if k > n {
		v = 0
	}
	return v
}

// 組み合わせ
func Combination(n int, k int) int {
	return Permutation(n, k) / Permutation(k, k-1)
}

// rのうち、n個を選ぶ組み合わせのパターンを列挙する
func ListCombination(r []int, n int) [][]int {
	var selectionMap [][]int
	// ある要素を選択するかどうかを再起で計算
	// s: 選択有無を示す配列
	// i: index
	// cnt: 選択した個数
	var choice func(s []int, i int, cnt int)
	choice = func(s []int, i int, cnt int) {
		if cnt == n {
			// 選択結果を格納して終了
			// NOTE: スライスを直接渡すと、s内のポインタがコピーされるため、呼び出し元に影響が出る
			// https://zenn.dev/_kazuya/articles/af582f7652271b90cd11#%E3%81%A7%E3%81%AF%E3%80%81%E4%BB%AE%E3%81%AB%E3%81%93%E3%81%86%E3%81%84%E3%81%86%E5%95%8F%E9%A1%8C%E3%82%92%E8%80%83%E3%81%88%E3%81%A6%E3%81%BF%E3%81%BE%E3%81%97%E3%82%87%E3%81%86
			var numbers []int
			for ind, v := range s {
				if v == 1 {
					numbers = append(numbers, r[ind])
				}
			}
			selectionMap = append(selectionMap, numbers)
			return
		}
		if i == len(s) {
			return
		}
		// 選択する
		s[i] = 1
		choice(s, i+1, cnt+1)

		// 選択しない
		s[i] = 0
		choice(s, i+1, cnt)
	}

	selection := make([]int, len(r))
	choice(selection, 0, 0)
	return selectionMap
}

// int形式のスライスをヒープソートする
// https://www.simpletraveler.jp/2021/10/27/python-sort-algorythm-heap/
// func HeapSort(array []int) {
// 	upheap := func(array []int, n int) {
// 		for n != 0 {
// 			parent := int((n - 1) / 2)
// 			if array[n] > array[parent] {

// 			}
// 			n = parent
// 		}
// 	}
// }
