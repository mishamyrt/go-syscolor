#import <colors_darwin.h>

rgba_t NSColorToRGBA(NSColor *color) {
    rgba_t result;
    result.r = color.redComponent * 255;
    result.g = color.greenComponent * 255;
    result.b = color.blueComponent * 255;
    result.a = color.alphaComponent * 255;
    return result;
}

rgba_t getSelectedControlColor() {
    return NSColorToRGBA(
        [[NSColor selectedControlColor] colorUsingColorSpace:[NSColorSpace genericRGBColorSpace]]
    );
}

rgba_t getSelectedControlTextColor() {
    return NSColorToRGBA(
        [[NSColor selectedControlTextColor] colorUsingColorSpace:[NSColorSpace genericRGBColorSpace]]
    );
}

// int getControlAccentColor() {
//     return nscolorToInt(
//         [[NSColor controlAccentColor] colorUsingColorSpace:[NSColorSpace genericRGBColorSpace]]
//     );
// }