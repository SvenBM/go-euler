package main

func euler15(gridsize int) int {
	// problem: verteile 2 ziffern auf 2n stellen, wobei jede n mal vorkommen darf
	// factorsProd := 1
	// divisorsProd := 1
	// for i := 1; i <= gridsize; i++ {
	// 	if 2*i <= gridsize {
	// 		divisorsProd *= i
	// 		factorsProd *= (2*gridsize - (2*i - 1))
	// 	} else {
	// 		factorsProd *= 2
	// 	}
	// }
	// return factorsProd / divisorsProd

	//ist auch binnomialkoeff: verteile k=n auf 2n Stellen
	res := 1
	for i := 1; i <= gridsize; i++ {
		res = res * (gridsize + i) / i
	}
	return res
}
