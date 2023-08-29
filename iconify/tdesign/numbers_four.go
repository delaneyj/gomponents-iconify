package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumbersFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 4v8h7V4h2v16h-2v-6h-7a2 2 0 0 1-2-2V4h2Z"/>`),
		g.Group(children),
	)
}