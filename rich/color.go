package rich

import (
	"fmt"
	"strconv"
	"strings"
)

var colorNames = map[string][4]int{
	"black":         {0, 0, 0, 0},
	"red":           {255, 0, 0, 1},
	"green":         {0, 255, 0, 2},
	"yellow":        {255, 255, 0, 3},
	"blue":          {0, 0, 255, 4},
	"magenta":       {255, 0, 255, 5},
	"cyan":          {0, 255, 255, 6},
	"white":         {255, 255, 255, 7},
	"brightblack":   {128, 128, 128, 8},
	"brightred":     {255, 0, 0, 9},
	"brightgreen":   {0, 255, 0, 10},
	"brightyellow":  {255, 255, 0, 11},
	"brightblue":    {0, 0, 255, 12},
	"brightmagenta": {255, 0, 255, 13},
	"brightcyan":    {0, 255, 255, 14},
	"brightwhite":   {255, 255, 255, 15},
}

type ColorMode int

const (
	Mode8Bit ColorMode = iota
	Mode256
	ModeTrueColor
)

type Color struct {
	Name      string
	R         int
	G         int
	B         int
	AnsiIndex int
	Mode      ColorMode
}

func FromColorName(name string) Color {
	rgb, ok := colorNames[name]
	if !ok {
		// Default to white if color name not found
		rgb = colorNames["white"]
	}
	return Color{
		Name:      name,
		R:         rgb[0],
		G:         rgb[1],
		B:         rgb[2],
		AnsiIndex: rgb[3],
		Mode:      ModeTrueColor, // Default to true color
	}
}

func FromHex(hex string) Color {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) != 6 {
		// Invalid hex format, default to white
		return FromColorName("white")
	}

	r, err1 := strconv.ParseInt(hex[0:2], 16, 0)
	g, err2 := strconv.ParseInt(hex[2:4], 16, 0)
	b, err3 := strconv.ParseInt(hex[4:6], 16, 0)
	if err1 != nil || err2 != nil || err3 != nil {
		// Parsing error, default to white
		return FromColorName("white")
	}

	return Color{
		R:    int(r),
		G:    int(g),
		B:    int(b),
		Mode: ModeTrueColor,
	}
}

func FromRGB(r, g, b int) Color {
	return Color{
		R:    r,
		G:    g,
		B:    b,
		Mode: ModeTrueColor,
	}
}

func (c Color) ToANSI() string {
	switch c.Mode {
	case Mode8Bit:
		return fmt.Sprintf("\033[3%dm", c.AnsiIndex)
	case Mode256:
		return fmt.Sprintf("\033[38;5;%dm", c.AnsiIndex)
	case ModeTrueColor:
		return fmt.Sprintf("\033[38;2;%d;%d;%dm", c.R, c.G, c.B)
	default:
		return ""
	}
}
