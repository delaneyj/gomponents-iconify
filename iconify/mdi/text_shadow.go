package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextShadow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 3h13v3h-5v12H8V6H3V3m9 4h2v2h-2V7m3 0h2v2h-2V7m3 0h2v2h-2V7m-6 3h2v2h-2v-2m0 3h2v2h-2v-2m0 3h2v2h-2v-2m0 3h2v2h-2v-2Z"/>`),
		g.Group(children),
	)
}