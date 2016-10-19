#!/bin/bash

while true; do 
  curl localhost:3010/cpu-short
  sleep `echo $RANDOM % 200 + 30 | bc`
done
