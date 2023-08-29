package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProjectorTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path d="M4 12a2 2 0 0 1 2-2h36a2 2 0 0 1 2 2v20a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V12Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 19h6m-6 6h4m-2 15v-6m20 6v-6"/><circle cx="31" cy="22" r="5"/></g>`),
		g.Group(children),
	)
}