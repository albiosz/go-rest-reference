DROP DATABASE IF EXISTS honeycombs;
CREATE DATABASE honeycombs;
USE honeycombs;

CREATE TABLE users (
  email VARCHAR(255) PRIMARY KEY,
  password VARCHAR(255) NOT NULL,
  nickname VARCHAR(255) NOT NULL
);

CREATE TABLE games (
  id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  created_by VARCHAR(255) NOT NULL REFERENCES users(email),
  state ENUM('CREATED', 'IN_PROGRESS', 'FINISHED') NOT NULL DEFAULT 'CREATED',
  playing_user VARCHAR(255) REFERENCES users(email)
);

CREATE TABLE user_games (
  user_email VARCHAR(255) REFERENCES users(email),
  game_id INT UNSIGNED REFERENCES games(id),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  player_no TINYINT UNSIGNED NOT NULL,
  state ENUM('ACTIVE', 'FINISHED') NOT NULL DEFAULT 'ACTIVE',
  CONSTRAINT PK_user_games PRIMARY KEY (user_email, game_id),
  UNIQUE (game_id, player_no)
);

CREATE TABLE turns (
  id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  user_email VARCHAR(255) NOT NULL,
  game_id INT UNSIGNED NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  points TINYINT UNSIGNED NOT NULL DEFAULT 0,
  CONSTRAINT FK_user_games FOREIGN KEY (user_email, game_id) REFERENCES user_games(user_email, game_id)
);

DELIMITER $$

CREATE TRIGGER assign_player_no
BEFORE INSERT
ON user_games FOR EACH ROW
BEGIN
  DECLARE max_player_no TINYINT UNSIGNED;
  SELECT MAX(player_no) INTO max_player_no FROM user_games WHERE game_id = NEW.game_id;
  SET NEW.player_no = IFNULL(max_player_no, 0) + 1;
END$$

CREATE TRIGGER create_user_game
AFTER INSERT
ON games FOR EACH ROW
BEGIN
  INSERT INTO user_games (user_email, game_id)
  VALUES (NEW.created_by, NEW.id);
END$$

DELIMITER ;
