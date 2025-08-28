#!/bin/bash

PROGRAM=$1
EXC=$2

while :
do
    ps aux | grep -v "grep" | grep "$PROGRAM" | grep -vE "$EXC"
    sleep 1s
done
