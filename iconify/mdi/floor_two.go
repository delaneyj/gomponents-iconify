package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloorTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M12 3L2 12h3v8h14v-8h3L12 3zM9 8h4a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2h-2v2h4v2H9v-4a2 2 0 0 1 2-2h2v-2H9V8z" fill="currentColor"/>`),
		g.Group(children),
	)
}