package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloorA(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M12 3L2 12h3v8h14v-8h3L12 3zm-1 5h2a2 2 0 0 1 2 2v8h-2v-3h-2v3H9v-8a2 2 0 0 1 2-2zm0 2v3h2v-3h-2z" fill="currentColor"/>`),
		g.Group(children),
	)
}