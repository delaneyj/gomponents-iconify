package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NailPolishOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="32" height="24" x="8" y="20" rx="2"/><path d="M17 4h14v16H17zm5 28h4l1 5h-6l1-5Zm2-12v12m7-12H17"/></g>`),
		g.Group(children),
	)
}