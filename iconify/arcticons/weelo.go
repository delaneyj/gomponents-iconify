package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Weelo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" d="m22.425 23.614l-5.07-6.941c-.129-.192-.988-1.423-2.132-1.567c-.883-.11-6.98 0-9.387 0c-1.177 0-1.609.318-1.166.982c.58.87 8.451 12.106 11.23 16.031c.616.87 2.264.859 2.773.178a36139.26 36139.26 0 0 1 11.956-15.95c.13-.192.726-1.064 1.868-1.207c.885-.11 9.523-.083 9.523-.083s1.978-.104 1.36.918c-.415.688-8.747 12.41-11.563 16.337c-.552.769-1.867.863-2.425.145l-4.185-4.996"/>`),
		g.Group(children),
	)
}