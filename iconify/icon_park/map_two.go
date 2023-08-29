package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M44 10H4V40H44V10Z"/><path stroke="#fff" stroke-linecap="round" d="M10 16L38 34"/><path stroke="#fff" stroke-linecap="round" d="M38 16L24 35"/><path stroke="#fff" stroke-linecap="round" d="M24 16L10 34"/></g>`),
		g.Group(children),
	)
}