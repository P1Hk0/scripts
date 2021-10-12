#!/bin/bash

# bash ./dwm-status.sh &

# maybe should modify the path in function updateBattery
#please link ./goblocks.go to ~/goblocks
cd ~/goblocks; ./goblocks & 
bash ~/scripts/chKeymap.sh
# bash ~/scripts/wp-change.sh
/bin/bash ~/scripts/wp-autochange.sh &
#picom -o 0.95 -i 0.88 --detect-rounded-corners --vsync --blur-background-fixed -f -D 5 -c -b
# picom -b
bash ~/scripts/tap-to-click.sh &
bash ~/scripts/inverse-scroll.sh &
nm-applet &
# xfce4-power-manager &
#xfce4-volumed-pulse &
# bash ./run-mailsync.sh &
bash ~/scripts/autostart_wait.sh &
