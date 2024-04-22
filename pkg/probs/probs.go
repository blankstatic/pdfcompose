package probs

import "errors"

var (
	errReadiness = errors.New("not ready")
	errLiveness  error
)

func SetLivenessErr(e error) {
	errLiveness = e
}

func GetLivenessErr() error {
	return errLiveness
}

func SetReadinessErr(e error) {
	errReadiness = e
}

func GetReadinessErr() error {
	return errReadiness
}
