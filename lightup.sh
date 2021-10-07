originBri=`cat /sys/class/backlight/intel_backlight/brightness`
maxBri=$(cat /sys/class/backlight/intel_backlight/max_brightness)
currentBri=$((`cat /sys/class/backlight/intel_backlight/brightness`+ 43))
if [ $currentBri -ge $maxBri ] ; then
    echo $maxBri > /sys/class/backlight/intel_backlight/brightness
    echo "Current brightness : $currentBri"
elif [ $currentBri -le $maxBri ] ; then
    echo $currentBri > /sys/class/backlight/intel_backlight/brightness
    echo "Current brightness : $originBri"
fi
~
~

