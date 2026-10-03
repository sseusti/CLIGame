package main

/*
код писать в этом файле
наверняка у вас будут какие-то структуры с методами, глобальные переменные ( тут можно ), функции
*/
type Player struct {
	inventory   []Item
	currentRoom *Room
	isBackpack  bool
}

type Item struct {
	name string
}

type Room struct {
	name  string
	items []Item
	exits []Exit
}

type Door struct {
	isOpened bool
}

type Exit struct {
	destination *Room
	door        *Door
}

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
	/*
		эта функция инициализирует игровой мир - все комнаты
		если что-то было - оно корректно перезатирается
	*/
}

func handleCommand(command string) string {
	/*
		данная функция принимает команду от "пользователя"
		и наверняка вызывает какой-то другой метод или функцию у "мира" - списка комнат
	*/
	return "not implemented"
}
