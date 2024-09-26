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
