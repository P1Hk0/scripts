currentLayout=$(setxkbmap -print | grep xkb_symbols | sed -r 's/.*\+(..)\+.*/\1/')
uskey="us"
dekey="de"
echo $currentLayout
if [[ $currentLayout == $uskey ]] 
then
    echo "change to de"
    setxkbmap -layout de
elif [[ $currentLayout == $dekey ]]
then
    echo "change to us"
    setxkbmap -layout us
fi
