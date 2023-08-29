package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShirtFoldedBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 44h-23l-12.51-12.49a12 12 0 0 0-8.22-3.51h-57.1a12 12 0 0 0-7.66 3.47L79 44H56a20 20 0 0 0-20 20v152a20 20 0 0 0 20 20h144a20 20 0 0 0 20-20V64a20 20 0 0 0-20-20ZM96 61l1.48-1.48L112.92 86L96 102.48Zm32 3.21L120.89 52h14.22ZM160 61v41.51L143.08 86l15.44-26.47ZM60 68h12v44a20 20 0 0 0 34.08 14.21l9.92-9.71V212H60Zm136 144h-56v-95.5l9.92 9.69A20 20 0 0 0 184 112V68h12Z"/>`),
		g.Group(children),
	)
}