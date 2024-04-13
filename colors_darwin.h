#import <AppKit/AppKit.h>\

typedef struct {
    uint8_t r;
    uint8_t g;
    uint8_t b;
    uint8_t a;
} rgba_t;

rgba_t getSelectedControlTextColor();
rgba_t getSelectedControlColor();

// int getControlAccentColor();

// Not implemented yet
// int getControlBackgroundColor();
// int getControlColor();
// int	controlTextColor();
// int getDisabledControlTextColor();
// int getFindHighlightColor();
// int getGridColor();
// int getHeaderTextColor();
// int getHighlightColor();