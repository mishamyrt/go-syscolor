package syscolor

//#cgo CFLAGS: -x objective-c
//#cgo LDFLAGS: -framework Foundation -framework AppKit
//#include "colors_darwin.h"
import "C"
import "image/color"

type providerImpl struct{}

var provider colorProvider = &providerImpl{}

func (providerImpl) SelectedBackground() (color.RGBA, error) {
	cColor := C.getSelectedControlColor()
	return color.RGBA{
		R: uint8(cColor.r),
		G: uint8(cColor.g),
		B: uint8(cColor.b),
		A: uint8(cColor.a),
	}, nil
}

func (providerImpl) SelectedForeground() (color.RGBA, error) {
	cColor := C.getSelectedControlTextColor()
	return color.RGBA{
		R: uint8(cColor.r),
		G: uint8(cColor.g),
		B: uint8(cColor.b),
		A: uint8(cColor.a),
	}, nil
}

// func (providerImpl) ControlAccent() color.RGBA {
// 	return RGBAFromInt(
// 		int(C.getControlAccentColor()),
// 	)
// }
