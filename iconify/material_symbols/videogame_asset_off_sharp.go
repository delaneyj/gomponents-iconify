package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideogameAssetOffSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 18V5.975h3.975L15 15h-2.85L1.375 4.225L2.8 2.8l18.4 18.4l-1.425 1.425L15.15 18H2ZM8.85 6H22v11.875h-1.275L8.85 6Zm8.65 3q-.625 0-1.063.438T16 10.5q0 .625.438 1.063T17.5 12q.625 0 1.063-.438T19 10.5q0-.625-.438-1.063T17.5 9ZM7 15h2v-2h2v-2H9V9H7v2H5v2h2v2Z"/>`),
		g.Group(children),
	)
}