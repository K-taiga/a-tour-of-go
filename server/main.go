package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
	// "math/rand"
	"math"
	// "math/cmplx"
	// "runtime"
	// "time"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil)
}

// ↓Hello World用
// func main() {
//   http.Handle("/", &templateHandler{filename: "template.html"})

//   if err := http.ListenAndServe(":8080", nil); err != nil {
//     log.Fatal("ListenAndServe", err)
//   }
// }

// 関数の２つ以上の引数が同じ型である場合には、最後の型を残して省略して記述
// 戻り値の型も指定できる
func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// 戻り値に名前をつけると関数の最初で変数を定義したものとして、関数の中でそのまm使える
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	// 名前をつけた戻り値の変数はreturnのみで返せる　naked return
	// 長いコードではnaked returnは可読性が悪い
	return
}

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func sqrt(x float64) string {
	// ifの()も不要
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// Pow = べき乗
	// 条件で評価する前にステートメントを記述できる
	// xをnべき乗したものとlimを比較
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func Sqrt(x float64) float64 {
	// z := float64(1)
	z := 1.0
	// iが直前の値
	// z -= i 直前に求めたzの値からiを引く
	// その値が限りなくなくなるまで続ける　1e - 10
	for i := 1.0; i*i > 1e-10; z -= i {
		i = (z*z - x) / (2 * z)
	}
	return z
}

// 変数のリスト　パッケージと関数内で利用できる　型は一緒なら最後だけでOK
// var c , python , java bool

// var i, j int = 1,2

// var (
// 	ToBe   bool       = false
// 	MaxInt uint64     = 1<<64 - 1
// 	// cmplx.Sqrtは平方根
// 	z      complex128 = cmplx.Sqrt(-5 + 12i)
// )

// const Pi = math.Pi

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func main() {
	// ・Packages, variables, and functions.

	// fmt.Println("My favorite number is", rand.Intn(10))

	// fmt.Printf("Now you have %g problems.", math.Sqrt(7))

	// 大文字で始まるのはパッケージがエクスポートしている名前　小文字では参照できない
	// fmt.Println(math.Pi);

	// fmt.Println(add(42,13))

	// a, b := swap("hello", "world")
	// fmt.Println(a,b)

	// fmt.Println(split(18))

	// var i int
	// fmt.Println(i, c, python, java)

	// 変数の初期値がそのまま型になる
	// var c , python , java = true, false, "no!"
	// fmt.Println(i, j, c, python, java)

	// 関数の外ではvar宣言必要
	// var i, j int = 1, 2
	// varの省略が := 暗黙的な型宣言　入れたものがそのまま型になる
	// 関数の中でのみこの書き方ができる
	// k := 3
	// c, python, java := true, false, "no!"
	// fmt.Println(i, j, k, c, python, java)

	// fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	// fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	// fmt.Printf("Type: %T Value: %v\n", z, z)

	// 初期値なしはzero value
	// var i int
	// var f float64
	// var b bool
	// var s string
	// 0 0 false ""
	// fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// 変数の型と入れるものの型が一緒じゃないと入れられない　明示的に型変換が必要
	// var x, y int = 3, 4
	// var f float64 = math.Sqrt(float64(x*x + y*y))
	// var z uint = uint(f)
	// fmt.Println(x, y, z)
	// i := 42
	// f := float64(i)
	// u := uint(f)
	// fmt.Println(i, f, u)

	// 型を指定しなければ入れたもので型が推論される
	// i := 42           // int
	// f := 3.142        // float64
	// g := 0.867 + 0.5i // complex128
	// fmt.Printf("i is of type %T\n", i)
	// fmt.Printf("f is of type %T\n", f)
	// fmt.Printf("g is of type %T\n", g)

	// const World = "世界"
	// fmt.Println("Hello", World)
	// fmt.Println("Happy", Pi, "Day")

	// const Truth = true
	// fmt.Println("Go rules?", Truth)

	// fmt.Println(needInt(Small))
	// fmt.Println(needFloat(Small))
	// fmt.Println(needFloat(Big))

	// ・Flow control statements: for, if, else, switch and defer

	// falseになるまで
	// sum := 0
	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)

	// 初期化と後処理の記述は任意
	// ;も省略可 goではwhileはなくforのみ
	// sum := 1
	// for sum < 1000 {
	// 	sum += sum
	// }
	// fmt.Println(sum)

	// 無限ループ
	// for {
	// }

	// fmt.Println(sqrt(2), sqrt(-4))

	// fmt.Println(
	// 	pow(3, 2, 10),
	// 	pow(3, 3, 20),
	// )

	// Exercise: Loops and Functions
	// fmt.Println(Sqrt(16))

	// fmt.Print("Go run on ")
	//どれかのcaseだけ実行し自動でbreak
	// switch os := runtime.GOOS; os {
	// case "darwin":
	// 	fmt.Println("OS x.")
	// case "linux":
	// 	fmt.Println("Linux.")
	// default:
	// 	fmt.Printf("%s.", os)
	// }

	// fmt.Println("When's Saturday?")
	// today := time.Now().Weekday()
	// switch time.Saturday {
	// case today + 0:
	// 	fmt.Println("Today")
	// case today + 1:
	// 	fmt.Println("Tomorrow")
	// case today + 2:
	// 	fmt.Println("In two days")
	// default:
	// 	fmt.Println("Too far away")
	// }

	// t := time.Now()
	// 条件のないswitchはswitch trueと同じ
	// switch {
	// case t.Hour() < 12:
	// 	fmt.Println("Good morning!")
	// case t.Hour() < 17:
	// 	fmt.Println("Good afternoon")
	// default:
	// 	fmt.Println("Good evening")
	// }

	// defer = 延期する
	// deferへ渡した関数の実行を呼び出し元の関数の終わりまで延期させる
	// deferで渡した関数への引数はすぐに評価されるが、それが実行されるのは最後
	// defer fmt.Println("world")
	// fmt.Println("hello")

	// deferへ渡した関数が複数あればスタックされ、last in first outで実行される(入れたもの順に実行される)
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
