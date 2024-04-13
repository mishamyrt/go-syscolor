#include "colors_linux.h"

void init() {
    gtk_init(NULL, NULL);
}

// Available colors:
// theme_fg_color
// theme_text_color
// theme_bg_color
// theme_base_color
// theme_selected_bg_color
// theme_selected_fg_color
// insensitive_bg_color
// insensitive_fg_color
// insensitive_base_color
// theme_unfocused_fg_color
// theme_unfocused_text_color
// theme_unfocused_bg_color
// theme_unfocused_base_color
// theme_unfocused_selected_bg_color
// theme_unfocused_selected_fg_color
// unfocused_insensitive_color
// borders
// unfocused_borders
// warning_color
// error_color
// success_color

rgba_t gdkRGBAToRGBA(GdkRGBA *color) {
    rgba_t rgba;
    rgba.r = color->red * 255;
    rgba.g = color->green * 255;
    rgba.b = color->blue * 255;
    rgba.a = color->alpha * 255;
    return rgba;
}

rgba_t getThemeColorByName(const char *color_name) {
    GtkStyleContext *context = gtk_widget_get_style_context(GTK_WIDGET(gtk_window_new(GTK_WINDOW_TOPLEVEL)));
    GdkRGBA color;
    gtk_style_context_lookup_color(context, color_name, &color);
    return gdkRGBAToRGBA(&color);
}

rgba_t getSelectedForegroundColor() {
    return getThemeColorByName("theme_selected_fg_color");
}

rgba_t getSelectedBackgroundColor() {
    return getThemeColorByName("theme_selected_bg_color");
}
