package inforoom

import "fmt"

type Rooms struct {
	Name string
	h    int
	w    int
	l    int
}

func RoomsItems() []Rooms {

	Room := Rooms{"комната", 4, 3, 5}

	return []Rooms{Room}
}

func ShowRoomDetails(rooms []Rooms) {
	fmt.Print("плошадь дома:\n")
	for _, room := range rooms {
		fmt.Printf("Комната: %s \nПлощадь: %d кв м\n", room.Name, room.w*room.l)
	}

}
