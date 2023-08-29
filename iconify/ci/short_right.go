package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShortRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.17 13l-3.58 3.59L14 18l6-6l-6-6l-1.41 1.41L16.17 11H4v2h12.17Z"/>`),
		g.Group(children),
	)
}