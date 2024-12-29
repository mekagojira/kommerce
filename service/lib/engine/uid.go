package engine

import "github.com/google/uuid"

func GetUid() *Result[string] {
	uid, err := uuid.NewV7()

	if err != nil {
		return NewResult[string]().WithError(err)
	}

	idString := uid.String()

	return NewResult[string]().WithPureData(idString)
}

func SetUid(id *string) *Result[string] {
	uid := GetUid()

	if !uid.IsOk() {
		return uid
	}

	*id = uid.PureData()

	return uid
}
