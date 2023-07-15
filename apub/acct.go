package apub

import (
	"errors"
	"strings"
)

var ErrMalformedAcct = errors.New("[apub.Acct] malformed style acct")

type Acct struct {
	acct string
	at   int
}

func ParseAcct(s string) Acct {
	s, _ = strings.CutPrefix(s, "acct:")
	s, _ = strings.CutPrefix(s, "@")
	at := strings.Index(s, "@")
	return Acct{s, at}
}

func ParseFullAcct(s string) (Acct, error) {
	a := ParseAcct(s)
	if a.at == -1 || a.at == len(a.acct)-1 {
		return Acct{}, ErrMalformedAcct
	}
	return a, nil
}

func (a Acct) Username() string {
	if a.at == -1 {
		return a.acct
	} else {
		return a.acct[:a.at]
	}
}

func (a Acct) Domain() string {
	if a.at == -1 {
		return ""
	} else {
		return a.acct[a.at+1:]
	}
}

func (a Acct) String() string {
	return a.acct
}
