package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Laptop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M8 8h32v20H8zm0 20L4 41h40l-4-13"/><path d="M19.9 35h8.2l1.9 6H18l1.9-6Z"/></g>`),
		g.Group(children),
	)
}