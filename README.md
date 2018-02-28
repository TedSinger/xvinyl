# xvinyl 
A window-switcher utility for a non-tiling window manager to fake some tiling actions

### Motivation:
I am not hardcore enough to use a *real* tiling window manager, but I like laying out my windows in a simple grid and switching by using arrow keys. 

## Usage:

Call:

`$ xvinyl left`

and the next window over to the left will be focused and raised. Likewise with right, up, or down.

I have keybindings in my desktop environment:
* Super+Left -> xvinyl left
* Super+Right -> xvinyl right
* Super+Down -> xvinyl down
* Super+Up -> xvinyl up

`$ xvinyl overlap`

will select a random window that overlaps the current window. I bind this to Alt+Tab. This works for me since I don't often put windows hiding behind each other, except as an intermediate state. If you tend to keep windows stacked on top of each other, maybe don't use this function.

Requires `wmctrl`

## Build:

$ go get github.com/TedSinger/xvinyl

$ cd ~/go/src/github.com/TedSinger/xvinyl

$ vgo build
