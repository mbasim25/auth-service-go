package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/mbasim25/ticketing-app-microservices/util"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.RandomPassword()
	require.NoError(t, err)

	arg := CreateUserParams{
		Username: util.RandomTestStrin(),
		Email:    util.RandomEmail(),
		Password: hashedPassword,
		Role:     "USER",
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Role, user.Role)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)

	sameUser, err := testQueries.GetUser(context.Background(), user.Email)
	require.NoError(t, err)
	require.NotEmpty(t, sameUser)

	require.Equal(t, sameUser.ID, user.ID)
	require.Equal(t, sameUser.Username, user.Username)
	require.Equal(t, sameUser.Email, user.Email)
	require.Equal(t, sameUser.Password, user.Password)
	require.Equal(t, sameUser.Role, user.Role)

	require.WithinDuration(t, sameUser.CreatedAt, user.CreatedAt, time.Second)
}

func TestDeleteUser(t *testing.T) {
	user := createRandomUser(t)

	err := testQueries.DeleteUser(context.Background(), user.ID)
	require.NoError(t, err)

	sameUser, err := testQueries.GetUser(context.Background(), user.Email)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, sameUser)
}

func TestListUsers(t *testing.T) {
	for i := 0; i < 4; i++ {
		createRandomUser(t)
	}

	arg := ListUsersParams{
		Limit:  2,
		Offset: 2,
	}

	users, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 2)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}
