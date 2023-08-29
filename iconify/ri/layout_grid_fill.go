package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutGridFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 12.999V20a1 1 0 0 1-1 1h-8v-8.001h9Zm-11 0V21H3a1 1 0 0 1-1-1v-7.001h9ZM11 3v7.999H2V4a1 1 0 0 1 1-1h8Zm10 0a1 1 0 0 1 1 1v6.999h-9V3h8Z"/>`),
		g.Group(children),
	)
}