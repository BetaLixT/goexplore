package controls

import (
	"bubbletea/domain/models"

	tea "github.com/charmbracelet/bubbletea"
)

type TodoControl struct {
	data     []models.Todo
	cursor   int
	selected map[int]struct{}
}

func (ctrl *TodoControl) Init() tea.Cmd {
	return nil
}

func (ctrl *TodoControl) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch(msg) := msg.(type) {
		case tea.KeyMsg:
			switch msg.String(){
			case "q", "ctrl+c":
				return ctrl, tea.Quit
			}
	}
}
