-- выборка фильмов с названием студии;
Select movies.movies.name, movies.studios.name
FROM movies.movies JOIN movies.studios ON movies.studios.id = movies.movies.studio_id;

-- выборка фильмов для некоторого актёра;
Select movies.movies.name, movies.participants.name
FROM movies.movies
         JOIN movies.movies_participants ON movies.movies.id = movies.movies_participants.movie_id
         JOIN movies.participants ON movies.participants.id = movies.movies_participants.participant_id
         JOIN movies.roles ON movies.roles.id = movies.movies_participants.role_id
WHERE movies.roles.name = 'Actor' and
    movies.participants.name = 'Bruce Lee'

-- подсчёт фильмов для некоторого режиссёра;
Select count(*)
FROM movies.movies
         JOIN movies.movies_participants ON movies.movies.id = movies.movies_participants.movie_id
         JOIN movies.participants ON movies.participants.id = movies.movies_participants.participant_id
         JOIN movies.roles ON movies.roles.id = movies.movies_participants.role_id
WHERE movies.roles.name = 'Director' and
    movies.participants.name = 'James Cameron'