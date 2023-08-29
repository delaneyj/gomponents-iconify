package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalculatorSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 1.5A1.5 1.5 0 0 1 2.5 0h10A1.5 1.5 0 0 1 14 1.5v12a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5v-12ZM4 5h7V4H4v1Zm0 4h1V8H4v1Zm4 0H7V8h1v1Zm2 0h1V8h-1v1Zm-5 3H4v-1h1v1Zm2 0h1v-1H7v1Zm4 0h-1v-1h1v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}