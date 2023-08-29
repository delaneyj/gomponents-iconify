package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="24" r="20"/><path fill="#2F88FF" d="M24 4C27.9021 4 31.719 5.14149 34.9805 7.28385C38.2419 9.42621 40.8054 12.4758 42.3551 16.057C43.9048 19.6382 44.3731 23.5946 43.7022 27.4386C43.0313 31.2826 41.2506 34.8464 38.5794 37.6909L24 24V4Z"/></g>`),
		g.Group(children),
	)
}