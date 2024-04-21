package syscolor

import "os"

// IsGNOME returns true if current desktop environment is GNOME.
// This function relies on XDG_CURRENT_DESKTOP, so result may be inaccurate
func IsGNOME() bool {
	return os.Getenv("XDG_CURRENT_DESKTOP") == "GNOME"
}
