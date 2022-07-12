package homework

func average(input [15]float32) (result float32) {
	input = [15]float32{1, 2, 3, 4, 5, 6}
	var sum float32
	N := 6
	for i := 0; i <= N; i++ {
		sum += input[i]
	}
	return (sum) / (float32(N))

}
