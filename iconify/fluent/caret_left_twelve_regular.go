package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretLeftTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 3.994c0-.887-1.07-1.334-1.7-.712L4.26 5.288a1 1 0 0 0 0 1.425L6.3 8.719c.63.621 1.7.174 1.7-.713V3.994ZM4.963 6L7 3.994v4.012L4.963 6Z"/>`),
		g.Group(children),
	)
}