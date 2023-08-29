package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneRemoveRoad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 4h2v9h-2zM4 4h2v16H4zm7 0h2v4h-2zm0 6h2v4h-2zm0 6h2v4h-2zm11.5.41L21.09 15L19 17.09L16.91 15l-1.41 1.41l2.09 2.09l-2.09 2.09L16.91 22L19 19.91L21.09 22l1.41-1.41l-2.09-2.09z"/>`),
		g.Group(children),
	)
}