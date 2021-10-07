originBri=`cat /sys/class/backlight/intel_backlight/brightness`
maxBri=$(cat /sys/class/backlight/intel_backlight/max_brightness)
currentBri=$((`cat /sys/class/backlight/intel_backlight/brightness`- 43))
if [ $currentBri -ge 0 ] ; then
    echo $currentBri > /sys/class/backlight/intel_backlight/brightness
    echo "Current brightness : $currentBri"
elif [ $currentBri -le 0 ] ; then
    echo 10 > /sys/class/backlight/intel_backlight/brightness
    echo "Current brightness : $originBri"
fi
