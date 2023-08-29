package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerUpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 8h10V6h2v2h2v2h-2v2h-2v-2H6v10H4V8h2zm10 4v2h-2v-2h2zm0-6V4h-2v2h2z"/>`),
		g.Group(children),
	)
}