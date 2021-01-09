package main

func main() {
	var balance int64 = 15_000_000_00 // 15 миллионов в копейках
	var transfer int64 = 10_000_000_00 // 10 миллионов в копейках
	total := balance + transfer
	println(total)
}
