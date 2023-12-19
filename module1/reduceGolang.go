package module1

func Recursive(val int) int {
	if val != 1 {
		return val + Recursive(val-1)
	} else {
		return 1
	}
}