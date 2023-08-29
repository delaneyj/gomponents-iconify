package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpCoffeeMaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 7V4h2V2H4v20h16v-2h-4.03A4.966 4.966 0 0 0 18 16v-5H8v5c0 1.64.81 3.09 2.03 4H6V4h2v3h10z"/><circle cx="13" cy="9" r="1" fill="currentColor"/>`),
		g.Group(children),
	)
}