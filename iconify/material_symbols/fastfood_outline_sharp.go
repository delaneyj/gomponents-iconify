package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastfoodOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 15q0-1.5.65-2.625t1.7-1.875q1.05-.75 2.4-1.125T8.5 9q1.4 0 2.75.375t2.4 1.125q1.05.75 1.7 1.875T16 15H1Zm17 8v-2h1.4l1.4-14h-9.55L11 5h5V1h2v4h5l-1.8 18H18Zm0-2h1.4H18Zm-4.675-8q-.725-1-2.037-1.5T8.5 11q-1.475 0-2.788.5T3.676 13h9.65ZM8.5 13ZM1 19v-2h15v2H1Zm0 4v-2h15v2H1Z"/>`),
		g.Group(children),
	)
}