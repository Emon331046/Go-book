package main

import "fmt"

func main() {
	//x := make(map[string]int)
	//x["KEY"] = 10
	//fmt.Println(x["KEY"])
	//x1 := make(map[int]int)
	//x1[1] = 10
	//fmt.Println(x1)
	//delete(x1, 1)
	//fmt.Println(x1)
	//elements := make(map[string]string)
	//elements["H"] = "Hydrogen"
	//elements["He"] = "Helium"
	//elements["Li"] = "Lithium"
	//elements["Be"] = "Beryllium"
	//elements["B"] = "Boron"
	//elements["C"] = "Carbon"
	//elements["N"] = "Nitrogen"
	//elements["O"] = "Oxygen"
	//elements["F"] = "Fluorine"
	//elements["Ne"] = "Neon"
	//
	//elements := map[string]string{
	//	"H":  "Hydrogen",
	//	"He": "Helium",
	//	"Li": "Lithium",
	//	"Be": "Beryllium",
	//	"B":  "Boron",
	//	"C":  "Carbon",
	//	"N":  "Nitrogen",
	//	"O":  "Oxygen",
	//	"F":  "Fluorine",
	//	"Ne": "Neon",
	//}
	//fmt.Println(elements["Li"])
	//name, ok := elements["Un"]
	//fmt.Println("ok1", name, ok)
	//name1, ok1 := elements["Li"]
	//fmt.Println("ok1", name1, ok1)
	//if name2, ok2 := elements["Li"]; ok2 {
	//	println("ok2", name2, ok2)
	//}
	xx := make(map[string]map[string]string)
	xx["H"] = map[string]string{
		"name":  "hyd",
		"state": "GAS",
	}
	xx1 := xx["H"]
	xx1["name"] = "hydro"
	fmt.Println(xx1["name"])
	//fmt.Print(&xx["H"], &xx1)
	fmt.Println(xx["H"])

	//elements := map[string]map[string]string{
	//	"H": map[string]string{
	//		"name":  "Hydrogen",
	//		"state": "gas",
	//	},
	//	"He": map[string]string{
	//		"name":  "Helium",
	//		"state": "gas",
	//	},
	//	"Li": map[string]string{
	//		"name":  "Lithium",
	//		"state": "solid",
	//	},
	//	"Be": map[string]string{
	//		"name":  "Beryllium",
	//		"state": "solid",
	//	},
	//	"B": map[string]string{
	//		"name":  "Boron",
	//		"state": "solid",
	//	},
	//	"C": map[string]string{
	//		"name":  "Carbon",
	//		"state": "solid",
	//	},
	//	"N": map[string]string{
	//		"name":  "Nitrogen",
	//		"state": "gas",
	//	},
	//	"O": map[string]string{
	//		"name":  "Oxygen",
	//		"state": "gas",
	//	},
	//	"F": map[string]string{
	//		"name":  "Fluorine",
	//		"state": "gas",
	//	},
	//	"Ne": map[string]string{
	//		"name":  "Neon",
	//		"state": "gas",
	//	},
	//}
	//
	//if el, ok := elements["Li"]; ok {
	//	fmt.Println(el["name"], el["state"])
	//}
}
