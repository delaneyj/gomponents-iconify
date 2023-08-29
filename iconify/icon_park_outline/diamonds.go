package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Diamonds(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M12 8h24l8 10l-20 24L4 18l8-10ZM4 18h40M24 42l-8-24m8 24l8-24"/><path d="m8 13l-4 5l20 24l20-24l-4-5"/></g>`),
		g.Group(children),
	)
}