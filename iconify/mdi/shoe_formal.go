package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoeFormal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.5 9V8h-1l-1 1H15l-1-1h-1l-6 4H4a2 2 0 0 0-2 2v2h8l3-1h2v1h6.5v-2s.5-1 .5-2.5s-.5-2.5-.5-2.5Z"/>`),
		g.Group(children),
	)
}