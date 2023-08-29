package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hq(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 7h2v4h4V7h2v10H9v-4H5v4H3V7zm10 2h2v6h-2V9zm6 6h-4v2h8v-2h-2V9h-2V7h-4v2h4v6z"/>`),
		g.Group(children),
	)
}