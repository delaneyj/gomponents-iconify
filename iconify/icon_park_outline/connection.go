package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Connection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="m39 34l5 5l-5 5"/><path d="M8 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path stroke-linecap="round" d="M12 8h8a4 4 0 0 1 4 4v23a4 4 0 0 0 4 4h16"/></g>`),
		g.Group(children),
	)
}