package repository

import (
	"errors"
	"zkfmapf123/src/base"
	"zkfmapf123/src/utils"
)

type transfers struct{
	FromAccountId int `json:"from_account_id"`
	ToAccountId int `json:"to_account_id"`
	Amount int `json:"account"`
	CreatedAt int64 `json:"created_at"`
}

func NewTransfer(fromAccountId int, toAccountId int, amount int) *transfers{
	now := utils.Now()
	return &transfers{FromAccountId: fromAccountId, ToAccountId: toAccountId, Amount: amount, CreatedAt: now}
}

func (t *transfers) Transfer(fromAccount *base.AccountModels, toAccount *base.AccountModels)error{

	if t.Amount > fromAccount.Balance {
		return errors.New("more money than you currently have")
	}

	querys := []string {
		"update accounts set balance = ? where id = ?",
		"update accounts set balance = ? where id = ?",
		"insert into transfers (from_account_id, to_account_id,amount,created_at) values(?,?,?,?)",
	}
	interfaces := [][]interface{} {
		{(fromAccount.Balance -  t.Amount), t.FromAccountId},
		{(toAccount.Balance + t.Amount), t.ToAccountId},
		{t.FromAccountId, t.ToAccountId, t.Amount, t.CreatedAt},
	}

	for i:= 0; i<3; i++ {
		err := Transaction(querys[i], interfaces[i])
		if err != nil{
			return err
		}
	}

	return nil
}

