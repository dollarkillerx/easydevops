
#! /bin/sh

lsof -i :8086 | awk '{print $2}'> tmp
pid=$(awk 'NR==2{print}' tmp);
kill -9 $pid

# 如果文件不存在 git clone 
if [ ! -d "/asdasd/asdasdasd/easyutils" ];then
	cd /asdasd/asdasdasd
	git clone -b es  https://github.sads
	cd easyutils
	./dollarki &	
else
# 如果文件存在
	cd /asdasd/asdasdasd/easyutils
	git pull
	./dollarki &

fi
