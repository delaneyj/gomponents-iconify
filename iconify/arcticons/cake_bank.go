package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CakeBank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.017 20.063v7.856m5.63 0l-3.563-3.928l3.563-3.901m-3.563 3.901h-2.067m-11.594 2.912a3.459 3.459 0 0 1-2.453 1.016h0a3.47 3.47 0 0 1-3.47-3.47v-.916a3.47 3.47 0 0 1 3.47-3.47h0c.958 0 1.825.389 2.453 1.017m19.84 2.911h3.414m1.823 3.928h-5.237v-7.856H39.5m-22.908 7.85l3.405-7.832l3.404 7.856"/>`),
		g.Group(children),
	)
}