package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToolsFlatHeadOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 21v-2h8v2H8Zm0-3l-1-7l2-8h6l2 8l-1 7H8Zm1.725-2h4.55l.55-4H9.15l.575 4ZM9.3 10h5.4l-1.25-5h-2.9L9.3 10Zm4.975 6h-4.55h4.55Z"/>`),
		g.Group(children),
	)
}