package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Casino(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M8.5 10A2.5 2.5 0 0 0 13 8.5c0-.564-.194-1.079-.509-1.497L12.5 7l-5-6l-5 6l.009.003A2.478 2.478 0 0 0 2 8.5A2.5 2.5 0 0 0 6.5 10l.5-.666V11.5C7 13 4.5 13 4.5 13a.5.5 0 1 0 0 1h6a.5.5 0 0 0 0-1S8 13 8 11.5V9.334l.5.666z"/>`),
		g.Group(children),
	)
}