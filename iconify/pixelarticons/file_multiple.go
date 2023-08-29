package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileMultiple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 18H7V2h8v2h2v2h-2v2h2V6h2v2h2v10zM9 4v12h10v-6h-6V4H9zM3 6h2v14h12v2H3V6z"/>`),
		g.Group(children),
	)
}