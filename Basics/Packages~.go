// NOTE: https://go-tour-jp.appspot.com/basics/1

// Goではmainパッケージから実行が始まる
package main

// MEMO: 1_Packages
// fmtとmath/randパッケージをインポート
// パッケージ名はインポートパスの最後の要素と同じ名前になる
import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

// main関数を定義
func main1() {
	// fmtパッケージのPrintln関数を呼び出し
	// コンソールに出力
	fmt.Println("My favorite number is", rand.Intn(10))
}

// MEMO: 2_Imports

// 以下でも可能
// import "fmt"
// import "math"

func main2() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}

// MEMO: 3_Exported names
func main3() {
	// 大文字で始まる名前は外部パッケージから参照可能(小文字はだめ)
	fmt.Println(math.Pi)
}

// MEMO: 4_Functions
// 後ろに型を書く
func add4(x int, y int) int {
	return x + y
}

func main4() {
	fmt.Println(add4(42, 13))
}

// MEMO: 5_FUnctions continued
// 複数の引数が同じ型の場合，最後に型を書く
func add5(x, y int) int {
	return x + y
}

func main5() {
	fmt.Println(add5(42, 13))
}

// MEMO: 6_Multiple results
// 複数の戻り値を返すことができる
func swap(x, y string) (string, string) {
	return y, x
}

func main6() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

// MEMO: 7_Named return values
// 関数の最初で定義した変数名が戻り値として使われる
// 可読性に影響があるため短い関数でのみ使うべき
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main7() {
	fmt.Println(split(17))
}

// MEMO: 8_Variables

var c, python, java bool

func main8() {
	// var宣言で変数を定義
	var i int
	fmt.Println(i, c, python, java)
}

// MEMO: 9_Variables with initializers
// 初期値を与えることができる
var i, j int = 1, 2

func main9() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

// MEMO: 10_Short variable declarations
// 関数内でのみ:=を使って暗黙的な方宣言，変数を定義できる
func main10() {
	var i, j int = 1, 2
	k := 3 // var，型宣言を省略できる
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}

// MEMO: 11_Basic types

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main11() {
	// %Tは変数の型を表示，%vは変数の値を表示
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

// MEMO: 12_Zero values
func main12() {
	// 初期化されていない変数はゼロ値が与えられる
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

// MEMO: 13_Type conversions
func main13() {
	var x, y int = 3, 4
	// 型変換はT(v)で行う
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

// MEMO: 14_Type inference
func main14() {
	v := 42 // 自動的に型が決定される
	fmt.Printf("v is of type %T\n", v)
}

// MEMO: 15_Constants
const Pi = 3.14

func main15() {
	// constを使って定数を宣言，再代入はできない
	// 定数は:=を使って宣言できない
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

// MEMO: 16_Numeric Constants
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main16() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
