package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandUpDownLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.207 9.043L12 2.836L5.793 9.043l1.414 1.414L12 5.664l4.793 4.793l1.414-1.414ZM5.793 14.957L12 21.165l6.207-6.208l-1.414-1.414L12 18.336l-4.793-4.793l-1.414 1.414Z"/>`),
		g.Group(children),
	)
}