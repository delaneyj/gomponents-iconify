package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Message(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M2 1h18v1h1v14h-1v1H5v1H4v1H3v1H2v1H1V2h1V1m1 2v13h1v-1h15V3H3Z"/>`),
		g.Group(children),
	)
}