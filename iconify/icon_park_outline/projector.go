package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Projector(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M38 16h6v16H4V16h20M6 38v-6h6v6H6Zm36 0v-6h-6v6h6ZM10 24h8"/><circle cx="31" cy="16" r="7" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/><circle cx="31" cy="16" r="3" fill="currentColor"/></g>`),
		g.Group(children),
	)
}