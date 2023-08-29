package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeLow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M4.34 0L3 2H1v4h2l1.34 2H5V0h-.66zM6 3v2c.09 0 .18-.01.25-.03c.43-.11.75-.51.75-.97a.997.997 0 0 0-1-1z"/>`),
		g.Group(children),
	)
}