package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TruckTrailer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 15v2H10a3 3 0 0 1-3 3a3 3 0 0 1-3-3H2V6a2 2 0 0 1 2-2h13a2 2 0 0 1 2 2v9h3M7 16a1 1 0 0 0-1 1a1 1 0 0 0 1 1a1 1 0 0 0 1-1a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}