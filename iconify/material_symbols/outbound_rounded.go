package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutboundRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 11.4v1.25q0 .425.288.713t.712.287q.425 0 .713-.288T16 12.65V9q0-.425-.288-.713T15 8h-3.65q-.425 0-.712.288T10.35 9q0 .425.288.713t.712.287h1.225L8.2 14.375q-.275.275-.275.688t.275.712q.3.3.713.3t.712-.3L14 11.4ZM12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22Z"/>`),
		g.Group(children),
	)
}