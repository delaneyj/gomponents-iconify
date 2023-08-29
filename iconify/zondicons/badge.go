package zondicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Badge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 12a6 6 0 1 1 0-12a6 6 0 0 1 0 12zm0-3a3 3 0 1 0 0-6a3 3 0 0 0 0 6zm4 2.75V20l-4-4l-4 4v-8.25a6.97 6.97 0 0 0 8 0z"/>`),
		g.Group(children),
	)
}