package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideogameAssetOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 16q0 .65-.35 1.15t-.925.725L18.85 16H20V8h-9.15l-2-2H20q.825 0 1.413.588T22 8v8Zm-4.5-4q-.625 0-1.063-.438T16 10.5q0-.625.438-1.063T17.5 9q.625 0 1.063.438T19 10.5q0 .625-.438 1.063T17.5 12Zm-8.35 0Zm5.7 0ZM7 15v-2H5v-2h2V9h2v2h2v2H9v2H7Zm-3 3q-.825 0-1.413-.588T2 16V8q0-.85.6-1.438t1.45-.587h1.925L8 8H4v8h9.15L1.375 4.225L2.8 2.8l18.4 18.4l-1.425 1.425L15.15 18H4Z"/>`),
		g.Group(children),
	)
}