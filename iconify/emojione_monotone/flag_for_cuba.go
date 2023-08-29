package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForCuba(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2m21.678 12.3a27.923 27.923 0 0 1 5.691 11.8H26.1L14.3 14.3h39.378M10.352 37.9l1.488-4.502l-3.932-2.873h4.917l1.476-4.426l1.519 4.426h4.873l-3.933 2.873l1.489 4.502l-3.948-2.768l-3.949 2.768m43.326 11.8H14.3l11.8-11.8h33.27a27.93 27.93 0 0 1-5.692 11.8"/>`),
		g.Group(children),
	)
}