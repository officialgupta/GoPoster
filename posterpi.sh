#!/bin/bash

#Start Loop Process
while true; do

#Delete old files
rm /home/pi/PosterPi.txt
rm /home/pi/PosterPi.tmp

#Find all jpg images, inc subdirs, add to PosterPi.txt
find /home/pi/posters -name "*jpg" > /home/pi/PosterPi.txt

#Launch FBI viewer, read from PosterPi.tmp
fbi -t 300 -a --noverbose -l /home/pi/PosterPi.txt

#When FBI reaches the end of the list, begin again.
clear
done