package uuid

import (
	"fmt"
	"github.com/satori/go.uuid"
)

func CreateUUID() string {
	u1 := uuid.Must(uuid.NewV4())
	return fmt.Sprint(u1)
}
