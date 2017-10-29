package main

import "os/exec"
import "strings"
import "strconv"
import "fmt"

func GetActiveWid() int {
	c := exec.Command("xdotool", "getactivewindow")
	out, err := c.Output()
	if err != nil {
		fmt.Println("`xdotool` failed. Is it installed?")
		return 0
	}
	wid, _ := strconv.Atoi(strings.Trim(string(out), "\n"))
	return wid
}
type Window struct {
	Wid int

	Xmin  int
	Width int
	Xmax  int
	Xmid  int

	Ymin   int
	Height int
	Ymax   int
	Ymid   int
}


func (w Window) Select() {
	cmds := []string{"windowraise", "windowfocus", "windowactivate"}
	for _, c := range cmds {
		exec.Command("xdotool", c, strconv.Itoa(w.Wid)).Run()	
	}
}

func (w Window) Area() int {
	return w.Width * w.Height
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func min(a int, b int) int {
	return -max(-a,-b)
}

func (w Window) HighOverlap(other Window) bool {
	overlapThreshold := 0.48
	xmin := max(w.Xmin, other.Xmin)
	xmax := min(w.Xmax, other.Xmax)
	ymin := max(w.Ymin, other.Ymin)
	ymax := min(w.Ymax, other.Ymax)

	overlapArea := float64((xmax - xmin) * (ymax - ymin))
	return overlapArea / float64(w.Area()) > overlapThreshold || overlapArea / float64(other.Area()) > overlapThreshold
}

// as a monadic nomad, i use a list of zero-or-one element to denote Option[T]
func makeWindow(wmctrl_out_line string) []Window {
	parts := strings.Split(wmctrl_out_line, " ")
	components := make([]string, 9)
	current_component := 0
	for _, text := range parts {
		if current_component == 8 {
			if text == "" {
				components[8] += " "
			} else {
				components[8] += text
			}
		} else if text != "" {
			components[current_component] = text
			current_component += 1
		}
	}

	wid, err0 := strconv.ParseInt(strings.TrimLeft(components[0], "x0"), 16, 64)
	xmin, err1 := strconv.Atoi(components[3])
	ymin, err2 := strconv.Atoi(components[4])
	width, err3 := strconv.Atoi(components[5])
	height, err4 := strconv.Atoi(components[6])
	gravity, err5 := strconv.Atoi(components[1])
	
	problem := err0 != nil || err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil
	
	if gravity == -1 {
		// Most people have a Desktop, which *is* a window, but we don't want to switch to it
		return []Window{}
	} else if problem {
		fmt.Println("Something did not parse correctly in this line of output from `wmctrl`. This window will be ignored: ")
		fmt.Println(wmctrl_out_line)
		return []Window{}
	} else {
		w := Window{int(wid), xmin, width, xmin + width, xmin + width/2, ymin, height, ymin + height, ymin + height/2}
		return []Window{w}
	}

}

func GetWindows() []Window {
	c := exec.Command("wmctrl", "-Gpl")
	out, err := c.Output()
	if err != nil {
		fmt.Println("`wmctrl` failed. Is it installed?")
		return []Window{}
	}
	lines := strings.Split(strings.Trim(string(out), "\n"), "\n")

	// does Go have an easier way to concat things?
	windows := make([]Window, len(lines))
	i := 0
	for _, line := range lines {
		for _, w := range makeWindow(line) {
			windows[i] = w
			i += 1
		}
	}
	return windows[0:i]
}

func GetActiveWindow(wid int, windows *[]Window) ([]Window) {
	for _, w := range *windows {
		if w.Wid == wid {
			return []Window{w}
		}
	}
	fmt.Println("Could not find the active window with WID=" + strconv.Itoa(wid))
	return []Window{}
}
