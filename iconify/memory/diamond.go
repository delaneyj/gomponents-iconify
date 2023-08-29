package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Diamond(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M6 2h10v1h1v1h1v1h1v1h1v6h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8v-1H7v-1H6v-1H5v-1H4v-1H3v-1H2V6h1V5h1V4h1V3h1V2m9 3V4h-1v2h1v1h2V6h-1V5h-1m-3 1V4h-2v2H9v1h4V6h-1M8 6V4H7v1H6v1H5v1h2V6h1m-4 5h1v1h1v1h1v1h1v-2H7V9H4v2m6 1v4h2v-4h1V9H9v3h1m4 0v2h1v-1h1v-1h1v-1h1V9h-3v3h-1Z"/>`),
		g.Group(children),
	)
}