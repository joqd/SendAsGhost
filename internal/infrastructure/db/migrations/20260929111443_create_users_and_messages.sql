-- +goose Up
CREATE TABLE users (
	id BIGINT PRIMARY KEY,
	
	first_name TEXT NOT NULL,
	last_name TEXT,
	username TEXT,

	joined_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	last_activity_at TIMESTAMPTZ
);

CREATE TABLE messages (
	id BIGINT PRIMARY KEY,
	
	from_user_id BIGINT NOT NULL REFERENCES users(id),
	to_user_id BIGINT NOT NULL REFERENCES users(id),

	sent_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	received_at TIMESTAMPTZ
);

CREATE INDEX idx_messages_from_user_id ON messages(from_user_id);
CREATE INDEX idx_message_to_user_id ON messages(to_user_id);

-- +goose Down
DROP TABLE messages;
DROP TABLE users;
