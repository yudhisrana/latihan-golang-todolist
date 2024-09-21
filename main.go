package main

import (
	"fmt"
	"os"

	"github.com/yudhisrana/latihan-golang-todolist/menu"
	"github.com/yudhisrana/latihan-golang-todolist/todos"
)

func main() {
	list := todos.TodoList{}

	for {
		menu.ClearTerminal() // Bersihkan terminal setiap kali menu ditampilkan
		input := menu.Menu()  // Tampilkan menu dan minta input pengguna

		switch input {
		case "1": // Menambahkan todo baru
			fmt.Print("Input Todo: ")
			var todo string
			fmt.Scanln(&todo)
			list.InsertTodo(todo)
			fmt.Printf("Todo %s has been added\n\n", todo)
			menu.Continue()

		case "2": // Menampilkan daftar todos
			if !list.ShowList() {
				menu.Continue()
			} else {
				menu.Continue()
			}
			
		case "3": // Menyelesaikan todo berdasarkan nomor
			if !list.ShowList() {
				menu.Continue()
			} else {
				fmt.Print("Input number of todo to complete: ")
				var number int
				fmt.Scanln(&number)
				list.CompleteTodo(number)
				menu.Continue()
			}

		case "4": // Menghapus todo berdasarkan nomor
			if !list.ShowList() {
				menu.Continue()
			} else {
				fmt.Print("Input number of todo to delete: ")
				var number int
				fmt.Scanln(&number)
				list.DeleteTodo(number)
				menu.Continue()
			}

		case "5": // Keluar dari program
			fmt.Println("Exiting program. Goodbye!")
			os.Exit(0)

		default: // Jika input tidak valid
			fmt.Println("Invalid input. Please try again.")
			menu.Continue()
		}
	}
}
