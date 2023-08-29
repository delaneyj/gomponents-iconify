package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VillaOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 19V9.375q0-.625.35-1.137t.925-.738L14.65 3.525q.5-.2.925.1T16 4.45V12h1q0-.825.588-1.412T19 10q.825 0 1.413.588T21 12v7q0 .825-.588 1.413T19 21H5q-.825 0-1.413-.588T3 19Zm2 0h4v-5q0-.825.588-1.413T11 12h3V5.9L5 9.375V19Zm6 0h3v-2q0-.425.288-.713T15 16q.425 0 .713.288T16 17v2h3v-5h-8v5Zm-1.5-6.55ZM15 16.5Zm0 0Z"/>`),
		g.Group(children),
	)
}