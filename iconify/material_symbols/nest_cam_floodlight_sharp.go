package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestCamFloodlightSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 19h2v3h-2v-3Zm4.5-1.9l2.1 2.1l-1.4 1.4l-2.1-2.1l1.4-1.4ZM0 0h2q.825 0 1.413.588T4 2v1.7q.8-1.25 2.125-1.975T9.025 1H16v6.275l-1.575-1.55q-1.05-1.05-2.45-1.637t-2.9-.588q-1.425 0-2.762.55T4 5.475v2.3q.675-1.125 2.075-1.95t3-.825q1.2 0 2.325.475T13.375 6.8l4.95 4.95l-8.6 8.575l-4.95-4.95q-.225-.225-.413-.45t-.362-.5V16q0 .825-.588 1.413T2 18H0V0Zm17 14h3v2h-3v-2Z"/>`),
		g.Group(children),
	)
}