package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosShareRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 23q-.825 0-1.413-.588T4 21V10q0-.825.588-1.413T6 8h2q.425 0 .713.288T9 9q0 .425-.288.713T8 10H6v11h12V10h-2q-.425 0-.713-.288T15 9q0-.425.288-.713T16 8h2q.825 0 1.413.588T20 10v11q0 .825-.588 1.413T18 23H6Zm5-18.175l-.9.9q-.3.3-.7.288T8.7 5.7q-.275-.3-.287-.7t.287-.7l2.6-2.6q.3-.3.7-.3t.7.3l2.6 2.6q.275.275.275.688T15.3 5.7q-.3.3-.713.3t-.712-.3L13 4.825V15q0 .425-.288.713T12 16q-.425 0-.713-.288T11 15V4.825Z"/>`),
		g.Group(children),
	)
}