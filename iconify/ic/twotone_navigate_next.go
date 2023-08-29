package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneNavigateNext(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.02 18l6-6l-6-6l-1.41 1.41L13.19 12l-4.58 4.59z"/>`),
		g.Group(children),
	)
}