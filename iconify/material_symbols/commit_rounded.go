package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommitRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 17q-1.825 0-3.188-1.137T7.1 13H3q-.425 0-.713-.288T2 12q0-.425.288-.713T3 11h4.1q.35-1.725 1.713-2.863T12 7q1.825 0 3.188 1.137T16.9 11H21q.425 0 .713.288T22 12q0 .425-.288.713T21 13h-4.1q-.35 1.725-1.712 2.863T12 17Zm0-2q1.25 0 2.125-.875T15 12q0-1.25-.875-2.125T12 9q-1.25 0-2.125.875T9 12q0 1.25.875 2.125T12 15Z"/>`),
		g.Group(children),
	)
}