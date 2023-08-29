package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignVertically(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 11h18v2H3v-2Zm15 7v3h-2v-3h-3l4-4l4 4h-3ZM8 18v3H6v-3H3l4-4l4 4H8ZM18 6h3l-4 4l-4-4h3V3h2v3ZM8 6h3l-4 4l-4-4h3V3h2v3Z"/>`),
		g.Group(children),
	)
}