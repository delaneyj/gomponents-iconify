package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HdmiCableSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 0h9v5H3V0Zm3 3H5V2h1v1Zm4 0H9V2h1v1Z" clip-rule="evenodd"/><path fill="currentColor" d="M2 6h11v3.809l-2 1V13h-1v2H9v-2H6v2H5v-2H4v-2.191l-2-1V6Z"/>`),
		g.Group(children),
	)
}