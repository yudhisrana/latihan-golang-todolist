package todos

import "fmt"

type Todo struct {
	Number     int
	Content    string
	IsComplete bool
}

type TodoList struct {
	List []Todo
}

func (t *TodoList) InsertTodo(content string) {
	todo := Todo{
		Number:     len(t.List) + 1,
		Content:    content,
		IsComplete: false,
	}
	t.List = append(t.List, todo)
}

func (t *TodoList) ShowList() bool {
	if len(t.List) == 0 {
		fmt.Println("Todo list is empty")
		return false
	} else {
		for _, todo := range t.List {
			fmt.Printf("No.%d - Todo : %s - Is Complete : %t\n", todo.Number, todo.Content, todo.IsComplete)
		}
	}
	return true
}

func (t *TodoList) CompleteTodo(number int){
	for index := range t.List {
		if t.List[index].Number == number {
			t.List[index].IsComplete = true
			fmt.Printf("\nTodo %s has been completed\n\n", t.List[index].Content)
			return
		} else {
			fmt.Printf("\nTodo with number %d not found\n\n", number)
			return
		}
	}
}

func (t *TodoList) DeleteTodo(number int) {
	for index := range t.List {
		if t.List[index].Number == number {
			// simpan data sebelumnya
			dataBeforeDelete := t.List[index].Content

			// hapus data
			t.List = append(t.List[:index], t.List[index+1:]...)
			fmt.Printf("\nTodo %s has been deleted\n\n", dataBeforeDelete)
			return
		} else {
			fmt.Printf("\nTodo with number %d not found\n\n", number)
			return
		}
	}
}