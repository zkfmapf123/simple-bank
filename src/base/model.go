package base

type AccountModels struct {
	Id int `json:"id" db:"id"`
	Owner string `json:"owner" db:"owner"`
	Balance int `json:"balance" db:"balance"`
	Currency string `json:"currency" db:"currency"`
	CreatedAt int `json:"created_at" db:"created_at"`
	UpdatedAt int `json:"updated_at" db:"updated_at"`
	Timezone string `json:"timezone" db:"timezone"`
}

type EntrieModels struct{
	Id int
	AccountId string `db:"account_id"`
	Amount int
}

type TranserModels struct {
	Id int
	FromAccountId int `db:"from_account_id"`
	ToAccountId int `db:"to_account_id"`
	Amount int
	CreatedAt int `db:"created_at"`
}