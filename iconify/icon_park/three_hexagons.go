package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeHexagons(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 27L14 21L4 27V39L14 45L24 39V27Z"/><path d="M44 27L34 21L24 27V39L34 45L44 39V27Z"/><path d="M34 9L24 3L14 9V21L24 27L34 21V9Z"/></g>`),
		g.Group(children),
	)
}