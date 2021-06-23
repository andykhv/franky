CREATE TABLE IF NOT EXISTS users.users_by_email_password (
    email text,
    password text,
    apiKey text,
    id text,
    creationDate timestamp,
    PRIMARY KEY ((email, password))
);