package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DriveFileMoveRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h5.175q.4 0 .763.15t.637.425L12 6h8q.825 0 1.413.588T22 8v10q0 .825-.588 1.413T20 20H4Zm8.2-6l-.925.925q-.275.275-.275.7t.275.7q.275.275.7.275t.7-.275L15.3 13.7q.3-.3.3-.7t-.3-.7l-2.625-2.625q-.275-.275-.7-.275t-.7.275q-.275.275-.275.7t.275.7L12.2 12H9q-.425 0-.713.288T8 13q0 .425.288.713T9 14h3.2Z"/>`),
		g.Group(children),
	)
}