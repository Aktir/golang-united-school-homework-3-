package homework

func average(input [15]float32) (result float32) {
	items := 0
	for _, v := range input {
		if v != 0.0 {
			items += 1
		}

	}

	var sum float32
	N := items
	for i := 0; i < N; i++ {
		sum += input[i]
	}
	return (sum) / (float32(N))

}
