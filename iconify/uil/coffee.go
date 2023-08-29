package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coffee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 17h4a5 5 0 0 0 5-5v-1h1a3 3 0 0 0 0-6h-1V4a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v8a5 5 0 0 0 5 5Zm9-10h1a1 1 0 0 1 0 2h-1ZM6 5h10v7a3 3 0 0 1-3 3H9a3 3 0 0 1-3-3Zm15 14H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}