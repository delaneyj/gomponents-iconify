package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FindInPage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 15q-.825 0-1.413-.588T10 13q0-.825.588-1.413T12 11q.825 0 1.413.588T14 13q0 .825-.588 1.413T12 15ZM6 2h8l6 6v12q0 .825-.588 1.413T18 22H6q-.825 0-1.413-.588T4 20V4q0-.825.588-1.413T6 2Zm6 15q.575 0 1.075-.138T14 16.45L17.575 20H18v-2.4L15.45 15q.275-.425.413-.925T16 13q0-1.65-1.175-2.825T12 9q-1.65 0-2.825 1.175T8 13q0 1.65 1.175 2.825T12 17Z"/>`),
		g.Group(children),
	)
}