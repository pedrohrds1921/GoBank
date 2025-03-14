
-- name: CreatePersonAccount :one
INSERT INTO person (monthly_income,age,full_name,phone,email,category,balance )
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;
-- name: GetPersonAcountById :one
SELECT * FROM person WHERE id = $1;

-- name: UpdatePersonBalance :exec
UPDATE person SET balance = $1 WHERE id = $2;

-- name: DeletePersonAccount :exec
DELETE FROM person WHERE id = $1;