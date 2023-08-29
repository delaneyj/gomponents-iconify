package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LongRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18.17 13l-2.58 2.59L17 17l5-5l-5-5l-1.41 1.41L18.17 11H2v2h16.17Z"/>`),
		g.Group(children),
	)
}