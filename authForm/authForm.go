package authForm


type AuthRequest struct {
	email string `json:"email"`
	password string `json:"password"`
	token string `json:"token"`
}

func New() *AuthRequest {
	return &AuthRequest{}
} 

func (r *AuthRequest) Fill(email string, password string, token string){
	r.email = email
	r.password = password
	r.token = token
}

func (r *AuthRequest) Unpack() map[string]string{
	var credentials map[string]string
	credentials = make(map[string]string)
	credentials["email"] = r.email
	credentials["password"] = r. password
	credentials["token"] = r.token
	return credentials
}