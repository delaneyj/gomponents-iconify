package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackHandOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.8 23q-2.05 0-3.85-.938T6 19.45L1.2 12.4l1.675-1.675L7 13.575V3h2v14.425L5.3 14.85l2.375 3.45q.875 1.275 2.225 1.988t2.9.712q2.575 0 4.388-1.813T19 14.8V4h2v10.8q0 3.425-2.388 5.813T12.8 23ZM11 12V1h2v11h-2Zm4 0V2h2v10h-2Zm-2.85 4.5Z"/>`),
		g.Group(children),
	)
}