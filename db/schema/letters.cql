CREATE KEYSPACE letters WITH REPLICATION = {
    'class' : 'SimpleStrategy',
    'replication_factor' : 1
};

CREATE TABLE letters.letters (
    "id" varchar,
    "title" varchar,
    "content" varchar,
    "date" varchar,
    "time" varchar,
    PRIMARY KEY("id")
);
