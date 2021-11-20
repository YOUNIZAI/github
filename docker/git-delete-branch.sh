# !/bin/bash

## Usage
# 执行时输入工作路径

## Command
# 删除当前分支外的所有分支
path=$1
cd $path
git branch | xargs git branch \-d $path
