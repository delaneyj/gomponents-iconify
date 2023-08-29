package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Income(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M31 34L43 34"/><path d="M43 26V10C43 8.34315 41.6569 7 40 7H8C6.34315 7 5 8.34315 5 10V38C5 39.6569 6.34315 41 8 41H28.4706"/><path d="M36 39L31 34L35.9996 29"/><path d="M15 15L20 21L25 15"/><path d="M14 27H26"/><path d="M14 21H26"/><path d="M20 21V33"/></g>`),
		g.Group(children),
	)
}