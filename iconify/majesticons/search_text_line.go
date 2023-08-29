package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SearchTextLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M4 6a1 1 0 0 1 1-1h14a1 1 0 1 1 0 2H5a1 1 0 0 1-1-1zm0 4a1 1 0 0 1 1-1h14a1 1 0 1 1 0 2H5a1 1 0 0 1-1-1zm0 4a1 1 0 0 1 1-1h5a1 1 0 1 1 0 2H5a1 1 0 0 1-1-1zm0 4a1 1 0 0 1 1-1h5a1 1 0 1 1 0 2H5a1 1 0 0 1-1-1zm11.5-4a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3zM12 15.5a3.5 3.5 0 1 1 6.58 1.665l1.127 1.128a1 1 0 0 1-1.414 1.414l-1.128-1.128A3.5 3.5 0 0 1 12 15.5z"/></g>`),
		g.Group(children),
	)
}