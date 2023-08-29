package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellSignalThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 20H4.731a.731.731 0 0 1-.517-1.249L18.751 4.214A.731.731 0 0 1 20 4.731V20zm-8 0v-9"/>`),
		g.Group(children),
	)
}