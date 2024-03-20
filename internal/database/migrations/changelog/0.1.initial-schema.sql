DROP DATABASE IF EXISTS honeycombs;
CREATE DATABASE honeycombs;
USE honeycombs;

CREATE TABLE users (
  id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  email VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  nickname VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE games (
  id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  created_by INT UNSIGNED NOT NULL REFERENCES users(id),
  state ENUM('CREATED', 'IN_PROGRESS', 'FINISHED') NOT NULL DEFAULT 'CREATED',
  playing_user INT UNSIGNED REFERENCES users(id)
);

CREATE TABLE user_games (
  user_id INT UNSIGNED REFERENCES users(id),
  game_id INT UNSIGNED REFERENCES games(id),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  player_no TINYINT UNSIGNED NOT NULL,
  state ENUM('ACTIVE', 'FINISHED') NOT NULL DEFAULT 'ACTIVE',
  CONSTRAINT PK_user_games PRIMARY KEY (user_id, game_id),
  UNIQUE (game_id, player_no)
);

CREATE TABLE turns (
  id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  user_id INT UNSIGNED NOT NULL,
  game_id INT UNSIGNED NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  points TINYINT UNSIGNED NOT NULL DEFAULT 0,
  CONSTRAINT FK_user_games FOREIGN KEY (user_id, game_id) REFERENCES user_games(user_id, game_id)
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
  INSERT INTO user_games (user_id, game_id)
  VALUES (NEW.created_by, NEW.id);
END$$

CREATE PROCEDURE clear_db()
BEGIN
  DECLARE done INT DEFAULT FALSE;
  DECLARE current_table VARCHAR(100) DEFAULT "";

  DECLARE cursor_table CURSOR FOR
  SELECT table_name FROM information_schema.tables WHERE table_schema = 'honeycombs';

  DECLARE CONTINUE HANDLER FOR NOT FOUND SET done = TRUE;

  SET FOREIGN_KEY_CHECKS = 0;

  OPEN cursor_table;
  truncate_loop: LOOP
    FETCH cursor_table INTO current_table;
    IF done THEN
      LEAVE truncate_loop;
    END IF;

    SET @sql = CONCAT('TRUNCATE TABLE ', current_table);
    PREPARE stmt FROM @sql;
    EXECUTE stmt;
    DEALLOCATE PREPARE stmt;
  END LOOP;
  CLOSE cursor_table;

  SET FOREIGN_KEY_CHECKS = 1;
END$$

DELIMITER ;
