package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyEuroCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12a4 4 0 0 0 6.4 3.2M8 12a4 4 0 0 1 6.4-3.2M8 12h3m-3 0H7m14 0a9 9 0 1 1-18 0a9 9 0 0 1 18 0Z"/>`),
		g.Group(children),
	)
}