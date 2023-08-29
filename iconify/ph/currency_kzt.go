package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyKzt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 96a8 8 0 0 1-8 8h-64v112a8 8 0 0 1-16 0V104H56a8 8 0 0 1 0-16h144a8 8 0 0 1 8 8ZM56 64h144a8 8 0 0 0 0-16H56a8 8 0 0 0 0 16Z"/>`),
		g.Group(children),
	)
}