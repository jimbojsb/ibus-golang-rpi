#!/bin/bash
pacman -Sy go git mpd dhcp bridge-utils vim
hostnamectl set-hostname propellerhead.local
ln -s /usr/bin/mpd /usr/local/bin/mpd
cd /tmp
curl -O https://aur.archlinux.org/packages/pa/package-query/package-query.tar.gz
tar zxvf package-query.tar.gz
cd package-query
makepkg -si
cd ..
curl -O https://aur.archlinux.org/packages/ya/yaourt/yaourt.tar.gz
tar zxvf yaourt.tar.gz
cd yaourt
makepkg -si
makepkg -si
makepkg -si --asroot
rm -fR package-query*
rm -fR yaourt*
yaourt hostapd-8192cu
mkdir /music

#https://www.dropbox.com/s/sal8knn5olomard/shairport.gz?dl=1
#https://www.dropbox.com/s/1sdui60bfqdsx8j/hostapd.gz?dl=1
#https://www.dropbox.com/s/zkfgexo87rrouwi/cp210x.ko.gz?dl=1