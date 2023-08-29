package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransitEnterexit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 18V8h3v4.75l6.8-6.8l2.2 2.2L11.15 15H16v3H6Z"/>`),
		g.Group(children),
	)
}