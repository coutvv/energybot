CREATE TABLE  IF NOT EXISTS USER (
			"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
			"tele_id" integer UNIQUE,
			"user_name" TEXT,
			"first_name" TEXT,
			"last_name" TEXT
		);


CREATE TABLE IF NOT EXISTS GAME (
    "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
    "chat_id" integer NOT NULL,
    "state" integer
);


CREATE TABLE IF NOT EXISTS GAME_PLAYER (
    "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
    "game_id" integer NOT NULL,
    "user_id" integer NOT NULL,
    FOREIGN KEY (game_id) REFERENCES game(id),
    FOREIGN KEY (user_id) REFERENCES user(id)
);