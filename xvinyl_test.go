package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestMakeWindow1(t *testing.T) {
	line := "0x04800003 -1 1725   0    0    1920 1080 ted-QAL51 Desktop"
	assert.Equal(t, len(makeWindow(line)), 0, "Windows with gravity -1 should be ignored")
}

func getTestWindow(t *testing.T) Window {
	line := "0x04600003  0 5567   990  58   939  1054 ted-QAL51 ~/go/src/github.com/TedSinger/xvinyl/window.go - Sublime Text"
	ws := makeWindow(line)
	assert.Equal(t, len(ws), 1, "Did not parse Window line")
	return ws[0]
}

func TestMakeWindow2(t *testing.T) {
	w := getTestWindow(t)
	assert.Equal(t, w.Wid, 73400323, "Wrong Wid")
	assert.Equal(t, w.Xmin, 990, "Wrong Xmin")
	assert.Equal(t, w.Width, 939, "Wrong Width")
	assert.Equal(t, w.Xmax, 990 + 939, "Wrong Xmax")
	assert.Equal(t, w.Xmid, 990 + 939 / 2, "Wrong Xmid")
	assert.Equal(t, w.Ymin, 58, "Wrong Ymin")
	assert.Equal(t, w.Height, 1054, "Wrong Height")
	assert.Equal(t, w.Ymax, 58 + 1054, "Wrong Ymax")
	assert.Equal(t, w.Ymid, 58 + 1054 / 2, "Wrong Ymid")
}