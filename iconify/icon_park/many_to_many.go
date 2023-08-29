package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ManyToMany(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M42 7H6C4.89543 7 4 7.89543 4 9V39C4 40.1046 4.89543 41 6 41H42C43.1046 41 44 40.1046 44 39V9C44 7.89543 43.1046 7 42 7Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M30 30V18L38 30V18"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M10 30V18L18 30V18"/><path stroke="#fff" stroke-linecap="round" d="M24 20V21"/><path stroke="#fff" stroke-linecap="round" d="M24 27V28"/></g>`),
		g.Group(children),
	)
}