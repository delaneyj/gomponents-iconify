package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuFoldLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 18v2H3v-2h18ZM6.596 3.903L8.01 5.318L4.828 8.5l3.182 3.18l-1.414 1.415L2 8.5l4.596-4.597ZM21 11v2h-9v-2h9Zm0-7v2h-9V4h9Z"/>`),
		g.Group(children),
	)
}