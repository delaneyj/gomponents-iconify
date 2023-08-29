package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DialogLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor"><path stroke-width="1.5" d="M10 22a8 8 0 1 0-7.22-4.55c.172.36.232.766.13 1.15l-.328 1.227a1.3 1.3 0 0 0 1.591 1.591L5.4 21.09a1.671 1.671 0 0 1 1.15.13c1.045.5 2.215.78 3.451.78Z"/><path stroke-width="1.5" d="M18 14.502a6.45 6.45 0 0 0 .198-.087c.362-.165.768-.227 1.153-.124l.476.127a1.3 1.3 0 0 0 1.591-1.591l-.127-.476c-.103-.385-.04-.791.125-1.153A6.5 6.5 0 1 0 9.5 5.996"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6.518 14h.01m3.481 0h.009m3.482 0h.009"/></g>`),
		g.Group(children),
	)
}