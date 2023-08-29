package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IceSkatingOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22v-2h4v-2H3V3h9v4.5q0 .675.413 1.225t1.062.725l2.625.725q1.275.35 2.087 1.412T19 14v4h-3v2h2q1.25 0 2.125-.875T21 17h2q0 2.075-1.463 3.538T18 22H2Zm3-6h12v-2q0-.675-.413-1.188t-1.062-.712l-2.625-.75q-.9-.275-1.575-.887T10.3 9H8.5q-.2 0-.35-.15T8 8.5q0-.2.15-.35T8.5 8h1.6q-.05-.3-.063-.5T10 7H8.5q-.2 0-.35-.15T8 6.5q0-.2.15-.35T8.5 6H10V5H5v11Zm3 4h6v-2H8v2Zm-3-4Z"/>`),
		g.Group(children),
	)
}