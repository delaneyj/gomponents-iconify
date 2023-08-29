package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M-2.031 202.672V452.36h1204.062V202.672H-2.031zm0 363.781v430.875h1204.062V566.453H-2.031zm131.562 189.25h453.688v136.938H129.531V755.703z"/>`),
		g.Group(children),
	)
}