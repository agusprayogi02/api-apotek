package service

type LoginService interface {
	LoginUser(email string, password string) bool
}

type LoginInformation struct {
	Email    string
	Password string
}

func StaticLoginService() LoginService {
	return &LoginInformation{
		Email:    "agus21apy@gmail.com",
		Password: "testing",
	}
}

func (e LoginInformation) LoginUser(email string, password string) bool {
	return e.Email == email && e.Password == password
}
