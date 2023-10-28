CREATE TABLE users (
    id serial PRIMARY KEY,
    name varchar(30),
    email varchar(30)
);

CREATE TABLE projects (
    id serial PRIMARY KEY,
    name varchar(30),
    description varchar(255)
);

CREATE TABLE user_projects (
    id serial PRIMARY KEY,
    user_id integer REFERENCES users(id),
    project_id integer REFERENCES projects(id)
);

CREATE TABLE tasks (
    id serial PRIMARY KEY,
    name varchar(30),
    description varchar(255),
    project_id integer REFERENCES projects(id),
    status boolean DEFAULT false,
    urgency int default 0,
    due_date timestamp
);

CREATE TABLE subtasks (
    id serial PRIMARY KEY,
    name varchar(30),
    description varchar(255),
    task_id integer REFERENCES tasks(id),
    status boolean DEFAULT false
);
