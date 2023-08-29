package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Checklist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.55 19L2 15.45l1.4-1.4l2.125 2.125l4.25-4.25l1.4 1.425L5.55 19Zm0-8L2 7.45l1.4-1.4l2.125 2.125l4.25-4.25l1.4 1.425L5.55 11ZM13 17v-2h9v2h-9Zm0-8V7h9v2h-9Z"/>`),
		g.Group(children),
	)
}