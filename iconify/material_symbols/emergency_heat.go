package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmergencyHeat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 14q0-2.625 1.25-4.675T7 5.875q1.5-1.4 2.75-2.138L11 3v3.3q0 .925.625 1.462t1.4.538q.425 0 .813-.175t.712-.575L15 7q1.8 1.05 2.9 2.912T19 14q0 2.2-1.075 4.013T15.1 20.874q.425-.6.663-1.313T16 18.05q0-1-.375-1.888t-1.075-1.587L11 11.1l-3.525 3.475q-.725.725-1.1 1.6T6 18.05q0 .8.238 1.513t.662 1.312q-1.75-1.05-2.825-2.863T3 14Zm8-.1l2.125 2.075q.425.425.65.95T14 18.05q0 1.225-.875 2.087T11 21q-1.25 0-2.125-.863T8 18.05q0-.575.225-1.113t.65-.962L11 13.9ZM21 11q-.425 0-.713-.288T20 10q0-.425.288-.713T21 9q.425 0 .713.288T22 10q0 .425-.288.713T21 11Zm-1-3V3h2v5h-2Z"/>`),
		g.Group(children),
	)
}