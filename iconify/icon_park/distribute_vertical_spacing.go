package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DistributeVerticalSpacing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path d="M5 8H43"/><path d="M5 40H43"/><rect width="20" height="8" x="14" y="20" fill="#2F88FF" stroke-linejoin="round"/></g>`),
		g.Group(children),
	)
}