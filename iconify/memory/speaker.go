package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Speaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M4 1h14v1h1v18h-1v1H4v-1H3V2h1V1m1 2v16h12V3H5m4 2h1V4h2v1h1v2h-1v1h-2V7H9V5m0 13v-1H8v-1H7v-4h1v-1h1v-1h4v1h1v1h1v4h-1v1h-1v1H9m1-5H9v2h1v1h2v-1h1v-2h-1v-1h-2v1Z"/>`),
		g.Group(children),
	)
}