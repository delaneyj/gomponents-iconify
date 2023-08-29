package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyFranc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 21v-3H5v-2h2V3h11v2H9v6h8v2H9v3h4v2H9v3H7Z"/>`),
		g.Group(children),
	)
}