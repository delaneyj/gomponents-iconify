package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextRotateUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 20h-2V6.8l-1.05 1.05l-1.4-1.4L18 3l3.5 3.45l-1.45 1.4L19 6.8V20Zm-5-2.9L3 13v-2l11-4.1v1.9l-2.8.95v4.45l2.8 1v1.9Zm-4.4-3.45v-3.3l-4.55 1.6v.1l4.55 1.6Z"/>`),
		g.Group(children),
	)
}