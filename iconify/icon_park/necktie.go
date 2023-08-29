package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Necktie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M15 36L21 13H27L33 36L24 44L15 36Z"/><path d="M21 4H27L30 6L27 13H21L18 6L21 4Z"/></g>`),
		g.Group(children),
	)
}