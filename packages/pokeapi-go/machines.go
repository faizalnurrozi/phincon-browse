package pokeapi

import (
	"fmt"

	"github.com/faizalnurrozi/phincon-browse/packages/pokeapi-go/structs"
)

// Machine returns a single machine (by ID).
func Machine(id string) (result structs.Machine, err error) {
	err = do(fmt.Sprintf("machine/%s", id), &result)
	return result, err
}
