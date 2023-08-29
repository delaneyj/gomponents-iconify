package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadioboxMarked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M8 2h6v1h2v1h1v1h1v1h1v2h1v6h-1v2h-1v1h-1v1h-1v1h-2v1H8v-1H6v-1H5v-1H4v-1H3v-2H2V8h1V6h1V5h1V4h1V3h2V2m1 2v1H7v1H6v1H5v2H4v4h1v2h1v1h1v1h2v1h4v-1h2v-1h1v-1h1v-2h1V9h-1V7h-1V6h-1V5h-2V4H9m4 3v1h1v1h1v4h-1v1h-1v1H9v-1H8v-1H7V9h1V8h1V7h4Z"/>`),
		g.Group(children),
	)
}