package db

const (
	INSERT = `
		INSERT INTO "%v" (
			%v
		)
		VALUES (
			%v
		);
	`
	READ_ONE = `
		SELECT *
		FROM "%v"
		WHERE id = %v
		LIMIT 1;
	`
	READ_ALL = `
		SELECT *
		FROM "%v"
		WHERE (
			%v
		);
	`
)
