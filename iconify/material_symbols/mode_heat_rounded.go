package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ModeHeatRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 14q0-2.825 1.675-5.425t4.6-4.55q.55-.375 1.137-.038T12 5v1.3q0 .85.588 1.425t1.437.575q.425 0 .813-.188t.687-.537q.2-.25.513-.313t.587.138Q18.2 8.525 19.1 10.275T20 14q0 2.2-1.075 4.013T16.1 20.874q.425-.6.663-1.313T17 18.05q0-1-.375-1.888t-1.075-1.587L12 11.1l-3.525 3.475q-.725.725-1.1 1.6T7 18.05q0 .8.238 1.513t.662 1.312q-1.75-1.05-2.825-2.863T4 14Zm8-.1l2.125 2.075q.425.425.65.95T15 18.05q0 1.225-.875 2.087T12 21q-1.25 0-2.125-.863T9 18.05q0-.575.225-1.113t.65-.962L12 13.9Z"/>`),
		g.Group(children),
	)
}