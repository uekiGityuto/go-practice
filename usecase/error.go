package usecase

import "golang.org/x/xerrors"

type Error struct {
	err    error
	reason string
}

func (e Error) Error() string {
	return e.reason
}

var errNotFound = xerrors.New("not_found")
var ErrNotFound = Error{err: errNotFound, reason: "検索しましたが存在しませんでした。"}
