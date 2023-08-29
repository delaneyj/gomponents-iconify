package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PagesLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 8v12h14V8H5Zm0-2h14V4H5v2Zm15 16H4a1 1 0 0 1-1-1V3a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1ZM7 10h4v4H7v-4Zm0 6h10v2H7v-2Zm6-5h4v2h-4v-2Z"/>`),
		g.Group(children),
	)
}