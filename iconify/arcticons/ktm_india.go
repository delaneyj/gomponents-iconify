package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KtmIndia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.52 25.2c0-1.66 1.34-3 3-3h0c1.66 0 3 1.34 3 3V30m-6-7.8V30m6-4.8c0-1.66 1.34-3 3-3h0c1.66 0 3 1.34 3 3V30M10.48 18v12m0-2.55l5.43-5.4m-3.7 3.68l4.27 4.25m-2.14-10.09h14.18m-7.96 10.09V19.89"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33c-1.1 0-2 .9-2 2v33c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2v-33c0-1.1-.9-2-2-2Z"/>`),
		g.Group(children),
	)
}