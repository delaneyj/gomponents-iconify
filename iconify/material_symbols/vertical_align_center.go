package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerticalAlignCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 22v-4.2l-1.6 1.6L8 18l4-4l4 4l-1.4 1.4l-1.6-1.6V22h-2Zm-7-9v-2h16v2H4Zm8-3L8 6l1.4-1.4L11 6.2V2h2v4.2l1.6-1.6L16 6l-4 4Z"/>`),
		g.Group(children),
	)
}