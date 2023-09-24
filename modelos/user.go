package modelos

import (
	"time"
)

type User struct {
	Id        int
	Name      string
	createdAt time.Time
	Status    bool
}

func (usuario *User) AddUser(id int, name string, createdAt time.Time, status bool) {
	usuario.Id = id
	usuario.Name = name
	usuario.createdAt = createdAt
	usuario.Status = status
}
