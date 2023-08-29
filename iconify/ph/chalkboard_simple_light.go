package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChalkboardSimpleLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M26 168V56a14 14 0 0 1 14-14h176a14 14 0 0 1 14 14v112a6 6 0 0 1-12 0V56a2 2 0 0 0-2-2H40a2 2 0 0 0-2 2v112a6 6 0 0 1-12 0Zm220 32a6 6 0 0 1-6 6H16a6 6 0 0 1 0-12h98v-26a6 6 0 0 1 6-6h64a6 6 0 0 1 6 6v26h50a6 6 0 0 1 6 6Zm-120-6h52v-20h-52Z"/>`),
		g.Group(children),
	)
}