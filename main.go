package main

type Player struct {
	inventory   []Item
	currentRoom *Room
	isBackpack  bool
}

type Item struct {
	name string
}

func initItems(names ...string) []Item {
	items := []Item{}

	for _, name := range names {
		items = append(items, Item{
			name: name,
		})
	}

	return items
}

type Room struct {
	name  string
	items []Item
	exits []*Exit
}

func initNamedRooms(names ...string) []*Room {
	rooms := []*Room{}

	for _, name := range names {
		rooms = append(rooms, &Room{
			name: name,
		})
	}

	return rooms
}

func (r *Room) addItems(items ...Item) {
		r.items = append(r.items, items...)
}

func (r *Room) addExits(exits ...*Exit) {
		r.exits = append(r.exits, exits...)
}

type Door struct {
	isOpened bool
}

func initDoor() *Door {
	return &Door{
		isOpened: false,
	}
}

type Exit struct {
	destination *Room
	door        *Door
}

func initExits(destinations ...*Room) []*Exit {
	exits := []*Exit{}

	for _, destination := range destinations {
		exits = append(exits, &Exit{
			destination: destination,
		})
	}

	return exits
}

//

func main() {
	/*
		в этой функции можно ничего не писать,
		но тогда у вас не будет работать через go run main.go
		очень круто будет сделать построчный ввод команд тут, хотя это и не требуется по заданию
	*/

	initGame()

	_ = handleCommand("some command")
}

func initGame() {
	rooms := initNamedRooms("кухня", "коридор", "комната", "улица")
	kitchen, corridor, sleepingRoom, street := rooms[0], rooms[1], rooms[2], rooms[3]

	items := initItems("чай", "ключи", "конспекты", "рюкзак")
	tea, keys, notes, backpack := items[0], items[1], items[2], items[3]

	kitchen.addItems(tea)
	sleepingRoom.addItems(keys, notes, backpack)

	door := initDoor()

	exits := initExits(kitchen, corridor, sleepingRoom, street)
	toKitchen, toCorridor, toSleepingRoom, toStreet := exits[0], exits[1], exits[2], exits[3]
	toStreet.door = door

	kitchen.addExits(toCorridor)
	corridor.addExits(toKitchen, toSleepingRoom, toStreet)
	sleepingRoom.addExits(toCorridor)
	street.addExits(toCorridor)
}

func handleCommand(command string) string {
	/*
		данная функция принимает команду от "пользователя"
		и наверняка вызывает какой-то другой метод или функцию у "мира" - списка комнат
	*/
	return "not implemented"
}
