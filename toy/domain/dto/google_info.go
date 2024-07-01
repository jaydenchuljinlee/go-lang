package dto

type GoogleInfo struct {
	ID           int64  `redis:"id"`
	Email        string `redis:"email"`
	AccessToken  string `redis:"accessToken"`
	ExpiresIn    string `redis:"expiresIn"`
	RefreshToken string `redis:"refreshToken"`
	TokenType    string `redis:"tokenType"`
	Scope        string `redis:"scope"`
	CreatedAt    string `redis:"createdAt"`
	UpdatedAt    string `redis:"updatedAt"`
}
