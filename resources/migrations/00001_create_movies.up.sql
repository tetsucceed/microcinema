CREATE TABLE IF NOT EXISTS movias (
   movie_id serial PRIMARY KEY,
   movie_name VARCHAR (50) UNIQUE NOT NULL,
   movie_poster VARCHAR (50) UNIQUE NOT NULL,
   movie_url VARCHAR (50) UNIQUE NOT NULL,
   movie_paid boolean NOT NULL DEFAULT FALSE,
   movie_release_date DATE DEFAULT NULL,
   movie_genre VARCHAR (50) UNIQUE NOT NULL
);

INSERT INTO movias (movie_id, movie_name, movie_poster, movie_url, movie_paid, movie_release_date, movie_genre)
VALUES
(1, "Бойцовский клуб", "/static/posters/fightclub.jpg", "https://youtu.be/qtRKdVHc-cE", TRUE, DATE("1999"), "triller"),
(2, "Крестный отец", "/static/posters/father.jpg", "https://youtu.be/ar1SHxgeZUc", FALSE, DATE("1988"), "drama"),
(3, "Криминальное чтиво", "/static/posters/pulpfiction.jpg","https://youtu.be/s7EdQ4FqbhY", TRUE, DATE("1996"), "comedy");
