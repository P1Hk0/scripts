#!/bin/bash

bash ./dwm-status.sh &
bash ./chKeymap.sh
bash ./wp-change.sh
# /bin/bash ~/scripts/wp-autochange.sh &
#picom -o 0.95 -i 0.88 --detect-rounded-corners --vsync --blur-background-fixed -f -D 5 -c -b
# picom -b
bash ./tap-to-click.sh &
bash ./inverse-scroll.sh &
nm-applet &
# xfce4-power-manager &
#xfce4-volumed-pulse &
# bash ./run-mailsync.sh &
bash ./autostart_wait.sh &
