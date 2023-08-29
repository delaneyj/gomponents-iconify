package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestCamFloodlightOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16 9.4l-2.625-2.6q-.85-.85-1.975-1.325T9.075 5q-1.625 0-3.013.825T4 7.775q.675-1.125 2.063-1.95T9.075 5q1.2 0 2.325.475T13.375 6.8L16 9.4ZM0 18V0h2q.825 0 1.413.588T4 2v1.7q.8-1.25 2.125-1.975T9.025 1H16v8.4l2.3 2.375l-8.575 8.55l-4.95-4.95q-.225-.225-.413-.45t-.362-.5V16q0 .825-.588 1.413T2 18H0ZM9.025 3q-1.5 0-2.638 1.038T5.05 6.55q.8-.675 1.85-1.112T9.075 5q1.2 0 2.325.475T13.375 6.8l.625.6V3H9.025ZM5 11.075q0 .8.313 1.55T6.2 13.95l3.525 3.55l5.775-5.75l-3.525-3.55q-.575-.575-1.337-.887T9.075 7q-1.7 0-2.887 1.213T5 11.075ZM20 16h-3v-2h3v2Zm-2.8 4.6l-2.1-2.1l1.4-1.4l2.1 2.1l-1.4 1.4ZM12 22v-3h2v3h-2Zm-1.75-9.75Z"/>`),
		g.Group(children),
	)
}