package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HailRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 6q-.825 0-1.413-.588T10 4q0-.825.588-1.413T12 2q.825 0 1.413.588T14 4q0 .825-.588 1.413T12 6Zm-2 16q-.425 0-.713-.288T9 21V10.1q-1.025.35-1.438 1.188T7.05 13.1q-.05.375-.325.638t-.65.262q-.45 0-.763-.325t-.262-.75Q5.375 10.2 7.213 8.6T12 7q2.225 0 3.463-.988T16.95 3q.05-.425.338-.713T18 2q.425 0 .713.3t.237.7q-.2 1.875-1.137 3.313T15 8.4V21q0 .425-.288.713T14 22q-.425 0-.713-.288T13 21v-5h-2v5q0 .425-.288.713T10 22Zm-5 0q-.425 0-.713-.288T4 21v-4q0-.425.288-.713T5 16h1q.425 0 .713.288T7 17v4q0 .425-.288.713T6 22H5Z"/>`),
		g.Group(children),
	)
}