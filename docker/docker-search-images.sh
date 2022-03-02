#!/bin/bash

#docker search
name=$1
docker search $name

##docker search with tag
#docker search $name |grep ./docker-search-tag.sh
