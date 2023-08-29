package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuFoldFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 18v2H3v-2h18ZM6.95 3.55v9.9L2 8.5l4.95-4.95ZM21 11v2h-9v-2h9Zm0-7v2h-9V4h9Z"/>`),
		g.Group(children),
	)
}