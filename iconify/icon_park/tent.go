package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path stroke="#000" d="M10 12L4 36H16"/><path fill="#2F88FF" stroke="#000" d="M38 12H10L16 36H44L38 12Z"/><path stroke="#fff" d="M12 18H39"/><path stroke="#000" d="M10 12L13 24"/><path stroke="#000" d="M38 12L41 24"/></g>`),
		g.Group(children),
	)
}