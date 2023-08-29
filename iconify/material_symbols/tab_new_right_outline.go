package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabNewRightOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 19V5v14v-.025V19Zm0 2q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v8.35q-.475-.175-.975-.263T19 13V5H5v14h8q0 .525.088 1.025t.262.975H5Zm14 2l-1.4-1.4l1.575-1.6H15v-2h4.175L17.6 16.4L19 15l4 4l-4 4Zm-8-6h2v-4h4v-2h-4V7h-2v4H7v2h4v4Z"/>`),
		g.Group(children),
	)
}