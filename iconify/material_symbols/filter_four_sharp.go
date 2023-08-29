package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterFourSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 18V2h16v16H6Zm-4 4V6h2v14h14v2H2Zm13-7h2V5h-2v4h-2V5h-2v6h4v4Z"/>`),
		g.Group(children),
	)
}