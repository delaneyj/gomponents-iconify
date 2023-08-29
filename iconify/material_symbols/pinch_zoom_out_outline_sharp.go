package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinchZoomOutOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 12V7h1.5v2.45l2.975-2.975l1.05 1.05L3.55 10.5H6V12H1Zm6.525-5.475l-1.05-1.05L9.45 2.5H7V1h5v5h-1.5V3.55L7.525 6.525ZM13.575 23L7.6 17l1.575-1.625l2.825.8V7h2v11.825l-2.425-.675l2.85 2.85H21v-7h2v9h-9.425ZM15 16v-5h2v5h-2Zm3 0v-4h2v4h-2Zm-.5 2Z"/>`),
		g.Group(children),
	)
}