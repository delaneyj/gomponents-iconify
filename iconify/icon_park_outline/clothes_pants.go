package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClothesPants(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M33 44h9L38 4H10L6 44h9l9-25l9 25ZM24 4v5.5"/><path d="M17 4s0 6-2 8s-6.1 3-6.1 3M31 4s0 6 2 8s6.1 3 6.1 3"/></g>`),
		g.Group(children),
	)
}