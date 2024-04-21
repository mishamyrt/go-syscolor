# 🎨 syscolor [![Quality assurance](https://github.com/mishamyrt/go-syscolor/actions/workflows/quality-assurance.yaml/badge.svg)](https://github.com/mishamyrt/go-syscolor/actions/workflows/quality-assurance.yaml) [![GitHub release](https://img.shields.io/github/v/tag/mishamyrt/go-syscolor)](https://GitHub.com/mishamyrt/go-syscolor/releases/)

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

