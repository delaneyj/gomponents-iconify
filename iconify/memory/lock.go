package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M10 12h2v1h1v2h-1v1h-2v-1H9v-2h1v-1M8 2h6v1h1v1h1v4h1v1h1v10h-1v1H5v-1H4V9h1V8h1V4h1V3h1V2m1 2v1H8v3h6V5h-1V4H9m7 6H6v8h10v-8Z"/>`),
		g.Group(children),
	)
}