package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewParallelOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 3H5v18h14V3m-2 16h-2V5h2v14m-4 0h-2V5h2v14M7 5h2v14H7V5Z"/>`),
		g.Group(children),
	)
}