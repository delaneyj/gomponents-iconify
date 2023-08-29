package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableLamp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M31 44L42 28L22 12"/><path fill="#2F88FF" d="M6 12L16 22L24 8L20 4L6 12Z"/><path d="M38 44H12"/><path d="M17 44V40"/></g>`),
		g.Group(children),
	)
}