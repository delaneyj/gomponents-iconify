package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Functions(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m5 18l7.68-6L5 6V4h14v2H8.263L16 12l-7.737 6H19v2H5v-2Z"/>`),
		g.Group(children),
	)
}