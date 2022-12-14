package screen

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type Credits struct {
	source      *widgets.List
	transfer    *widgets.List
	destination *widgets.List
}

func NewCredits() *Credits {
	c := Credits{}
	c.source = widgets.NewList()
	c.transfer = widgets.NewList()
	c.destination = widgets.NewList()
	return &c
}

func (c *Credits) Run() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	setList("Andrew", c.source)
	setList("github.com/andrewarrow", c.transfer)
	c.transfer.Border = false
	setList("Arrow", c.destination)

	c.source.Rows = append(c.source.Rows, "CLICK THAT LIKE BUTTON")
	c.destination.Rows = append(c.destination.Rows, "SMASH THAT SUBSCRIBE BUTTON")

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.0,
			ui.NewCol(0.33, c.source),
			ui.NewCol(0.33, c.transfer),
			ui.NewCol(0.33, c.destination),
		),
	)

	ui.Render(grid)
	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Millisecond * 144).C
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				grid.SetRect(0, 0, payload.Width, payload.Height)
				ui.Clear()
			case "<Enter>":
				c.handleEnter()
			}
		case <-ticker:
			c.advanceRiver()
		}
		ui.Render(grid)
	}
}

func (c *Credits) handleEnter() {
}

func (c *Credits) advanceRiver() {
	list := []string{"red", "green", "yellow", "blue", "magenta", "cyan", "white"}
	pick := rand.Intn(len(list))
	c.transfer.Rows = append([]string{fmt.Sprintf("[ ](bg:%s)", list[pick])}, c.transfer.Rows...)
	if len(c.transfer.Rows) > 23 {
		c.transfer.Rows = c.transfer.Rows[0:23]
	}
}
