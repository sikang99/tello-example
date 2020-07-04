#!/bin/bash - 
#===============================================================================
#
#          FILE: wifi.sh
# 
#         USAGE: ./wifi.sh 
# 
#   DESCRIPTION: 
# 
#       OPTIONS: ---
#  REQUIREMENTS: ---
#          BUGS: ---
#         NOTES: ---
#        AUTHOR: bsdpunk, Dusty Carver
#  ORGANIZATION: 
#       CREATED: 04/04/2020 23:37
#      REVISION:  ---
#===============================================================================

WLAN="wlp0s20f3"


    while [ -z "$TELLO" ]; do 
    
sudo ifconfig $WLAN up
    TELLO=$(sudo iwlist scan | egrep -o "TELLO-[0-9A-F]+" )
    echo $TELLO
done
echo $WLAN
sudo nmcli device wifi connect $TELLO
sudo dhclient $WLAN
mv this.flv new.flv
#./main &
#./video &
#./minus ~/Downloads/res10_300x300_ssd_iter_140000.caffemodel ~/Downloads/deploy.prototxt
#./key ~/Downloads/res10_300x300_ssd_iter_140000.caffemodel ~/Downloads/deploy.prototxt
./key ~/Downloads/res10_300x300_ssd_iter_140000.caffemodel ~/Downloads/deploy.prototxt
