package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionsWalk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 23L9.8 8.9L8 9.6V13H6V8.3l5.05-2.15q.35-.15.738-.175t.737.1q.35.125.663.35T13.7 7l1 1.6q.65 1.05 1.763 1.725T19 11v2q-1.75 0-3.125-.725t-2.35-1.85L12.9 13.5l2.1 2V23h-2v-6.5l-2.1-1.6L9.1 23H7Zm6.5-17.5q-.825 0-1.413-.588T11.5 3.5q0-.825.588-1.413T13.5 1.5q.825 0 1.413.588T15.5 3.5q0 .825-.588 1.413T13.5 5.5Z"/>`),
		g.Group(children),
	)
}