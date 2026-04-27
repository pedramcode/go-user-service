package repository

import (
	"database/sql"
	"time"
)

func ToNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{
		String: *s,
		Valid:  true,
	}
}

func ToNullBool(s bool) sql.NullBool {
	return sql.NullBool{
		Bool:  s,
		Valid: true,
	}
}

func ToNullTime(s *time.Time) sql.NullTime {
	if s == nil {
		return sql.NullTime{Valid: false}
	}
	return sql.NullTime{
		Time:  *s,
		Valid: true,
	}
}
