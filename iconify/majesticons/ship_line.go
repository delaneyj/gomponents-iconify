package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShipLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 16l2-4l-4-.889M12 10v5m0-5l-5 1.111M12 10l5 1.111M5 16l-2-4l4-.889m0 0L6 6h4m7 5.111L18 6h-4m0 0V3h-4v3m4 0h-4M3 20l1.245-.498a2.57 2.57 0 0 1 2.38.248v0a2.57 2.57 0 0 0 3.36-.446l.035-.04a2.631 2.631 0 0 1 3.96 0l.036.04a2.57 2.57 0 0 0 3.36.446v0a2.57 2.57 0 0 1 2.38-.248L21 20"/>`),
		g.Group(children),
	)
}