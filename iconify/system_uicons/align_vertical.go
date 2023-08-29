package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m15.5 7.5l-3 3l3 3m4-3h-7m-7-3l3 3l-3 3m3-3h-7m9-7v14"/>`),
		g.Group(children),
	)
}