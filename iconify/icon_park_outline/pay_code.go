package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PayCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M44 4H4v40h40V4Z"/><path stroke-linecap="round" d="M12 16v16m8-16v16m8-16v16m8-16v16"/></g>`),
		g.Group(children),
	)
}