package widgets

import (
	"fmt"
	termuiw "github.com/gizak/termui/v3/widgets"
	"github.com/shirou/gopsutil/v3/process"
	"gonitor/util"
	"log"
	"sort"
	"time"
)

type SortType string

const (
	PIDSort = "pid"
	CPUSort = "cpu"
	MemSort = "mem"
)

var sortMethod = CPUSort

type ProcessInfo struct {
	Process    *process.Process
	Name       string
	CPUUsed    float64
	MemoryUsed float32
}

type ProcessList []ProcessInfo

type ProcessListWidget struct {
	*termuiw.List
	interval     time.Duration
	processInfos ProcessList
}

func (l ProcessList) Len() int {
	return len(l)
}
func (l ProcessList) Less(i, j int) bool {
	switch sortMethod {
	case PIDSort:
		return l[i].Process.Pid < l[j].Process.Pid
	case MemSort:
		return l[i].MemoryUsed < l[j].MemoryUsed
	case CPUSort:
		return l[i].CPUUsed < l[j].CPUUsed
	}
	// Unreachable
	return l[i].CPUUsed < l[j].CPUUsed
}

func (l ProcessList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (w *ProcessListWidget) init() {
	w.Title = "Processes"
	w.TitleStyle = util.ClearBold

	if w.interval > 0 {
		go func() {
			for range time.NewTicker(w.interval).C {
				w.update()
			}
		}()
	}
}

func FetchProcessInfos() []ProcessInfo {
	var processInfos []ProcessInfo
	processes, err := process.Processes()
	if err != nil {
		log.Fatalf("Error getting process list: %v", err)
	}
	for _, proc := range processes {
		name, err := proc.Name()
		if err != nil {
			name = "?????"
		}
		cpu, err := proc.CPUPercent()
		if err != nil {
			cpu = 0.0
		}
		mem, err := proc.MemoryPercent()
		if err != nil {
			mem = 0.0
		}
		processInfos = append(processInfos, ProcessInfo{proc, name, cpu, mem})
	}
	sort.Sort(sort.Reverse(ProcessList(processInfos)))
	return processInfos
}

func (w *ProcessListWidget) update() {
	width := w.Block.Dx()
	nameWidth := width - (8 + 5 + 5 + 5 + 2 + 2) // length of each field (but name) + spaces + border + %

	w.processInfos = FetchProcessInfos()
	var rows []string
	rows = append(rows, fmt.Sprintf("%-8s %-*s %6s %6s  ",
		"PID", nameWidth, "Name", "CPU", "MEM"))
	for _, proc := range w.processInfos {
		rows = append(rows, fmt.Sprintf(
			"%8d %-*s %5.1f%% %5.1f%%  ",
			proc.Process.Pid,
			nameWidth,
			proc.Name,
			proc.CPUUsed,
			proc.MemoryUsed,
		))
	}
	w.Rows = rows
}

func NewProcessListWidget(interval time.Duration) *ProcessListWidget {
	widget := ProcessListWidget{
		List:         termuiw.NewList(),
		interval:     interval,
		processInfos: FetchProcessInfos(),
	}
	widget.init()
	return &widget
}
