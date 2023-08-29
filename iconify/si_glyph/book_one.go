package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 3.007C4.691 1.444 2.11 2.32 0 3.61v10.325c2.105-1.298 5.248-2.364 8-.946v-9.98zm0 9.995c3.629-1.463 5.919-.353 8 .937V3.621c-2.081-1.285-4.118-2.438-8-.845v10.226z"/>`),
		g.Group(children),
	)
}