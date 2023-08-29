package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForPuertoRico(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2m21.677 12.3a27.913 27.913 0 0 1 5.692 11.8H27.514l-11.8-11.8h37.963M10.352 37.9l1.488-4.502l-3.932-2.873h4.917L14.3 26.1l1.52 4.426h4.872l-3.934 2.873l1.491 4.502l-3.949-2.768l-3.948 2.767m43.325 11.8H15.714l11.8-11.8h31.855a27.907 27.907 0 0 1-5.692 11.8"/>`),
		g.Group(children),
	)
}