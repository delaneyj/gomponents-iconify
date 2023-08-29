package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpDoorFront(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 19V3H5v16H3v2h18v-2h-2zm-4-6h-2v-2h2v2z"/>`),
		g.Group(children),
	)
}