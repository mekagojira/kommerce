package engine

import "github.com/google/uuid"

func GetUid() Result[string] {
	uid, err := uuid.NewV7()

	if err != nil {
		return Result[string]{
			Error: err,
		}
	}

	idString := uid.String()

	return Result[string]{
		Data: &idString,
	}
}
