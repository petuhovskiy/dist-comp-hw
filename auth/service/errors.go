package service

import "errors"

var ErrLoginIsRequired = errors.New("provide email or phone")
var ErrUnknownNotifyType = errors.New("unknown notify type")