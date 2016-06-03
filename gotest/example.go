package gotest

func Add(a, b int) int {
	return a + b
}

func Mod(a, b int) int {
	if a == 0 || b == 0 { // ignore this
		return 0
	}
	return a % b
}
