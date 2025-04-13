package model

import "fmt"

type Account struct{
	AccountNo string
	Pwd string
	Balance float64
}

func NewAccount(AccountNO string, pwd string, balance float64) *Account{
	if len(AccountNO) < 6 || len(AccountNO) > 10{
		fmt.Println("账号长度不正确")
		return nil
	}
	if len(pwd) != 6{
		fmt.Println("密码长度不正确")
		return nil
	}
	if balance < 20{
		fmt.Println("余额不足")
		return nil
	}
	return &Account{
		AccountNo: AccountNO,
		Pwd: pwd,
		Balance: balance,
	}
}