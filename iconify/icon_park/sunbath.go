package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sunbath(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path d="M11 36V40"/><path stroke-linejoin="round" d="M6 36L40 36"/><path stroke-linejoin="round" d="M8 30H24"/><path stroke-linejoin="round" d="M37 11L34 16"/><path d="M35 36V40"/><path stroke-linejoin="round" d="M44 10L29 36"/><path stroke-linejoin="round" d="M39 19L44 28"/></g>`),
		g.Group(children),
	)
}