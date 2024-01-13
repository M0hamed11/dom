package infoappliances

import "fmt"

type appliancesStruct struct {
	name string
}

func AppliancesItems() []appliancesStruct {
	fridge := appliancesStruct{"холодильник"}
	oven := appliancesStruct{"духовка"}
	teapot := appliancesStruct{"чайник"}
	hairdryer := appliancesStruct{"фен"}
	iron := appliancesStruct{"утюг"}
	return []appliancesStruct{fridge, oven, teapot, hairdryer, iron}
}

func ShowAppliances(appliances []appliancesStruct) {
	fmt.Println("Бытовая техника:")
	for _, everyItem := range appliances {
		fmt.Print(everyItem.name, "\n")
	}

}
