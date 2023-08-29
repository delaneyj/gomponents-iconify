package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="24" r="20"/><path fill="#2F88FF" d="M24 4C26.7322 4 29.4355 4.55981 31.943 5.64491C34.4505 6.73 36.709 8.31736 38.5794 10.3091L24 24V4Z"/></g>`),
		g.Group(children),
	)
}