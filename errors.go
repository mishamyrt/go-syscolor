package syscolor

import "errors"

// ErrNotSupported is returned if the requested color type is not supported on this platform
var ErrNotSupported = errors.New("requested color type is not supported on this platform")

// ErrNotImplemented is returned if the requested color type is not implemented
var ErrNotImplemented = errors.New("not implemented")
