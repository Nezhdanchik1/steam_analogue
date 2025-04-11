CREATE TABLE IF NOT EXISTS games (
     id SERIAL PRIMARY KEY,
     title TEXT NOT NULL,
     description TEXT,
     developer TEXT,
     price NUMERIC(10, 2) NOT NULL
);
