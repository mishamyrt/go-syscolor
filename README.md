# 🎨 syscolor

A library that allows you to get system colors to build interfaces that look native.

The library implements support for:

- macOS
- Linux
    - GTK
    - ~~KDE~~
- ~~Windows~~

## Installation

```bash
go get -u github.com/mishamyrt/syscolor
```

## Usage

If your application is planned to run in Linux, first check if Unity or GNOME is used. I haven't found any stable methods of getting color in Plasma.

```go
func GetAccent() color.RGBA {
    var systemColor color.RGBA
    if runtime.GOOS == "darwin" || syscolor.IsGNOME() {
        systemColor, _ = syscolor.SelectedBackground()
    }
    return systemColor
}
```

