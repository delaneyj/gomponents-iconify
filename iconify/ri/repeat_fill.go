package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RepeatFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 4h15a1 1 0 0 1 1 1v7h-2V6H6v3L1 5l5-4v3Zm12 16H3a1 1 0 0 1-1-1v-7h2v6h14v-3l5 4l-5 4v-3Z"/>`),
		g.Group(children),
	)
}