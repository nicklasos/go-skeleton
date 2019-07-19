package main

import "net/http"

func testHandler(w http.ResponseWriter, r *http.Request) {
	statement := `
		INSERT INTO
		users (name)
		VALUES	(?)
	`

	_, err := db.Exec(statement, "Nicklasos")
	if err != nil {
		logger.Error("Error on creating new user", err)

		bodyMarshal(w, response{"success": false, "message": "Cannot create new user"})
	}

	bodyMarshal(w, response{"success": true, "message": "Hello, world!"})
}
