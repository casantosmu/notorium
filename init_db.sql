CREATE TABLE notes (
    note_id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    content VARCHAR NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE INDEX idx_notes_created_at ON notes(created_at);

INSERT INTO notes (title, content, created_at, expires_at)
VALUES
    (
        'An old silent pond',
        'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n- Matsuo Bashō',
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP + INTERVAL '365 days'
    ),
    (
        'Over the wintry forest',
        'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n- Natsume Soseki',
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP + INTERVAL '365 days'
    ),
    (
        'First autumn morning',
        'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n- Murakami Kijo',
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP + INTERVAL '365 days'
    );

CREATE TABLE sessions (
    token TEXT PRIMARY KEY,
    data BYTEA NOT NULL,
    expiry TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE INDEX sessions_expiry_idx ON sessions (expiry);

CREATE TABLE users (
    user_id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    hashed_password VARCHAR NOT NULL,
    created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);
