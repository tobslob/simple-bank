package db

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tobslob/simple-bank/utils"
)

func createRandomUser(t *testing.T) User {
	var userArg = CreateUserParams{
		Username:       utils.RandomString(),
		HashedPassword: "secret",
		FullName:       utils.RandomString(),
		Email:          utils.RandomEmail(),
	}
	user, err := testQueries.CreateUser(context.Background(), userArg)
	if err != nil {
		log.Fatal("Error while creating account:", err)
	}

	require.NotEmpty(t, user)

	require.Equal(t, userArg.FullName, user.FullName)
	require.Equal(t, userArg.Email, user.Email)
	require.Equal(t, userArg.Username, user.Username)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetAUser(context.Background(), user1.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.CreatedAt, user2.CreatedAt)
}
