package repository

import (
	"context"
	"dovenet/user-service/internal/domain"
	sql_repository "dovenet/user-service/internal/infrastructure/persistent/repository/sqlc"

	"github.com/jmoiron/sqlx"
)

type OtpRepository struct {
	queries *sql_repository.Queries
}

func NewOtpRepository(db *sqlx.DB) *OtpRepository {
	return &OtpRepository{
		queries: sql_repository.New(db.DB),
	}
}

func (r *OtpRepository) GetByID(ctx context.Context, id int32) (*domain.Otp, error) {
	otp, err := r.queries.OtpGetByID(ctx, id)
	// TODO handle error
	if err != nil {
		return nil, err
	}
	return &domain.Otp{
		Entity: domain.Entity{
			Id:        otp.ID,
			CreatedAt: otp.CreatedAt.Time,
			UpdatedAt: otp.UpdatedAt.Time,
		},
		UserID: otp.UserID,
		Reason: domain.OtpReason(otp.Reason),
		Medium: domain.OtpMedium(otp.Medium),
		Code:   otp.Code,
		UsedAt: &otp.UsedAt.Time,
	}, nil
}

func (r *OtpRepository) GetValidOtp(ctx context.Context, userID int32, code string, reason domain.OtpReason, medium domain.OtpMedium) (*domain.Otp, error) {
	otp, err := r.queries.OtpGetValidOtp(ctx, sql_repository.OtpGetValidOtpParams{
		UserID: userID,
		Code:   code,
		Reason: string(reason),
		Medium: string(medium),
	})
	// TODO handle error
	if err != nil {
		return nil, err
	}
	return &domain.Otp{
		Entity: domain.Entity{
			Id:        otp.ID,
			CreatedAt: otp.CreatedAt.Time,
			UpdatedAt: otp.UpdatedAt.Time,
		},
		UserID: otp.UserID,
		Reason: domain.OtpReason(otp.Reason),
		Medium: domain.OtpMedium(otp.Medium),
		Code:   otp.Code,
		UsedAt: &otp.UsedAt.Time,
	}, nil
}

func (r *OtpRepository) DeleteByID(ctx context.Context, id int32) error {
	_, err := r.queries.OtpDeleteByID(ctx, id)
	// TODO handle error
	if err != nil {
		return err
	}
	return nil
}

func (r *OtpRepository) Create(ctx context.Context, otp *domain.Otp) error {
	result, err := r.queries.OtpCreate(ctx, sql_repository.OtpCreateParams{
		UserID: otp.UserID,
		Code:   otp.Code,
		Reason: string(otp.Reason),
		Medium: string(otp.Medium),
		UsedAt: ToNullTime(otp.UsedAt),
	})
	// TODO handle error
	if err != nil {
		return err
	}
	otp.Entity.Id = result.ID
	otp.Entity.CreatedAt = result.CreatedAt.Time
	otp.Entity.UpdatedAt = result.UpdatedAt.Time
	return nil
}

func (r *OtpRepository) Update(ctx context.Context, otp *domain.Otp) error {
	result, err := r.queries.OtpUpdate(ctx, sql_repository.OtpUpdateParams{
		ID:     otp.Entity.Id,
		UserID: otp.UserID,
		Code:   otp.Code,
		Reason: string(otp.Reason),
		Medium: string(otp.Medium),
		UsedAt: ToNullTime(otp.UsedAt),
	})
	// TODO handle error
	if err != nil {
		return err
	}
	otp.UpdatedAt = result.Time
	return nil
}
