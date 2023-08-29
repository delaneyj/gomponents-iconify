package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridOnSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21h4.675v-4.675H3V21Zm6.675 0h4.65v-4.675h-4.65V21Zm6.65 0H21v-4.675h-4.675V21ZM3 14.325h4.675v-4.65H3v4.65Zm6.675 0h4.65v-4.65h-4.65v4.65Zm6.65 0H21v-4.65h-4.675v4.65ZM3 7.675h4.675V3H3v4.675Zm6.675 0h4.65V3h-4.65v4.675Zm6.65 0H21V3h-4.675v4.675Z"/>`),
		g.Group(children),
	)
}