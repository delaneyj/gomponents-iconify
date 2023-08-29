package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AodTabletOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.75 11.5q-.325 0-.537-.213T8 10.75q0-.325.213-.537T8.75 10h6.5q.325 0 .537.213t.213.537q0 .325-.213.537t-.537.213h-6.5Zm1 3q-.325 0-.537-.213T9 13.75q0-.325.213-.537T9.75 13h4.5q.325 0 .537.213t.213.537q0 .325-.213.537t-.537.213h-4.5ZM3 20q-.825 0-1.413-.588T1 18V6q0-.825.588-1.413T3 4h18q.825 0 1.413.588T23 6v12q0 .825-.588 1.413T21 20H3ZM4 6H3v12h1V6Zm2 12h12V6H6v12ZM20 6v12h1V6h-1Zm0 0h1h-1ZM4 6H3h1Z"/>`),
		g.Group(children),
	)
}