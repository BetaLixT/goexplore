package controls

import (
	"bubbletea/domain/models"
	"bubbletea/domain/services"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type TodoControl struct {
	data     []models.Todo
	cursor   int
	selected map[int]struct{}
	service  services.TodoService
}

func (ctrl *TodoControl) Init() tea.Cmd {
	ctrl.data, _ = ctrl.service.ListTodos()
	return nil
}

func (ctrl *TodoControl) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return ctrl, tea.Quit
		}
	}
	return ctrl, nil
}

func (ctrl *TodoControl) View() string {
	s := "Pending TODOs\n"
	for i, todo := range ctrl.data {
		cursor := " "
		if ctrl.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, todo.Title)
	}
	s += "\npress q to quit\n"
	return s
}

func NewTodoControl() *TodoControl {
	return &TodoControl{}
}
