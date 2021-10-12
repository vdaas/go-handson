package main

// FizzBuzzを実行する中で、"Fizz", "Buzz", "FizzBuzz", それ以外の数値 がそれぞれ何回出力されたか
// カウントし、出力してください、ただし、それ以外の数値は "others" としてカウントします
// 例: 1以上16未満のFizzBuzzの場合
//
//   Fizz 4
//   Buzz 2
//   FizzBuzz 1
//   others 8
//
// という結果になります。
func main() {
	// TODO 数値をカウントするためのmapを 宣言 & 初期化 (ヒント: make を使用)
	// counter := ...

	for i := 1; i < 100; i++ {
		switch {
		case i%15 == 0:
			// TODO counterの"FizzBuzz"をカウントアップ
		case i%5 == 0:
			// TODO counterの"Buzz"をカウントアップ
		case i%3 == 0:
			// TODO counterの"Fizz"をカウントアップ
		default:
			// TODO counterの"others"をカウントアップ
		}
	}
	// TODO counterのすべての要素を表示
}
