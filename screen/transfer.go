package screen

import (
	"ack/files"
	"ack/tcp"
	"fmt"
	"log"
	"strings"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type Transfer struct {
	source      *widgets.List
	transfer    *widgets.List
	destination *widgets.List
	text        string
	step        int
	segments    []string
}

func NewTransfer() *Transfer {
	t := Transfer{}
	t.source = widgets.NewList()
	t.transfer = widgets.NewList()
	t.destination = widgets.NewList()
	t.text = files.ReadFile("data/antony_and_cleopatra.txt")
	t.segments = tcp.BreakIntoSegments(t.text)
	return &t
}

func (t *Transfer) Run() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	setList("Source (server)", t.source)
	setList("", t.transfer)
	t.transfer.Border = false
	setList("Destination (client)", t.destination)

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.0,
			ui.NewCol(0.33, t.source),
			ui.NewCol(0.33, t.transfer),
			ui.NewCol(0.33, t.destination),
		),
	)

	ui.Render(grid)
	uiEvents := ui.PollEvents()
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
				t.handleEnter()
			}
		}
		ui.Render(grid)
	}
}

func (t *Transfer) handleEnter() {
	if t.step == 0 {
		t.source.Rows = append(t.source.Rows, "LISTEN on port 80")
	} else if t.step == 1 {
		t.destination.Rows = append(t.destination.Rows, "SYN-SENT")
		t.transfer.Rows = append(t.transfer.Rows, "<-- <100,SYN>")
	} else if t.step == 2 {
		t.source.Rows = append(t.source.Rows, "SYN-RECEIVED")
		t.source.SelectedRow = 1
		t.transfer.Rows = append(t.transfer.Rows, "<300,101,SYN,ACK> -->")
		t.transfer.SelectedRow = 1
	} else if t.step == 3 {
		t.destination.Rows = append(t.destination.Rows, "ESTABLISHED")
		t.destination.SelectedRow = 1
		t.transfer.Rows = append(t.transfer.Rows, "<-- <101,301,ACK>")
		t.transfer.SelectedRow = 2
	} else if t.step == 4 {
		t.source.Rows = append(t.source.Rows, "ESTABLISHED")
		t.source.SelectedRow = 2
	} else if t.step == 5 {
		t.source.Rows = []string{"ESTABLISHED"}
		t.transfer.Rows = []string{}
		t.destination.Rows = []string{"ESTABLISHED"}
		t.source.SelectedRow = 0
		t.destination.SelectedRow = 0
	} else if t.step == 6 {
		t.source.Rows = append(t.source.Rows, strings.Split(t.text, "\n")...)
		t.source.Title = fmt.Sprintf("Source %d / %d = %.2f", len(t.text), tcp.MAX_SEGMENT_SIZE, float64(len(t.text))/float64(tcp.MAX_SEGMENT_SIZE))
	} else if t.step == 7 {
		t.source.Rows = []string{"ESTABLISHED"}
		for i, segment := range t.segments {
			t.source.Rows = append(t.source.Rows, fmt.Sprintf("%d %s", 301+i, segment))
		}
	} else if t.step == 8 {
		t.source.Rows = []string{"ESTABLISHED"}
		t.segments = t.segments[1:]
		for i, segment := range t.segments {
			t.source.Rows = append(t.source.Rows, fmt.Sprintf("%d %s", 302+i, segment))
		}
		t.transfer.Rows = append(t.transfer.Rows, "<301,512> -->")
		t.transfer.SelectedRow = 0
	}
	t.step++
}

func setList(title string, l *widgets.List) {
	l.SelectedRowStyle.Fg = ui.ColorWhite
	l.SelectedRowStyle.Bg = ui.ColorMagenta
	l.TextStyle.Fg = ui.ColorWhite
	l.TextStyle.Bg = ui.ColorBlack
	l.Title = title
}
