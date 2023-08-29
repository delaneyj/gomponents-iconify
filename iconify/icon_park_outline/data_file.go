package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataFile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M8 44V4h23l9 10.5V44H8Z"/><path d="M15 28h6v7h-6zm-1 7h20M21 23h6v12h-6zm6-5h6v17h-6z"/></g>`),
		g.Group(children),
	)
}