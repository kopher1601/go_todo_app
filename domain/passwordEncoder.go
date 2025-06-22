package domain

type PasswordEncoder interface {
	Encode(password string) (string, error)
	Matches(password, encodedPassword string) bool
}

type passwordEncoder struct{}
