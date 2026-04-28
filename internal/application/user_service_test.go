package application

import (
	"context"
	"errors"
	"testing"
	"time"

	"dovenet/user-service/internal/domain"
	"dovenet/user-service/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// Test suite setup helper
func setupUserServiceTest(t *testing.T) (*UserService, *mocks.UserRepository, *mocks.CredentialRepository, *mocks.OtpRepository) {
	userRepo := mocks.NewUserRepository(t)
	credRepo := mocks.NewCredentialRepository(t)
	otpRepo := mocks.NewOtpRepository(t)

	service := NewUserService(userRepo, credRepo, otpRepo)

	return service, userRepo, credRepo, otpRepo
}

// Test CreateSuperuser - Success case
func TestUserService_CreateSuperuser_Success(t *testing.T) {
	// Arrange
	service, userRepo, credRepo, _ := setupUserServiceTest(t)

	ctx := context.Background()
	username := "admin"
	email := "admin@example.com"
	password := "securepassword123"

	// Mock user repository expectations
	var createdUser *domain.User
	userRepo.On("Create", ctx, mock.AnythingOfType("*domain.User")).
		Run(func(args mock.Arguments) {
			// Capture the user that was passed to Create
			createdUser = args.Get(1).(*domain.User)
			// Simulate setting the ID after creation
			createdUser.Entity.Id = 1
			createdUser.Entity.CreatedAt = time.Now()
			createdUser.Entity.UpdatedAt = time.Now()
		}).
		Return(nil)

	// Mock credential repository expectations (called twice)
	credRepo.On("Create", ctx, mock.AnythingOfType("*domain.Credential")).Return(nil).Twice()

	// Act
	result, err := service.CreateSuperuser(ctx, username, email, password)

	// Assert
	assert.NoError(t, err)
	require.NotNil(t, result)

	// Verify user fields
	assert.Equal(t, email, result.Email)
	assert.Equal(t, username, result.Username)
	assert.True(t, result.IsSuperuser)
	assert.True(t, result.IsVerified)
	assert.NotZero(t, result.Id)

	// Verify the user passed to repository had correct fields
	require.NotNil(t, createdUser)
	assert.Equal(t, email, createdUser.Email)
	assert.Equal(t, username, createdUser.Username)
	assert.True(t, createdUser.IsSuperuser)
	assert.True(t, createdUser.IsVerified)

	// Verify all expectations were met
	userRepo.AssertExpectations(t)
	credRepo.AssertExpectations(t)
}

// Test CreateSuperuser - User Creation Fails
func TestUserService_CreateSuperuser_UserCreationFails(t *testing.T) {
	// Arrange
	service, userRepo, credRepo, _ := setupUserServiceTest(t)

	ctx := context.Background()
	username := "admin"
	email := "admin@example.com"
	password := "securepassword123"

	expectedErr := errors.New("database connection failed")
	userRepo.On("Create", ctx, mock.AnythingOfType("*domain.User")).Return(expectedErr)

	// Act
	result, err := service.CreateSuperuser(ctx, username, email, password)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Nil(t, result)

	// Verify credential Create was never called
	credRepo.AssertNotCalled(t, "Create")
	userRepo.AssertExpectations(t)
}

// Test CreateSuperuser - First Credential Creation Fails
func TestUserService_CreateSuperuser_FirstCredentialCreationFails(t *testing.T) {
	// Arrange
	service, userRepo, credRepo, _ := setupUserServiceTest(t)

	ctx := context.Background()
	username := "admin"
	email := "admin@example.com"
	password := "securepassword123"

	// Mock successful user creation
	userRepo.On("Create", ctx, mock.AnythingOfType("*domain.User")).
		Run(func(args mock.Arguments) {
			user := args.Get(1).(*domain.User)
			user.Entity.Id = 1
		}).
		Return(nil)

	// Mock first credential creation fails
	expectedErr := errors.New("credential creation failed")
	credRepo.On("Create", ctx, mock.AnythingOfType("*domain.Credential")).
		Return(expectedErr).Once()

	// Act
	result, err := service.CreateSuperuser(ctx, username, email, password)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Nil(t, result)

	// Verify second credential creation was not attempted
	credRepo.AssertNumberOfCalls(t, "Create", 1)
	userRepo.AssertExpectations(t)
}

// Test CreateSuperuser - Second Credential Creation Fails
func TestUserService_CreateSuperuser_SecondCredentialCreationFails(t *testing.T) {
	// Arrange
	service, userRepo, credRepo, _ := setupUserServiceTest(t)

	ctx := context.Background()
	username := "admin"
	email := "admin@example.com"
	password := "securepassword123"

	// Mock successful user creation
	userRepo.On("Create", ctx, mock.AnythingOfType("*domain.User")).
		Run(func(args mock.Arguments) {
			user := args.Get(1).(*domain.User)
			user.Entity.Id = 1
		}).
		Return(nil)

	// Mock first credential creation succeeds, second fails
	expectedErr := errors.New("second credential creation failed")
	credRepo.On("Create", ctx, mock.AnythingOfType("*domain.Credential")).
		Return(nil).Once()
	credRepo.On("Create", ctx, mock.AnythingOfType("*domain.Credential")).
		Return(expectedErr).Once()

	// Act
	result, err := service.CreateSuperuser(ctx, username, email, password)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Nil(t, result)

	// Verify both credential creations were attempted
	credRepo.AssertNumberOfCalls(t, "Create", 2)
	userRepo.AssertExpectations(t)
}

// Test CreateSuperuser - Verify credentials have correct data
func TestUserService_CreateSuperuser_VerifyCredentialData(t *testing.T) {
	// Arrange
	service, userRepo, credRepo, _ := setupUserServiceTest(t)

	ctx := context.Background()
	username := "admin"
	email := "admin@example.com"
	password := "securepassword123"

	// Capture credentials
	var capturedCredentials []*domain.Credential
	userRepo.On("Create", ctx, mock.AnythingOfType("*domain.User")).
		Run(func(args mock.Arguments) {
			user := args.Get(1).(*domain.User)
			user.Entity.Id = 1
		}).
		Return(nil)

	credRepo.On("Create", ctx, mock.AnythingOfType("*domain.Credential")).
		Run(func(args mock.Arguments) {
			cred := args.Get(1).(*domain.Credential)
			capturedCredentials = append(capturedCredentials, cred)
		}).
		Return(nil).Twice()

	// Act
	_, err := service.CreateSuperuser(ctx, username, email, password)

	// Assert
	assert.NoError(t, err)
	require.Len(t, capturedCredentials, 2)

	// Find username and email credentials (order may vary)
	var usernameCred, emailCred *domain.Credential
	for _, cred := range capturedCredentials {
		switch cred.Key {
		case "username":
			usernameCred = cred
		case "email":
			emailCred = cred
		}
	}

	// Verify username credential
	require.NotNil(t, usernameCred, "Username credential not found")
	assert.Equal(t, int32(1), usernameCred.UserID)
	assert.Equal(t, domain.Password, usernameCred.Type)
	assert.Equal(t, "username", usernameCred.Key)
	assert.Equal(t, password, usernameCred.Value)

	// Verify email credential
	require.NotNil(t, emailCred, "Email credential not found")
	assert.Equal(t, int32(1), emailCred.UserID)
	assert.Equal(t, domain.Password, emailCred.Type)
	assert.Equal(t, "email", emailCred.Key)
	assert.Equal(t, password, emailCred.Value)

	userRepo.AssertExpectations(t)
	credRepo.AssertExpectations(t)
}

// Table-driven test for CreateSuperuser
func TestUserService_CreateSuperuser_TableDriven(t *testing.T) {
	tests := []struct {
		name          string
		username      string
		email         string
		password      string
		setupMocks    func(*mocks.UserRepository, *mocks.CredentialRepository)
		expectedError error
		validateUser  func(*testing.T, *domain.User)
	}{
		{
			name:     "successful superuser creation",
			username: "admin",
			email:    "admin@example.com",
			password: "password123",
			setupMocks: func(u *mocks.UserRepository, c *mocks.CredentialRepository) {
				u.On("Create", mock.Anything, mock.AnythingOfType("*domain.User")).
					Run(func(args mock.Arguments) {
						user := args.Get(1).(*domain.User)
						user.Entity.Id = 1
					}).
					Return(nil)
				c.On("Create", mock.Anything, mock.AnythingOfType("*domain.Credential")).Return(nil).Twice()
			},
			expectedError: nil,
			validateUser: func(t *testing.T, user *domain.User) {
				assert.True(t, user.IsSuperuser)
				assert.True(t, user.IsVerified)
			},
		},
		{
			name:     "empty username",
			username: "",
			email:    "admin@example.com",
			password: "password123",
			setupMocks: func(u *mocks.UserRepository, c *mocks.CredentialRepository) {
				u.On("Create", mock.Anything, mock.AnythingOfType("*domain.User")).
					Run(func(args mock.Arguments) {
						user := args.Get(1).(*domain.User)
						user.Entity.Id = 1
					}).
					Return(nil)
				c.On("Create", mock.Anything, mock.AnythingOfType("*domain.Credential")).Return(nil).Twice()
			},
			expectedError: nil, // Domain might allow empty? Adjust based on your validation
			validateUser: func(t *testing.T, user *domain.User) {
				assert.Equal(t, "", user.Username)
			},
		},
		{
			name:     "invalid email format",
			username: "admin",
			email:    "invalid-email",
			password: "password123",
			setupMocks: func(u *mocks.UserRepository, c *mocks.CredentialRepository) {
				u.On("Create", mock.Anything, mock.AnythingOfType("*domain.User")).
					Run(func(args mock.Arguments) {
						user := args.Get(1).(*domain.User)
						user.Entity.Id = 1
					}).
					Return(nil)
				c.On("Create", mock.Anything, mock.AnythingOfType("*domain.Credential")).Return(nil).Twice()
			},
			expectedError: nil, // Domain might allow? Add validation if needed
			validateUser: func(t *testing.T, user *domain.User) {
				assert.Equal(t, "invalid-email", user.Email)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			service, userRepo, credRepo, _ := setupUserServiceTest(t)
			tt.setupMocks(userRepo, credRepo)

			ctx := context.Background()

			// Act
			result, err := service.CreateSuperuser(ctx, tt.username, tt.email, tt.password)

			// Assert
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				require.NotNil(t, result)
				assert.Equal(t, tt.email, result.Email)
				assert.Equal(t, tt.username, result.Username)
				if tt.validateUser != nil {
					tt.validateUser(t, result)
				}
			}

			userRepo.AssertExpectations(t)
			credRepo.AssertExpectations(t)
		})
	}
}

// Test with context cancellation
func TestUserService_CreateSuperuser_ContextCancellation(t *testing.T) {
	// Arrange
	service, userRepo, credRepo, _ := setupUserServiceTest(t)

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	username := "admin"
	email := "admin@example.com"
	password := "password123"

	// Mock should respect context cancellation
	userRepo.On("Create", ctx, mock.AnythingOfType("*domain.User")).
		Return(context.Canceled)

	// Act
	result, err := service.CreateSuperuser(ctx, username, email, password)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, context.Canceled, err)
	assert.Nil(t, result)

	userRepo.AssertExpectations(t)
	credRepo.AssertNotCalled(t, "Create")
}
