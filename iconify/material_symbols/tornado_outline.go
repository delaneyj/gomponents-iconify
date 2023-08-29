package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TornadoOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 3h22L12 22L1 3Zm3.475 2L6.2 8h11.6l1.725-3H4.475Zm2.9 5L9.1 13h5.8l1.725-3h-9.25Zm2.9 5L12 18l1.725-3h-3.45Z"/>`),
		g.Group(children),
	)
}