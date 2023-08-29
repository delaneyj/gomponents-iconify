package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatPaintOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 22q-.825 0-1.413-.588T9 20v-4H6q-.825 0-1.413-.588T4 14V7q0-1.65 1.175-2.825T8 3h12v11q0 .825-.588 1.413T18 16h-3v4q0 .825-.588 1.413T13 22h-2ZM6 10h12V5h-1v4h-2V5h-1v2h-2V5H8q-.825 0-1.413.588T6 7v3Zm0 4h12v-2H6v2Zm0-2v2v-2Z"/>`),
		g.Group(children),
	)
}