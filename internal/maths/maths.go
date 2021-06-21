package maths

func LazinessGen(n string, c string, w float32) int {
	lazy := ((len(n) * len(c)) / int(w*w)) * 100
	return lazy
}
