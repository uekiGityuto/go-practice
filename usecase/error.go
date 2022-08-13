package usecase

type Error struct {
	reason string
}

func (e Error) Error() string {
	return e.reason
}

var NotFoundErr = Error{reason: "検索しましたが存在しませんでした。"}
