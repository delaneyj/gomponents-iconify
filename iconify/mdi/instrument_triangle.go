package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InstrumentTriangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 2h-1v4.2c-.1.1-.2.1-.3.3L2.1 20.7c-.3.6.1 1.3.8 1.3H16v-2H4.8L11 9.2l5.7 10l1.7-1l-6.6-11.8l-.3-.3V2M21 6h-1v12l-.5 4h2l-.5-4V6Z"/>`),
		g.Group(children),
	)
}