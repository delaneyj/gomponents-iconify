package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FluidMedOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 22q-.825 0-1.413-.588T17 20v-1.1q-1.725-.35-2.863-1.713T13 14V7h10v7q0 1.825-1.137 3.188T19 18.9V20h3v2h-3Zm.75-9H21V9h-6v2h.75q.825 0 1.563.375T18.55 12.4q.2.3.525.45t.675.15ZM18 17q.975 0 1.75-.563T20.825 15H19.75q-.825 0-1.563-.363T16.95 13.6q-.225-.275-.537-.437T15.75 13H15v1q0 1.275.863 2.138T18 17Zm-1.05-4.6ZM8 22.5L6 21v-4H5q-.825 0-1.413-.588T3 15V7.5H2v-2h4V4H4.5V2h5v2H8v1.5h4v2h-1V15q0 .825-.588 1.413T9 17H8v5.5ZM5 15h4v-1.5H6.5V12H9v-1.5H6.5V9H9V7.5H5V15Z"/>`),
		g.Group(children),
	)
}