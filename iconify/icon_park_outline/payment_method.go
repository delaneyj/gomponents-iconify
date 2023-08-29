package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaymentMethod(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M38 10L10 38M6 6l6 8l6-8M5 14h14M5 20h14m-7-6v12m20.846 0H42v16H22v-5.85"/>`),
		g.Group(children),
	)
}