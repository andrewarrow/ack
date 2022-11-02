package screen

import (
	"fmt"
	"log"
	"sync"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var mutex sync.Mutex

type TransferWithOptions struct {
	source             *widgets.List
	transfer           *widgets.List
	destinationBuffer  *widgets.List
	destinationProgram *widgets.List
	step               int
	bufferSize         int
	wireSpeed          int
	processSpeed       int
	totalSegments      int
	segments           int
}

func NewTransferWithOptions(bufferSize, wireSpeed, processSpeed int) *TransferWithOptions {
	t := TransferWithOptions{}
	t.source = widgets.NewList()
	t.transfer = widgets.NewList()
	t.destinationBuffer = widgets.NewList()
	t.destinationProgram = widgets.NewList()
	t.step = 301
	t.bufferSize = bufferSize
	t.wireSpeed = wireSpeed
	t.processSpeed = processSpeed
	return &t
}

func (t *TransferWithOptions) Run() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	setList("Source (server)", t.source)
	t.source.Rows = append(t.source.Rows, "102.7 MB")
	t.source.Rows = append(t.source.Rows, "102,729,562 bytes")
	t.source.Rows = append(t.source.Rows, "102,729,562 / 512")
	t.source.Rows = append(t.source.Rows, "200,643.67578125")
	t.source.Rows = append(t.source.Rows, "200,644 segments")
	t.source.Rows = append(t.source.Rows, "200644")
	t.totalSegments = 200644
	t.segments = 200644
	setList("", t.transfer)
	t.transfer.Border = false
	for i := 0; i < 10; i++ {
		t.transfer.Rows = append(t.transfer.Rows, "")
	}
	t.transfer.Rows = append(t.transfer.Rows, "        --------->")

	setList("Destination (buffer)", t.destinationBuffer)
	setList("Destination (program)", t.destinationProgram)

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(1.0,
			ui.NewCol(0.33, t.source),
			ui.NewCol(0.33, t.transfer),
			ui.NewCol(0.33,
				ui.NewRow(0.5, t.destinationBuffer),
				ui.NewRow(0.5, t.destinationProgram),
			),
		),
	)

	ui.Render(grid)
	uiEvents := ui.PollEvents()
	wireSpeedTicker := time.NewTicker(time.Millisecond * time.Duration(t.wireSpeed)).C
	processSpeedTicker := time.NewTicker(time.Millisecond * time.Duration(t.processSpeed)).C
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
			}
		case <-wireSpeedTicker:
			t.advanceTransfer()
		case <-processSpeedTicker:
			t.readBuffer()
		}
		ui.Render(grid)
	}
}

func (t *TransferWithOptions) advanceTransfer() {
	mutex.Lock()
	defer mutex.Unlock()
	if len(t.destinationBuffer.Rows) == t.bufferSize {
		return
	}
	t.segments--
	percent := (float64(t.segments) / float64(t.totalSegments)) * 100.0
	t.source.Rows[5] = fmt.Sprintf("%d %.2f%%", t.segments, percent)
	t.destinationBuffer.Rows = append([]string{fmt.Sprintf("Seg %03d", t.step)}, t.destinationBuffer.Rows...)
	fullString := ""
	if len(t.destinationBuffer.Rows) == t.bufferSize {
		fullString = "FULL"
	}
	t.destinationBuffer.Title = fmt.Sprintf("Size %d %s", len(t.destinationBuffer.Rows), fullString)
	t.step++
}

func (t *TransferWithOptions) readBuffer() {
	mutex.Lock()
	defer mutex.Unlock()
	item := ""
	if len(t.destinationBuffer.Rows) > 0 {
		item = t.destinationBuffer.Rows[len(t.destinationBuffer.Rows)-1]
		t.destinationBuffer.Rows = t.destinationBuffer.Rows[0 : len(t.destinationBuffer.Rows)-1]
		t.destinationBuffer.Title = fmt.Sprintf("Size %d", len(t.destinationBuffer.Rows))
	}
	if item != "" {
		t.destinationProgram.Rows = append([]string{item}, t.destinationProgram.Rows...)
	}
	t.destinationProgram.Title = fmt.Sprintf("Size %d", len(t.destinationProgram.Rows))
}
