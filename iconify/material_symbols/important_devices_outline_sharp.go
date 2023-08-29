package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImportantDevicesOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 21v-9.95h6V21h-6Zm1-2h4v-5.95h-4V19Zm-9 2v-2h2v-2H2V3h18v6.05h-2V5H4v10h10v2h-2v2h2v2H8Zm.7-7.25L11 12l2.3 1.75l-.85-2.85l2.3-1.85H11.9l-.9-2.8l-.9 2.8H7.25l2.3 1.85l-.85 2.85ZM11 10Z"/>`),
		g.Group(children),
	)
}