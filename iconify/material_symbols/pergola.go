package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pergola(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3q0-.375.275-.688T4 2q.45 0 .725.313T5 3v1h14V3q0-.375.313-.688T20 2q.375 0 .688.313T21 3v18h-2V10H5v11H3Zm8 0v-3H8v-2h8v2h-3v3h-2Z"/>`),
		g.Group(children),
	)
}