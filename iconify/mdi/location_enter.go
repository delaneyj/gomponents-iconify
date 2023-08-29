package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationEnter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14 12l-4-4v3H2v2h8v3m12-4a10 10 0 0 1-19.54 3h2.13a8 8 0 1 0 0-6H2.46A10 10 0 0 1 22 12Z"/>`),
		g.Group(children),
	)
}