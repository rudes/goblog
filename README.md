# goblog

Nofrills blog written in go with an open api for posting (so we can post form other things)

# Installation

Before you run anything, edit the config
```bash
cp static/example.toml static/config.toml
$EDITOR static/config.toml
```

Set environment to prod
```bash
sed -i '0,/dev/s/dev/prod/' db/scripts/load-db.sh
```

Also set a volume for persisting data
```bash
mkdir -p /var/lib/cassandra
tee docker-compose.yml <<-'EOF'
    volumes:
	- /var/lib/cassandra:/var/lib/cassandra
EOF
```

goblog is best served with [Docker Compose](https://github.com/docker/compose)
```bash
docker-compose up -d --build
```
