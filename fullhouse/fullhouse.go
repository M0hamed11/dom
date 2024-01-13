package fullhouse

import (
	"dom/infoappliances"
	"dom/infofamily"
	"dom/infofurniture"
	"dom/inforoom"
)

func Dom() {
	appliances := infoappliances.AppliancesItems()
	family := infofamily.FamilyMembers()
	furniture := infofurniture.FurnitureItems()
	room := inforoom.RoomsItems()

	inforoom.ShowRoomDetails(room)
	infofurniture.ShowFurniture(furniture)
	infofamily.ShowFamily(family)
	infoappliances.ShowAppliances(appliances)
}
