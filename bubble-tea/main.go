package main

import (
	"fmt"
	"os"

	"bubbletea/app/controls"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(controls.NewTodoControl())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
