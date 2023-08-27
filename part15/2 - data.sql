INSERT INTO movies.studios (id, name) VALUES (1, 'Columbia Pictures');
INSERT INTO movies.studios (id, name) VALUES (2, 'Paramount Pictures');
ALTER SEQUENCE movies.studios_id_seq RESTART WITH 100;

INSERT INTO movies.roles (id, name) VALUES (1, 'Actor');
INSERT INTO movies.roles (id, name) VALUES (2, 'Director');
ALTER SEQUENCE movies.roles_id_seq RESTART WITH 100;

INSERT INTO movies.participants (id, name, birthday) VALUES (1, 'Alen Delon', '01-01-1950');
INSERT INTO movies.participants (id, name, birthday) VALUES (2, 'Bruce Lee', '01-01-1950');
INSERT INTO movies.participants (id, name, birthday) VALUES (3, 'James Cameron', '01-01-1950');
ALTER SEQUENCE movies.participants_id_seq RESTART WITH 100;

INSERT INTO movies.movies (id, name, year, fee, studio_id, rank) VALUES (1, 'Movie name', 1950, 1000, 1, 'PG-10');
INSERT INTO movies.movies (id, name, year, fee, studio_id, rank) VALUES (2, 'Film name', 1950, 2000, 2, 'PG-13');
INSERT INTO movies.movies (id, name, year, fee, studio_id, rank) VALUES (3, 'Cinema name', 1950, 3000, 2, 'PG-18');
ALTER SEQUENCE movies.movies_id_seq RESTART WITH 100;


INSERT INTO movies.movies_participants (movie_id, participant_id, role_id) VALUES (1, 1, 1);
INSERT INTO movies.movies_participants (movie_id, participant_id, role_id) VALUES (2, 2, 1);
INSERT INTO movies.movies_participants (movie_id, participant_id, role_id) VALUES (3, 1, 1);
INSERT INTO movies.movies_participants (movie_id, participant_id, role_id) VALUES (3, 3, 2);
ALTER SEQUENCE movies.movies_participants_id_seq RESTART WITH 100;