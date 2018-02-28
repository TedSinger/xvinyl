package main

import (
	"os/exec"
	"strings"
	"strconv"
	"fmt"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil/ewmh"
	"github.com/BurntSushi/xgbutil"
)

func GetActiveWid(X *xgbutil.XUtil) xproto.Window {
	w, _ := ewmh.ActiveWindowGet(X)
	return w
}

type Window struct {
	Wid xproto.Window

	Desktop int
	Xmin  int
	Width int
	Xmax  int
	Xmid  int

	Ymin   int
	Height int
	Ymax   int
	Ymid   int
}


func (w Window) Select(X *xgbutil.XUtil) {
	ewmh.ActiveWindowReq(X, w.Wid)
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
	desktop, err1 := strconv.Atoi(components[1])
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
		w := Window{xproto.Window(wid), desktop, xmin, width, xmin + width, xmin + width/2, ymin, height, ymin + height, ymin + height/2}
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

	windows := make([]Window, 0, len(lines))
	for _, line := range lines {
		windows = append(windows, makeWindow(line)...)
	}
	return windows
}

func GetActiveWindow(wid xproto.Window, windows *[]Window) ([]Window) {
	for _, w := range *windows {
		if w.Wid == wid {
			return []Window{w}
		}
	}
	fmt.Println("Could not find the active window with WID=" + strconv.Itoa(int(wid)))
	return []Window{}
}
