package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NextTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M23.5 3.75a.75.75 0 0 1 1.5 0v20.5a.75.75 0 0 1-1.5 0V3.75ZM3 5.254C3 3.438 5.041 2.37 6.533 3.406l12.504 8.68a2.25 2.25 0 0 1 .013 3.688l-12.504 8.81C5.056 25.634 3 24.57 3 22.745V5.254Z"/>`),
		g.Group(children),
	)
}