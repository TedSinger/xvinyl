package main

// import "fmt"

func (from Window) nearness(to Window, xdir int, ydir int) int {
	ydiff := to.Ymid - from.Ymid
	xdiff := to.Xmid - from.Xmid
	// TODO: scale directions by aspect ratio
	dot := xdir*xdiff + ydir*ydiff
	other := xdir*ydiff + -ydir*xdiff
	if other < 0 {
		other = -other
	}
	if dot > other {
		return ydiff*ydiff + xdiff*xdiff
	} else {
		return 9999999
	}
}


func (w Window) getNextBy (windows *[]Window, xdir int, ydir int) (Window) {
	minSeen := 9999998
	var ret Window
	for _, other := range *windows {

		f := w.nearness(other, xdir, ydir)
		// fmt.Println(w, f)
		if f < minSeen {
			minSeen = f
			ret = other
		}
	}
	return ret
	
}