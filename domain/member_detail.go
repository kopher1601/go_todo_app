package domain

import "time"

type MemberDetail struct {
	ID            int
	Profile       string
	Introduction  string
	RegisteredAt  time.Time
	ActivatedAt   time.Time
	DeactivatedAt time.Time
}
