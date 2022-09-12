package main

import "fmt"

func main() {
	var x [5]int
	x[4] = 100
	fmt.Println(x[4])

	var x2 [5]float64
	x2[0] = 98
	x2[1] = 93
	x2[2] = 77
	x2[3] = 82
	x2[4] = 83

	var total float64 = 0
	for i := 0; i < len(x2); i++ {
		total += x2[i]
	}
	fmt.Println(total / float64(len(x2)))

	var total2 float64 = 0
	for _, value := range x2 {
		total2 += value
	}
	fmt.Println(total2 / float64(len(x)))

	x3 := [5]float64{98, 93, 77, 82, 83}
	fmt.Println(x3[2])

	x4 := [5]float64{
		1,
		93,
		77,
		82,
		83,
	}
	fmt.Println(x4[4])

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

	name, ok := elements["Li"]
	fmt.Println(name, ok)

	elements2 := map[string]map[string]string{
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
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}

	if el, ok := elements2["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}
