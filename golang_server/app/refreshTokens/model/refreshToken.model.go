package model

type RefreshToken struct {
	ID        int    `db:"id"`
	Token     string `db:"token"`
	ExpiresAt int64  `db:"expires_at"`
	UserID    int    `db:"user_id"`
}
