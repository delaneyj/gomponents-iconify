package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BabyChangingStation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 22V12l1.575-4.625q.2-.65.75-1.012T6.5 6q.2 0 .4.038t.4.137L11.45 8H14v2h-3L8.3 8.825L7 12.75V22H3Zm6-3v-2h12v2H9Zm10.5-3q-.625 0-1.063-.438T18 14.5q0-.625.438-1.063T19.5 13q.625 0 1.063.438T21 14.5q0 .625-.438 1.063T19.5 16ZM13 16q-.825 0-1.413-.588T11 14v-1H9v-2h3q.425 0 .713.288T13 12v1h2v-2h2v3q0 .825-.588 1.413T15 16h-2ZM8 5q-.825 0-1.413-.588T6 3q0-.825.588-1.413T8 1q.825 0 1.413.588T10 3q0 .825-.588 1.413T8 5Z"/>`),
		g.Group(children),
	)
}