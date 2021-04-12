#!/usr/bin/env bash

# go get github.com/githubnemo/CompileDaemon

MODE=$1

if [ "$MODE" = "dev" ] ; then
   CompileDaemon -command="./hoax ${@:2}" \
    -exclude-dir=.git
    -color=true \
    -graceful-kill=true
else
  rm main > /dev/null
  go build main.go
  ./main ${@:1}
fi
