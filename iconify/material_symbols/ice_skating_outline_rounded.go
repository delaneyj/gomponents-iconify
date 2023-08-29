package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IceSkatingOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 22q-.425 0-.713-.288T2 21q0-.425.288-.713T3 20h3v-2H5q-.825 0-1.413-.588T3 16V5q0-.825.588-1.413T5 3h5q.825 0 1.413.588T12 5v2.5q0 .675.413 1.225t1.062.725l2.625.725q1.275.35 2.087 1.412T19 14v2q0 .825-.588 1.413T17 18h-1v2h2q1.25 0 2.125-.875T21 17q0-.425.288-.713T22 16q.425 0 .713.288T23 17q0 2.075-1.463 3.538T18 22H3Zm2-6h12v-2q0-.675-.413-1.188t-1.062-.712l-2.625-.75q-.9-.275-1.575-.887T10.3 9H8.5q-.2 0-.35-.15T8 8.5q0-.2.15-.35T8.5 8h1.6q-.05-.3-.063-.5T10 7H8.5q-.2 0-.35-.15T8 6.5q0-.2.15-.35T8.5 6H10V5H5v11Zm3 4h6v-2H8v2Zm-3-4Z"/>`),
		g.Group(children),
	)
}