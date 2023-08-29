package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideogameAssetOffOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.725 17.875L18.85 16H20V8h-9.15l-2-2H22v11.875h-1.275ZM17.5 12q-.625 0-1.063-.438T16 10.5q0-.625.438-1.063T17.5 9q.625 0 1.063.438T19 10.5q0 .625-.438 1.063T17.5 12Zm-8.35 0Zm5.7 0ZM2 18V5.975h3.975L8 8H4v8h9.15L1.375 4.225L2.8 2.8l18.4 18.4l-1.425 1.425L15.15 18H2Zm5-3v-2H5v-2h2V9h2v2h2v2H9v2H7Z"/>`),
		g.Group(children),
	)
}