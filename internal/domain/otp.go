package domain

import (
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
