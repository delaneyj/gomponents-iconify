package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyKztLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M206 96a6 6 0 0 1-6 6h-66v114a6 6 0 0 1-12 0V102H56a6 6 0 0 1 0-12h144a6 6 0 0 1 6 6ZM56 62h144a6 6 0 0 0 0-12H56a6 6 0 0 0 0 12Z"/>`),
		g.Group(children),
	)
}