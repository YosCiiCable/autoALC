#main image
FROM ubuntu:latest
RUN apt update -y 

# - - -

#dev kit image
FROM anatolelucet/neovim
FROM purefish/docker-fish
#RUN touch ~/.config/fish/config.fish
#RUN sed -i 2a'if status is-interactive\n    \# Commands to run in interactive sessions can go here\n	~/.config/fish/config.fish oh-my-posh init fish --config "/mnt/c/1system/fishBaseTheme.omp.json" | .\nend' ~/.config/fish/config.fish
