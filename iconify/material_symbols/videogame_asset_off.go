package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideogameAssetOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 18q-.825 0-1.413-.588T2 16V8q0-.85.6-1.438t1.45-.587h1.925L15 15h-2.85L1.375 4.225L2.8 2.8l18.4 18.4l-1.425 1.425L15.15 18H4ZM8.85 6H20q.825 0 1.413.588T22 8v8q0 .65-.35 1.15t-.925.725L8.85 6Zm8.65 3q-.625 0-1.063.438T16 10.5q0 .625.438 1.063T17.5 12q.625 0 1.063-.438T19 10.5q0-.625-.438-1.063T17.5 9ZM7 15h2v-2h2v-2H9V9H7v2H5v2h2v2Z"/>`),
		g.Group(children),
	)
}