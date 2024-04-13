#include <gtk-3.0/gtk/gtk.h>

typedef struct {
    uint8_t r;
    uint8_t g;
    uint8_t b;
    uint8_t a;
} rgba_t;

void init();

rgba_t getSelectedBackgroundColor();
rgba_t getSelectedForegroundColor();
