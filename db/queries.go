package db

const(
	insertRoom = "INSERT INTO rooms(created_at, updated_at, deleted_at, room, name) VALUES (current_timestamp, current_timestamp, NULL ,?, ?)"

)
