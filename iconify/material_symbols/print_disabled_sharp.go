package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrintDisabledSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.775 22.625L18 20.85V21H6v-4H2V8h3.15L1.375 4.225L2.8 2.8l18.4 18.4l-1.425 1.425ZM8 19h8v-.15L12.15 15H8v4Zm11.85-2H22V8H10.85l9 9Zm-10-10l-4-4H18v4H9.85ZM18 12.5q-.425 0-.713-.288T17 11.5q0-.425.288-.713T18 10.5q.425 0 .713.288T19 11.5q0 .425-.288.713T18 12.5Z"/>`),
		g.Group(children),
	)
}