package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ManageSearchRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 19q-.425 0-.713-.288T2 18q0-.425.288-.713T3 17h8q.425 0 .713.288T12 18q0 .425-.288.713T11 19H3Zm0-5q-.425 0-.713-.288T2 13q0-.425.288-.713T3 12h3q.425 0 .713.288T7 13q0 .425-.288.713T6 14H3Zm0-5q-.425 0-.713-.288T2 8q0-.425.288-.713T3 7h3q.425 0 .713.288T7 8q0 .425-.288.713T6 9H3Zm11 7q-2.075 0-3.538-1.463T9 11q0-2.075 1.463-3.538T14 6q2.075 0 3.538 1.463T19 11q0 .725-.213 1.438t-.637 1.312l3.15 3.15q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275l-3.15-3.15q-.6.425-1.313.638T14 16Zm0-2q1.25 0 2.125-.875T17 11q0-1.25-.875-2.125T14 8q-1.25 0-2.125.875T11 11q0 1.25.875 2.125T14 14Z"/>`),
		g.Group(children),
	)
}