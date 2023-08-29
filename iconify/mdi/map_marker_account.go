package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapMarkerAccount(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C8.14 2 5 5.14 5 9c0 5.25 7 13 7 13s7-7.75 7-13c0-3.86-3.14-7-7-7m0 2a2 2 0 1 1 .001 4.001A2 2 0 0 1 12 4m0 10c-1.67 0-3.14-.85-4-2.15c0-1.32 2.67-2.05 4-2.05s4 .73 4 2.05A4.783 4.783 0 0 1 12 14Z"/>`),
		g.Group(children),
	)
}