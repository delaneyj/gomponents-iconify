package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CastConnected(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 3H3c-1.11 0-2 .89-2 2v3h2V5h18v14h-7v2h7a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2M1 10v2a9 9 0 0 1 9 9h2c0-6.08-4.93-11-11-11m18-3H5v1.63c3.96 1.28 7.09 4.41 8.37 8.37H19M1 14v2a5 5 0 0 1 5 5h2a7 7 0 0 0-7-7m0 4v3h3a3 3 0 0 0-3-3Z"/>`),
		g.Group(children),
	)
}