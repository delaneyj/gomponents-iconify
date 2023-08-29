package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackHandSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.8 23q-2.05 0-3.85-.938T6 19.45L1.2 12.4l1.675-1.675L7 13.575V3h2v9h2V1h2v11h2V2h2v10h2V4h2v10.8q0 3.425-2.388 5.813T12.8 23Z"/>`),
		g.Group(children),
	)
}