package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AodTabletRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 20q-.825 0-1.413-.588T1 18V6q0-.825.588-1.413T3 4h18q.825 0 1.413.588T23 6v12q0 .825-.588 1.413T21 20H3Zm3-2h12V6H6v12Zm2.75-6.5q-.325 0-.537-.213T8 10.75q0-.325.213-.537T8.75 10h6.5q.325 0 .537.213t.213.537q0 .325-.213.537t-.537.213h-6.5Zm1 3q-.325 0-.537-.213T9 13.75q0-.325.213-.537T9.75 13h4.5q.325 0 .537.213t.213.537q0 .325-.213.537t-.537.213h-4.5Z"/>`),
		g.Group(children),
	)
}