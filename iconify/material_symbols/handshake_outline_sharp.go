package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandshakeOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.85 23l-9.875-9.875l7.55-7.55l3.65 3.65l.7-.7l-5.35-5.35L1.7 10l.975.975L1.25 12.4l-2.35-2.35L8.575.375l3.1 3.1l3.1-3.1l9.85 9.85L11.85 23Zm.025-2.8l9.925-9.975l-7.05-7.05l-1.675 1.675l3.65 3.65l-3.55 3.55l-3.65-3.65L4.8 13.125l.725.725l3.725-3.725l1.4 1.4l-3.725 3.725l.7.7l3.725-3.725l1.425 1.425l-3.725 3.725l.7.7l3.725-3.725l1.4 1.4l-3.725 3.725l.725.725Zm-2.9-12.825Z"/>`),
		g.Group(children),
	)
}