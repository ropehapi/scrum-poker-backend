create table players(
    id serial primary key,
    name varchar,
    email varchar,
    password varchar,
    role varchar
);

INSERT INTO players(name, email, password, role) VALUES("root", "root@root.com", "root", "ROOT")