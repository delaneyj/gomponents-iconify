package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MovieEditOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 19q-.825 0-1.413-.588T2 17V5q0-.825.588-1.413T4 3l2 4h3L7 3h2l2 4h3l-2-4h2l2 4h3l-2-4h3q.825 0 1.413.588T22 5v4H4v8h8v2H4Zm14.3-6.475l1.075 1.075l-3.875 3.85v1.05h1.05l3.875-3.85l1.05 1.05l-4 4q-.15.15-.338.225T16.75 20H14.5q-.2 0-.35-.15T14 19.5v-2.25q0-.2.075-.387t.225-.338l4-4Zm3.175 3.175L18.3 12.525l1.45-1.45q.275-.275.7-.275t.7.275l1.775 1.775q.275.275.275.7t-.275.7l-1.45 1.45Z"/>`),
		g.Group(children),
	)
}