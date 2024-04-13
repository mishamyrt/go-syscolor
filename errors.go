package syscolor

import "errors"

var ErrNotSupported = errors.New("requested color type is not supported on this platform")

var ErrNotImplemented = errors.New("not implemented")
