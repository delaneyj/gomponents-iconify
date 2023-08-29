package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionsFork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 4v8.5l3-3L9 13c1 1 1 2 1 2v6h4v-7s0-1-.53-2S12 10 12 10L9 6.58L11.5 4M18 4l-4.46 4.47L14 9s.93 1 1.47 2c.21.4.33.79.4 1.13L21 7"/>`),
		g.Group(children),
	)
}