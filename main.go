package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	board [][]int
	solution [][]int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model) View() string {
	return ""
}

func initialModel() model {
	// example where filling in the middle square is the solution
	return model{
		board: [][]int{
			{0, 0, 0},
		 	{0, 0, 0},
			{0, 0, 0}},
		
		solution: [][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0}},
	}
}

func main(){	
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}
}