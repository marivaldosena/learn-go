package main

import "fmt"

func main() {
	var x [5]float64

	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0

	for _, value := range x {
		total += value
	}

	fmt.Println(total / float64(len(x)))

	y := [5]float64{98, 93, 77, 82, 83}
	fmt.Println(y)

	slice := make([]float64, 5)
	fmt.Println(slice)

	arr := [5]float64{1, 2, 3, 4, 5}
	sl := arr[0:5]
	fmt.Println(sl)

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)

	map1 := make(map[string]int)
	map1["key"] = 10
	fmt.Println(map1["key"])

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["Li"])

	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}

	elementos := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},

		"F": map[string]string{
			"name":  "Fluoride",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}

	if el, ok := elementos["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}
