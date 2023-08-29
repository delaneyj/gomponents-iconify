package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KickFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 3h18v18H3V3Zm7.564 2.536h-4.31v12.928h4.31V15.59H12v1.436h1.436v1.436h4.31v-4.309h-1.437v-1.436h-1.436v-1.436h1.436V9.845h1.436V5.536h-4.309v1.436H12V8.41h-1.436V5.536Z"/>`),
		g.Group(children),
	)
}