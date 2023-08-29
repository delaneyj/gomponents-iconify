package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Account(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M9 3h4v1h1v1h1v4h-1v1h-1v1H9v-1H8V9H7V5h1V4h1V3m1 5v1h2V8h1V6h-1V5h-2v1H9v2h1m-3 4h8v1h2v1h1v1h1v4H3v-4h1v-1h1v-1h2v-1m-1 4H5v1h12v-1h-1v-1h-2v-1H8v1H6v1Z"/>`),
		g.Group(children),
	)
}