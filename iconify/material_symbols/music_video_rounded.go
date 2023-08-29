package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicVideoRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.5 17q1.05 0 1.775-.725T13 14.5V9h2q.425 0 .713-.288T16 8q0-.425-.288-.713T15 7h-2q-.425 0-.713.288T12 8v4.5q-.325-.225-.7-.363T10.5 12q-1.05 0-1.775.725T8 14.5q0 1.05.725 1.775T10.5 17ZM4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Z"/>`),
		g.Group(children),
	)
}