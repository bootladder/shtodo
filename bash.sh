#!/bin/bash

# Put the following line in your .bashrc
# source ~/path/to/shtodo/bash.sh

# Execute shtodo if installed, else install
if [ -e /usr/bin/shtodo ]; then
    /usr/bin/shtodo
else
    echo Installing shtodo in /usr/bin/
    cp shtodo /usr/bin/shtodo
    /usr/bin/shtodo
fi

#Aliases 
alias vitodo='vi ~/Documents/todo.txt'
alias cattodo='cat ~/Documents/todo.txt'
pushtodo() {
  cd ~/Documents/todo/
  cp ../todo.txt .
  git add todo.txt
  git commit -m "new todo $(date)"
  git push origin master
  cd -
}
pulltodo() {
  cd ~/Documents/todo/
  git pull
  echo pwd is $(pwd)
  cp todo.txt ~/Documents/todo.txt
  cd -
}

