package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExclusiveGateway(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linejoin="round" d="M22.8 4.201L4.413 22.586a2 2 0 0 0 0 2.828L22.8 43.8a2 2 0 0 0 2.829 0l18.384-18.385a2 2 0 0 0 0-2.828L25.628 4.2a2 2 0 0 0-2.829 0Z"/><path stroke-linecap="round" d="m18.043 29.987l12-11.962m-11.99-.009l11.98 11.98"/></g>`),
		g.Group(children),
	)
}