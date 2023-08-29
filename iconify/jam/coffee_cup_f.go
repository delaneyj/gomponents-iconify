package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoffeeCupF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 0h17a3 3 0 0 1 0 6h-1.252A8 8 0 0 1 0 4V0zm16 4h1a1 1 0 0 0 0-2h-1v2z"/>`),
		g.Group(children),
	)
}