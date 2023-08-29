package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResizeHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m9.28 6.78l-8.5 8.5l-.686.72l.687.72l8.5 8.5l1.44-1.44L3.936 17h24.125l-6.78 6.78l1.437 1.44l8.5-8.5l.686-.72l-.687-.72l-8.5-8.5l-1.44 1.44L28.063 15H3.938l6.782-6.78l-1.44-1.44z"/>`),
		g.Group(children),
	)
}