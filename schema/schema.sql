 CREATE TABLE authors (
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL,
  bio  text not NULL,
  uid BIGSERIAL references users(userid)
);

CREATE TABLE users (
  userid BIGSERIAL PRIMARY KEY,
  username text not null,
  pass text not null
);