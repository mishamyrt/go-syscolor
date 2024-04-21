package syscolor_test

import (
	"os"
	"testing"

	"github.com/mishamyrt/go-syscolor"
)

func TestIsGNOME(t *testing.T) {
	t.Parallel()
	previousEnv := os.Getenv("XDG_CURRENT_DESKTOP")
	os.Setenv("XDG_CURRENT_DESKTOP", "GNOME")
	isGNOME := syscolor.IsGNOME()
	if isGNOME != true {
		t.Errorf("expected true, got %v", isGNOME)
	}
	os.Setenv("XDG_CURRENT_DESKTOP", "KDE")
	isGNOME = syscolor.IsGNOME()
	if isGNOME != false {
		t.Errorf("expected false, got %v", isGNOME)
	}
	os.Setenv("XDG_CURRENT_DESKTOP", previousEnv)
}
