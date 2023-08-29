package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuUnfoldLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 18v2H3v-2h18ZM17.404 3.903L22 8.5l-4.596 4.596l-1.414-1.415L19.172 8.5L15.99 5.318l1.414-1.415ZM12 11v2H3v-2h9Zm0-7v2H3V4h9Z"/>`),
		g.Group(children),
	)
}