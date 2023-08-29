package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatTextdirectionRToL(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 15v-5q-1.65 0-2.825-1.175T5 6q0-1.65 1.175-2.825T9 2h8v2h-2v11h-2V4h-2v11H9Zm-2.2 4l1.6 1.6L7 22l-4-4l4-4l1.4 1.4L6.8 17H21v2H6.8Z"/>`),
		g.Group(children),
	)
}