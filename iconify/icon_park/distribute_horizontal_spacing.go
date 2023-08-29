package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DistributeHorizontalSpacing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path d="M8 43L8 5"/><path d="M40 43L40 5"/><rect width="8" height="20" x="20" y="14" fill="#2F88FF" stroke-linejoin="round"/></g>`),
		g.Group(children),
	)
}