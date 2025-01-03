package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukan jumlah ikan dan kapasitas wadah ikan:")
	fmt.Scan(&x, &y)

	weights := make([]float64, x)

	fmt.Printf("Masukan berat %d ikan:\n", x)
	for i := 0; i < x; i++ {
		fmt.Scan(&weights[i])
	}

	totalWeights := []float64{}
	sum := 0.0
	for i := 0; i < x; i++ {
		sum += weights[i]
		if (i+1)%y == 0 || i == x-1 {
			totalWeights = append(totalWeights, sum)
			sum = 0.0
		}
	}

	fmt.Println("Total berat ikan yang ada di setiap wadah:")
	for _, total := range totalWeights {
		fmt.Printf("%.2f ", total)
	}
	fmt.Println()

	totalSum := 0.0
	for _, total := range totalWeights {
		totalSum += total
	}
	averageWeight := totalSum / float64(len(totalWeights))

	fmt.Printf("Berat rata-rata ikan di dalam setiap wadah: %.2f\n", averageWeight)
}
