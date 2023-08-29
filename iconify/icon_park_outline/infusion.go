package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Infusion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M38 30c0 7.732-6.268 14-14 14s-14-6.268-14-14S24 4 24 4s14 18.268 14 26Z"/><path stroke-linecap="round" d="M18 30h12m-6-6v12"/></g>`),
		g.Group(children),
	)
}