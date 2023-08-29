package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PostEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M10 5.5V9a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1V5.5a.5.5 0 0 1 .5-.5a.49.49 0 0 1 .21 0L5.5 7l3.8-2a.488.488 0 0 1 .2 0a.5.5 0 0 1 .5.5zM1.25 2.92l.08.08L5.5 5l4.19-2h.06a.49.49 0 0 0 .25-.5a.5.5 0 0 0-.5-.5h-8a.5.5 0 0 0-.5.5a.49.49 0 0 0 .25.42z" fill="currentColor"/>`),
		g.Group(children),
	)
}