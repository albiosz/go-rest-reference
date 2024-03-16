package seed

import "database/sql"

func insertUsers(db *sql.DB) {
	result, err := db.Exec(
		`INSERT INTO honeycombs.users (email, password, nickname)
		VALUES ('user1@mail.com', 'password1', 'user1'),
			('user2@mail.com', 'password2', 'user2');
	`)
	if err != nil {
		panic(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected != 2 {
		panic("Unexpected state of the DB!")
	}
}

func insertGames(db *sql.DB) {
	result, err := db.Exec(
		`INSERT INTO honeycombs.games (created_by)
		VALUES ('user1@mail.com');
	`)
	if err != nil {
		panic(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected != 1 {
		panic("Unexpected state of the DB!")
	}

}

func InsertAll(db *sql.DB) {
	insertUsers(db)
	insertGames(db)
}
