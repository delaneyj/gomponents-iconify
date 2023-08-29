package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IceSkatingOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22v-2h4v-2H3V3h9v4.5q0 .675.413 1.225t1.062.725L19 10.975V18h-3v2h2q1.25 0 2.125-.875T21 17h2q0 2.075-1.463 3.538T18 22H2Zm3-6h12v-3.475l-4.1-1.175q-.9-.275-1.575-.888T10.3 9H8V8h2.1q-.05-.3-.063-.5T10 7H8V6h2V5H5v11Zm3 4h6v-2H8v2Zm-3-4Z"/>`),
		g.Group(children),
	)
}