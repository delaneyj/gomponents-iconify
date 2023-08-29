package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarberClippers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M10 8h28v9l-5 7v12s0 8-9 8s-9-8-9-8V24l-5-7V8Zm5-4v6m6-6v6m6-6v6"/><rect width="6" height="10" x="21" y="28" rx="3"/><path d="M10 17h28M33 4v6"/></g>`),
		g.Group(children),
	)
}