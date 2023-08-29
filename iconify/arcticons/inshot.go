package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inshot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 9.611h28.889a4.012 4.012 0 0 1 4 4v18.097m4.111 6.681H13.611a4.012 4.012 0 0 1-4-4v-18.2m0-6.578V5.5m28.778 32.889V42.5"/>`),
		g.Group(children),
	)
}