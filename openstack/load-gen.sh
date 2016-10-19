#!/bin/bash

while true; do 
  curl localhost:3010/cpu-short
  sleep `echo $RANDOM % 100 | bc`
done
