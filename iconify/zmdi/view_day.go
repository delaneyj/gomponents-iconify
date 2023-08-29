package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewDay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 384v-64h405v64H0zm384-277q9 0 15 6t6 15v128q0 9-6 15t-15 6H21q-8 0-14.5-6T0 256V128q0-9 6.5-15t14.5-6h363zM0 0h405v64H0V0z"/>`),
		g.Group(children),
	)
}