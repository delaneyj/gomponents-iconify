package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func XiaoduHome(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 20L24 6l20 14v22H4V20Z"/><path d="M12.687 26.686a16.002 16.002 0 0 1 22.627 0m-16.971 5.657a8 8 0 0 1 11.314 0"/></g>`),
		g.Group(children),
	)
}