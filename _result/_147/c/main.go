package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var wtr = bufio.NewWriter(os.Stdout)

type testimony struct {
	x int
	y int
}

func main() {

	defer flush()

	n := ni()
	tt := make([][]testimony, n)
	for i := 0; i < n; i++ {
		a := ni()
		for j := 0; j < a; j++ {
			x, y := ni2()
			tt[i] = append(tt[i], testimony{x, y})
		}
	}
	//  iがパターン
	//  iを検証
	ans := 0
	for i := 0; i < (1 << uint(n)); i++ {
		flg := true
		//  j人目を検証
		for j := 0; j < n; j++ {
			honest := (i>>uint(j))&1 == 1
			if !honest {
				// j人目は嘘つき
				continue
			}
			for _, t := range tt[j] {
				if (i>>uint(t.x-1))&1 != t.y {
					//   矛盾
					flg = false
					break
				}
			}
			if !flg {
				break
			}
		}
		if flg {
			//   このパターンは成功
			count := 0
			for j := 0; j < n; j++ {
				count += (i >> uint(j)) & 1
			}
			if ans < count {
				ans = count
			}
		}
	}
	out(ans)
}

// ==================================================
// init
// ==================================================

const inf = math.MaxInt64
const mod1000000007 = 1000000007
const mod998244353 = 998244353
const mod = mod1000000007

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
	if len(os.Args) > 1 && os.Args[1] == "i" {
		b, e := ioutil.ReadFile("./input")
		if e != nil {
			panic(e)
		}
		sc = bufio.NewScanner(strings.NewReader(strings.Replace(string(b), " ", "\n", -1)))
	}
}

// ==================================================
// io
// ==================================================

// 標準入力をintにしてかえす
// next int
func ni() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
func ni2() (int, int) {
	return ni(), ni()
}
func ni3() (int, int, int) {
	return ni(), ni(), ni()
}

func nis(arg ...int) []int {
	n := arg[0]
	t := 0
	if len(arg) == 2 {
		t = arg[1]
	}

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = ni() - t
	}
	return a
}

func nf() float64 {
	sc.Scan()
	f, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return f
}

// 標準入力をそのままで返す
func ns() string {
	sc.Scan()
	return sc.Text()
}

func out(v ...interface{}) {
	_, e := fmt.Fprintln(wtr, v...)
	if e != nil {
		panic(e)
	}
}

// これで出力する
func flush() {
	e := wtr.Flush()
	if e != nil {
		panic(e)
	}
}

func atoi(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return i
}

func itoa(i int) string {
	return strconv.Itoa(i)
}
func istoas(is []int) []string {
	r := make([]string, len(is))
	for i, v := range is {
		r[i] = itoa(v)
	}
	return r
}

func btoi(b byte) int {
	return atoi(string(b))
}

// ==================================================
// num
// ==================================================

func max(a []int) int {
	x := -int(1e+18)
	for i := 0; i < len(a); i++ {
		if x < a[i] {
			x = a[i]
		}
	}
	return x
}
func min(a []int) int {
	x := int(1e+18)
	for i := 0; i < len(a); i++ {
		if x > a[i] {
			x = a[i]
		}
	}
	return x
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
func sqrt(i int) int {
	return int(math.Sqrt(float64(i)))
}

func larger(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
func smaller(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// sort integer
func sorti(sl []int) {
	sort.Sort(sort.IntSlice(sl))
}

// 大きい順
func sortDescending(arr []int) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
}
func sortSearch(ii []int, t int) int {
	return sort.SearchInts(ii, t)
}

func reversei(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

//   二分探索のtemplate
// for i := 0; i < (1 << uint(n)); i++ {
// 	out(i)
// 	for j := 0; j < n; j++ {
// 		if i>>uint(j)&1 == 1 {
// 			// ビットが立っていた場合の処理

// 		}
// 	}
// }
