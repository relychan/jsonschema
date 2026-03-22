package jsonschema

import "errors"

// ErrIDInvalidHTTPScheme error occurs when the scheme of the ID URL isn't http(s).
var ErrIDInvalidHTTPScheme = errors.New("unexpected HTTP scheme")
