-- name: CreateCompanyAccount :one
INSERT INTO company (annual_revenue ,years_in_business,trade_name ,phone ,
corporate_email ,category ,balance )
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetCompanyAcountById :one
SELECT * FROM company WHERE id = $1;

-- name: UpdateCompanyBalance :exec
UPDATE company SET balance = $1 WHERE id = $2;

-- name: DeleteCompanyAccount :exec
DELETE FROM company WHERE id = $1;