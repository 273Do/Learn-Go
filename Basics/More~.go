package main

import (
	"fmt"
	"math"
	"strings"
	// "golang.org/x/tour/pic"
)

// MEMO: 1_Pointers
func main1b() {
	i, j := 42, 2701

	// &オペレータで変数のアドレスを取得
	p := &i
	// *オペレータでポインタが指す先の変数を示す
	fmt.Println(p)  // iのメモリアドレス
	fmt.Println(*p) // メモリアドレス(ポインタ)の中の値
	*p = 21         // ポインタを通じてiに値を代入
	fmt.Println(i)

	p = &j       // jのアドレスをpに代入
	*p = *p / 37 // jの値を37で割る
	fmt.Println(j)
}

// MEMO: 2_Structs
// struct (構造体)は，フィールド( field )の集まり
// クラスに似ているが，フィールドのみを持つ
type Vertex struct {
	X int
	Y int
}

func main2b() {
	fmt.Println(Vertex{1, 2})
}

// MEMO: 3_Struct fields

func main3b() {
	v := Vertex{1, 2}
	v.X = 4 // フィールドへのアクセス
	fmt.Println(v.X)
}

// MEMO: 4_Pointers to structs
// ポインタを使って構造体のフィールドにアクセスする
func main4b() {
	v := Vertex{1, 2}
	p := &v
	// ポインタを取得
	fmt.Println(p)
	fmt.Println(*p)
	p.X = 1e9 // (*p).X と書いても同じ
	fmt.Println(v)

	var test string = "test"
	var q = &test
	fmt.Println(q)
	fmt.Println(*q)
}

// MEMO: 5_Struct literals
// フィールドの一部だけを列挙することができる．
var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main5b() {
	fmt.Println(v1, p, v2, v3)
}

// MEMO: 6_Arrays
func main6b() {
	// [n]T 型は，型 T の n 個の変数の配列( array )を表す．
	var a [2]string // 文字列の2個の配列
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// 長さ6のint配列
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

// MEMO: 7_Slices
// NOTE: 配列は固定長であるが，スライスは可変長
func main7b() {

	primes := [6]int{2, 3, 5, 7, 11, 13}

	// 配列の一部を参照するデータ構造
	// インデックス1から4手前までの要素を参照
	var s []int = primes[1:4]
	fmt.Println(s)
}

// MEMO: 8_Slices are like references to arrays
func main8b() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// スライスの要素を変更すると，元の配列の対応する要素も変更される
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

// MEMO: 9_Slice literals
func main9b() {
	// 配列リテラルのようにスライスリテラルを作成できる
	// 配列リテラルと違い，長さを指定しない
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

// MEMO: 10_Slice defaults
func main10b() {
	s := []int{2, 3, 5, 7, 11, 13}

	// 開始インデックスが省略された場合は0
	s = s[1:4]
	fmt.Println(s)

	// 終了インデックスが省略された場合はスライスの長さ
	s = s[:2]
	fmt.Println(s)

	// 両方省略された場合は元のスライス
	s = s[1:]
	fmt.Println(s)
}

// MEMO: 11_Slice length and capacity

func printSlice11(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main11b() {
	// len() は配列やスライスの長さ(length)、cap() は容量(capacity)を求めます。
	// 長さは実際に使用されている数、容量はメモリ上に確保されている数
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice11(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice11(s)

	// Extend its length.
	s = s[:4]
	printSlice11(s)

	// Drop its first two values.
	s = s[2:]
	printSlice11(s)

}

// MEMO: 12_Nil slices
func main12b() {
	var s []int // nilスライス

	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

// MEMO: 13_Creating a slice with make
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x) // len=長さ，cap=容量
}

func main13b() {
	// make関数でスライスを作成
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

// MEMO: 14_Slices of slices
// スライスは他のスライスを含むことができる(二重配列のようなもの)
func main14b() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

// MEMO: 15_Append to a slice
// append関数でスライスに要素を追加できる
func main15b() {
	var s []int
	printSlice11(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice11(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice11(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice11(s)
}

// MEMO: 16_Range
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main16b() {
	// rangeはスライスやマップを反復処理するために使う
	// インデックスと値を返す(jsのmap関数と同じ)
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

// MEMO: 17_Range continued
// indexやvalueを使わない場合は，_を使って省略できる
func main17b() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i==(2^i)
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

// MEMO: 18_Exercise: Slices

func Pic(dx, dy int) [][]uint8 {
	// スライスのスライスを作成
	// 二次元配列のようなもの
	pic := make([][]uint8, dy)
	for y := range pic {
		pic[y] = make([]uint8, dx)
		for x := range pic[y] {
			pic[y][x] = uint8((x + y) / 2)
		}
	}
	return pic
}

func main18b() {
	// pic.Show(Pic)
}

// MEMO: 19_Maps
// mapはキーと値のペアを保持するデータ構造
type Vertex19 struct {
	Lat, Long float64
}

// mapはキーと値のペアを保持するデータ構造
// mはstring型のキーとVertex型の値を持つmap(連想配列のようなもの)
var m map[string]Vertex19

func main19b() {
	m = make(map[string]Vertex19)
	m["Bell Labs"] = Vertex19{
		40.68433, -74.39967,
	}
	m["Test"] = Vertex19{1, 2}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m)
}

// MEMO: 20_Map literals
var m20 = map[string]Vertex19{
	// mapリテラルはキーが必要
	"Bell Labs": Vertex19{
		40.68433, -74.39967,
	},
	"Google": Vertex19{
		37.42202, -122.08408,
	},
}

func main20b() {
	fmt.Println(m20)
}

// MEMO: 21_Map literals continued
// mapリテラルの要素の型が省略された場合，最後に指定された型から推定される
var m21 = map[string]Vertex19{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main21b() {
	fmt.Println(m21)
}

// MEMO: 22_Mutating Maps
func main22b() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"] // 2つ目の値でキーが存在するかどうかを確認できる
	fmt.Println("The value:", v, "Present?", ok)
}

// MEMO: 23_Exercise: Maps
// skip

// MEMO: 24_Function values
// 関数も変数に代入できる(jsでいうコールバック関数)
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main24b() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot)) // 関数はcomputeで呼び出される
	fmt.Println(compute(math.Pow))

}

// Memo: 25_Function closures
func adder() func(int) int { // intを引数に取り，intを返す関数を返す
	sum := 0 // 関数内で変数を保持することができる(クロージャの形成)
	return func(x int) int {
		sum += x
		return sum
	}
}

func main25b() {
	// クロージャを作成，posとnegはそれぞれ独立したsumを持つ
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

// MEMO: 26_Exercises: Fibonacci closure
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
