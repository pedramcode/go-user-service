package port

import (
	"context"
	"dovenet/user-service/internal/domain"
)

//go:generate mockery --name=OtpRepository --output=../../../mocks --outpkg=mocks --filename=mock_otp_repository.go
type OtpRepository interface {
	GetByID(ctx context.Context, id int32) (*domain.Otp, error)
	GetValidOtp(ctx context.Context, userID int32, code string, reason domain.OtpReason, medium domain.OtpMedium) (*domain.Otp, error)
	DeleteByID(ctx context.Context, id int32) error
	Create(ctx context.Context, otp *domain.Otp) error
	Update(ctx context.Context, otp *domain.Otp) error
}
