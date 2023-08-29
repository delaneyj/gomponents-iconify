package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NineOclock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2m6 32h-2.555c-.35.6-.845 1.096-1.445 1.445V38h-4v-2.555A4.004 4.004 0 0 1 28.555 34H14v-4h14.555c.35-.6.846-1.096 1.445-1.445V6h4v22.554c.601.35 1.096.845 1.445 1.446H38v4"/><circle cx="32" cy="32" r="3" fill="currentColor"/>`),
		g.Group(children),
	)
}