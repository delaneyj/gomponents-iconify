package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoffeeCup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 4a6 6 0 1 0 12 0V2H2v2zm14-4h1a3 3 0 0 1 0 6h-1.252A8 8 0 0 1 0 4V0h16zm0 4h1a1 1 0 0 0 0-2h-1v2z"/>`),
		g.Group(children),
	)
}