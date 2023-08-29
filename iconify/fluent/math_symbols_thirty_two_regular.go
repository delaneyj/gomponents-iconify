package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MathSymbolsThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.998 2a1 1 0 0 1 1 1v4H13a1 1 0 1 1 0 2H8.998v4a1 1 0 1 1-2 0V9H3a1 1 0 1 1 0-2h3.998V3a1 1 0 0 1 1-1Zm5.71 25.293L9.413 23l4.293-4.293a1 1 0 0 0-1.414-1.414L8 21.586l-4.293-4.293a1 1 0 0 0-1.414 1.414L6.586 23l-4.293 4.293a1 1 0 1 0 1.414 1.414L8 24.414l4.293 4.293a1 1 0 0 0 1.414-1.414ZM19 7a1 1 0 1 0 0 2h10a1 1 0 1 0 0-2H19Zm0 15a1 1 0 1 0 0 2h10a1 1 0 1 0 0-2H19Zm7-3a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm-2 10a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/>`),
		g.Group(children),
	)
}