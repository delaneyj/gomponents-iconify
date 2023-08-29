package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Expand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22v-2h16v2H4Zm8-3l-4-4l1.4-1.4l1.6 1.55v-6.3L9.4 10.4L8 9l4-4l4 4l-1.4 1.4L13 8.85v6.3l1.6-1.55L16 15l-4 4ZM4 4V2h16v2H4Z"/>`),
		g.Group(children),
	)
}