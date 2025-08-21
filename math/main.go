package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {

	fmt.Println(PowerAndSqrt(9, 3))
	fmt.Println(CompareNumber(5.5, -2.3))
	fmt.Println(Trig(30))
	fmt.Println(RoundDemo(3.6))
	Game()
}

func PowerAndSqrt(x float64, n int) (float64, float64) {
	return math.Pow(x, float64(n)), math.Sqrt(float64(x))
}

func CompareNumber(a, b float64) (float64, float64, float64) {
	return math.Max(a, b), math.Min(a, b), math.Abs(a - b)
}

func Trig(n float64) (float64, float64, float64) {
	return math.Sin(n * math.Pi / 180), math.Cos(n * math.Pi / 180), math.Tan(n * math.Pi / 180)
}

func RoundDemo(n float64) (float64, float64, float64) {
	return math.Round(n), math.Floor(n), math.Ceil(n)
}

func Game() {
	rand.Seed(time.Now().UnixNano())

	target := rand.Intn(100)
	fmt.Println("target is: ", target)
	var guess int
	for {
		fmt.Println("请输入你的数字:")
		fmt.Scan(&guess)
		if guess < target {
			fmt.Println("小了")
		} else if guess > target {
			fmt.Println("大了")
		} else {
			fmt.Println("恭喜你猜中了，数字就是:", target)
			break
		}
	}
}
