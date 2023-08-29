package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sort(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" stroke="#000" stroke-linejoin="round" stroke-width="4"><path d="M24 42L15 29H33L24 42Z"/><path d="M24 6L15 19H33L24 6Z"/></g>`),
		g.Group(children),
	)
}