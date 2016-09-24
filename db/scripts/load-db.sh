#!/bin/bash

bash cassandra

sleep 10

cqlsh -f /tmp/schema/letters.cql
cqlsh -f /tmp/load/letters.cql

pkill -f CassandraDaemon
