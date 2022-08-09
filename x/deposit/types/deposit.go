package types

import (
	deposittypes "github.com/sentinel-official/hub/x/deposit/types"

	clienttypes "github.com/sentinel-official/cli-client/types"
)

type Deposit struct {
	Address string            `json:"address"`
	Amount  clienttypes.Coins `json:"amount"`
}

func NewDepositFromRaw(v *deposittypes.Deposit) Deposit {
	return Deposit{
		Address: v.Address,
		Amount:  clienttypes.NewCoinsFromRaw(v.Coins),
	}
}

type Deposits []Deposit

func NewDepositsFromRaw(v deposittypes.Deposits) Deposits {
	items := make(Deposits, 0, len(v))
	for i := 0; i < len(v); i++ {
		items = append(items, NewDepositFromRaw(&v[i]))
	}

	return items
}
