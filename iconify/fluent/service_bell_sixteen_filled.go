package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServiceBellSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 1.5a2 2 0 0 0-2 2v.341C3.67 4.665 2 6.888 2 9.5a.5.5 0 0 0 .5.5h11a.5.5 0 0 0 .5-.5a6.002 6.002 0 0 0-4-5.659V3.5a2 2 0 0 0-2-2ZM7.5 6a.5.5 0 0 1 .5-.5c1.36 0 2.56.679 3.283 1.714a.5.5 0 1 1-.82.572A2.996 2.996 0 0 0 8 6.5a.5.5 0 0 1-.5-.5Zm-5 5a1.5 1.5 0 0 0 0 3h11a1.5 1.5 0 0 0 0-3h-11Z"/>`),
		g.Group(children),
	)
}