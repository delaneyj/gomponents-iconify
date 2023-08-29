package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CooperativeHandshake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 40L36 28L32 32L28 36L24 40ZM24 40L4 20L16 8L24 16"/><path d="M17 23L32 8L44 20L36 28L28 20L22 26L17 23ZM17 23L24 16"/><path d="M28 36L25 33"/><path d="M32 32L29 29"/></g>`),
		g.Group(children),
	)
}