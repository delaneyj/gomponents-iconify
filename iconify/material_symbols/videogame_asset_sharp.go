package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideogameAssetSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 18V6h20v12H2Zm5-3h2v-2h2v-2H9V9H7v2H5v2h2v2Zm7.5 0q.625 0 1.063-.438T16 13.5q0-.625-.438-1.063T14.5 12q-.625 0-1.063.438T13 13.5q0 .625.438 1.063T14.5 15Zm3-3q.625 0 1.063-.438T19 10.5q0-.625-.438-1.063T17.5 9q-.625 0-1.063.438T16 10.5q0 .625.438 1.063T17.5 12Z"/>`),
		g.Group(children),
	)
}