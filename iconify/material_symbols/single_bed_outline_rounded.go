package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SingleBedOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 17v-5.025q0-.825.588-1.4T6 10V7q0-.825.588-1.413T8 5h8q.825 0 1.413.588T18 7v3q.825 0 1.413.588T20 12v5h-1.35l-.475 1.5q-.075.225-.263.363T17.5 19q-.25 0-.425-.163t-.25-.387L16.35 17h-8.7l-.475 1.5q-.075.225-.262.363T6.5 19q-.25 0-.425-.163t-.25-.387L5.35 17H4Zm9-7h3V7h-3v3Zm-5 0h3V7H8v3Zm-2 5h12v-3H6v3Zm12 0H6h12Z"/>`),
		g.Group(children),
	)
}