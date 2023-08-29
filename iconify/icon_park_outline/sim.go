package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sim(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M8 4h24.889L40 11.273V44H8V4Z"/><path d="M33 26H15v10h18V26Z"/><path stroke-linecap="round" d="M15 12v6m6-6v6m6-6v6"/></g>`),
		g.Group(children),
	)
}