package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApprovalRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 22q-.825 0-1.413-.588T4 20v-4q0-.825.588-1.413T6 14h12q.825 0 1.413.588T20 16v4q0 .825-.588 1.413T18 22H6Zm1-4h10q.425 0 .713-.288T18 17q0-.425-.288-.713T17 16H7q-.425 0-.713.288T6 17q0 .425.288.713T7 18Zm4.175-5.15l-3.7-5.175q-.225-.325-.312-.712t-.038-.788q.3-1.825 1.65-3T12 2q1.875 0 3.225 1.175t1.65 3q.05.4-.038.788t-.312.712l-3.7 5.175q-.3.425-.825.425t-.825-.425Z"/>`),
		g.Group(children),
	)
}