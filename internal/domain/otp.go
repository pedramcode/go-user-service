package domain

import (
	"context"
	"time"
)

type OtpReason string

const (
	Register       OtpReason = "register"
	Login          OtpReason = "login"
	ResetPassword  OtpReason = "reset_password"
	ChangePassword OtpReason = "change_password"
	CriticalOp     OtpReason = "critical_operation"
)

type OtpMedium string

const (
	Email  OtpMedium = "email"
	SMS    OtpMedium = "sms"
	Notif  OtpMedium = "notif"
	Manual OtpMedium = "manual"
)

type Otp struct {
	Entity
	UserID int32
	Reason OtpReason
	Medium OtpMedium
	Code   string
	UsedAt *time.Time
}

type OtpRepository interface {
	GetByID(ctx context.Context, id int32) (*Otp, error)
	GetValidOtp(ctx context.Context, userID int32, code string, reason OtpReason, medium OtpMedium) (*Otp, error)
	DeleteByID(ctx context.Context, id int32) error
	Create(ctx context.Context, otp *Otp) error
	Update(ctx context.Context, otp *Otp) error
}
