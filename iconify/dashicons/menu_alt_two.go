package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuAltTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 15h10v-2H5v2zM5 5v2h10V5H5zm0 6h10V9H5v2z"/>`),
		g.Group(children),
	)
}