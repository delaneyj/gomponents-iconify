package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyKztBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M212 100a12 12 0 0 1-12 12h-60v100a12 12 0 0 1-24 0V112H56a12 12 0 0 1 0-24h144a12 12 0 0 1 12 12ZM56 64h144a12 12 0 0 0 0-24H56a12 12 0 0 0 0 24Z"/>`),
		g.Group(children),
	)
}