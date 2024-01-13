package infofamily

import "fmt"

type familyStruct struct {
	name string
}

func FamilyMembers() []familyStruct {
	mother := familyStruct{"мама"}
	father := familyStruct{"папа"}
	sister := familyStruct{"сестра"}
	return []familyStruct{mother, father, sister}
}

func ShowFamily(family []familyStruct) {
	fmt.Println("\nСемья:")
	for _, person := range family {
		fmt.Print(person.name, "\n ")
	}
}
