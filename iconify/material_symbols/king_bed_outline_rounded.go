package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KingBedOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 17v-5.025q0-.825.588-1.4T4 10V7q0-.825.588-1.413T6 5h12q.825 0 1.413.588T20 7v3q.825 0 1.413.588T22 12v5h-1.35l-.475 1.5q-.075.225-.263.363T19.5 19q-.25 0-.425-.163t-.25-.387L18.35 17H5.65l-.475 1.5q-.075.225-.262.363T4.5 19q-.25 0-.425-.163t-.25-.387L3.35 17H2Zm11-7h5V7h-5v3Zm-7 0h5V7H6v3Zm-2 5h16v-3H4v3Zm16 0H4h16Z"/>`),
		g.Group(children),
	)
}