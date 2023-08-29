package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutFourFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 13v8H4a1 1 0 0 1-1-1v-7h8Zm2-10h7a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1h-7V3ZM3 4a1 1 0 0 1 1-1h7v8H3V4Z"/>`),
		g.Group(children),
	)
}