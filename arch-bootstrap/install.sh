#!/bin/sh
mkdir /propellerhead
cd /propellerhead
mkdir bin
mkdir mpd
mkdir shairport
cd ~/propellerhead
go build src/propellerhead.go
cp propellerhead /propellerhead
cd ~/shairport
make
cp shairport /propellerhead