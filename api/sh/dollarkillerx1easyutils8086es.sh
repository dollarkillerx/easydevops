
#! /bin/sh

lsof -i :8086 | awk '{print $2}'> tmp
pid=$(awk 'NR==2{print}' tmp);
kill -9 $pid

cd /asdasd/asdasdasd

git pull https://github.sads


./dollarki &
