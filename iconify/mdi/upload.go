package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Upload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 16v-6H5l7-7l7 7h-4v6H9m-4 4v-2h14v2H5Z"/>`),
		g.Group(children),
	)
}