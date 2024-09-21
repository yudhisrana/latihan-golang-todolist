package menu

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func Menu() string {
	fmt.Println("----- Program Todo List -----")
	fmt.Println("1. Add Todo")
	fmt.Println("2. Show Todo")
	fmt.Println("3. Complete Todo")
	fmt.Println("4. Delete Todo")
	fmt.Println("5. Exit")
	fmt.Print("\nEnter your choice: ")
	var input string
	fmt.Scanln(&input)
	fmt.Println("")
	return input
}

func ClearTerminal() {
	var clearCMD *exec.Cmd

	switch runtime.GOOS {
		case "windows":
			clearCMD = exec.Command("cmd", "/c", "cls")
		default:
			clearCMD = exec.Command("clear")
	}

	clearCMD.Stdout = os.Stdout
	clearCMD.Run()
}

func Continue() {
	fmt.Print("Press enter to continue...")
	fmt.Scanln()
}