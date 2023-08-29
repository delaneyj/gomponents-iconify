package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DetectorAlarmOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 21v-5h2v5h-2Zm8.075-2.525l-3.55-3.525l1.425-1.425l3.525 3.55l-1.4 1.4Zm-14.15 0l-1.4-1.4l3.525-3.55l1.425 1.425l-3.55 3.525ZM5 5v1h14V5H5Zm3.1 3l.3 1h7.2l.3-1H8.1Zm-1.15 3L6 8H3V3h18v5h-3l-1.15 3h-9.9ZM5 5v1v-1Z"/>`),
		g.Group(children),
	)
}