package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ModeHeat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 14q0-2.625 1.25-4.675T8 5.875q1.5-1.4 2.75-2.138L12 3v3.3q0 .925.625 1.462t1.4.538q.425 0 .813-.175t.712-.575L16 7q1.8 1.05 2.9 2.912T20 14q0 2.2-1.075 4.013T16.1 20.874q.425-.6.663-1.313T17 18.05q0-1-.375-1.888t-1.075-1.587L12 11.1l-3.525 3.475q-.725.725-1.1 1.6T7 18.05q0 .8.238 1.513t.662 1.312q-1.75-1.05-2.825-2.863T4 14Zm8-.1l2.125 2.075q.425.425.65.95T15 18.05q0 1.225-.875 2.087T12 21q-1.25 0-2.125-.863T9 18.05q0-.575.225-1.113t.65-.962L12 13.9Z"/>`),
		g.Group(children),
	)
}