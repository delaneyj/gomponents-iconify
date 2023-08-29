package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlightLandRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.85 15.9L3.725 11.7q-.325-.1-.525-.35t-.2-.6V6.875q0-.325.263-.513t.562-.087L4.2 6.4q.15.05.25.15t.15.25l.6 1.8L10 9.95V3.125q0-.425.338-.7t.762-.15l.425.1q.225.05.388.213t.237.387l2.6 8.275l5 1.4q.575.15.913.613T21 14.3q0 .825-.675 1.325t-1.475.275ZM4 21q-.425 0-.713-.288T3 20q0-.425.288-.713T4 19h16q.425 0 .713.288T21 20q0 .425-.288.713T20 21H4Z"/>`),
		g.Group(children),
	)
}