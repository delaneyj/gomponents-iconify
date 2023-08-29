package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneNavigateBefore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14.2 6l-6 6l6 6l1.41-1.41L11.03 12l4.58-4.59z"/>`),
		g.Group(children),
	)
}