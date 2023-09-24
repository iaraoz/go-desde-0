package users

import (
	"fmt"
	"time"

	"github.com/iaraoz/go-desde-0/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(10, "israel", time.Now(), true)
	fmt.Println(u)
}
