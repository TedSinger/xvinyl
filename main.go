package main


import "fmt"
import "os"

func main() {
	if len(os.Args) != 2 {
		printUsage()
		return
	}
	windows := GetWindows()
	if len(windows) == 0 {
		fmt.Println("Unrecoverable error - couldn't find any Windows")
		return
	}
	wid := GetActiveWid()
	if wid == 0 {
		fmt.Println("Unrecoverable error - couldn't determine the active WID")
		return
	}

	maybeActive := GetActiveWindow(wid, &windows)
	if len(maybeActive) == 1 {
		var nextWindow Window
		active := maybeActive[0]
		switch {
		case os.Args[1] == "left":
			nextWindow = active.getNextBy(&windows, -1, 0)
		case os.Args[1] == "right":
			nextWindow = active.getNextBy(&windows, 1, 0)
		case os.Args[1] == "up":
			nextWindow = active.getNextBy(&windows, 0, -1)
		case os.Args[1] == "down":
			nextWindow = active.getNextBy(&windows, 0, 1)
		case true:
			printUsage()
			return
		}
		nextWindow.Select()
	} else {
		fmt.Println("Unrecoverable error - couldn't find any Window with the active WID")
	}
}

func printUsage () () {
	fmt.Println("usage: xvinyl <direction>\n        left, right, up, or down")
}