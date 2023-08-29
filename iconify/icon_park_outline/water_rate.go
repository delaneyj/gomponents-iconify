package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterRate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M24 44c8.837 0 16-6.512 16-14.546C40 21.576 34.667 13.091 24 4C13.333 13.09 8 21.576 8 29.454C8 37.489 15.163 44 24 44Z"/><path stroke-linecap="round" d="M18.857 19L24 24.368L29.143 19M18 26.158h12m-12 5.368h12m-6-5.368V36"/></g>`),
		g.Group(children),
	)
}