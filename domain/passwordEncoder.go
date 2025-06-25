package domain

type PasswordEncoder interface {
	Encode(password string) (string, error)
	Matches(password, encodedPassword string) bool
}

type passwordEncoder struct{}

func (p passwordEncoder) Encode(password string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (p passwordEncoder) Matches(password, encodedPassword string) bool {
	//TODO implement me
	panic("implement me")
}
