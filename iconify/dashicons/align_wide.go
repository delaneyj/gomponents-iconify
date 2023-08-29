package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignWide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 5h10V3H5v2zm12 8V7H3v6h14zM5 17h10v-2H5v2z"/>`),
		g.Group(children),
	)
}