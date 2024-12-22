package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func ratarata(arrBerat arrBalita, n int) float64 {
	total := 0.0
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	var balita arrBalita

	fmt.Print("Masukkan banyaknya data berat balita: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data berat balita ke-%d: ", i+1)
		fmt.Scan(&balita[i])
	}

	var bMin, bMax float64
	hitungMinMax(balita, n, &bMin, &bMax)
	rata := ratarata(balita, n)

	fmt.Printf("data Berat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("data Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("data Ratarata berat balita: %.2f kg\n", rata)
}
