package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NextWeek(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21V6h6V2h8v4h6v15H2Zm8-15h4V4h-4v2Zm1 11.5l4-4l-4-4l-1.4 1.4l2.6 2.6l-2.6 2.6l1.4 1.4Z"/>`),
		g.Group(children),
	)
}