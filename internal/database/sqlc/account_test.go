package database

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {

	arg := CreateAccountParams{
		Owner:    "eben",
		Currency: "USD",
		Balance:  100,
	}

	a, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, a)

	require.Equal(t, arg.Owner, a.Owner)
	require.Equal(t, arg.Currency, a.Currency)
	require.Equal(t, arg.Balance, a.Balance)

	require.NotZero(t, a.ID)
	require.NotZero(t, a.CreatedAt)
}

func TestGetAccount(t *testing.T) {
	acc, err := testQueries.GetAccount(context.Background(), 1)
	require.NoError(t, err)
	fmt.Println(acc.Owner)
	require.NotEmpty(t, acc)
}
