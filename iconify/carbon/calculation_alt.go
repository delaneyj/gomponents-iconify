package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalculationAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M19 13h6v2h-6zm-6 8h-2v-2H9v2H7v2h2v2h2v-2h2v-2zM7 9h6v2H7zm12 8h6v2h-6z"/><path fill="currentColor" d="M27 3H5a2.002 2.002 0 0 0-2 2v22a2.002 2.002 0 0 0 2 2h22a2.002 2.002 0 0 0 2-2V5a2.002 2.002 0 0 0-2-2ZM15 5v10H5V5ZM5 17h10v10H5Zm12 10V5h10v22Z"/>`),
		g.Group(children),
	)
}