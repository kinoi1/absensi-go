package models

// Request data saat user submit username & password
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Request data saat Next.js meminta rotasi token
type RefreshRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

// Struktur data Profil User
type UserProfile struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Response standar yang dibaca oleh NextAuth authorize()
type LoginResponse struct {
	AccessToken  string      `json:"accessToken"`
	RefreshToken string      `json:"refreshToken"`
	ExpiresIn    int64       `json:"expiresIn"` // dalam detik (misal: 900 = 15 menit)
	Profile      UserProfile `json:"profile"`
}

// Response saat sukses melakukan refresh token
type RefreshResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken,omitempty"` // Opsional jika ingin merotasi refresh token juga
	ExpiresIn    int64  `json:"expiresIn"`
}
