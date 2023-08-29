package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MathPercent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.243 6.343a1 1 0 1 1 1.414 1.414l-9.9 9.9a1 1 0 0 1-1.414-1.414l9.9-9.9ZM9.879 9.879A2 2 0 1 1 7.05 7.05a2 2 0 0 1 2.83 2.83Zm4.242 7.071a2 2 0 1 0 2.829-2.829a2 2 0 0 0-2.829 2.829Z"/>`),
		g.Group(children),
	)
}