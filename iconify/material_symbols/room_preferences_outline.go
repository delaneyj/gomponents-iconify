package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoomPreferencesOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-2h2V3h10v1h4v6h-2V6h-2v6h-2V5H7v14h5v2H3Zm14 1l-.3-1.5q-.3-.125-.563-.275t-.512-.35l-1.45.5l-1-1.725l1.125-1q-.05-.375-.05-.638t.05-.637l-1.125-1l1-1.725l1.45.5q.25-.2.513-.363t.562-.287L17 12h2l.3 1.5q.3.125.563.275t.512.35l1.45-.5l1 1.725l-1.125 1q.05.375.05.638t-.05.637l1.125 1l-1 1.725l-1.45-.475q-.25.2-.513.35t-.562.275L19 22h-2Zm1-3q.825 0 1.413-.588T20 17q0-.825-.588-1.413T18 15q-.825 0-1.413.588T16 17q0 .825.588 1.413T18 19Zm-7-6q-.425 0-.713-.288T10 12q0-.425.288-.713T11 11q.425 0 .713.288T12 12q0 .425-.288.713T11 13Zm-4 6V5v14Z"/>`),
		g.Group(children),
	)
}