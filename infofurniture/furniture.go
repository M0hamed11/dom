package infofurniture

import "fmt"

type furnitureStruct struct {
	name string
}

func FurnitureItems() []furnitureStruct {

	bed := furnitureStruct{"кровать"}
	closet := furnitureStruct{"шкаф"}
	couch := furnitureStruct{"диван"}
	table := furnitureStruct{"стол"}
	armchair := furnitureStruct{"кресло"}
	return []furnitureStruct{bed, closet, couch, table, armchair}
}

func ShowFurniture(furniture []furnitureStruct) {

	fmt.Println("\nМебель в комнате:")
	for _, item := range furniture {
		fmt.Println(item.name, " ")
	}
	fmt.Print("\n")
}
