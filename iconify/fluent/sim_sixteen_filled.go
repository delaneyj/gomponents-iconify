package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SimSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.5 1A2.5 2.5 0 0 0 3 3.5v9A2.5 2.5 0 0 0 5.5 15h5a2.5 2.5 0 0 0 2.5-2.5V5.328a2.5 2.5 0 0 0-.732-1.767l-1.829-1.829A2.5 2.5 0 0 0 8.672 1H5.5Zm1 7H8v5H6.5A1.5 1.5 0 0 1 5 11.5v-2A1.5 1.5 0 0 1 6.5 8ZM9 8h.5A1.5 1.5 0 0 1 11 9.5v2A1.5 1.5 0 0 1 9.5 13H9V8Z"/>`),
		g.Group(children),
	)
}