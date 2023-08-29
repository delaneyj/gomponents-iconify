package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalculatorFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 24H56a16 16 0 0 0-16 16v176a16 16 0 0 0 16 16h144a16 16 0 0 0 16-16V40a16 16 0 0 0-16-16ZM88 200a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm0-40a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm40 40a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm0-40a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm40 40a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm0-40a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm16-56a8 8 0 0 1-8 8H80a8 8 0 0 1-8-8V64a8 8 0 0 1 8-8h96a8 8 0 0 1 8 8Z"/>`),
		g.Group(children),
	)
}