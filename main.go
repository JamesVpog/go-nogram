package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type pos struct {
	x int 
	y int
}
type model struct {
	board [][]int
	solution [][]int
	cursor pos
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// what messages were sent to update?
	case tea.KeyMsg:
		switch msg.String(){			
		// what was the keypress?
		case "ctrl+c", "q":
			return m, tea.Quit
		case "k", "up":
			// move the cursor up to the next row if possible
			if m.cursor.y > 0 {	
				m.cursor.y--
			}
		case "l", "right":
			// move right to next col if possible
			if m.cursor.x < len(m.board[0]) - 1 {
				m.cursor.x++
			}
		case "j", "down":
			if m.cursor.y < len(m.board) - 1 {
				m.cursor.y++
			}
		case "h", "left":
			if m.cursor.x > 0 {
				m.cursor.x--
			}
		case " ", "enter":
			// fill in the square at that position
			m.board[m.cursor.x][m.cursor.y] = 1	
		case "x", "X":
			// draw x in the square at that position 
			m.board[m.cursor.x][m.cursor.y] = 2
		}
	}
	return m, nil
}


func (m model) View() string {
	//TODO: make grid
	var s  strings.Builder	
	fmt.Fprintf(&s, "You are at: (%d, %d)\n", m.cursor.x, m.cursor.y)
	return s.String()
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