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
	k := ni()
	cnt := 0
	for a := 1; a*a*a <= k; a++ {
		if k%a == 0 {
			for b := a; b*b <= k/a; b++ {
				if (k/a)%b == 0 {
					cnt++
				}
			}
		}
	}
	out(cnt)
}

// ==================================================
// init
// ==================================================

const inf = math.MaxInt64
const mod1000000007 = 1000000007
const mod998244353 = 998244353
const mod = mod998244353

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
func MinIn2Values(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
func MaxIn2Values(a int, b int) int {
	if a < b {
		return b
	}
	return a
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

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
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

// 順列
// aaに並べ替えたいやつ
// 例えば[0,1,2,3,4]
// これを
// 　==========================
//
//	for nextPermutation(pp) {
//		out(pp)
//	}
//
// ===========================
// みたいに使うと全てのppが見れる
func nextPermutation(aa []int) bool {
	n := len(aa)
	l := n - 2
	for l >= 0 && aa[l] > aa[l+1] {
		l--
	}
	if l < 0 {
		return false
	}
	r := n - 1
	for l < r && aa[l] > aa[r] {
		r--
	}
	aa[l], aa[r] = aa[r], aa[l]
	l++
	r = n - 1
	for l < r {
		aa[l], aa[r] = aa[r], aa[l]
		l++
		r--
	}
	return true
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

// ==================================================
// mod
// ==================================================

func madd(a, b int) int {
	a += b
	if a >= mod || a <= -mod {
		a %= mod
	}
	if a < 0 {
		a += mod
	}
	return a
}
func mmul(a, b int) int {
	return a * b % mod
}
func mpow(a, n int) int {
	if mod == 1 {
		return 0
	}
	r := 1
	for n > 0 {
		if n&1 == 1 {
			r = r * a % mod
		}
		a, n = a*a%mod, n>>1
	}
	return r
}
