package syscolor

import "image/color"

type colorProvider interface {
	SelectedBackground() (color.RGBA, error)
	SelectedForeground() (color.RGBA, error)
}

func SelectedBackground() (color.RGBA, error) {
	return provider.SelectedBackground()
}

func SelectedForeground() (color.RGBA, error) {
	return provider.SelectedForeground()
}
