package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReviewsRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6 18l-3.15 3.15q-.25.25-.55.125T2 20.8V4q0-.825.588-1.413T4 2h16q.825 0 1.413.588T22 4v12q0 .825-.588 1.413T20 18H6Zm6-5.525l1.9 1.15q.275.175.55-.013t.2-.512l-.5-2.175l1.7-1.475q.25-.225.15-.537t-.45-.338L13.325 8.4l-.875-2.05q-.125-.3-.45-.3t-.45.3l-.875 2.05l-2.225.175Q8.1 8.6 8 8.913t.15.537l1.7 1.475l-.5 2.175q-.075.325.2.513t.55.012l1.9-1.15Z"/>`),
		g.Group(children),
	)
}