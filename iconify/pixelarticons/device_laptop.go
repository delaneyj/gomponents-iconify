package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceLaptop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 4H4v12h16V4H6zm12 2v8H6V6h12zm4 12H2v2h20v-2z"/>`),
		g.Group(children),
	)
}