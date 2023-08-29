package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LongLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m5.83 11l2.58-2.59L7 7l-5 5l5 5l1.41-1.41L5.83 13H22v-2H5.83Z"/>`),
		g.Group(children),
	)
}