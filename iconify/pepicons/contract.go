package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Contract(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><path d="M11.707 9.707a1 1 0 0 1-1.414-1.414l4-4a1 1 0 1 1 1.414 1.414l-4 4Z"/><path d="M11 10a1 1 0 1 1 0-2h4a1 1 0 1 1 0 2h-4Z"/><path d="M12 9a1 1 0 1 1-2 0V5a1 1 0 1 1 2 0v4Zm-6.293 6.707a1 1 0 0 1-1.414-1.414l4-4a1 1 0 1 1 1.414 1.414l-4 4Z"/><path d="M10 15a1 1 0 1 1-2 0v-4a1 1 0 1 1 2 0v4Z"/><path d="M5 12a1 1 0 1 1 0-2h4a1 1 0 1 1 0 2H5Z"/></g>`),
		g.Group(children),
	)
}