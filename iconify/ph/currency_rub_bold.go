package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyRubBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M148 156a64 64 0 0 0 0-128H88a12 12 0 0 0-12 12v92H56a12 12 0 0 0 0 24h20v16H56a12 12 0 0 0 0 24h20v20a12 12 0 0 0 24 0v-20h44a12 12 0 0 0 0-24h-44v-16ZM100 52h48a40 40 0 0 1 0 80h-48Z"/>`),
		g.Group(children),
	)
}