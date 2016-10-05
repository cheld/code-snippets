#!/bin/bash

service monasca-api status
service monasca-log-api status
service zookeeper status
service kafka status
service monasca-thresh status
service monasca-notofication status
service monasca-persister status
service monasca-log-transformer status
service monasca-log-persister status
service influxdb status
service elasticsearch status
service grafana-server status
service kibana status
service monasca-agent status
service monasca-log-agent status

service --status-all
