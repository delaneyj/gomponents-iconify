package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OnlineMeeting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M44 8H4v30h15l5 5l5-5h15V8Z"/><path d="M12 19v8m8-11v14m8-9v4m8-6v8"/></g>`),
		g.Group(children),
	)
}