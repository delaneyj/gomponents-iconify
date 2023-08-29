package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GameTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M20 15L24 19L28 15V4H20V15Z"/><path d="M20 33L24 29L28 33V44H20V33Z"/><path d="M33 28L29 24L33 20L44 20L44 28L33 28Z"/><path d="M15 20L19 24L15 28L4 28L4 20L15 20Z"/></g>`),
		g.Group(children),
	)
}