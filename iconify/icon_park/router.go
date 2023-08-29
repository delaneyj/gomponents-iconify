package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Router(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="40" height="14" x="4" y="28" fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="2"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 35L22 35"/><rect width="4" height="4" x="30" y="33" fill="#fff" rx="2"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M12 28L12 8"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M36 28V8"/></g>`),
		g.Group(children),
	)
}