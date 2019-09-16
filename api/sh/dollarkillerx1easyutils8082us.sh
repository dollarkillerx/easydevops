
#! /bin/sh

lsof -i :8082 | awk '{print $2}'> tmp
pid=$(awk 'NR==2{print}' tmp);
kill -9 $pid

cd /asdasd/asdasdasd/us

git pull https://github.sads

cd /asdsaadsad

./dollarki &
