package syscolor

import "image/color"

type providerImpl struct{}

var provider colorProvider = &providerImpl{}

func (providerImpl) SelectedBackground() (color.RGBA, error) {
	return color.RGBA{}, ErrNotImplemented
}

func (providerImpl) SelectedForeground() (color.RGBA, error) {
	return color.RGBA{}, ErrNotImplemented
}
