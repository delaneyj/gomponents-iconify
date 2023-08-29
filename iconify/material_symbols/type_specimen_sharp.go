package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TypeSpecimenSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.8 14.5h1.6l.8-2.3h3.65l.8 2.3h1.55l-3.4-9h-1.6l-3.4 9Zm2.85-3.6l1.3-3.75h.1l1.3 3.75h-2.7ZM6 18V2h16v16H6Zm-4 4V6h2v14h14v2H2Z"/>`),
		g.Group(children),
	)
}