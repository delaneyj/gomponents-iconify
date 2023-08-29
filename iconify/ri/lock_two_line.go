package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 8V7a6 6 0 1 1 12 0v1h2a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1h2Zm13 2H5v10h14V10Zm-8 5.732A2 2 0 0 1 12 12a2 2 0 0 1 1 3.732V18h-2v-2.268ZM8 8h8V7a4 4 0 0 0-8 0v1Z"/>`),
		g.Group(children),
	)
}