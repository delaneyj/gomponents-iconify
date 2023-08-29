package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatBrush(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M34 5H6v15h28V5Z"/><path stroke-linecap="round" d="M34.025 12H43v16.101l-24 3.1V43"/></g>`),
		g.Group(children),
	)
}