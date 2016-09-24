FROM spotify/cassandra:latest

COPY schema /tmp/schema
COPY load /tmp/load
COPY scripts/load-db.sh /usr/local/bin/load-db.sh

RUN chmod +x /usr/local/bin/load-db.sh
RUN load-db.sh

ENTRYPOINT cassandra-singlenode
