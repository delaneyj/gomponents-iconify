package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnvelopeOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M36 16V8H20H4V20V32H12"/><path d="M12 40H44V28V16H28H12V28V40Z"/><path stroke-linecap="round" d="M12 16L28 28L44 16"/><path stroke-linecap="round" d="M32 16H12V31"/><path stroke-linecap="round" d="M44 31V16H24"/></g>`),
		g.Group(children),
	)
}