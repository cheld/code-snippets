#!/bin/bash
  

  if [[ $(docker ps | grep "dashboard" | awk '{print $1}' | wc -l) != 0 ]]; then
    docker rm -f $(docker ps | grep "dashboard" | awk '{print $1}')
  fi

  if [[ $(docker ps | grep "heapster" | awk '{print $1}' | wc -l) != 0 ]]; then
    docker rm -f $(docker ps | grep "heapster" | awk '{print $1}')
  fi

  if [[ $(docker ps | grep "logstash" | awk '{print $1}' | wc -l) != 0 ]]; then
    docker rm -f $(docker ps | grep "logstash" | awk '{print $1}')
  fi
