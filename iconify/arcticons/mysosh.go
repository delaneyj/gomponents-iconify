package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mysosh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m43.5 31.15l-37.8 2.7l-1.2-17l37.8-2.7Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21 29.55h0a2.69 2.69 0 0 1-2.7-2.7v-1.8a2.69 2.69 0 0 1 2.7-2.7h0a2.69 2.69 0 0 1 2.7 2.7v1.8a2.75 2.75 0 0 1-2.7 2.7Zm13.4-11v11m0-4.5a2.69 2.69 0 0 1 2.7-2.7h0a2.69 2.69 0 0 1 2.7 2.7v4.5m-31.5-1.2a3 3 0 0 0 2.7 1.2h1.6a2.69 2.69 0 0 0 2.7-2.7h0a2.69 2.69 0 0 0-2.7-2.7h-1.8a2.69 2.69 0 0 1-2.7-2.7h0a2.69 2.69 0 0 1 2.7-2.7h1.6A3 3 0 0 1 15.1 20m11.7 9a3.17 3.17 0 0 0 2.3.6h.6a1.79 1.79 0 0 0 1.8-1.8h0a1.79 1.79 0 0 0-1.8-1.8h-1.2a1.79 1.79 0 0 1-1.8-1.8h0a1.79 1.79 0 0 1 1.8-1.8h.6a3.34 3.34 0 0 1 2.3.6"/>`),
		g.Group(children),
	)
}