#!/bin/bash

bash cassandra

sleep 10

cqlsh -f /tmp/schema/letters.cql
if [ "$LETTERS" -eq "dev" ];then
    cqlsh -f /tmp/load/letters.cql
fi

pkill -f CassandraDaemon
