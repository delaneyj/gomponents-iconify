package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Liqueur(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M23 31L7.41 13h31.177L22.999 31Zm0 1v10"/><path d="M17 44h12"/><path stroke-linejoin="round" d="M16 23h14m5-5h3a6 6 0 1 0-5.917-5M11 17.144L19.534 27M35 17.143l-8.79 10.15"/></g>`),
		g.Group(children),
	)
}