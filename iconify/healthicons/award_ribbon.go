package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AwardRibbon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M13 20c0-6.075 4.925-11 11-11s11 4.925 11 11s-4.925 11-11 11s-11-4.925-11-11Z"/><path fill-rule="evenodd" d="M31 34.391c5.328-2.597 9-8.065 9-14.391c0-8.837-7.163-16-16-16c-8.836 0-16 7.163-16 16c0 6.327 3.672 11.796 9.001 14.392V43a1 1 0 0 0 1.555.832L24 40.202l5.445 3.63A1 1 0 0 0 31 43v-8.609ZM24 34c7.732 0 14-6.268 14-14S31.732 6 24 6s-14 6.268-14 14s6.268 14 14 14Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}