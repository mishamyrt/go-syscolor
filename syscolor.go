// Package syscolor provides access to system colors
package syscolor

import "image/color"

type colorProvider interface {
	SelectedBackground() (color.RGBA, error)
	SelectedForeground() (color.RGBA, error)
}

// SelectedBackground returns native selected background color
func SelectedBackground() (color.RGBA, error) {
	return provider.SelectedBackground()
}

// SelectedForeground returns native selected foreground color
func SelectedForeground() (color.RGBA, error) {
	return provider.SelectedForeground()
}
