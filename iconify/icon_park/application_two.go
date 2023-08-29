package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApplicationTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" stroke="#000" stroke-width="4"><circle cx="34.5" cy="13.5" r="6.5"/><circle cx="34.5" cy="34.5" r="6.5"/><circle cx="13.5" cy="13.5" r="6.5"/><circle cx="13.5" cy="34.5" r="6.5"/></g>`),
		g.Group(children),
	)
}