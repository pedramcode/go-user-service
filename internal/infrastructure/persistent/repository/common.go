package repository

import "database/sql"

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
