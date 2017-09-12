# xvinyl 
A utility for a non-tiling window manager to fake some tiling actions

Call:

`$ xvinyl left`

and the next window over to the left will take focus. Similar with right, up, or down.

I've set keybindings in my desktop environment:
* Super+Left -> xvinyl left
* Super+Right -> xvinyl right
* Super+Down -> xvinyl down
* Super+Up -> xvinyl up

Requires `wmctrl` and `xdotool`

TODO:
address shameful lack of testing