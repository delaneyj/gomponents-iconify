package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContentPaste(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 1a2.5 2.5 0 0 1 2.45 2H17a3 3 0 0 1 3 3v13a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6a3 3 0 0 1 3-3h3.05a2.5 2.5 0 0 1 2.45-2m1.41 2a1.495 1.495 0 0 0-2.82 0h2.82M6 4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2h11a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2h-1v3H7V4H6m2 0v2h7V4H8Z"/>`),
		g.Group(children),
	)
}