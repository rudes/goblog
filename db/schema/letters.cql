CREATE KEYSPACE letters WITH REPLICATION = {
    'class' : 'SimpleStrategy',
    'replication_factor' : 1
};

CREATE TABLE letters.letters (
    "id" text,
    "title" text,
    "content" text,
    "date" timestamp,
    PRIMARY KEY("id")
);
