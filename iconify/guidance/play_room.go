package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayRoom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M7.5 8v5m2.5-2.5H5m-.909 10H5.5s1.5-5 6.5-5s6.5 5 6.5 5h1.409a3.591 3.591 0 0 0 3.499-4.398L20.5 3.5h-.25S17 4.5 12 4.5s-8.25-1-8.25-1H3.5L.592 16.102A3.59 3.59 0 0 0 4.091 20.5ZM16.5 8.5a2 2 0 1 1 0 4a2 2 0 0 1 0-4Z"/>`),
		g.Group(children),
	)
}