package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewCozy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 10.5V3h7.5v7.5H3ZM3 21v-7.5h7.5V21H3Zm10.5-10.5V3H21v7.5h-7.5Zm0 10.5v-7.5H21V21h-7.5Z"/>`),
		g.Group(children),
	)
}