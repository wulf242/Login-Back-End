package authForm

type AuthRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Token    string `json:"token" binding:"required"`
}

func New() *AuthRequest {
	return &AuthRequest{}
}

func (r *AuthRequest) Fill(email string, password string, token string) {
	r.Email = email
	r.Password = password
	r.Token = token
}

func (r *AuthRequest) Unpack() map[string]string {
	var credentials map[string]string
	credentials = make(map[string]string)
	credentials["email"] = r.Email
	credentials["password"] = r.Password
	credentials["token"] = r.Token
	return credentials
}
