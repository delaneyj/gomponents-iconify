package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewGridLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M5 7a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1H5zM2 8a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V8z"/><path d="M9 5a1 1 0 0 1 1 1v3h4V6a1 1 0 1 1 2 0v3h5a1 1 0 1 1 0 2h-5v2h5a1 1 0 1 1 0 2h-5v3a1 1 0 1 1-2 0v-3h-4v3a1 1 0 1 1-2 0v-3H3a1 1 0 1 1 0-2h5v-2H3a1 1 0 1 1 0-2h5V6a1 1 0 0 1 1-1zm1 6v2h4v-2h-4z"/></g>`),
		g.Group(children),
	)
}