-- DROP SCHEMA IF EXISTS movies CASCADE;
-- CREATE SCHEMA movies;

DROP TABLE IF EXISTS movies.movies_participants;
DROP TABLE IF EXISTS movies.participants;
DROP TABLE IF EXISTS movies.movies;
DROP TABLE IF EXISTS movies.studios;
DROP TABLE IF EXISTS movies.roles;
DROP TYPE IF EXISTS movies.ranks;

CREATE TABLE movies.roles (    id SERIAL PRIMARY KEY,
    name varchar(200) NOT NULL DEFAULT '',
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    UNIQUE(name)
);

CREATE TABLE movies.participants (
    id SERIAL PRIMARY KEY,
    name varchar(200) NOT NULL DEFAULT '',
    birthday date NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    UNIQUE(name)
);

CREATE TYPE movies.ranks AS ENUM ('PG-10', 'PG-13', 'PG-18');

CREATE TABLE movies.studios (
  id SERIAL PRIMARY KEY,
  name varchar(200) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  UNIQUE(name)
);

CREATE TABLE movies.movies (
    id BIGSERIAL PRIMARY KEY,
    name varchar(200) NOT NULL,
    year INTEGER NOT NULL DEFAULT 0,
    fee INTEGER DEFAULT 0,
    studio_id INTEGER REFERENCES movies.studios(id) ON DELETE CASCADE ON UPDATE CASCADE DEFAULT 0,
    rank ranks NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    UNIQUE(name, year),
    CONSTRAINT year_not_less_1800 check (year >= 1800)
);

CREATE TABLE movies.movies_participants (
    id BIGSERIAL PRIMARY KEY,
    movie_id BIGINT NOT NULL REFERENCES movies.movies(id),
    participant_id INTEGER NOT NULL REFERENCES movies.participants(id),
    role_id INTEGER NOT NULL REFERENCES movies.roles(id),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    UNIQUE(movie_id, participant_id, role_id)
);
