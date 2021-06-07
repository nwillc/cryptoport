package model

import "github.com/nwillc/cryptoport/pkg/externalapi/crypto"

// Config persisted configuration.
type Config struct {
	AppID     crypto.AppID
	Portfolio Portfolio
}
