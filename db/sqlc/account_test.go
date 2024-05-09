package db

import (
	"context"
	"log"
	"testing"

	"github.com/losuch/fc-order/util"
	"github.com/stretchr/testify/require"
)


var account Account

// createRandomAccount creates arandom account for testing purposes.
// It generates random values for the account's attributes and inserts the account into the database.
// It returns the created account
func createRandomAccount(t *testing.T) Account {
    arg := CreateAccountParams{
        Email:       util.RandomEmail(), 
        HashedPassword:    util.RandomString(10),
        Role:       "ADMIN",
    }

    log.Printf("arg: %v", arg)

    account, err := testQueries.CreateAccount(context.Background(), arg)
    require.NoError(t, err)
    require.NotEmpty(t, account)
    require.Equal(t, arg.Email, account.Email)
    require.Equal(t, arg.Role, account.Role)

    require.NotZero(t, account.AccountID)
    require.NotZero(t, account.CreatedAt)
    return account
}

// TestCreateAccount is a unit test for the createRandomAccount function.
func TestCreateAccount(t *testing.T) {
    account = createRandomAccount(t)
}

// TestGetAccountByEmail is a unit test for the GetAccount function. 
func TestGetAccountByEmail(t *testing.T) {
    account, err := testQueries.GetAccountByEmail(context.Background(), account.Email)
    require.NoError(t, err)
    require.NotEmpty(t, account)
    require.Equal(t, account.Email, account.Email)
}

// TestGetAccount is a unit test for the GetAccount function.
func TestGetAccount(t *testing.T) {
    account, err := testQueries.GetAccount(context.Background(), account.AccountID)
    require.NoError(t, err)
    require.NotEmpty(t, account)
    require.Equal(t, account.Email, account.Email)
}

// TestGetAccountList is a unit test for the GetAccountList function.
func TestGetAccountList(t *testing.T) {
    accounts, err := testQueries.GetAccountList(context.Background())
    require.NoError(t, err)
    require.NotEmpty(t, accounts)
}

// TestUpdateAccount is a unit test for the UpdateAccount function.
func TestUpdateAccount(t *testing.T) {
    arg := UpdateAccountParams{
        AccountID: account.AccountID,
        HashedPassword: util.RandomString(10),
        Role: "ADMIN",
    }

    updatedAccount, err := testQueries.UpdateAccount(context.Background(), arg)
    require.NoError(t, err)
    require.NotEmpty(t, updatedAccount)
    require.Equal(t, arg.HashedPassword, updatedAccount.HashedPassword)
}

// TestDeleteAccount is a unit test for the DeleteAccount function.
func TestDeleteAccount(t *testing.T) {
    err := testQueries.DeleteAccount(context.Background(), account.AccountID)
    require.NoError(t, err)

    account, err := testQueries.GetAccount(context.Background(), account.AccountID)
    require.Error(t, err)
    require.Empty(t, account)
}