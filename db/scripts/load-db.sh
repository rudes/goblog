#!/bin/bash

bash cassandra

sleep 10

cqlsh -f /tmp/schema/letters.cql
ENVI="dev"
if [  "$ENVI" = "dev" ];then
    cqlsh -f /tmp/load/letters.cql
fi

pkill -f CassandraDaemon
