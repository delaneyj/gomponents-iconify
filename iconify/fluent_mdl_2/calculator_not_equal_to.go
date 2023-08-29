package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalculatorNotEqualTo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1432 640l-747 640h1235v128H536l-430 369l-84-98l317-271H0v-128h488l747-640H0V512h1384l430-369l84 98l-317 271h339v128h-488z"/>`),
		g.Group(children),
	)
}