package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuUnfoldFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 18v2H3v-2h18ZM17.05 3.55L22 8.5l-4.95 4.95v-9.9ZM12 11v2H3v-2h9Zm0-7v2H3V4h9Z"/>`),
		g.Group(children),
	)
}