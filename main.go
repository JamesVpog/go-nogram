package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type pos struct {
	col int 
	row int
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
			if m.cursor.row > 0 {	
				m.cursor.row--
			}
		case "l", "right":
			// move right to next col if possible
			if m.cursor.col < len(m.board[0]) - 1 {
				m.cursor.col++
			}
		case "j", "down":
			if m.cursor.row < len(m.board) - 1 {
				m.cursor.row++
			}
		case "h", "left":
			if m.cursor.col > 0 {
				m.cursor.col--
			}
		case " ", "enter":
			// fill in the square at that position
			m.board[m.cursor.row][m.cursor.col] = 1	
		case "x", "X":
			// draw x in the square at that position 
			m.board[m.cursor.row][m.cursor.col] = 2
		}
	}
	return m, nil
}

// styles for the view
var (	
	cellStyle = lipgloss.NewStyle().
	    Width(20).
	    Height(5).
	    Align(lipgloss.Center, lipgloss.Center).
	    Border(lipgloss.NormalBorder())
	
	
	highlightedCell = cellStyle.Copy().
		Background(lipgloss.Color("1"))
)


func (m model) View() string {
	
	// loop through the m.board and render each cell, store in renderedCells
	// use joinHorizontal to make a row, stored in rows
	var rows []string
	for r := range m.board {
		var renderedCells []string
		for c := range m.board[r] {
			strForm := strconv.Itoa(m.board[r][c])
			// if the cell is currently where the user is at highlight it!
			var cell string
			if r == m.cursor.row && c == m.cursor.col {
				cell = highlightedCell.Render(strForm)
			} else {
				cell = cellStyle.Render(strForm)
			}	
			renderedCells = append(renderedCells, cell)
		}
		row := lipgloss.JoinHorizontal(lipgloss.Top, renderedCells...)
		rows = append(rows, row)
	}
	
	// send rows to joinVertical to make the grid
	grid := lipgloss.JoinVertical(lipgloss.Left, rows...)
	var s  strings.Builder		
	fmt.Fprintf(&s, "You are at: (%d, %d)\n", m.cursor.col, m.cursor.row)
	fmt.Fprintf(&s, "\n%s\n", grid)
	
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