package helpers

import (
	"log"
	"strconv"

	"emperror.dev/errors"
)

func GetIdFromString(id *string) (int, error) {
	log.Fatalf("Id: %v", id)
	ID, err := strconv.Atoi(*id)
	if err != nil {
		return ID, errors.Wrap(err, "Something went wrong!")
	}
	return ID, nil
}
