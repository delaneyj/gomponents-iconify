package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppBlockingSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 16q-1.65 0-2.825-1.175T14 12q0-1.65 1.175-2.825T18 8q1.65 0 2.825 1.175T22 12q0 1.65-1.175 2.825T18 16Zm0-1.5q.3 0 .588-.075t.562-.225l-3.35-3.35q-.15.275-.225.563T15.5 12q0 1.05.725 1.775T18 14.5Zm2.2-1.35q.15-.275.225-.563T20.5 12q0-1.05-.725-1.775T18 9.5q-.3 0-.588.075t-.562.225l3.35 3.35ZM5 23V1h14v6h-2V6H7v12h10v-1h2v6H5Z"/>`),
		g.Group(children),
	)
}