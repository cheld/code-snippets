#!/bin/bash


  # kill & remove all containers
  if [[ $(docker ps | grep "nsenter" | awk '{print $1}' | wc -l) != 0 ]]; then
    docker rm -f $(docker ps | grep "nsenter" | awk '{print $1}')
  fi
  if [[ $(docker ps | grep "k8s_" | awk '{print $1}' | wc -l) != 0 ]]; then
    docker rm -f $(docker ps | grep "k8s_" | awk '{print $1}')
  fi

  if [[ -d /var/lib/kubelet ]]; then
        # umount if there are mounts in /var/lib/kubelet
        if [[ ! -z $(mount | grep "/var/lib/kubelet" | awk '{print $3}') ]]; then
          # The umount command may be a little bit stubborn sometimes, so run the commands twice to ensure the mounts are gone
          mount | grep "/var/lib/kubelet/*" | awk '{print $3}' | xargs umount 1>/dev/null 2>/dev/null
          mount | grep "/var/lib/kubelet/*" | awk '{print $3}' | xargs umount 1>/dev/null 2>/dev/null
          umount /var/lib/kubelet 1>/dev/null 2>/dev/null
          umount /var/lib/kubelet 1>/dev/null 2>/dev/null
        fi
        # Delete the directory
        rm -rf /var/lib/kubelet
  fi

  # delete k8s states & configuration
  if [[ -d /var/run/kubernetes ]]; then
    rm -rf /var/run/kubernetes
  fi

  if [[ -d /etc/kubernetes ]]; then
    rm -rf /etc/kubernetes
  fi

  if [[ -d /etc/cni ]]; then
    rm -rf /etc/cni
  fi

