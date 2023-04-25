#!/usr/bin/env bash
(
go build -o releasebot_mac
scp releasebot_mac 2macmini:~/.local/bin/releasebot
) &

(
GOOS=linux GOARCH=amd64 go build -o releasebot_linux
scp releasebot_linux arne:~/.local/bin/releasebot
) &

wait
