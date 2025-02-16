package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

// MEMO: 1_For
func main1a() {
	sum := 0
	// jsのように括弧はいらない
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// MEMO: 2_For continued
func main2a() {
	sum := 1
	// イテレート変数は省略可能
	// for ; sum < 1000; {
	// 	sum += sum
	// }
	fmt.Println(sum)
}

// MEMO: 3_For is Go's "while"
func main3a() {
	sum := 1
	// セミコロンも省略可能
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// MEMO: 4_Forever
func main4a() {
	// 無限ループになる
	for {
		fmt.Println("Hello, world")
	}
}

// MEMO: 5_If
func sqrt(x float64) string {
	// jsのように括弧はいらない
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main5a() {
	fmt.Println(sqrt(2), sqrt(-4))
}

// MEMO: 6_If with a short statement
func pow6a(x, n, lim float64) float64 {
	// if文の前に短いステートメントを書くことができる
	// そのifスコープ内でのみ有効
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main6a() {
	fmt.Println(
		pow6a(3, 2, 10),
		pow6a(3, 3, 20),
	)
}

// MEMO: 7_If and else
func pow7a(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// %gは浮動小数点数を出力
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// vはここでは使えない
	return lim
}

func main7a() {
	fmt.Println(
		pow7a(3, 2, 10),
		pow7a(3, 3, 20),
	)
}

// MEMO: 8_Exercise: Loops and Functions
func Sqrt(x float64) float64 {
	z := 1.0
	// 10回繰り返す
	for i := 0; i < 10; i++ {
		// ニュートン法
		// z = z - (z^2 - x) / 2z
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main8a() {
	fmt.Println(Sqrt(2))
}

// MEMO: 9_Switch
func main9a() {
	fmt.Print("Go runs on ")
	// 自動でbreakされる
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

// MEMO: 10_Switch evaluation order
func main10a() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday() // 4
	switch time.Saturday {        // 6
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

// MEMO: 11_Switch with no condition
func main11a() {
	t := time.Now()
	// switch trueと同じ
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

// MEMO: 12_Defer
func main12a() {
	defer fmt.Println("world") // この関数(main)が終了するときに実行される

	fmt.Println("hello")
}

// MEMO: 13_Stacking defers
func main13a() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i) // 複数のdeferはスタック(LIFO)して実行される
	}

	fmt.Println("done")
}
