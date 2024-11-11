DROP TABLE IF EXISTS users;

CREATE TABLE public.users (
    user_id  SERIAL PRIMARY KEY,
    first_name  VARCHAR(255),
    last_name  VARCHAR(255),
    google  TEXT CHECK (google IS NOT NULL OR apple IS NOT NULL) UNIQUE,
    apple  TEXT CHECK (google IS NOT NULL OR apple IS NOT NULL) UNIQUE,
    role  SERIAL REFERENCES roles NOT NULL,
    trainer  SERIAL REFERENCES public.users,
    created  TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated  TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);
