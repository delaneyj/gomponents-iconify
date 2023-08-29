package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M112 171q-37 0-64 21v-21q0-53 37.5-90.5T176 43q21 0 21-22q0-21-21-21q-70 0-120.5 50T5 171v106q0 45 31 76t76 31t76-31t31-76q0-44-31-75t-76-31zm256 0q-37 0-64 21v-21q0-53 37.5-90.5T432 43q21 0 21-22q0-21-21-21q-70 0-120.5 50T261 171v106q0 45 31 76t76 31t76-31t31-76q0-44-31-75t-76-31z"/>`),
		g.Group(children),
	)
}