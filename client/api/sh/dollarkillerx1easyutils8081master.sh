
#! /bin/sh

lsof -i :8081 | awk '{print $2}'> tmp
pid=$(awk 'NR==2{print}' tmp);
kill -9 $pid

# 如果文件不存在 git clone 
if [ ! -d "/home/s/easyutils" ];then
	cd /home/s
	git clone -b master  https://github
	cd easyutils
	./api &	
else
# 如果文件存在
	cd /home/s/easyutils
	git pull
	./api &

fi
