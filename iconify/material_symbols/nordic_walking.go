package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NordicWalking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m4 23l2-9h1.525l-2 9H4Zm3 0L9.8 8.9L8 9.6V13H6V8.3l5.05-2.15q.725-.3 1.475-.063T13.7 7l1 1.6q.65 1.05 1.763 1.725T19 11v2q-1.65 0-3.088-.688T13.5 10.5l-.6 3l2.1 2V23h-2v-6l-2.1-2l-1.8 8H7Zm6.5-17.5q-.825 0-1.413-.588T11.5 3.5q0-.825.588-1.413T13.5 1.5q.825 0 1.413.588T15.5 3.5q0 .825-.588 1.413T13.5 5.5Zm4 17.5v-9H19v9h-1.5Z"/>`),
		g.Group(children),
	)
}