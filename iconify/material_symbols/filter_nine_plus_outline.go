package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterNinePlusOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 14h2q.825 0 1.413-.588T14 12V8q0-.825-.588-1.413T12 6h-1q-.825 0-1.413.588T9 8v1q0 .825.588 1.413T11 11h1v1h-2v2Zm2-5h-1V8h1v1Zm-4 9q-.825 0-1.413-.588T6 16V4q0-.825.588-1.413T8 2h12q.825 0 1.413.588T22 4v12q0 .825-.588 1.413T20 18H8Zm0-2h12V4H8v12Zm-4 6q-.825 0-1.413-.588T2 20V6h2v14h14v2H4ZM8 4v12V4Zm8.5 9h2v-2H20V9h-1.5V7h-2v2h-2v2h2v2Z"/>`),
		g.Group(children),
	)
}