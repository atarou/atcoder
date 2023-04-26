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

func main() {

	defer flush()

	n := ni()

	ss := make(map[string]bool)
	for i := 0; i < n; i++ {
		s := ns()
		if ss[s] {

		} else {
			out(i + 1)
			ss[s] = true
		}
	}
}

// ==================================================
// init
// ==================================================

const inf = math.MaxInt64
const mod1000000007 = 1000000007
const mod998244353 = 998244353

// const mod = mod1000000007

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
func nss(arg ...int) []string {
	n := arg[0]

	a := make([]string, n)
	for i := 0; i < n; i++ {
		a[i] = ns()
	}
	return a
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
func mod(x, y int) int {
	m := x % y
	if m < 0 {
		return m + y
	}
	return m
}
func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

// sort integer
func sorti(sl []int) {
	sort.Sort(sort.IntSlice(sl))
}
func sorts(sl []string) {
	sort.Strings(sl)
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

// 順列
func permute(arr []int, l, r int, result *[][]int) {
	if l == r {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		*result = append(*result, tmp)
	} else {
		for i := l; i <= r; i++ {
			arr[l], arr[i] = arr[i], arr[l]
			permute(arr, l+1, r, result)
			arr[l], arr[i] = arr[i], arr[l]
		}
	}
}

func Permutations(arr []int) [][]int {
	var result [][]int
	permute(arr, 0, len(arr)-1, &result)
	return result
}

// 階乗
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// 最大公約数
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

type vertex struct {
	d int
}
type edge struct {
	to int
}
type graph struct {
	v []vertex
	e [][]edge
}

// NewGraph creates a new graph with n vertices.
func NewGraph(n int) *graph {
	g := &graph{
		v: make([]vertex, n),
		e: make([][]edge, n),
	}
	return g
}

// AddEdge adds a edge connects vertex u to v and v to u.
func (g *graph) AppendEdge(from, to int) {
	g.e[from] = append(g.e[from], edge{
		to: to,
	})
}
