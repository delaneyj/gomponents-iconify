package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloorL(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M12 3L2 12h3v8h14v-8h3L12 3zM9 8h2v8h4v2H9V8z" fill="currentColor"/>`),
		g.Group(children),
	)
}