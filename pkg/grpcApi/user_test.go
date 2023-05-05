package grpcApi

import (
	"context"
	"testing"
	"time"

	"github.com/Damione1/portfolio-playground/db/models"
	"github.com/Damione1/portfolio-playground/pkg/pb"
	"github.com/Damione1/portfolio-playground/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// MockDB is a mock implementation of the database.
type MockDB struct {
	mock.Mock
}

func (m *MockDB) Insert(ctx context.Context, user *models.User, columns boil.Columns) error {
	args := m.Called(ctx, user, columns)
	return args.Error(0)
}

// ... Add more mock methods for other database operations as needed

func TestCreateUser(t *testing.T) {
	// Create a mock database
	mockDB := &MockDB{}

	// Create a server config with the mock database
	config := &Config{
		DB:                   mockDB,
		AccessTokenDuration:  time.Hour,
		RefreshTokenDuration: 24 * time.Hour,
	}

	// Create a new server with the config
	server := NewServer(config)

	// Prepare a request for creating a user
	req := &pb.CreateUserRequest{
		Email:    "test@example.com",
		Password: "Test@123",
		Name:     "Test User",
	}

	// Set the expected behavior of the mock database
	hashedPassword, _ := util.HashPassword(req.GetPassword())
	mockDB.On("Insert", mock.Anything, &models.User{
		Email:    req.GetEmail(),
		Password: hashedPassword,
		Name:     req.GetName(),
	}, boil.Infer()).Return(nil)

	// Call the CreateUser method on the server
	resp, err := server.CreateUser(context.Background(), req)

	// Assert the response and error
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, resp.User)
	assert.Equal(t, req.Email, resp.User.Email)
	assert.Equal(t, req.Name, resp.User.Name)

	// Assert the mock database expectations
	mockDB.AssertExpectations(t)
}
