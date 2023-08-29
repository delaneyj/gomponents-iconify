package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Psdown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m542 52l279 295q23 22 48.5 71t25.5 79v376q0 63-43.5 107T746 1024H149q-62 0-105.5-44T0 873V497q0-30 25-82t50-68L354 52q40-52 94-52t94 52z"/>`),
		g.Group(children),
	)
}