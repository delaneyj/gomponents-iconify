package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lightbulb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M8 19h6v2H8v-2m0-1v-4H7v-1H6v-1H5v-1H4V5h1V4h1V3h1V2h1V1h6v1h1v1h1v1h1v1h1v6h-1v1h-1v1h-1v1h-1v4H8m5-6h1v-1h1v-1h1V6h-1V5h-1V4h-1V3H9v1H8v1H7v1H6v4h1v1h1v1h1v1h1v3h2v-3h1v-1Z"/>`),
		g.Group(children),
	)
}