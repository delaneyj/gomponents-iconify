package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PergolaOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3q0-.375.275-.688T4 2q.45 0 .725.313T5 3v1h14V3q0-.375.313-.688T20 2q.375 0 .688.313T21 3v18h-2V10H5v11H3ZM5 8h14V6H5v2Zm6 13v-3H8v-2h8v2h-3v3h-2ZM5 8V6v2Z"/>`),
		g.Group(children),
	)
}