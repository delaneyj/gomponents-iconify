package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenSearchDesktopRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 13.5q.525 0 .988-.138t.912-.412l1.525 1.575q.2.225.513.213t.537-.213q.225-.225.238-.537T16 13.45l-1.55-1.55q.275-.425.413-.9T15 10q0-1.475-1.038-2.488T11.5 6.5q-1.425 0-2.463 1.012T8 10q0 1.475 1.038 2.488T11.5 13.5Zm0-1.5q-.825 0-1.413-.588T9.5 10q0-.825.588-1.413T11.5 8q.8 0 1.4.588T13.5 10q0 .825-.588 1.413T11.5 12ZM2 21q-.425 0-.713-.288T1 20q0-.425.288-.713T2 19h20q.425 0 .713.288T23 20q0 .425-.288.713T22 21H2Zm2-3q-.825 0-1.413-.588T2 16V5q0-.825.588-1.413T4 3h16q.825 0 1.413.588T22 5v11q0 .825-.588 1.413T20 18H4Z"/>`),
		g.Group(children),
	)
}