package user

type User struct {
	Name string
	Mail string
	Nick string
	Pass string
}

func NewUser(name, mail, nick, pass string) *User {

	user := User {
		Name: name,
		Mail: mail,
		Nick: nick,
		Pass: pass,
		}

	return &user
}
