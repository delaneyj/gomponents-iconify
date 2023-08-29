package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Octagon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 16V9l5-5h7l5 5v7.03L15.03 21H8l-5-5M8.39 5L4 9.39v6.21L8.4 20h6.21L19 15.61V9.39L14.61 5H8.39Z"/>`),
		g.Group(children),
	)
}