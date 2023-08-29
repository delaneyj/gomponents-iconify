package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineEndSquareOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 15.5h7v-7h-7v7Zm-2 2V13H2v-2h9V6.5h11v11H11Zm5.5-5.5Z"/>`),
		g.Group(children),
	)
}