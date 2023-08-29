package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coffee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M1 20v-2h16v2H1M2 3h17v1h1v6h-1v1h-3v3h-1v1h-1v1H4v-1H3v-1H2V3m14 2v4h2V5h-2M4 5v8h1v1h8v-1h1V5H4Z"/>`),
		g.Group(children),
	)
}