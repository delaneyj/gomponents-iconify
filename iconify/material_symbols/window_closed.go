package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowClosed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21V3h16v18H4Zm3-10h4v-1h2v1h4V6H7v5Zm0 7h10v-5H7v5Zm-1 1h12V5H6v14Z"/>`),
		g.Group(children),
	)
}