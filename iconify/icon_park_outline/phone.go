package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Phone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linejoin="round" d="M8 30h32v12a2 2 0 0 1-2 2H10a2 2 0 0 1-2-2V30Zm32 0V6a2 2 0 0 0-2-2H10a2 2 0 0 0-2 2v24"/><path stroke-linecap="round" d="M22 37h4"/></g>`),
		g.Group(children),
	)
}