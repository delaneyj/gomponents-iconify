package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkdownFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 3h18a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1Zm4 12.5v-4l2 2l2-2v4h2v-7h-2l-2 2l-2-2H5v7h2Zm11-3v-4h-2v4h-2l3 3l3-3h-2Z"/>`),
		g.Group(children),
	)
}