package main

import (
	"bufio"
	"container/heap"
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

	n, m := ni2()
	aaa := make([][]int, n)
	for i := 0; i < n; i++ {
		aaa[i] = make([]int, m)
		aa := nis(m)
		aaa[i] = aa
	}
	ans := -math.MaxInt32
	for i := 0; i < m; i++ {
		for j := i + 1; j < m; j++ {
			sum := 0
			for k := 0; k < n; k++ {
				sum += MaxIn2Values(aaa[k][i], aaa[k][j])
			}
			if ans < sum {
				ans = sum
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
func MaxIn2Values(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
func MinIn2Values(a int, b int) int {
	if a < b {
		return a
	}
	return b
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

// nPk
func permutation(n int, k int) int {
	if k > n || k <= 0 {
		panic(fmt.Sprintf("invalid param n:%v k:%v", n, k))
	}
	v := 1
	for i := 0; i < k; i++ {
		v *= (n - i)
	}
	return v
}

/*
for {

		// Do something

		if !nextPermutation(sort.IntSlice(x)) {
			break
		}
	}
*/
func nextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

type combFactorial struct {
	fac    []int
	facinv []int
}

func newcombFactorial(n int) *combFactorial {

	//facには階乗
	fac := make([]int, n)
	//facinvには逆元が入る
	facinv := make([]int, n)
	fac[0] = 1
	facinv[0] = minvfermat(1, mod)

	for i := 1; i < n; i++ {
		fac[i] = mmul(i, fac[i-1])
		facinv[i] = minvfermat(fac[i], mod)
	}

	return &combFactorial{
		fac:    fac,
		facinv: facinv,
	}
}

func (c *combFactorial) factorial(n int) int {
	return c.fac[n]
}

func (c *combFactorial) combination(n, r int) int {
	if r > n {
		return 0
	}
	return mmul(mmul(c.fac[n], c.facinv[r]), c.facinv[n-r])
}

func (c *combFactorial) permutation(n, r int) int {
	if r > n {
		return 0
	}
	return mmul(c.fac[n], c.facinv[n-r])
}

func (c *combFactorial) homogeousProduct(n, r int) int {
	return c.combination(n-1+r, r)
}

// 最大公約数
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// 最小公倍数
func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

type vertex struct {
	seen bool
	memo int
}
type edge struct {
	to   int
	cost int
}
type graph struct {
	v     []vertex
	nodes [][]edge
}

// NewGraph creates a new graph with n vertices.
func NewGraph(n int) *graph {
	g := &graph{
		v:     make([]vertex, n),
		nodes: make([][]edge, n),
	}
	return g
}

// AddEdge adds a edge connects vertex u to v and v to u.
func (g *graph) AppendEdge(from, to, cost int) {
	g.nodes[from] = append(g.nodes[from], edge{
		to:   to,
		cost: cost,
	})
}

type PriorityQueueItem struct {
	node, distance int
	index          int //for heap interface
}

type PriorityQueue []*PriorityQueueItem

// 以下heapのインターフェースを満たすために必要な要素
// キューの長さを返す
func (pq PriorityQueue) Len() int { return len(pq) }

// i番目の要素がj番目の要素よりも優先度が高い場合にtrueを返す
// 優先度の高さはここで識別される
// 今回はdistance
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

// i番目の要素とj番目の要素を入れ替える
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// 要素を追加する
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*PriorityQueueItem)
	item.index = n
	*pq = append(*pq, item)
}

// 優先度の高い要素を取り出す
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
func (graph *graph) dijkstra(start int) []int {
	//各ノードにかかるコスト
	dist := make([]int, len(graph.nodes))
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[start] = 0

	pq := make(PriorityQueue, 1)
	pq[0] = &PriorityQueueItem{node: start, distance: 0, index: 0}
	//pqをpriotiry queueとして機能するように初期化
	heap.Init(&pq)

	for pq.Len() > 0 {
		//優先度の高いやつを"一つだけ"取り出す
		item := heap.Pop(&pq).(*PriorityQueueItem)
		//vは優先度の高いノード
		//最初はスタート
		v := item.node
		if dist[v] < item.distance {
			continue
		}
		for _, edge := range graph.nodes[v] {
			nextDist := dist[v] + edge.cost
			if nextDist < dist[edge.to] {
				dist[edge.to] = nextDist
				heap.Push(&pq, &PriorityQueueItem{node: edge.to, distance: nextDist})
			}
		}
	}
	return dist
}

// 向きを変えた
func (g *graph) reverseGraph() *graph {
	rg := NewGraph(len(g.v))
	for i, edges := range g.nodes {
		for _, e := range edges {
			rg.AppendEdge(e.to, i, e.cost)
		}
	}
	return rg
}

// 順番に末端から数字を入れていく作業
func (g *graph) DFS1(v int, stack *[]int) {
	g.v[v].seen = true
	for _, adjEdge := range g.nodes[v] {
		if !g.v[adjEdge.to].seen {
			g.DFS1(adjEdge.to, stack)
		}
	}
	*stack = append(*stack, v)
}

func (g *graph) DFS2(v int, scc *[]int) {
	g.v[v].seen = true
	*scc = append(*scc, v)
	for _, adjEdge := range g.nodes[v] {
		if !g.v[adjEdge.to].seen {
			g.DFS2(adjEdge.to, scc)
		}
	}
}

// 強連結成分分解
func (g *graph) stronglyConnectedComponents() [][]int {
	stack := make([]int, 0)

	for i := 0; i < len(g.v); i++ {
		if !g.v[i].seen {
			g.DFS1(i, &stack)
		}
	}

	rg := g.reverseGraph()

	for i := range rg.v {
		rg.v[i].seen = false
	}

	var sccList [][]int

	for len(stack) > 0 {
		//一番番号の遅いやつから取っていく
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !rg.v[node].seen {
			scc := make([]int, 0)
			rg.DFS2(node, &scc)
			sccList = append(sccList, scc)
		}
	}
	return sccList
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

func mdiv(a, b int) int {
	a %= mod
	return a * minvfermat(b, mod) % mod
}

func mpow(a, n, m int) int {
	if m == 1 {
		return 0
	}
	r := 1
	for n > 0 {
		if n&1 == 1 {
			r = r * a % m
		}
		a, n = a*a%m, n>>1
	}
	return r
}

func minv(a, m int) int {
	p, x, u := m, 1, 0
	for p != 0 {
		t := a / p
		a, p = p, a-t*p
		x, u = u, x-t*u
	}
	x %= m
	if x < 0 {
		x += m
	}
	return x
}

// m only allow prime number
func minvfermat(a, m int) int {
	return mpow(a, m-2, mod)
}
