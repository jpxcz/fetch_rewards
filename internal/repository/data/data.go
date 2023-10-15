package data

import (
	"errors"

	"github.com/google/uuid"
)

var Database = make(map[string]int)

func GetNewId() string {
	u := uuid.NewString()
	if _, ok := Database[u]; ok {
		return GetNewId()
	}

	return u
}

func SetId(id string, val int) bool {
	Database[id] = val
	return true
}

func GetId(id string) (int, error) {
	if v, ok := Database[id]; ok {
		return v, nil
	}

	return 0, errors.New("id not found:" + id)
}
