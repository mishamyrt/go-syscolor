package syscolor

//#cgo pkg-config: gdk-3.0 gtk+-3.0
//#include "colors_linux.h"
import "C"
import "image/color"

type providerImpl struct{}

var provider colorProvider = &providerImpl{}

func (providerImpl) SelectedBackground() (color.RGBA, error) {
	cColor := C.getSelectedBackgroundColor()
	return color.RGBA{
		R: uint8(cColor.r),
		G: uint8(cColor.g),
		B: uint8(cColor.b),
		A: uint8(cColor.a),
	}, nil
}

func (providerImpl) SelectedForeground() (color.RGBA, error) {
	cColor := C.getSelectedForegroundColor()
	return color.RGBA{
		R: uint8(cColor.r),
		G: uint8(cColor.g),
		B: uint8(cColor.b),
		A: uint8(cColor.a),
	}, nil
}

func init() {
	C.init()
}
