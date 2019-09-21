#! /bin/sh

kill -9 $(grep webserver)

cd ~/newweb/

git pull https:/github.com/sdsd

cd webserver

./webserver &
