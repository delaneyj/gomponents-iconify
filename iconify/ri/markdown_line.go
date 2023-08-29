package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkdownLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 3h18a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1Zm1 2v14h16V5H4Zm3 10.5H5v-7h2l2 2l2-2h2v7h-2v-4l-2 2l-2-2v4Zm11-3h2l-3 3l-3-3h2v-4h2v4Z"/>`),
		g.Group(children),
	)
}