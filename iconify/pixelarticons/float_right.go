package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloatRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 4h6v8h-8V4h2zm4 6V6h-4v4h4zm-8-4H2v2h10V6zm0 4H2v2h10v-2zm10 4v2H2v-2h20zm0 6v-2H2v2h20z"/>`),
		g.Group(children),
	)
}