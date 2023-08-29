package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrintDisabledOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.775 22.625L18 20.85V21H6v-4H2V8h3.15L1.375 4.225L2.8 2.8l18.4 18.4l-1.425 1.425ZM8 19h8v-.15L12.15 15H8v4Zm11.85-2l-2-2H20v-5h-7.15l-2-2H22v9h-2.15ZM4 15h2v-2h4.15l-3-3H4v5Zm12-7V5H7.85l-2-2H18v5h-2Zm2 4.5q.425 0 .713-.288T19 11.5q0-.425-.288-.713T18 10.5q-.425 0-.713.288T17 11.5q0 .425.288.713T18 12.5ZM12.85 10H20h-7.15ZM4 10h3.15H4Z"/>`),
		g.Group(children),
	)
}