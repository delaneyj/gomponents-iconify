package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AtmSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.375 15v-4.5h-2.25V9h6v1.5h-2.25V15h-1.5ZM2 15V9h5v6H5.5v-1.5h-2V15H2Zm1.5-3h2v-1.5h-2V12Zm12 3V9H22v6h-1.5v-4.5h-1V14H18v-3.5h-1V15h-1.5Z"/>`),
		g.Group(children),
	)
}