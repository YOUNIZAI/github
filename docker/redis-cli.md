#!/bin/bash
nameid=$(./dcomp ps |awk '{print $1}' |grep redis)
ip=$(docker inspect $nameid |grep IPAdd)
redis-cli $ip




