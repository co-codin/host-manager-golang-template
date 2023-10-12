#!/bin/zsh


go build -o host-manager cmd/web/*.go && ./host-manager \
-dbuser='postgres' \
-pusherHost='localhost' \
-pusherSecret='somesecret' \
-pusherKey='somekey' \
-pusherApp="1"
