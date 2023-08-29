package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m5.928 7.976l4.357 4.357l-.618.62L5 8.284v-.618L9.667 3l.618.619l-4.357 4.357z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}