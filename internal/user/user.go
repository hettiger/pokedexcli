package user

type User struct {
	Pokedex pokedex
}

func NewUser() User {
	return User{
		Pokedex: newPokedex(),
	}
}
