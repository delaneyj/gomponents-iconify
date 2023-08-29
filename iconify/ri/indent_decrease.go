package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IndentDecrease(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 4h18v2H3V4Zm0 15h18v2H3v-2Zm8-5h10v2H11v-2Zm0-5h10v2H11V9Zm-8 3.5L7 9v7l-4-3.5Z"/>`),
		g.Group(children),
	)
}