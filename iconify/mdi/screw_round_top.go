package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScrewRoundTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.5 17v2L12 22l-1.5-3l3-2m1-10.7l-1 .7V6h-3v3l-1 .7v1l5-3.3V6.3m0 4l-1 .7V9l-3 2v2l-1 .7v1l5-3.3v-1.1m0 4l-1 .7v-2l-3 2v2l-1 .7v1l5-3.3v-1.1M7 5h10s-1-3-5-3s-5 3-5 3Z"/>`),
		g.Group(children),
	)
}