package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalculatorOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M4 4.5h7m-7 4h1m2 0h1m2 0h1m-7 3h1m2 0h1m2 0h1m-8.5 3h10a1 1 0 0 0 1-1v-12a1 1 0 0 0-1-1h-10a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1Z"/>`),
		g.Group(children),
	)
}