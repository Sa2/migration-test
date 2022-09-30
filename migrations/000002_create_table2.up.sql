BEGIN;
CREATE TABLE IF NOT EXISTS users2(
    id serial PRIMARY KEY,
    username VARCHAR (50) UNIQUE NOT NULL,
    age INT NOT NULL
);
CREATE INDEX users2_age on users2 (age);
COMMIT;