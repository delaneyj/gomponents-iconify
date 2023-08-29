package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GenderFemaleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M200 96a72 72 0 1 1-72-72a72 72 0 0 1 72 72Z" opacity=".2"/><path d="M208 96a80 80 0 1 0-88 79.6V200H88a8 8 0 0 0 0 16h32v24a8 8 0 0 0 16 0v-24h32a8 8 0 0 0 0-16h-32v-24.4A80.11 80.11 0 0 0 208 96ZM64 96a64 64 0 1 1 64 64a64.07 64.07 0 0 1-64-64Z"/></g>`),
		g.Group(children),
	)
}