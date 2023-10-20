CREATE SCHEMA IF NOT EXISTS quartz_user;

CREATE TABLE IF NOT EXISTS quartz_user.user
(
    id            serial PRIMARY KEY,
    email         varchar(255)                NOT NULL,
    password_hash bytea                       NOT NULL,
    status        smallint                    NOT NULL,
    created_at    timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at    timestamp(0) with time zone NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS quartz_user.user_profile
(
    user_id      serial
        PRIMARY KEY
        CONSTRAINT fk__UserProfile_userId__User_id
            REFERENCES quartz_user.user (id)
            ON UPDATE CASCADE ON DELETE CASCADE,
    firstname    varchar(255) NOT NULL DEFAULT '',
    middlename   varchar(255) NOT NULL DEFAULT '',
    lastname     varchar(255) NOT NULL DEFAULT '',
    phone_number varchar(255) NOT NULL DEFAULT ''
);

INSERT INTO quartz_user.user (email, password_hash, status)
VALUES ('adminfmq@freematiq.com', '$2a$12$JhPOa8rco4g8e6arqtRYvuybSkYb/DdJAcxTwWQG6PMFXnq2nQNKi', 10); -- passw0rd

INSERT INTO quartz_user.user_profile (user_id, firstname)
SELECT id, 'Admin'
FROM quartz_user.user
WHERE email = 'adminfmq@freematiq.com';