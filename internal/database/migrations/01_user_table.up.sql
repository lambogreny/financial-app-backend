CREATE EXTENTSION IF NOT EXITSTS "uuid-ossp";

CREATE TABLE users(
    user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email TEXT NOT NULL,
    password_hash bytes,
    create at TIMESTAMP NOT NULL DEFAULT NOw(),
    deleted_at TIMESTAMP
);

CREATE UNIQUE INDEX  user_email
    on users (email);