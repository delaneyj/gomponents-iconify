package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Disney(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 11.15c-2.68-.3-2.5-1.15-2.5-1.58s.87-2.08 8.91-2.08c9.63 0 30.09 7.26 30.09 20s-17.88 9.82-21.42 9S9.85 32 9.85 28.18c0-2.78 7.13-4.35 14.59-4.35s10 2.09 10 4c0 .84-.4 1.84-2.3 2.38"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.66 15.3a190.17 190.17 0 0 0 0 25.21"/>`),
		g.Group(children),
	)
}