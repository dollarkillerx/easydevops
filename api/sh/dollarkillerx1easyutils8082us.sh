
#! /bin/sh

lsof -i :8082 | awk '{print $2}'> tmp
pid=$(awk 'NR==2{print}' tmp);
kill -9 $pid

# 如果文件不存在 git clone 
if [ ! -d "/asdasd/asdasdasd/us/us" ];then
	cd /asdasd/asdasdasd/us
	git clone -b us  https://github.sads
	cd us
	cd /asdsaadsad
	./dollarki &	
else
# 如果文件存在
	cd /asdasd/asdasdasd/us/us
	git pull
	cd /asdsaadsad
	./dollarki &

fi
