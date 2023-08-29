package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloorOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M12 3L2 12h3v8h14v-8h3L12 3zm-2 5h4v10h-2v-8h-2V8z" fill="currentColor"/>`),
		g.Group(children),
	)
}