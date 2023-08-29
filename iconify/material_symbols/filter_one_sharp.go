package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterOneSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 15h2V5h-4v2h2v8Zm-8 3V2h16v16H6Zm-4 4V6h2v14h14v2H2Z"/>`),
		g.Group(children),
	)
}