#!/bin/sh
cp -ra /datas/* /app
rm -r /datas
ln -s /app /datas
/app/main