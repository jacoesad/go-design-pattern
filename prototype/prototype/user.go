package prototype

type User struct {
	email  string
	mobile string
}

func NewUser(email string, mobile string) User {
	return User{
		email:  email,
		mobile: mobile,
	}
}

// WithEmail creates a copy of User with the provided email
func (u User) WithEmail(email string) User {
	u.email = email
	return u
}

// WithPhone creates a copy of User with the provided phone
func (u User) WithMobile(mobile string) User {
	u.mobile = mobile
	return u
}
