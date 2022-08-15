package usecase

type Error struct {
	reason string
}

func (e Error) Error() string {
	return e.reason
}

// Unwrap(), As(target interface{}), Is(target error)メソッドをErrorに実装すべき？

var ErrNotFound = Error{reason: "検索しましたが存在しませんでした。"}
