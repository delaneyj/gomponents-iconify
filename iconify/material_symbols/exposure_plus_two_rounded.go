package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExposurePlusTwoRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 19h-7.175q-.3 0-.512-.213t-.213-.512V17.2q0-.075.2-.525l4.8-4.875q.825-.875 1.163-1.462T18.6 9q0-.725-.563-1.313T16.35 7.1q-.65 0-1.113.263t-.787.687q-.25.35-.637.475T13.05 8.5q-.425-.175-.612-.563t.012-.737q.525-.875 1.513-1.538T16.4 5q2.075 0 3.238 1.188T20.8 9q0 1.125-.525 2.05T18.65 13.1L15 16.9l.05.1H20q.425 0 .713.288T21 18q0 .425-.288.713T20 19ZM6 14H4q-.425 0-.713-.288T3 13q0-.425.288-.713T4 12h2v-2q0-.425.288-.713T7 9q.425 0 .713.288T8 10v2h2q.425 0 .713.288T11 13q0 .425-.288.713T10 14H8v2q0 .425-.288.713T7 17q-.425 0-.713-.288T6 16v-2Z"/>`),
		g.Group(children),
	)
}