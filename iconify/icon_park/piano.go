package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Piano(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="40" height="16" x="4" y="8" fill="#2F88FF"/><rect width="40" height="16" x="4" y="24"/><path d="M10 24V32"/><path d="M16 24V32"/><path d="M26 24V32"/><path d="M32 24V32"/><path d="M38 24V32"/></g>`),
		g.Group(children),
	)
}