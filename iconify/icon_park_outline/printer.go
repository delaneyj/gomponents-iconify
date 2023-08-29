package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Printer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M37 32H11v12h26V32Z"/><path stroke-linecap="round" d="M4 20h40v18h-6.983v-6H10.98v6H4V20Z" clip-rule="evenodd"/><path d="M38 4H10v16h28V4Z"/></g>`),
		g.Group(children),
	)
}