package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataExplorationSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.5 20.5q.425 0 .713-.288t.287-.712q0-.425-.288-.713T19.5 18.5q-.425 0-.713.288t-.287.712q0 .425.288.713t.712.287ZM12 22q-2.8 0-4.938-1.3t-3.337-3.1l4.4-4.4L11.4 16l4.6-4.575V13h2V8h-5v2h1.575L11.3 13.275L8 10.5l-5.275 5.275q-.35-.9-.538-1.863T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12v10H12Z"/>`),
		g.Group(children),
	)
}