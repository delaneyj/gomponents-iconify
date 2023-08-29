package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 14.5a3.5 3.5 0 0 1-3.5 3.5H7V5h4.5A3.5 3.5 0 0 1 15 8.5c0 1.1-.5 2.07-1.29 2.71C15.05 11.71 16 13 16 14.5M11.5 6H8v5h3.5A2.5 2.5 0 0 0 14 8.5A2.5 2.5 0 0 0 11.5 6m1 6H8v5h4.5a2.5 2.5 0 0 0 2.5-2.5a2.5 2.5 0 0 0-2.5-2.5Z"/>`),
		g.Group(children),
	)
}