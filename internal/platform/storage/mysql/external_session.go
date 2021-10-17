package mysql

const (
	sqlExternalSessionTable = "external_sessions"
)

type sqlExternalSession struct {
	ID           string `db:"id"`
	UserID       string `db:"user_id"`
	AccessToken  string `db:"access_token"`
	RefreshToken string `db:"refresh_token"`
	Scope        string `db:"scope"`
	TokenType    string `db:"token_type"`
}
