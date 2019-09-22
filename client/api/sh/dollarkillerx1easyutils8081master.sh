
#! /bin/sh

lsof -i :8081 | awk '{print $2}'> tmp
pid=$(awk 'NR==2{print}' tmp);
kill -9 $pid

cd /home/s/easydevops
./api
