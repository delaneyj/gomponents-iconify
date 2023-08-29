package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Intersection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M31 31H40C41.1046 31 42 30.1046 42 29V8C42 6.89543 41.1046 6 40 6H19C17.8954 6 17 6.89543 17 8V17"/><path d="M17 17H8C6.89543 17 6 17.8954 6 19V40C6 41.1046 6.89543 42 8 42H29C30.1046 42 31 41.1046 31 40V31"/><rect width="14" height="14" x="17" y="17" fill="#2F88FF" rx="2"/></g>`),
		g.Group(children),
	)
}