package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chargepoint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.606 24H5.5m35-18.5h-33c-1.1 0-2 .9-2 2v33c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2v-33c0-1.1-.9-2-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.394 19.394A6.394 6.394 0 0 0 24 13h0a6.394 6.394 0 0 0-6.394 6.394v9.212A6.394 6.394 0 0 0 24 35h0a6.394 6.394 0 0 0 6.394-6.394"/>`),
		g.Group(children),
	)
}