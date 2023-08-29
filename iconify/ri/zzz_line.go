package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZzzLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 11v2l-5.327 6H11v2H3v-2l5.326-6H3v-2h8Zm10-8v2l-5.327 6H21v2h-8v-2l5.326-6H13V3h8Z"/>`),
		g.Group(children),
	)
}