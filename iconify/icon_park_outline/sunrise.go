package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sunrise(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 24h3m3-14l2 2m12-8v3M14 24c0-5.522 4.478-10 10-10s10 4.478 10 10a9.987 9.987 0 0 1-4.215 8.158M38 10l-2 2m8 12h-3m-3.019 13.982l-1.62-1.62"/><path d="M23.5 28c-3 0-9.5.2-9.5 3s4.606 2.79 7 3c2 .175 5.462 1.688 5 4c-1 5-17 4-21 4"/></g>`),
		g.Group(children),
	)
}