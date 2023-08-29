package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.69 2a1 1 0 0 0-1 1v10.37a2 2 0 0 1-2 2h-8l2.92-2.92A1 1 0 0 0 9.24 11l-4.63 4.66a1.19 1.19 0 0 0-.22.33a1 1 0 0 0 0 .76a1 1 0 0 0 .22.33l4.63 4.63a1 1 0 0 0 .7.29a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.42l-2.92-2.92h8a4 4 0 0 0 4-4V3a1 1 0 0 0-1.04-1Z"/>`),
		g.Group(children),
	)
}