package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Change(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M18 31H38V5"/><path d="M30 21H10V43"/><path d="M44 11L38 5L32 11"/><path d="M16 37L10 43L4 37"/></g>`),
		g.Group(children),
	)
}