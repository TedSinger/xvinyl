# xvinyl 
A utility for a non-tiling window manager to fake some tiling actions

Motivation:
I am not hardcore enough to use a *real* tiling window manager, but I like laying out my windows in a simple grid and switching by using arrow keys. 

Call:

`$ xvinyl left`

and the next window over to the left will be focused, raised, and activated. Similar with right, up, or down.

I have keybindings in my desktop environment:
* Super+Left -> xvinyl left
* Super+Right -> xvinyl right
* Super+Down -> xvinyl down
* Super+Up -> xvinyl up

Requires `wmctrl` and `xdotool`

TODO:
* Include current visibility status in selection criteria
* Add placement of windows
* Add switching between colocated windows