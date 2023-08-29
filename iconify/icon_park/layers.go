package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Layers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M4 11.9143L24 19L44 11.9143L24 5L4 11.9143Z"/><path stroke-linecap="round" d="M4 20L24 27L44 20"/><path stroke-linecap="round" d="M4 28L24 35L44 28"/><path stroke-linecap="round" d="M4 36L24 43L44 36"/></g>`),
		g.Group(children),
	)
}