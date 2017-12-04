package tweet_user

type User struct {
	Name, Mail, Nick, Pass string
}

func NewUser(name, mail, nick, pass string) *User {

	user := User{
		name,
		mail,
		nick,
		pass,
	}

	return &user
}
