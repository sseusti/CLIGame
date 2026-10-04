package main

import (
	"fmt"
	"slices"
	"strings"
)

func createAvailableRooms(exits []*Exit, resp string) string {
	resp += "можно пройти - "
	names := []string{}
	for _, exit := range exits {
		names = append(names, exit.destination.name)
	}
	resp += strings.Join(names, ", ")

	return resp
}

type Player struct {
	inventory   []Item
	currentRoom *Room
	isBackpack  bool
}

func initPlayer(room *Room) Player {
	return Player{
		inventory:   []Item{},
		currentRoom: room,
		isBackpack:  false,
	}
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
			door: &Door{
				isOpened: true,
			},
		})
	}

	return exits
}

type Handler struct {
	player *Player
}

func (h *Handler) handleCommand(command string) string {
	parts := strings.Split(command, " ")

	switch parts[0] {
	case "осмотреться":
		return h.handleLook()
	case "идти":
		return h.handleGo(parts[1])
	case "взять":
		return h.handleTake(parts[1])
	case "надеть":
		return h.handleWear(parts[1])
	case "применить":
		return h.handleApply(parts[1], parts[2])
	default:
		return "неизвестная команда"
	}
}

func (h *Handler) handleApply(item string, applyTo string) string {
	items := []string{}
	for _, i := range h.player.inventory {
		items = append(items, i.name)
	}

	if !slices.Contains(items, item) {
		return "нет предмета в инвентаре - " + item
	}

	if item == "ключи" {
		if h.player.currentRoom.name == "коридор" && applyTo == "дверь" {
			return "дверь открыта"
		}
	}

	return "не к чему применить"
}

func (h *Handler) handleWear(item string) string {
	if item != "рюкзак" {
		return "нет такого"
	}

	if h.player.currentRoom.name != "комната" || h.player.isBackpack {
		return "нет такого"
	}

	for idx, currentItem := range h.player.currentRoom.items {
		if currentItem.name == "рюкзак" {
			h.player.isBackpack = true

			h.player.inventory = append(
				h.player.inventory,
				currentItem,
			)

			h.player.currentRoom.items = append(
				h.player.currentRoom.items[:idx],
				h.player.currentRoom.items[idx+1:]...,
			)

			return "вы надели: рюкзак"
		}
	}

	return "нет такого"
}

func (h *Handler) handleTake(item string) string {
	items := []string{}
	for _, i := range h.player.currentRoom.items {
		items = append(items, i.name)
	}

	idx := slices.Index(items, item)
	if idx != -1 {
		if h.player.isBackpack {
			h.player.inventory = append(h.player.inventory, h.player.currentRoom.items[idx])
			h.player.currentRoom.items = append(h.player.currentRoom.items[:idx], h.player.currentRoom.items[idx+1:]...)
			return "предмет добавлен в инвентарь: " + item
		} else {
			return "некуда класть"
		}
	}

	return "нет такого"
}

func (h *Handler) handleLook() string {
	switch h.player.currentRoom.name {
	case "кухня":
		return LookKitchen(h.player.isBackpack)
	case "комната":
		return LookSleepingRoom(h.player.currentRoom.items, h.player.isBackpack, h.player.currentRoom.exits)
	}

	return ""
}

var uniqueStatus = map[string]string{
	"коридор": "ничего интересного.",
	"комната": "ты в своей комнате.",
	"кухня":   "кухня, ничего интересного.",
	"улица":   "на улице весна.",
}

func (h *Handler) handleGo(roomName string) string {
	for _, exit := range h.player.currentRoom.exits {
		if exit.destination.name == roomName {
			if exit.door.isOpened == false {
				return "дверь закрыта"
			}

			h.player.currentRoom = exit.destination
			return createAvailableRooms(h.player.currentRoom.exits, uniqueStatus[exit.destination.name]+" ")
		}
	}

	return "нет пути в " + roomName
}

func LookKitchen(isBackpack bool) string {
	resp := "ты находишься на кухне, на столе: чай, надо "
	if !isBackpack {
		resp += "собрать рюкзак и "
	}
	resp += "идти в универ. можно пройти - коридор"

	return resp
}

func LookSleepingRoom(items []Item, isBackpack bool, exits []*Exit) string {
	resp := ""
	if len(items) == 0 && isBackpack {
		resp += "пустая комната. "
		goto EXITS
	}

	if len(items) != 0 {
		names := []string{}
		for _, item := range items[:len(items)-1] {
			names = append(names, item.name)
		}
		resp += "на столе: " + strings.Join(names, ", ")
	}

	if !isBackpack {
		resp += ", на стуле: рюкзак"
	}

	resp += ". "

EXITS:
	resp += "можно пройти -"
	for _, exit := range exits {
		resp = resp + " " + exit.destination.name
	}

	return resp
}

var handler Handler

func main() {
	initGame()

	res := handleCommand("осмотреться")
	fmt.Println(res)
	res = handleCommand("идти коридор")
	fmt.Println(res)
	res = handleCommand("идти улица")
	fmt.Println(res)
	res = handleCommand("идти комната")
	fmt.Println(res)
	res = handleCommand("осмотреться")
	fmt.Println(res)
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

	player := initPlayer(kitchen)

	handler = Handler{
		player: &player,
	}
}

func handleCommand(command string) string {
	return handler.handleCommand(command)
}
