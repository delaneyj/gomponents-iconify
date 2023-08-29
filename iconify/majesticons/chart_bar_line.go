package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartBarLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M19 3a3 3 0 0 0-3 3v12a3 3 0 1 0 6 0V6a3 3 0 0 0-3-3zm-1 3a1 1 0 1 1 2 0v12a1 1 0 1 1-2 0V6zm-9 6a3 3 0 1 1 6 0v6a3 3 0 1 1-6 0v-6zm3-1a1 1 0 0 0-1 1v6a1 1 0 1 0 2 0v-6a1 1 0 0 0-1-1zM2 18a3 3 0 1 1 6 0a3 3 0 0 1-6 0zm3-1a1 1 0 1 0 0 2a1 1 0 0 0 0-2z"/></g>`),
		g.Group(children),
	)
}