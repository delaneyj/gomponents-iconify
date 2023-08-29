package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m192 5l192 86v128q0 89-55 162.5T192 475q-82-20-137-93.5T0 219V91zm-43 342l171-171l-30-30l-141 140l-55-55l-30 30z"/>`),
		g.Group(children),
	)
}