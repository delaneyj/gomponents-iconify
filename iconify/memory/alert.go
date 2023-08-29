package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M20 20H2v-1H1v-4h1v-2h1v-2h1V9h1V7h1V5h1V3h1V2h6v1h1v2h1v2h1v2h1v2h1v2h1v2h1v4h-1v1M9 6H8v2H7v2H6v2H5v2H4v2H3v2h16v-2h-1v-2h-1v-2h-1v-2h-1V8h-1V6h-1V4H9v2m1 1h2v6h-2V7m0 7h2v2h-2v-2Z"/>`),
		g.Group(children),
	)
}