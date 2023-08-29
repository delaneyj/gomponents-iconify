package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CssSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.5 15v-2H11v.5h2v-1H9.5V9h5v2H13v-.5h-2v1h3.5V15h-5Zm6.5 0v-2h1.5v.5h2v-1H16V9h5v2h-1.5v-.5h-2v1H21V15h-5ZM3 15V9h5v2H6.5v-.5h-2v3h2V13H8v2H3Z"/>`),
		g.Group(children),
	)
}