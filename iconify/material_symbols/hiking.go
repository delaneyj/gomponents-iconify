package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hiking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7 23l3.075-15.55q.15-.725.675-1.088T11.85 6q.575 0 1.063.25T13.7 7l1 1.6q.45.725 1.163 1.313t1.637.862V9H19v14h-1.5V12.85q-1.2-.275-2.225-.875T13.5 10.5l-.6 3l2.1 2V23h-2v-6l-2.1-2l-1.8 8H7Zm.425-9.875l-2.125-.4q-.4-.075-.625-.413t-.15-.762l.75-3.925q.15-.8.85-1.263t1.5-.312l1.15.225l-1.35 6.85ZM13.5 5.5q-.825 0-1.413-.588T11.5 3.5q0-.825.588-1.413T13.5 1.5q.825 0 1.413.588T15.5 3.5q0 .825-.588 1.413T13.5 5.5Z"/>`),
		g.Group(children),
	)
}