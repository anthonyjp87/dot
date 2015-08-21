package dot

func Dot(x, y []float64) (r float64) {

	for i := range x {
		r += x[i] * y[i]
	}
	return
}
