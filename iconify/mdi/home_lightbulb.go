package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeLightbulb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3L2 12h3v8h14v-8h3m-9 6h-2v-1h2m.5-2.42V16h-3v-1.42a3 3 0 1 1 3 0Z"/>`),
		g.Group(children),
	)
}