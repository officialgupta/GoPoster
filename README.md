# GoPoster

An application built for the Raspberry Pi that fetches the latest movie posters and displays them. Meant to be a digital poster display.

## Installation

* Install [raspian lite](https://www.raspberrypi.org/documentation/installation/installing-images/) on raspberry pi

* Change the default password on the default Pi account ([GUIDE HERE](https://www.raspberrypi.org/documentation/linux/usage/users.md))

* Ensure network connection is active, either by connecting an ethernet cable or through [WiFI](https://www.raspberrypi.org/documentation/configuration/wireless/README.md)

* Enable autologin to skip the login on boot
    * Run: `sudo raspi-config`
    * Choose option 3: Boot Options
    * Choose option B1: Desktop / CLI
    * Choose option B2: Console Autologin
    * Select Finish, and reboot the pi.

* If you need to access the pi remotely (which shouldn't be necessary) enable [SSH](https://www.raspberrypi.org/documentation/remote-access/ssh/README.md)

* Disable screen blanking so that the Raspberry Pi’s display doesn’t go to sleep
    * type `sudo nano /boot/cmdline.txt`
    * verify that the following line exists: `consoleblank=0`
    * if the line does not exist, add it

* Install the photo viewer (FBI), run `sudo apt-get update` then `sudo apt-get install fbi`

* Rotate the display output for the raspberry pi
    * type `sudo nano /boot/config.txt`
    * Change `display_rotate=0` to 1, 2, or 3 depending on which way your display is rotated
    * display_rotate=1 (this will rotate your display 90 degrees)
    * display_rotate=2 (this will rotate your display 180 degrees)
    * display_rotate=3 (this will rotate your display 270 degrees)
    * If there is no entry for “display_rotate”, you can add it

* At this point you can clone this repository containing the scripts required into the raspberry pi by running `git clone https://github.com/officialgupta/GoPoster.git` in the base directory (/home/pi)

* You will need a TMDB API key to fetch the appropriate posters, to do so you must create a TMDB account and then navigate to the [API Page](https://www.themoviedb.org/settings/api) to obtain the API key.
It is **very important** that you place this api key into a file named "key" in the base directory (/home/pi).

* Make the project run automatically on startup
    * Change to your home directory by typing `cd ~`
    * type `sudo nano .bashrc`
    * add a line to the end of the file that reads ./PosterPi.sh

* You must schedule the raspberry pi to fetch the latest posters. To schedule the poster to refresh *every day* type `0 0 * * * /home/pi/getposters.sh`. You can use an [alternate cron entry](https://www.raspberrypi.org/documentation/linux/usage/cron.md) for a different interval.


## Configuration

* The amount of time a poster displays for can be changed in the posterpi.sh file by changing the -t parameter within the fbi call. `fbi -t <TIME IN SECONDS> -a --noverbose -l /home/pi/PosterPi.txt`

## Notes

* You may need to make the shell files executable using `chmod +x <filename>` 

* If you change the Go code you will need to rebuild it in a linux environment, the easiest way is to install go on the raspberry pi using `sudo apt-get install golang` and then running `go build`

* The photo display elements of this project were adapted from Rob O'Hara's PosterPi, the writeup and some extra tips which can be found [at his blog post](http://www.robohara.com/?p=12528#two).
