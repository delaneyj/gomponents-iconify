package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatOverline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 5h14V3H5v2m4.62 11L12 9.67L14.37 16M11 7L5.5 21h2.25l1.12-3h6.25l1.13 3h2.25L13 7h-2Z"/>`),
		g.Group(children),
	)
}