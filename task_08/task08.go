package main

func main() {

	//turning 0 into 128
	println(setBitTo1(int64(0), 7))

	//turning 64 into 0
	println(setBitTo0(int64(64), 6))

}

func setBitTo0(num int64, pos int) int64 {
	return num &^ (1 << pos)
}

func setBitTo1(num int64, pos int) int64 {
	return num | (1 << pos)
}
