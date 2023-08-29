package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackedBarChartRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 20q-.825 0-1.413-.588T4 18V9h4v9q0 .825-.588 1.413T6 20ZM4 8V6q0-.825.588-1.413T6 4q.825 0 1.413.588T8 6v2H4Zm8 12q-.825 0-1.413-.588T10 18v-6h4v6q0 .825-.588 1.413T12 20Zm-2-9V9q0-.825.588-1.413T12 7q.825 0 1.413.588T14 9v2h-4Zm8 9q-.825 0-1.413-.588T16 18v-3h4v3q0 .825-.588 1.413T18 20Zm-2-6v-2q0-.825.588-1.413T18 10q.825 0 1.413.588T20 12v2h-4Z"/>`),
		g.Group(children),
	)
}