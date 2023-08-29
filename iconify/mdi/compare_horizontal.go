package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompareHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 14H2v2h7v3l4-4l-4-4v3m6-1v-3h7V8h-7V5l-4 4l4 4Z"/>`),
		g.Group(children),
	)
}