package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShortLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7.83 11l3.58-3.59L10 6l-6 6l6 6l1.41-1.41L7.83 13H20v-2H7.83Z"/>`),
		g.Group(children),
	)
}