package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Atm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.375 15v-4.5h-2.25V9h6v1.5h-2.25V15h-1.5ZM2 15v-5q0-.425.288-.713T3 9h3q.425 0 .713.288T7 10v5H5.5v-1.5h-2V15H2Zm1.5-3h2v-1.5h-2V12Zm12 3v-5q0-.425.288-.713T16.5 9H21q.425 0 .713.288T22 10v5h-1.5v-4.5h-1V14H18v-3.5h-1V15h-1.5Z"/>`),
		g.Group(children),
	)
}