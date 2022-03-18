#!/bin/sh

if [ "`ls -A /conf`" = "" ];
then 
        echo "检测到 /conf 目录为空，准备从 /datas 目录复制初始文件"
        cp -ra /datas/conf/* /conf
fi
if [ "`ls -A /img`" = "" ];
then 
        echo "检测到 /img 目录为空，准备从 /datas 目录复制初始文件"
        cp -ra /datas/img/* /img
fi
/Gin_HomeNavigation