package main

import "math/rand"
import "time"

const MaxInt int = 2147483648

func (from Window) distanceScore(to Window, xdir int, ydir int) int {
	// If `to` is mostly in the correct direction, return the square distance between `from` and `to`
	// Otherwise, return MaxInt
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
		return MaxInt
	}
}


func (w Window) getNextBy (windows *[]Window, xdir int, ydir int) (Window) {
	minSeen := MaxInt
	var ret Window = w
	for _, other := range *windows {
		if w.Desktop == other.Desktop {
			f := w.distanceScore(other, xdir, ydir)
			if f < minSeen {
				minSeen = f
				ret = other
			}
		}
	}
	return ret
}


func (w Window) getRandomOverlap(windows *[]Window) Window {
	choices := make([]Window, len(*windows))
	i := 0
	for _, wi := range *windows {
		if w != wi && w.HighOverlap(wi)  {
			choices[i] = wi
			i += 1
		}
	}
	if i == 0 {
		return w
	} else {
		rand.Seed(int64(time.Now().Nanosecond()))
		return choices[rand.Intn(i)]
	}
}