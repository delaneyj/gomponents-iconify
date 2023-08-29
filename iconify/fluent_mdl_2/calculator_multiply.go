package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalculatorMultiply(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1805 205l-755 755l755 755l-90 90l-755-755l-755 755l-90-90l755-755l-755-755l90-90l755 755l755-755l90 90z"/>`),
		g.Group(children),
	)
}