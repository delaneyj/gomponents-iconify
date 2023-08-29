package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TerminalFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11 12l-7.071 7.07l-1.414-1.413L8.172 12L2.514 6.343l1.414-1.414l7.07 7.07Zm0 7h10v2H11v-2Z"/>`),
		g.Group(children),
	)
}