package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutMasonryLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 20a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v16Zm-11-5H4v4h7v-4Zm9-4h-7v8h7v-8Zm-9-6H4v8h7V5Zm9 0h-7v4h7V5Z"/>`),
		g.Group(children),
	)
}