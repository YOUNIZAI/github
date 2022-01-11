# !/bin/bash

filedir="./"
srcStr=$1
dstStr=$2
if [ "$srcStr" == "" ];then
	echo "cmd: ./xxx.sh srcStr dstStr  dir"
	echo "e.g: ./grep-rl-sed.sh wgh1  wgh2 ./"
else

sed -i "s#$srcStr#$dstStr#g" `grep  "$srcStr" -rl $filedir`

fi


