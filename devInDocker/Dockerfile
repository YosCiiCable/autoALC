FROM ubuntu:latest
#dev
	FROM anatolelucet/neovim
	FROM purefish/docker-fish
FROM golang:1.20

# Robotgo
	RUN apt update -y &&\
		# gcc\
			apt install gcc libc6-dev -y &&\
		# x11\
			apt install libx11-dev xorg-dev libxtst-dev -y &&\
		# Bitmap\
			apt install libpng++-dev -y &&\
		# Hook\
			apt install xcb libxcb-xkb-dev x11-xkb-utils libx11-xcb-dev libxkbcommon-x11-dev libxkbcommon-dev -y &&\
		# Clipboard\
			apt install xsel xclip -y ;
