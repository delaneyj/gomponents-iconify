package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Database(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M7 2h8v1h2v1h1v1h1v12h-1v1h-1v1h-2v1H7v-1H5v-1H4v-1H3V5h1V4h1V3h2V2m1 14v-1H6v-1H5v2h1v1h2v1h6v-1h2v-1h1v-2h-1v1h-2v1H8m0-5v-1H6V9H5v3h2v1h2v1h4v-1h2v-1h2V9h-1v1h-2v1H8m1-3v1h4V8h2V7h2V6h-1V5h-2V4H8v1H6v1H5v1h2v1h2Z"/>`),
		g.Group(children),
	)
}