package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LooksOneOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 17h2V7h-4v2h2v8Zm9 4H3V3h18v18ZM5 19h14V5H5v14ZM5 5v14V5Z"/>`),
		g.Group(children),
	)
}