package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResizeHeight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256.236 504L120 367.764v-38.156h72v-144h-72V146.98L255.766 11.216L392 147.452v38.156h-72v144h72v38.627ZM159.1 361.608l97.137 97.137l97.137-97.137H288v-208h64.9l-97.134-97.137l-97.138 97.137H224v208Z"/>`),
		g.Group(children),
	)
}