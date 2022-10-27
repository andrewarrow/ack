package screen

import (
	"ack/files"
	"fmt"
	"log"
	"strings"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var shutdown bool
var source = widgets.NewList()
var transfer = widgets.NewList()
var destination = widgets.NewList()

func Setup() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	SetList("Source", source)
	SetList("", transfer)
	transfer.Border = false
	SetList("Destination", destination)

	text := files.ReadFile("data/antony_and_cleopatra.txt")
	source.Rows = strings.Split(text, "\n")
	source.Title = fmt.Sprintf("Source %d", len(text))

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.0,
			ui.NewCol(0.33, source),
			ui.NewCol(0.33, transfer),
			ui.NewCol(0.33, destination),
		),
	)

	ui.Render(grid)
	uiEvents := ui.PollEvents()
	for {
		if shutdown {
			break
		}
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				grid.SetRect(0, 0, payload.Width, payload.Height)
				ui.Clear()
			}
		}
		ui.Render(grid)
	}
}

func SetList(title string, l *widgets.List) {
	l.SelectedRowStyle.Fg = ui.ColorWhite
	l.SelectedRowStyle.Bg = ui.ColorMagenta
	l.TextStyle.Fg = ui.ColorWhite
	l.TextStyle.Bg = ui.ColorBlack
	l.Title = title
}
